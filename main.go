package main

import (
	"github.com/bouloou/GDCapp/src/appBackend"
	"github.com/bouloou/GDCapp/src/appFrontend"
)

func main() {
	appBackend.StartBackend()
	appFrontend.StartFrontend()
}
