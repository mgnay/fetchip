package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	host:= os.Args[1]
	host=strings.ReplaceAll(host, "http://", "")
	host=strings.ReplaceAll(host, "https://", "")
	if strings.Contains(host,"/"){
		temp :=strings.Index(host,"/")
		host =host[:temp]
	}else if strings.Contains(host, "\\") {
		temp:=strings.Index(host,"\\")
		host =host[:temp]
	}
	ip,err:= hostIP(host)
	if err != nil {
		fmt.Println("Girdiğiniz adres hatalıdır...")
	}else {
		fmt.Println(ip)
	}

}

func hostIP(host string) (net.IP, error) {
	ips, err := net.LookupIP(host)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4, nil
		}
	}
	return nil, err
}