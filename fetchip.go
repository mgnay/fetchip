package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		getPublicIP()
	} else {
		getHostIP(os.Args[1])
	}
}

func getPublicIP() {

	var publicIP string
	res, err := http.Get("http://ip.me/")
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if res.StatusCode != http.StatusOK {
		fmt.Println(err)
	}

	publicIP, _ = doc.Find("#ip-lookup").Attr("value")
	clipboard.WriteAll(publicIP)
	fmt.Println("Your public IP: ", publicIP)
}

func getHostIP(host string) {
	host = strings.ReplaceAll(host, "http://", "")
	host = strings.ReplaceAll(host, "https://", "")
	if strings.Contains(host, "/") {
		temp := strings.Index(host, "/")
		host = host[:temp]
	} else if strings.Contains(host, "\\") {
		temp := strings.Index(host, "\\")
		host = host[:temp]
	}
	ip, err := hostIP(host)
	if err != nil {
		fmt.Println(err)
	} else {
		clipboard.WriteAll(fmt.Sprint(ip))
		fmt.Println(host, " IPv4: ", ip)
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
