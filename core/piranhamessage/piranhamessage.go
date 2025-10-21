package piranhamessage

import (
	"LaserGo/core/network"
	"LaserGo/datastream"
	"LaserGo/utils"
	"encoding/binary"
)

type PiranhaMessage struct {
	*datastream.ByteStream
	messageId       int32
	messageVersion  int32
	client          *network.Client
	messageTypeName string
}

func NewPiranhaMessage(buf []byte, client *network.Client, msgTypeName string) *PiranhaMessage {
	pm := &PiranhaMessage{
		ByteStream:      datastream.NewByteStream(),
		messageId:       20000,
		messageVersion:  0,
		client:          client,
		messageTypeName: msgTypeName,
	}
	pm.ReplaceBuffer(buf)
	utils.DebuggerInst.Info(msgTypeName)
	utils.DebuggerInst.Info("PiranhaMessage.Send -> " + msgTypeName)
	return pm
}

func (p *PiranhaMessage) SetMessageType(id int32) {
	p.messageId = id
}

func (p *PiranhaMessage) GetMessageType() int32 {
	return p.messageId
}

func (p *PiranhaMessage) GetClient() *network.Client {
	return p.client
}

func (p *PiranhaMessage) GetMessageTypeName() string {
	return p.messageTypeName
}

func (p *PiranhaMessage) Send() {
	// Messaging::writeHeader
	header := make([]byte, 7)
	binary.BigEndian.PutUint16(header[0:2], uint16(p.messageId))

	bufLen := p.GetLength()
	header[2] = byte((bufLen >> 16) & 0xFF)
	header[3] = byte((bufLen >> 8) & 0xFF)
	header[4] = byte(bufLen & 0xFF)
	binary.BigEndian.PutUint16(header[5:7], uint16(p.messageVersion))

	final := append(header, p.GetBuffer()...)
	final = append(final, []byte{0xFF, 0xFF, 0, 0, 0, 0, 0}...)
	p.client.GetSocket().Write(final)
}

func (p *PiranhaMessage) Decode() {
	utils.DebuggerInst.Info("PiranhaMessage.Decode -> Decode is not implemented in this PiranhaMessage!")
}
func (p *PiranhaMessage) Encode() {
	utils.DebuggerInst.Info("PiranhaMessage.Encode -> Encode is not implemented in this PiranhaMessage!")
}
func (p *PiranhaMessage) Process() {}
