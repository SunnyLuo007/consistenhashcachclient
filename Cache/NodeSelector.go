package Cache

import (
	"awesomeProject/consistenhashcachclient/utils/hashcode"
)

type NodeSelector interface {
	Client
}

type SimpleNodeSelector struct {
	nodeList []Server
}

func (sns *SimpleNodeSelector) AddNode(svr Server, opt ...interface{}) error {
	sns.nodeList = append(sns.nodeList, svr) // 增加物理节点
	return nil
}

func (sns *SimpleNodeSelector) RemoveNode(index int) error {
	// 移除数组元素
	if index >= len(sns.nodeList) {
		return nil
	}
	sns.nodeList = append(sns.nodeList[0:index], sns.nodeList[index+1:]...)
	return nil
}

func (sns *SimpleNodeSelector) ListNode() []Server {
	return sns.nodeList
}

func (sns *SimpleNodeSelector) selectNode(key string) Server {
	if len(sns.nodeList) == 0{
		panic("there is not any cache server node to select")
	}
	hashCode := hashcode.String(key)
	return sns.nodeList[hashCode%len(sns.nodeList)]
}
