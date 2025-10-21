package main

import (
	ServerConnection "LaserGo/core/serverconnection"
	Utils "LaserGo/utils"
)

type LaserServerEmulator struct{}

func (Server *LaserServerEmulator) Init() {
	Utils.DebuggerInst.Info("LaserServerEmulator.InitServer -> Initialising Server")
	ServerConnection.Init()
	ServerConnection.Listen(ServerIp, ServerPort) // check config to change ip
}

func main() {
	ServerEmulator := &LaserServerEmulator{}
	ServerEmulator.Init()
}
