package main

import (
	"fmt"
	"reflect"
)

// type Doctor struct {
// 	name    string
// 	age     int
// 	hobbies []string
// 	parent  []string
// }

type Animal struct {
	Name   string `json: name`
	Origin string `json: origin`
}

type Bird struct {
	Animal
	Speed  float64
	CanFly bool
}

func main() {
	// general struct
	// aDoctor := Doctor{
	// 	name: "max",
	// 	age:  29,
	// 	hobbies: []string{
	// 		"game",
	// 		"movie",
	// 	},
	// }
	// fmt.Println(aDoctor)
	// fmt.Println(aDoctor.name)
	// fmt.Println(aDoctor.age)

	// embedded struct
	// b := Bird{}
	// b.Name = "max"
	// b.Origin = "taipei"
	// b.CanFly = true
	// b.Speed = 60
	// b := Bird{
	// 	Animal: Animal{Name: "max", Origin: "taipei"},
	// 	Speed:  60,
	// 	CanFly: true,
	// }
	// fmt.Println(b)

	// tag
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Origin")
	fmt.Println(field.Name, field.Tag) // Origin json: origin

}
