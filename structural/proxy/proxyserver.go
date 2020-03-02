package main

type ProxyServer struct {
	server Server
}

func NewProxyServer(s Server) *ProxyServer {
	return &ProxyServer{server:s}
}

func (p *ProxyServer) GetContent(userid string) string {
	if !p.IsAllowed(userid) {
		return "forbidden"
	}
	return p.server.GetContent(userid)
}

func (p *ProxyServer) IsAllowed(userid string) bool {
	if userid == "dicky" || userid == "angel" {
		return true
	}
	return false
}