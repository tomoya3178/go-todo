package main

import "go-todo/presentation"

func main() {
	presentation.GetRouter().Run(":3000")
}
