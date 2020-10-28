package main

import (
	"github.com/Kmotiko/gofc"
	"github.com/Kmotiko/gofc/sample/app"
)

func main() {
	// regist app
	ssw := app.NewSimpleSwitch()
	smon := app.NewSimpleMonitor()
	gofc.GetAppManager().RegistApplication(ssw)
	gofc.GetAppManager().RegistApplication(smon)

	// start server
	gofc.ServerLoop(gofc.DEFAULT_PORT)
}
