package main

import "fmt"

func main() {
	webServer := &WebServer{}
	proxyServer := NewProxyServer(webServer)
	fmt.Println(proxyServer.GetContent("dicky"))
	fmt.Println(proxyServer.GetContent("angel"))
	fmt.Println(proxyServer.GetContent("asdf"))
}
