package main

import (
	"fmt"
	"reflect"
)

var (
	name, course string
	module       float64
)

const myConstant = 123

var inferredTypeVar1, inferredTypeVar2 = "testing", 3.3

func main() {
	name = "testing"
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Module is set to", inferredTypeVar1, "and is of type", reflect.TypeOf(inferredTypeVar1))
	fmt.Println("Module is set to", inferredTypeVar2, "and is of type", reflect.TypeOf(inferredTypeVar2))
	fmt.Println("constant is set to", myConstant, "and is of type", reflect.TypeOf(myConstant))

	a := 10.000000000
	b := 3
	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))

	c := int(a) + b
	d, e := 3, 4
	fmt.Println("\nC has value:", c, "and is of type", reflect.TypeOf(c))
	fmt.Println("\nD and E has value:", d, e, "and is of type", reflect.TypeOf(d), reflect.TypeOf(e))

	//& to references a pointer
	//* to de-references a pointer
	localVar := "Tony"
	pointerVarToLocalVar := &localVar
	fmt.Println("Memory address of *localVar* variable is", pointerVarToLocalVar, "and the value of *module* is", *pointerVarToLocalVar)

	myCourse := "my course"
	changeCourse(myCourse)
	fmt.Println("course name changed?", myCourse)

	changeCourse2(&myCourse)
	fmt.Println("course name changed2?", myCourse)
}

func changeCourse(course string) string {
	course = "changed course name"

	fmt.Println("\nTrying to change course name to: ", course)

	return course
}

func changeCourse2(course *string) string {
	*course = "OMG"
	fmt.Println("Passed by reference:", course, " -> ", *course)
	return *course
}
