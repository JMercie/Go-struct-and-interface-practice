package main

import (
	"fmt"

	"github.com/JMercie1/pluralsight/organization"
)

func main() {
	p := organization.NewPerson(organization.NewEuropeanUnionID("222-22-2222", "Alemania"), "joseph", "mercie")
	err := p.SetTwitterHanlder("@JosephMLtc")
	if err != nil {
		fmt.Printf("twitter hanlder couldnt be set: %s\n", err)
	}

	// name1 := Name{First: "joseph", Last: "Mercie"}
	// name2 := Name{First: "joseph", Last: "Mercie"}

	// if name1 == name2 {
	// 	fmt.Println("we match")
	// }

	eui := organization.NewEuropeanUnionID("222-22-2222", "Alemania")
	ssn := organization.NewSocialSecurityNumber("222-22-333")

	if eui == ssn {
		fmt.Println("match")
	}
}

// slices and functions have not defined alocated memory so you can't not compare then at runtime with == or != ///
type Name struct {
	First string
	Last  string
}

// structs cant not compare to other structs even if their properties are similar or the same
