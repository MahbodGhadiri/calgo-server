package main

import (
	router "calgo/src/routes"
	"fmt"
	"net"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		l.Close()
		panic(err)
	}

	fmt.Println("listening on port:", l.Addr().(*net.TCPAddr).Port)

	if err := http.Serve(l, r); err != nil {
		panic(err)
	}

}
