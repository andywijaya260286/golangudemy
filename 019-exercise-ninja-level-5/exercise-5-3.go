package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	f150 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourwheel: true,
	}

	civic := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(f150)
	fmt.Println("f150")
	fmt.Println(f150.color)
	fmt.Println(f150.fourwheel)
	fmt.Println()
	fmt.Println(civic)
	fmt.Println("civic")
	fmt.Println(civic.doors)
	fmt.Println(civic.color)
	fmt.Println(civic.luxury)
}
