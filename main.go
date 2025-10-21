package main

import (
	"LaserGo/core"
	"LaserGo/utils"
)

type LaserServerEmulator struct{}

func (Server *LaserServerEmulator) Init() {
	utils.DebuggerInst.Info("LaserServerEmulator.InitServer -> Initialising Server")
	core.Init()
	core.Listen(ServerIp, ServerPort) // check config to change ip
}

func main() {
	ServerEmulator := &LaserServerEmulator{}
	ServerEmulator.Init()
}
