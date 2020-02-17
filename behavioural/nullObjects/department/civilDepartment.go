package department

type CivilDepartment struct {
	Name         string
	NumProfessor int
}

func (d *CivilDepartment) GetName() string {
	return d.Name
}

func (d *CivilDepartment) GetNumProfessor() int {
	return d.NumProfessor
}