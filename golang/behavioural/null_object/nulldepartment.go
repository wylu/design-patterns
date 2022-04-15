package main

type nullDepartment struct {
	numberOfProfessors int
}

func (n *nullDepartment) getNumberOfProfessors() int {
	return 0
}

func (n *nullDepartment) getName() string {
	return "nullDepartment"
}
