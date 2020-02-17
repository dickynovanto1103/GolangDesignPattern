package department

type NullDepartment struct {

}

func (d *NullDepartment) GetName() string {
	return "nullDepartment"
}

func (d *NullDepartment) GetNumProfessor() int {
	return 0
}