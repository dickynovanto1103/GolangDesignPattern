package main

import "fmt"

func main() {
	obj1 := NewObject(2)
	obj2 := NewObject(4)
	pack1 := NewPackage()
	pack2 := NewPackage()
	pack2.AddComponent(obj1)
	pack1.AddComponent(obj2)
	pack1.AddComponent(pack2)

	fmt.Println(pack1.Count())
}
