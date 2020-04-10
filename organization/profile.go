package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "america"
}

type europeanUnionID struct {
	id      string
	country string
}

func NewEuropeanUnionID(id, country string) Citizen {
	return europeanUnionID{
		id:      id,
		country: country,
	}
}

func (euid europeanUnionID) ID() string {
	return euid.id
}

func (euid europeanUnionID) Country() string {
	return euid.country
}

type TwitterHandler string

func (th TwitterHandler) RedirectURL() string {
	cleanHanlder := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHanlder)
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}
type Person struct {
	Name
	twitterHanlder TwitterHandler
	Citizen
}

func NewPerson(citizen Citizen, firstName, lastName string) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

func (p *Person) ID() string {
	return fmt.Sprintf("Social Security number is: %s", p.Citizen.ID())
}

func (p *Person) SetTwitterHanlder(hanlder TwitterHandler) error {
	if len(hanlder) == 0 {
		p.twitterHanlder = hanlder
	} else if !strings.HasPrefix(string(hanlder), "@") {
		return errors.New("Must start with @ symbol")
	}

	p.twitterHanlder = hanlder
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHanlder
}
