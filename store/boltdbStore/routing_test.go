package boltdbStore

import (
	"fmt"
	"github.com/polarismesh/polaris-server/common/model"
	"strconv"
	"testing"
	"time"
)

const (
	routeCount = 5
)

func TestRoutingStore_CreateRoutingConfig(t *testing.T){
	handler, err := NewBoltHandler(&BoltConfig{FileName: "./table.bolt"})
	if nil != err {
		t.Fatal(err)
	}

	defer handler.Close()

	rStore := &routingStore{handler: handler}

	for i := 0; i < routeCount; i++ {
		rStore.CreateRoutingConfig(&model.RoutingConfig{
			ID: "testid" + strconv.Itoa(i),
			InBounds: "v1" + strconv.Itoa(i),
			OutBounds: "v2" + strconv.Itoa(i),
			Revision: "revision" + strconv.Itoa(i) ,
			Valid: true,
			CreateTime: time.Now(),
			ModifyTime: time.Now(),
		})
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestRoutingStore_GetRoutingConfigs(t *testing.T){
	handler, err := NewBoltHandler(&BoltConfig{FileName: "./table.bolt"})
	if nil != err {
		t.Fatal(err)
	}

	defer handler.Close()

	rStore := &routingStore{handler: handler}

	totalCount, rs, err := rStore.GetRoutingConfigs(nil, 0, 20)
	if err != nil {
		t.Fatal(err)
	}
	if totalCount != routeCount {
		t.Fatal(fmt.Sprintf("routing total count not match, expect %d, got %d", routeCount, totalCount))
	}
	if len(rs) != routeCount {
		t.Fatal(fmt.Sprintf("routing count not match, expect %d, got %d", routeCount, len(rs)))
	}
	for _, r := range rs {
		fmt.Printf("routing conf is %+v\n", r.Config)
	}
}

func TestRoutingStore_UpdateRoutingConfig(t *testing.T){
	handler, err := NewBoltHandler(&BoltConfig{FileName: "./table.bolt"})
	if nil != err {
		t.Fatal(err)
	}

	defer handler.Close()

	rStore := &routingStore{handler: handler}

	for i := 0; i < routeCount; i++ {

		conf := &model.RoutingConfig{
			ID: "testid" + strconv.Itoa(i),
			InBounds: "vv1" + strconv.Itoa(i),
			OutBounds: "vv2" + strconv.Itoa(i),
			Revision: "revi" + strconv.Itoa(i),
		}

		err := rStore.UpdateRoutingConfig(conf)
		if err != nil {
			t.Fatal(err)
		}
	}

	// check update result
	totalCount, rs, err := rStore.GetRoutingConfigs(nil, 0, 20)
	if err != nil {
		t.Fatal(err)
	}
	if totalCount != routeCount {
		t.Fatal(fmt.Sprintf("routing total count not match, expect %d, got %d", routeCount, totalCount))
	}
	if len(rs) != routeCount {
		t.Fatal(fmt.Sprintf("routing count not match, expect %d, got %d", routeCount, len(rs)))
	}
	for _, r := range rs {
		fmt.Printf("routing conf is %+v\n", r.Config)
	}
}

func TestRoutingStore_GetRoutingConfigsForCache(t *testing.T){
	handler, err := NewBoltHandler(&BoltConfig{FileName: "./table.bolt"})
	if nil != err {
		t.Fatal(err)
	}

	defer handler.Close()

	rStore := &routingStore{handler: handler}

	// get create modify time
	totalCount, rs, err := rStore.GetRoutingConfigs(nil, 0, 20)
	if err != nil {
		t.Fatal(err)
	}
	if totalCount != routeCount {
		t.Fatal(fmt.Sprintf("routing total count not match, expect %d, got %d", routeCount, totalCount))
	}
	if len(rs) != routeCount {
		t.Fatal(fmt.Sprintf("routing count not match, expect %d, got %d", routeCount, len(rs)))
	}

	rss, err := rStore.GetRoutingConfigsForCache(rs[2].Config.ModifyTime, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(rss) != routeCount - 2 {
		t.Fatal(fmt.Sprintf("routing config count mismatch, except %d, got %d", routeCount - 2, len(rss)))
	}
}

func TestRoutingStore_GetRoutingConfigWithService(t *testing.T){
	// todo service first


}

func TestRoutingStore_GetRoutingConfigWithID(t *testing.T){
	handler, err := NewBoltHandler(&BoltConfig{FileName: "./table.bolt"})
	if nil != err {
		t.Fatal(err)
	}

	defer handler.Close()

	rStore := &routingStore{handler: handler}

	rc, err := rStore.GetRoutingConfigWithID("testid0")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("get routing conf %+v\n", rc)
}


func TestRoutingStore_DeleteRoutingConfig(t *testing.T){
	handler, err := NewBoltHandler(&BoltConfig{FileName: "./table.bolt"})
	if nil != err {
		t.Fatal(err)
	}

	defer handler.Close()

	rStore := &routingStore{handler: handler}
	for i := 0; i < routeCount; i++ {
		rStore.DeleteRoutingConfig("testid" + strconv.Itoa(i))
	}
}






