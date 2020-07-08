package Cache

type Server interface {
	Set(string, interface{}) error
	Get(string) (interface{}, bool)
	Del(string) error
}

type SimpleServer struct {
	Name string
	data map[string]interface{}
}

func (ss *SimpleServer) Set(k string, v interface{}) error {
	ss.data[k] = v
	return nil
}

func (ss *SimpleServer) Get(k string) (interface{}, bool) {
	v, ok := ss.data[k]
	return v, ok
}

func (ss *SimpleServer) Del(k string) error {
	delete(ss.data, k)
	return nil
}

func (ss *SimpleServer) String() string {
	return ss.Name
}

func (ss *SimpleServer) GetData() map[string]interface{} {
	return ss.data
}

func NewSimpleServer(name string) Server {
	return &SimpleServer{Name: name, data: make(map[string]interface{})}
}

