package main

import "fmt"

type Person struct {
	//Name    string
	//Address string
	Name, Address string
	Credit        int
}

type SuperBeing struct {
	//humanPart Person
	Person     // this is now embedded "
	SuperPower string
}

// func (p Person) asText() string {
func (p Person) String() string {
	return fmt.Sprintf("I'm a person name is %s, credit is %d", p.Name, p.Credit)
}

//func (p Person) increaseCredit(amount int) {
//	fmt.Printf("parameter p is at: %p\n", &p)
//	p.Credit += amount
//	fmt.Println("after increase, p is ", p)
//}

//func (p Person) increaseCredit(amount int) Person {
//	fmt.Printf("parameter p is at: %p\n", &p)
//	p.Credit += amount
//	fmt.Println("after increase, p is ", p)
//	return p
//}

func (p *Person) increaseCredit(amount int) {
	// not &p, this time (that would be the address of the POINTER)
	fmt.Printf("parameter p is at: %p\n", p)
	// Go uses & to extract a pointer, but does not require (or use) -> to indirect a pointer
	p.Credit += amount
	fmt.Println("after increase, p is ", p)
}

type Business struct {
	Name      string
	MaxCredit int
}

func (b *Business) increaseCredit(amount int) {
	b.MaxCredit += amount
}

type Creditable interface {
	increaseCredit(amount int)
}

func main() {
	//var ayo Person
	//ayo.Name = "Ayo"
	//ayo.Address = "On the hill"
	//ayo := Person{Name: "Ayo", Address: "On the hill"}
	//ayo := Person{Address: "On the hill", Name: "Ayo"}
	//ayo := Person{Address: "On the hill"}
	ayo := Person{"Ayo", "On the hill", 1_000}
	fmt.Printf("%#v\n", ayo)
	ayo2 := ayo
	fmt.Println("-------------------------")
	fmt.Printf("ayo is %#v,\nayo2 is %#v,\nayo is at %p, ayo2 is at %p, ayo == ayo2? is %t\n",
		ayo, ayo2, &ayo, &ayo2, ayo == ayo2)
	ayo2.Credit = 10_000
	fmt.Printf("ayo is %#v,\nayo2 is %#v,\nayo is at %p, ayo2 is at %p, ayo == ayo2? is %t\n",
		ayo, ayo2, &ayo, &ayo2, ayo == ayo2)

	fmt.Println(ayo)
	//fmt.Println(asText(ayo))
	//fmt.Println(ayo.asText())
	fmt.Println(ayo)
	fmt.Printf("ayo is at: %p\n", &ayo)
	//ayo = ayo.increaseCredit(10_000)
	ayo.increaseCredit(10_000) // pointer is taken automatically--can be confusing
	fmt.Println("back in main ", ayo)

	//clarkKent := SuperBeing{Person: Person{Name: "Clark Kent", Address: "Krypton", Credit: 1_000_000},
	//	SuperPower: "x-ray vision"}
	clarkKent := SuperBeing{Person{"Clark Kent", "Krypton", 1_000_000}, "x-ray vision"}

	// THIS IS NOT PERMITTED!!!
	//var clarkKent Person = SuperBeing{Person{"Clark Kent","Krypton",1_000_000},"x-ray vision"}
	fmt.Println(clarkKent)
	fmt.Println(clarkKent.Person.Name)
	fmt.Println(clarkKent.Name)

	// Problem #1 (which I think someone was trying to tell me in the chat and
	// I failed to interpret correctly.
	// The interface does not get declared with any indication that a pointer
	// is used, but because the methods that satisfy the interface take pointers,
	// so a pointer is needed to satisfy that interface
	// therefore, the only change made here (aside from putting the pointer asterisks
	// back into the method declarations, is the addition of these three ampersands
	//                                    v     v           v
	var creds []Creditable = []Creditable{&ayo, &clarkKent, &Business{"SuperACME", 10000}}
	fmt.Println(creds)
	for _, v := range creds {
		v.increaseCredit(10)
	}

}
