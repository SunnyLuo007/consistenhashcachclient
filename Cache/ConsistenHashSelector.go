package Cache

import (
	"awesomeProject/consistenhashcachclient/utils/hashcode"
	"fmt"
	"sort"
	"strconv"
)

type ConsistenHashNodeSelector struct {
	nodeList        []Server
	virtualNodeList []int // 虚拟节点
	nodeMap         map[int]Server
	splitNums       int // 一个物理节点分解物理节点数量
}

// 添加虚拟节点
func (chns *ConsistenHashNodeSelector) addVirtualNode(svr Server) {
	// 获得svr的内存地址作为初始值
	prefix := fmt.Sprintf("%p", svr)
	// 默认为1个节点
	if chns.splitNums == 0 {
		chns.splitNums = 1
	}
	// 生成虚拟节点、存入列表、映射
	for i := 0; i < chns.splitNums; i++ {
		hc := hashcode.String(prefix + strconv.Itoa(i))
		chns.virtualNodeList = append(chns.virtualNodeList, hc)
		chns.nodeMap[hc] = svr
	}
	// 排序列表
	sort.Ints(chns.virtualNodeList)
}

func (chns *ConsistenHashNodeSelector) AddNode(svr Server, opt ...interface{}) error {
	chns.nodeList = append(chns.nodeList, svr) // 增加物理节点
	// 添加虚拟节点
	chns.addVirtualNode(svr)
	return nil
}

func (chns *ConsistenHashNodeSelector) RemoveNode(index int) error {
	// 移除数组元素
	if index >= len(chns.nodeList) {
		return nil
	}
	// todo: 移除虚拟节点
	chns.nodeList = append(chns.nodeList[0:index], chns.nodeList[index+1:]...)
	return nil
}

func (chns *ConsistenHashNodeSelector) ListNode() []Server {
	return chns.nodeList
}

func (chns *ConsistenHashNodeSelector) selectNode(key string) Server {
	if len(chns.nodeList) == 0 {
		panic("there is not any cache server node to select")
	}
	hc := hashcode.String(key)
	// 查找hc所在属的虚拟节点
	vNode := chns.findNode(hc)
	return chns.nodeMap[vNode]
}

func (chns *ConsistenHashNodeSelector) findNode(hc int) int {
	// 二分查找
	l, r, mid := 0, len(chns.virtualNodeList)-1, 0
	for r >= l {
		mid = l + (r-l)/2 // 防止加法溢出
		if hc == chns.virtualNodeList[mid] { // 找到
			return hc
		} else if hc > chns.virtualNodeList[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 没找到，考察最后mid指向元素的大小
	if hc > chns.virtualNodeList[mid] { // 考察元素比mid指向的值要大，则取mid的下一个元素
		if (mid + 1) == len(chns.virtualNodeList) { // 防止越界
			mid = 0
		}
		return chns.virtualNodeList[mid+1]
	} else {
		if mid == 0 {
			mid = len(chns.virtualNodeList) - 1
		}
		return chns.virtualNodeList[mid-1]
	}
}

// @params: vNodeNum int  一个物理节点对应的虚拟节点数量
func NewConsistenHashNodeSelector(vNodeNum int) NodeSelector {
	return &ConsistenHashNodeSelector{make([]Server, 0),
		make([]int, 0),
		make(map[int]Server),
		vNodeNum,
	}
}
