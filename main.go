package main

import "fmt"

func main() {
	var celcius int
	var pilihan int
	fmt.Print("Enter temperature in celcius: ")
	fmt.Scanln(&celcius)
	for{
		fmt.Println("\n1. Fahrenheit")
		fmt.Println("2. Kelvin")
		fmt.Println("3. Reamur")
		fmt.Printf("Select temperature do you want to convert to: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fahrenheit := celcius*9/5 + 32
			fmt.Println("Temperature in fahrenheit is: ", fahrenheit, " F")
		case 2:
			kelvin := celcius + 273
			fmt.Println("Temperature in kelvin is: ", kelvin, " K")
		case 3:
			reamur := celcius * 4 / 5
			fmt.Println("Temperature in reamur is: ", reamur, " R")
		default:
			fmt.Println("Your input is invalid, please input a valid option (1, 2, or 3).")
	}
	}
	
}
