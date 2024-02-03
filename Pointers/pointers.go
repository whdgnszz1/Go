package main

import "fmt"

func main() {
	age := 32 // regular variable
	
	var agePointer *int // pointer
	agePointer = &age // pointer

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult years:", age)
}

func editAgeToAdultYears(age *int)  {
	// return *age - 18
	*age = *age - 18
}