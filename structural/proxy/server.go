package main

type Server interface {
	GetContent(userid string) string
}


