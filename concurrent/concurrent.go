package concurrent

type ActionFunc func(data interface{})

var Conncurrent Concurrenter

type Concurrenter struct {
	ConChan chan int // 并发控制

}

func (con Concurrenter) Goroutine(data interface{}, dofunc ActionFunc) {
	dofunc(data)
	<-con.ConChan
}



func (con Concurrenter) Run(data []interface{}, conNum int, do ActionFunc) {
	con.ConChan = make(chan int, conNum)
	for i, item := range data {
		go con.Goroutine(item, do)
		con.ConChan<-i
	}
	return

}


