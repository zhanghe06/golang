package main

import "fmt"

type Person struct {
	name string
}

type Dept struct {
	persons [4]Person
}

func main() {

	persons := [4]Person{}
	persons[0] = Person{"Tom"}
	d1 := Dept{persons}
	// DeptOne:        main.Dept{persons:[4]main.Person{main.Person{name:"Tom"}, main.Person{name:""}, main.Person{name:""}, main.Person{name:""}}}
	fmt.Printf("DeptOne:\t%#v\n", d1)

	ch := make(chan Dept, 1)
	ch <- d1
	d1Copy := <-ch
	d1.persons[0].name = "Jim"
	// DeptOneUpdate:  main.Dept{persons:[4]main.Person{main.Person{name:"Jim"}, main.Person{name:""}, main.Person{name:""}, main.Person{name:""}}}
	fmt.Printf("DeptOneUpdate:\t%#v\n", d1)
	// DeptOneCopy:    main.Dept{persons:[4]main.Person{main.Person{name:"Tom"}, main.Person{name:""}, main.Person{name:""}, main.Person{name:""}}}
	fmt.Printf("DeptOneCopy:\t%#v\n", d1Copy)
}
