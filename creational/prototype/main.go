package main

func main() {
	file1 := &file{name: "satu"}
	file2 := &file{name: "dua"}
	folder1 := &folder{name: "fol1", children: []inode{file1, file2}}
	file1.print("   ")
	file2.print("   ")
	folder1.print("   ")
}
