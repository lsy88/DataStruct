package MessageQueue

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

type Message struct {
	ID       int64
	TopicLen int64
	Topic    string //栏目
	MesType  int64  //消息类型1-consumer 2-producer 3-consumer-ack 4-error
	Len      int64  //消息长度
	Data     []byte //信息内容
}

//消息解码
func (mq *Message) Decode(reader io.Reader) *Message {
	var buf [128]byte
	_, err := reader.Read(buf[:])
	if err != nil {
		log.Fatalln(err)
	}
	//ID
	buff := bytes.NewBuffer(buf[:8])
	binary.Read(buff, binary.LittleEndian, &mq.ID)
	// topiclen
	buff = bytes.NewBuffer(buf[8:16])
	binary.Read(buff, binary.LittleEndian, &mq.TopicLen)
	// topic
	msgLastIndex := 16 + mq.TopicLen
	mq.Topic = string(buf[16:msgLastIndex])
	// mestype
	buff = bytes.NewBuffer(buf[msgLastIndex : msgLastIndex+8])
	binary.Read(buff, binary.LittleEndian, &mq.MesType)
	
	buff = bytes.NewBuffer(buf[msgLastIndex : msgLastIndex+16])
	binary.Read(buff, binary.LittleEndian, &mq.Len)
	
	if mq.Len <= 0 {
		return mq
	}
	
	mq.Data = buf[msgLastIndex+16:]
	return mq
	
}

//消息编码
func (mq *Message) Encode() []byte {
	mq.TopicLen = int64(len([]byte(mq.Topic)))
	mq.Len = int64(len([]byte(mq.Data)))
	
	var data []byte
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, mq.ID)
	data = append(data, buf.Bytes()...)
	
	buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, mq.TopicLen)
	data = append(data, buf.Bytes()...)
	
	data = append(data, []byte(mq.Topic)...)
	
	buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, mq.MesType)
	data = append(data, buf.Bytes()...)
	
	buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, mq.Len)
	data = append(data, buf.Bytes()...)
	data = append(data, mq.Data...)
	
	return data
}
