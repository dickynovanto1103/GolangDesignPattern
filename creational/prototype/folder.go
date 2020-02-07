package main

import "fmt"

type folder struct {
	children []inode
	name string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	var tempChildren []inode
	for _, child := range f.children {
		tempChildren = append(tempChildren, child.clone())
	}

	return &folder{
		children: tempChildren,
		name:     f.name,
	}
}