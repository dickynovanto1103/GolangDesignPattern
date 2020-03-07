package main

import "fmt"

func main() {
	file1 := &file{name: "satu"}
	file2 := &file{name: "dua"}
	folder1 := &folder{name: "fol1", children: []inode{file1, file2}}
	file1.print("   ")
	file2.print("   ")
	folder1.print("   ")
	folder2 := &folder{name: "fol2", children: []inode{folder1, file1, file2}}
	folder2.print("   ")
	fmt.Println("-----------printing folder 3 = clone of folder 2----------")
	folder3 := folder2.clone()
	folder3.print("   ")
}
