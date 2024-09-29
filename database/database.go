package database

import "fmt"

func CollectDb() string {
	fmt.Println("I'm in CollectDb package")
	return "Hi, I'm from database.go package"
}
