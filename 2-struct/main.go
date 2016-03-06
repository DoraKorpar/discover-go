package main

import "fmt"
import "time"

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func (x *user) sayHello() {
	fmt.Println("Hello", x.Name)
}

func (x *user) getAge() (int, error) {
	nbHourYear := 24 * 365
	layout := "January 2, 2006"
	t, err := time.Parse(layout, x.DOB)
	if err != nil {
		return 0, err
	}
	return (int(time.Since(t).Hours()) / nbHourYear), nil
}

func main() {
	u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
	u.sayHello()
	age, err := u.getAge()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Name, "who was born in", u.City, "would be", age, "years old today.")
}
