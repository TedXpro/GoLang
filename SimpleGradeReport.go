package main

import (
	"fmt"
)

type Student struct{
	Name string
	Grades map[string] int
}

func claculateAverage(student Student) float64{
	var total int
	for _, grade := range student.Grades{
		total += grade
	}

	return float64(total) / float64(len(student.Grades))
}

func main(){
	var name string
	var numSubjects int

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter the number of subjects: ")
	fmt.Scanln(&numSubjects)

	student := Student{Name: name, Grades: make(map[string]int)}
	for i:= 0; i < numSubjects; i++ {
		var subjectName string
		var grade int
		fmt.Println("Enter subject ", i + 1, " name")
		fmt.Scanln(&subjectName)

		fmt.Println("Enter your grade for ", subjectName, " :")
		fmt.Scanln(&grade)

		if(grade < 0 || grade > 100) {
			fmt.Println("Invalid grade. Please enter again")
			i--
			continue
		}

		student.Grades[subjectName] = grade
	}

	averageGrade := claculateAverage(student)

	fmt.Println("/n Student Name:", student.Name)
	fmt.Println("Subject | Grades:")
	for subject, grade := range student.Grades {
		fmt.Println(subject, ": ", grade)
	}

	fmt.Println("Average Grade result is: ", averageGrade)

}