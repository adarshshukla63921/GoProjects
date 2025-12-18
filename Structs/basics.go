package main

import "fmt"

// syntax : 
// type structName struct

type Student struct {
	RollNumber int
	Name string
	Age int
	Grade string
}

func updateGrade(s *Student, newGrade string){
	(*s).Grade = newGrade
}

// embedded struct
type AdmissionDetails struct{
	AdmissionNumber int
	StudentInformation Student
}
func main() {


	// creating student instance
	var student0 Student = Student{10,"Yashiro",20,"A+"}
	fmt.Println(student0)


	// creating a student instance
	// using struct litreals.
	student1 := Student{
		RollNumber: 1,
		Name: "Adarsh",
		Age: 22,
		Grade: "A"}
	fmt.Println(student1)

	// empty structs store zero values for fields
	student2 := Student{}
	fmt.Println(student2) // {0 "" 0 ""}

	// accessing structs using pointers.
	ptr := &student2
	// both notations work
	(*ptr).Age = 21
	(*ptr).RollNumber=2
	ptr.Name="Akshat"
	ptr.Grade="S"
	fmt.Println("Updated Student 2 Age : ",student2)

	// accessing feilds using dot operator
	fmt.Println("Student 1 Name : ",student1.Name)

	// Structs with missing values.
	student4 := Student{
		RollNumber: 5,
		Name: "Saksham"}

	fmt.Println("Student4 : ",student4) // missing fields get zero values


	// using embedded structs
	admissionInfo1 := AdmissionDetails{
		AdmissionNumber: 1001,
		StudentInformation: Student{
			RollNumber: 11,
			Name: "Rohan",
			Age: 19,
			Grade: "B+"}}
	fmt.Println("Admission Info 1 : ",admissionInfo1)

	// more on using embedded structs
	admissionInfo2 := AdmissionDetails{
		AdmissionNumber: 1002,
		StudentInformation: student0}
	fmt.Println("Admission Info 2 : ",admissionInfo2)

	// using structs with pointers and functions.
	updateGrade(&student1, "A++")
	fmt.Println("Updated Grade : ",student1)
}