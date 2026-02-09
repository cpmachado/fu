package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	port := 1337
	msg := "This is up"

	flag.IntVar(&port, "port", port, "port to serve")
	flag.StringVar(&msg, "msg", msg, "message to show")
	flag.Parse()

	if port < 0 || port > 60999 {
		log.Fatalf("invalid port: %d", port)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(msg))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	var ips []string

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ips = append(ips, ip.String())
		}
	}

	fmt.Println("Listening in:")
	for _, ip := range ips {
		fmt.Printf("- http://%s:%d\n", ip, port)
	}

	server.ListenAndServe()
}
