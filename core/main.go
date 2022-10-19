package main

import (
	"fmt"
	"net"
	"net/http"
)

func reportIP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Update2. reportIP, my ip: %v.\n", getLocalIP())
}

func main() {
	http.HandleFunc("/reportip", reportIP)
	http.ListenAndServe(":9000", nil)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
