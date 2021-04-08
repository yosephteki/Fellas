package main

import (
	"Fellas/Routes"
	"fmt"
)

func main() {
	fmt.Println("HELLO WORLD!")

	r := Routes.SetupRouter()
	r.Run()
}
