package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello, we are Holberton School")
	t := time.Now()
	fmt.Println("the date is", time.Now())
	fmt.Println("the year is", t.Year())
	fmt.Println("the month is", t.Month())
	fmt.Println("the day is", t.Day())
}
