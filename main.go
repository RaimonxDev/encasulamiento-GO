package main

import (
	"fmt"
	"github.com/RaimonxDev/encasulamiento-GO/course"
)

func main() {

	Go := course.New("GO desde 0", 0, false)

	Go.SetClasses(map[uint]string{
		1: "Que es POO",
		2: "Primero pasos en Go",
	})
	Go.SetPrice(40)

	fmt.Printf("%v", Go)

	Go.SetName("POO con Go")
	fmt.Println(Go.Name() + "\n")
	fmt.Println(Go.Price())
}
