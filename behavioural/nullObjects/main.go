package main

import (
	"GolangDesignPattern/behavioural/nullObjects/college"
	"fmt"
)

func main() {
	col1 := createCollege1()
	col2 := createCollege2()

	fmt.Println("tot prof col1: ", getNumProf(col1))
	fmt.Println("tot prof col2: ", getNumProf(col2))
}

func createCollege1() *college.College {
	col := &college.College{}
	col.AddDepartment("CS", 1)
	col.AddDepartment("civil", 3)
	return col
}

func createCollege2() *college.College {
	col := &college.College{}
	col.AddDepartment("CS", 2)
	col.AddDepartment("civil", 4)
	return col
}

func getNumProf(college *college.College) int {
	totProf := 0
	departmentArray := []string{"CS", "civil", "mec", "chemistry"}
	for _, depName := range departmentArray {
		dep := college.GetDepartment(depName)
		totProf += dep.GetNumProfessor()
	}

	return totProf
}