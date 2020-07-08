package Cache

type Client interface {
	AddNode(Server, ...interface{}) error
	RemoveNode(index int) error
	ListNode() []Server
	selectNode(string) Server
}

type CommonClient struct {
	//nodeList     []Server
	nodeSelector NodeSelector
}

func (cc *CommonClient) AddNode(svr Server, opt ...interface{}) error {
	//cc.nodeList = append(cc.nodeList, svr)
	cc.nodeSelector.AddNode(svr, opt)
	return nil
}

func (cc *CommonClient) RemoveNode(index int) error {
	cc.nodeSelector.RemoveNode(index)
	//// 移除数组元素
	//cc.nodeList = append(cc.nodeList[0:index], cc.nodeList[index+1:]...)
	return nil
}

func (cc *CommonClient) ListNode() []Server {
	return cc.nodeSelector.ListNode()
}

func (cc *CommonClient) selectNode(key string) Server {
	return cc.nodeSelector.selectNode(key)
}

func (cc *CommonClient) Set(k string, v interface{}) error {
	ser := cc.selectNode(k)
	ser.Set(k, v)
	return nil
}

func (cc *CommonClient) Get(k string) (interface{}, bool) {
	ser := cc.selectNode(k)
	return ser.Get(k)
}

func (cc *CommonClient) Del(k string) error {
	ser := cc.selectNode(k)
	return ser.Del(k)
}

func NewCommonClient(selector NodeSelector) *CommonClient {
	return &CommonClient{selector}
}