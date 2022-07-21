package main

import (
	"authMicroservice/app"
	"authMicroservice/app/config"
	"flag"
	"log"
)

func main() {
	config.Conn()
	hostPointer := flag.String("host", config.GetHost(), "set for host")
	portPointer := flag.String("port", config.GetPort(), "set for port")
	flag.Parse()
	application := app.App("Authentication Server")
	log.Fatal(application.Listen(*hostPointer + ":" + *portPointer))
}
