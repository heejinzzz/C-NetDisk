package main

import (
	"C-NetDisk/CLI/operation"
	"fmt"
	"os"
)

func main() {
	addr := GetServerAddr()
	err := operation.ConnectCNetDiskServer(addr)
	if err != nil {
		fmt.Println("Connect To C-NetDisk Server Failed. Error: ", err)
		os.Exit(1)
	}
	Welcome()
	username := Login()
	fmt.Println("Hello, ", username, ". Let's Start Exploring!")
	fmt.Println()
	ShowOptionalCommands()
	handler := NewCommandHandler(username)
	handler.WaitForCommand()
}
