package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print also")
	case (3 == 3):
		fmt.Println("Print 3==3")
		fallthrough // to include next case, no matter what is the condition
	case (4 == 4):
		fmt.Println("Print 4==4")
		fallthrough
	case (5 == 4):
		fmt.Println("noPrint 5==4")
		fallthrough
	case (6 == 6):
		fmt.Println("Print 6==6")
	}

	fmt.Println("------------")

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print also")
	case (5 == 4):
		fmt.Println("noPrint 5==4")
	default:
		fmt.Println("Print Default")
	}

	fmt.Println("------------")

	switch "car" {
	case "bike":
		fmt.Println("this is bike")
	case "car":
		fmt.Println("this is car")
	default:
		fmt.Println("Nothing")
	}

	fmt.Println("------------")

	year := 2000
	switch year {
	case 2001:
		fmt.Println("this is 2001")
	case 2000:
		fmt.Println("this is 2000")
	default:
		fmt.Println("Nothing")
	}

	fmt.Println("------------")

	switch "horse" {
	case "bike", "car", "plane":
		fmt.Println("this is vehicle")
	case "horse", "bear":
		fmt.Println("this is animal")
	default:
		fmt.Println("Nothing")
	}

}
