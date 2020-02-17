package college

import "GolangDesignPattern/behavioural/nullObjects/department"

type College struct {
	departments []department.Department
}

func (c *College) GetDepartment(name string) department.Department {
	for _, dep := range c.departments {
		if name == dep.GetName() {
			return dep
		}
	}

	return &department.NullDepartment{}
}

func (c *College) AddDepartment(name string, numProfessor int) {
	var dep department.Department

	switch name {
	case "CS":
		dep = &department.CSDepartment{
			Name: name,
			NumProfessor: numProfessor,
		}
	case "civil":
		dep = &department.CivilDepartment{
			Name: name,
			NumProfessor: numProfessor,
		}
	}

	c.departments = append(c.departments, dep)
}