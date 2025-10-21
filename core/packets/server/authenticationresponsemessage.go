package server

import (
	"LaserGo/core/network"
	"LaserGo/core/piranhamessage"
)

type AuthenticationResponseMessage struct {
	*piranhamessage.PiranhaMessage
}

func NewAuthenticationResponseMessage(buf []byte, client *network.Client) *AuthenticationResponseMessage {
	msg := &AuthenticationResponseMessage{
		PiranhaMessage: piranhamessage.NewPiranhaMessage(buf, client, "AuthenticationResponseMessage"),
	}
	msg.SetMessageType(20104)
	return msg
}

func (a *AuthenticationResponseMessage) Encode() *AuthenticationResponseMessage {
	a.WriteInt(int32(a.GetClient().GetHighId())) // wheelchaired code
	a.WriteInt(int32(a.GetClient().GetLowId()))  // wheelchaired code
	a.WriteInt(int32(a.GetClient().GetHighId())) // wheelchaired code
	a.WriteInt(int32(a.GetClient().GetLowId()))  // wheelchaired code
	a.WriteString("")
	a.WriteString("")
	a.WriteString("")
	a.WriteInt(63)
	a.WriteInt(0)
	a.WriteInt(0)
	a.WriteString("dev")
	a.WriteInt(0)
	a.WriteInt(0)
	a.WriteInt(0)
	a.WriteString("")
	a.WriteString("")
	a.WriteString("")
	a.WriteInt(0)
	a.WriteString("")
	a.WriteString("BE")
	a.WriteString("")
	a.WriteInt(1)
	a.WriteString("")
	a.WriteInt(0)
	a.WriteInt(0)
	a.WriteVInt(0)
	a.WriteString("")
	a.WriteVInt(1)
	a.WriteVInt(1)
	a.WriteString("")
	return a
}

func (a *AuthenticationResponseMessage) Process() *AuthenticationResponseMessage {
	return a
}
