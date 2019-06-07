package main

import (
	"grpcAssignment/server"
)

func main() {
	//go func() {
		server.Start("50051")
	//}()
	//time.Sleep(2 * time.Second)
	//client.Start("50050")
}
