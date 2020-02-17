package department

type CSDepartment struct {
	Name         string
	NumProfessor int
}

func (d *CSDepartment) GetName() string {
	return d.Name
}

func (d *CSDepartment) GetNumProfessor() int {
	return d.NumProfessor
}