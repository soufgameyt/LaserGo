package serverconnection

import (
	MessageManager "LaserGo/core/messagemanager"
	"LaserGo/core/network"
	"LaserGo/utils"
	"fmt"
	"net"
)

type serverconnection struct{}

var serverconnectinstance *serverconnection
var serverinst net.Listener // litterally took from https://go.dev/src/net/example_test.go lol

func Init() *serverconnection {
	if serverconnectinstance == nil {
		serverconnectinstance = &serverconnection{}
		utils.DebuggerInst.Info("ServerConnection.Init ->  Created ServerConnection Instance!")
	}

	return serverconnectinstance
}

func Listen(server_ip string, server_port int) {
	ln, err := net.Listen("tcp", server_ip+":"+itoa(server_port))
	if err != nil {
		utils.DebuggerInst.Error("Failed to start server:", err)
		return
	}
	serverinst = ln
	utils.DebuggerInst.Info("ServerConnection.Listen -> Server started on", server_ip, "with port", server_port)

	for {
		conn, err := serverinst.Accept()
		if err != nil {
			utils.DebuggerInst.Error("wheelchaired error", err)
			continue
		}

		utils.DebuggerInst.Info("ServerConnection.Listen -> New client connected!")
		HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	client := network.NewClient(conn)
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			utils.DebuggerInst.Warn("ServerConnection.Listen -> Client Disconnected!")
			return
		}

		MessageManager.ReceiveMessage(buf[:n], client)
	}
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
