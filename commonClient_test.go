package ConstHash

import (
	Cache "awesomeProject/consistenhashcachclient/Cache"
	"awesomeProject/consistenhashcachclient/utils"
	"strconv"
	"testing"
)

// 缓存服
var (
	cacheSvrList []Cache.Server = []Cache.Server{
		Cache.NewSimpleServer("svr1"),
		Cache.NewSimpleServer("svr2"),
		Cache.NewSimpleServer("svr3"),
		Cache.NewSimpleServer("svr4"),
		//Cache.NewSimpleServer("svr5"),
		//Cache.NewSimpleServer("svr6"),
		//Cache.NewSimpleServer("svr7"),
	}
)

// 客户端
var (
	//nodeSelector Cache.NodeSelector = &Cache.SimpleNodeSelector{}
	consistenHahsSlector Cache.NodeSelector = Cache.NewConsistenHashNodeSelector(200)
	cacheClient  *Cache.CommonClient      = Cache.NewCommonClient(consistenHahsSlector)
)

// 添加物理节点
func TestAddNode(t *testing.T) {
	for _, s := range cacheSvrList {
		ok := cacheClient.AddNode(s)
		if ok != nil {
			t.Fatal("add node fail:", ok)
		}
	}
	t.Log(cacheClient.ListNode())
}

//存储键值对
func TestSetAndGetKV(t *testing.T) {
	for i := 1; i <= 1000000; i++ {
		k := strconv.Itoa(i)
		err := cacheClient.Set(k, i)
		if err != nil {
			t.Fatal("set kv error:", err)
		}
	}
	//获取键值对
	for i := 1; i <= 1000000; i++ {
		k := strconv.Itoa(i)
		v, ok := cacheClient.Get(k)
		//t.Log(k,v)
		v = v.(int)
		if !(ok && v == i) {
			t.Fatalf("get kv error at:%v,%T", v, v)
		}

	}
}

func TestAvg(t *testing.T) {
	dataSizeList := make([]float64, len(cacheSvrList))
	for i, s := range cacheSvrList {
		dataSize := len(s.(*Cache.SimpleServer).GetData())
		//t.Log("svr:",s,"data_size:",dataSize)
		dataSizeList[i] = float64(dataSize)
	}
	t.Log("data size list:", dataSizeList)
	t.Log("stdev:", utils.Stdev(dataSizeList...))
}

//移除缓存服务物理节点
func TestDelNode(t *testing.T) {
	ok := cacheClient.RemoveNode(1)
	if ok != nil {
		t.Fatal("client remove node fail:", ok)
	}
	ok = cacheClient.RemoveNode(0)
	if ok != nil {
		t.Fatal("client remove node fail:", ok)
	}
	ok = cacheClient.RemoveNode(1)
	if ok != nil {
		t.Fatal("client remove node fail:", ok)
	}
	svrList := cacheClient.ListNode()
	s := svrList[0]
	if s.(*Cache.SimpleServer).Name != "svr3" {
		t.Fatal("delete node fail!")
	}
	t.Log(cacheClient.ListNode())
}
