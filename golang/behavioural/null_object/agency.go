package main

import "fmt"

func main() {
	college1 := createCollege1()
	college2 := createCollege2()
	departmentArray := []string{"computerscience", "mechanical", "civil", "electronics"}

	totalProfessors := 0
	for _, departmentName := range departmentArray {
		d := college1.getDepartment(departmentName)
		totalProfessors += d.getNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)

	// Reset the professor count
	totalProfessors = 0
	for _, departmentName := range departmentArray {
		d := college2.getDepartment(departmentName)
		totalProfessors += d.getNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}

func createCollege1() *college {
	college := &college{}
	college.addDepartment("computerscience", 4)
	college.addDepartment("mechanical", 5)
	return college
}

func createCollege2() *college {
	college := &college{}
	college.addDepartment("computerscience", 2)
	return college
}
