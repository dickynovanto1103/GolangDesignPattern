package main

type WebServer struct{}

func (w *WebServer) GetContent(userid string) string {
	return "content for " + userid
}
