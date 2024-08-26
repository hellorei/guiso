package main

import (
	"guiso/app/framework"
	"guiso/app/routes"
)

func main() {
	e := framework.SetupEcho()
	routes.RegisterRoutes(e)
	framework.StartServer(e)
}
