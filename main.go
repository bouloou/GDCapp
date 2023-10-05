package main

import (
	"appBackend"
	"appFrontend"
)

func main() {
	appBackend.StartBackend()
	appFrontend.StartFrontend()
}
