package server

type Fiber struct{}

func NewFiberServer() Server {
	return &Fiber{}
}

func (s Fiber) Run() {

}
