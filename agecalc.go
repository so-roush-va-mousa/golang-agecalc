package main

import (
	"fmt"
	"time"
)

func main()  {
	var birthDay int	
	fmt.Print("Enter your birth year:")
	fmt.Scanln(&birthDay)

	currentYear := time.Now().Year()
	yourAge := currentYear - birthDay

	fmt.Printf("your age is %d years old \n", yourAge)
}