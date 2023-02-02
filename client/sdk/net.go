package sdk

type connect struct {
	serverAddr         string
	sendChan, recvChan chan *Message
}

func newConnect(serverAddr string) *connect {
	return &connect{
		serverAddr: serverAddr,
		sendChan:   make(chan *Message),
		recvChan:   make(chan *Message),
	}
}

func (c *connect) send(data *Message) {
	c.recvChan <- data
}

//go的管道操作
// ch := make(chan int, 100)
// defer close(ch)
//

func (c *connect) recv() <-chan *Message { //"<-chan *Message" 是只读channel
	return c.recvChan
}

func (c *connect) close() {

}
