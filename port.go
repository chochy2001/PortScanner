package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 100; i++ {

		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))

		if err != nil {
			continue
		}
		errCon := conn.Close()
		if errCon != nil {
			return
		}
		fmt.Printf("El puerto %d esta abierto", i)
	}
	fmt.Println("vim-go")
}
