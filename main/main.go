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
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectURL())
	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println(p.FullName())
}
