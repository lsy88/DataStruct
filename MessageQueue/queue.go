package MessageQueue

import "time"

//消息队列
type MessageQueue interface {
	Send(message interface{})                           //向消息队列中发送
	Pull(size int, timeout time.Duration) []interface{} //从消息队列拉取信息
	Length() int                                        //消息队列的大小
	Capacity() int                                      //消息队列的容量
}

type MesQueue struct {
	Queue chan interface{}
	Cap   int //容量
	Size  int //大小
}

func (this *MesQueue) Send(mes Message) {
	select {
	case this.Queue <- mes:
		this.Size++
	default:
	
	}
}

func (this *MesQueue) Pull(size int, timeout time.Duration) []interface{} {
	mes := make([]interface{}, 0)
	for i := 0; i < size; i++ {
		select {
		case msg := <-this.Queue:
			mes = append(mes, msg.(Message))
		case <-time.After(timeout):
			return mes
		}
	}
	return mes
}

func (this *MesQueue) Length() int {
	return this.Size
}

func (this *MesQueue) Capacity() int {
	return this.Cap
}
