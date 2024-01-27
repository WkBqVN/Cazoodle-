package main

import (
	"fmt"

	"cazoodle.com/controller"
)

func main() {
	c := controller.GetInstance()
	fmt.Println(c)
	c.Init()
	err := c.StartOnPort(":6000")
	if err != nil {
		fmt.Println(err)
	}
}
