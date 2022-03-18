package main

import (
	"fmt"
)

type Student struct {
	Name string
	ID uint64
	GPA float32
}

func main () { 
	a := -36
	b := "Dusan"
	c := true
 	d := 3.14159
	e := Student {
		Name: b,
		ID: 5646879,
		GPA: 4.6,
	}
	f := 22

	fmt.Printf("hello %s\n", b)
	fmt.Printf("a is  %d\n", a)
	fmt.Printf("f is %+d\n", f)
	fmt.Printf("c is %t\n", c)
	fmt.Printf("d is %f, with 3 decimals is %.3f\n", d, d)
	fmt.Printf("student: %v\n", e)
	fmt.Printf("student: %+v\n", e)
	fmt.Printf("Type of student is %T", e)

	str := fmt.Sprintf("%s is a student\n with a student ID of %d \nand a GPA of %.3f\n",
		e.Name, e.ID , e.GPA)

		fmt.Println(str)
 }