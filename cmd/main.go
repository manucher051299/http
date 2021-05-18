package main

import (
	"log"
	"net"

	"github.com/manucher051299/http/pkg/server"
)

func main() {
	execute("0.0.0.0", "9999")
}

func execute(host string, port string) (err error) {
	srv := server.NewServer(net.JoinHostPort(host, port))
	srv.Register("/api/payments/1", func(req *server.Request) {
		id := req.PathParams["id"]
		log.Print(id)
		return
	})
	return srv.Start()
}
