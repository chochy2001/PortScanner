package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))

			if err != nil {
				return
			}
			errCon := conn.Close()
			if errCon != nil {
				return
			}
			fmt.Printf("El puerto %d esta abierto\n", i)

		}(i)
	}
	wg.Wait()
	fmt.Println("Puertos escaneados con exito")
}
