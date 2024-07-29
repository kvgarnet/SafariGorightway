package main

import "fmt"

func main() {
	//var names [3]string
	//var names [3]string = [3]string{"Ayo", "Hua"}
	// ooops, this is an array of 4 strings NOT compatible with array of 3 strings
	//var names [3]string = [...]string{"Ayo", "Hua", "Ishan", "Siobhan"}
	var names [3]string = [...]string{"Ayo", "Hua", "Ishan"}
	//names[0] = "Ayo"
	//names[1] = "Hua"
	fmt.Printf("names is %#v, and is of type %T\n", names, names)
	var names2 [3]string
	names2 = names
	fmt.Printf("names2 is %#v, and is of type %T\n", names2, names2)

	fmt.Printf("names == names2? %t\n", (names == names2))
	names[0] = "Siobhan"
	fmt.Printf("names is %#v, and is of type %T\n", names, names)
	fmt.Printf("names2 is %#v, and is of type %T\n", names2, names2)
	fmt.Printf("names == names2? %t\n", (names == names2))

}
