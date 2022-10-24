package main

import (
	"C-NetDisk/CLI/operation"
	pb "C-NetDisk/CLI/proto"
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/term"
	"os"
	pt "path"
	"strconv"
	"syscall"
)

func GetServerAddr() string {
	host := flag.String("host", "localhost", "C-NetDisk Server host")
	port := flag.String("port", "6732", "C-NetDisk Server port")
	flag.Parse()
	addr := *host + ":" + *port
	return addr
}

func Welcome() {
	fmt.Println("{ Welcome to use C-NetDisk }")
	fmt.Println()
	fmt.Println("written by: heejinzzz")
	fmt.Println("follow me on github: https://github.com/heejinzzz")
	fmt.Println("contact me by email: 1273860443@qq.com")
	fmt.Println()
}

func Login() string {
	fmt.Println("[LOGIN]  (if you haven't got an account, input \"register\" to create your account now!)")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Username: ")
	usernameBytes, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	username := string(usernameBytes)
	if username == "register" {
		fmt.Println()
		return Register()
	}
	fmt.Print("Input Password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	password := string(passwordBytes)
	fmt.Println()
	fmt.Println()
	err = operation.UserLogin(username, password)
	if err != nil {
		fmt.Println("Login Failed. Error: ", err)
		os.Exit(1)
	}
	return username
}

func Register() string {
	fmt.Println("[REGISTER]")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Username: ")
	usernameBytes, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	username := string(usernameBytes)
	if username == "register" {
		fmt.Println()
		return Register()
	}
	fmt.Print("Input Password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	password := string(passwordBytes)
	fmt.Println()
	fmt.Print("Confirm Password: ")
	confirmBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	confirm := string(confirmBytes)
	fmt.Println()
	fmt.Println()
	if password != confirm {
		fmt.Println("Input Password and Confirm Password do not match!")
		fmt.Println()
		return Register()
	}
	if !isPasswordValid(password) {
		fmt.Println("Password is invalid! Please make sure the password length is between " + strconv.Itoa(PasswordMinLength) + " and " + strconv.Itoa(PasswordMaxLength) + ".")
		fmt.Println()
		return Register()
	}
	err = operation.UserRegister(username, password)
	if err != nil {
		fmt.Println("Register Failed. Error: ", err)
		os.Exit(1)
	}
	return username
}

func isPasswordValid(password string) bool {
	if len(password) < PasswordMinLength || len(password) > PasswordMaxLength {
		return false
	}
	return true
}

func ShowOptionalCommands() {
	fmt.Println("Optional Command:")
	for i := 0; i < len(OptionalCommandList); i += 4 {
		for j := 0; j < 4 && i+j < len(OptionalCommandList); j++ {
			fmt.Printf("* %-15s", OptionalCommandList[i+j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func ShowChildren(parentName string) error {
	parentInfo, err := operation.GetItemInfo(parentName)
	if err != nil {
		return err
	}
	totalSize := float32(parentInfo.Size)
	totalUnitIndex := 0
	for totalSize >= 1024 && totalUnitIndex < len(operation.SizeUnits)-1 {
		totalSize = totalSize / 1024
		totalUnitIndex++
	}
	fmt.Printf("%s [item: %d  size: %.2f%s]\n", parentName, len(parentInfo.Children), totalSize, operation.SizeUnits[totalUnitIndex])
	fmt.Println()
	var files, directories []*pb.ItemBasicInfo
	for _, child := range parentInfo.Children {
		if child.Type == pb.ItemType_File {
			files = append(files, child)
		} else {
			directories = append(directories, child)
		}
	}
	id := 1
	if len(directories) > 0 {
		fmt.Println("{ Folders }")
		fmt.Printf("%-10s%-40s%-10s\n", "Id", "Name", "Size")
		for _, v := range directories {
			name, size := pt.Base(v.Name), float32(v.Size)
			if len(name) > operation.ItemNameMaxInvisibleLength {
				name = name[:operation.ItemNameMaxInvisibleLength] + "..."
			}
			unitIndex := 0
			for size >= 1024 && unitIndex < len(operation.SizeUnits)-1 {
				size = size / 1024
				unitIndex++
			}
			sizeString := fmt.Sprintf("%.2f%s", size, operation.SizeUnits[unitIndex])
			fmt.Printf("%-10d%-40s%-10s\n", id, name, sizeString)
			id++
		}
		if len(files) > 0 {
			fmt.Println()
		}
	}
	if len(files) > 0 {
		fmt.Println("{ Files }")
		fmt.Printf("%-10s%-40s%-10s\n", "Id", "Name", "Size")
		for _, v := range files {
			name, size := pt.Base(v.Name), float32(v.Size)
			if len(name) > operation.ItemNameMaxInvisibleLength {
				name = name[:operation.ItemNameMaxInvisibleLength] + "..."
			}
			unitIndex := 0
			for size >= 1024 && unitIndex < len(operation.SizeUnits)-1 {
				size = size / 1024
				unitIndex++
			}
			sizeString := fmt.Sprintf("%.2f%s", size, operation.SizeUnits[unitIndex])
			fmt.Printf("%-10d%-40s%-10s\n", id, name, sizeString)
			id++
		}
	}
	return nil
}

func ShowFullName(parentName string, id int) error {
	parentInfo, err := operation.GetItemInfo(parentName)
	if err != nil {
		return err
	}
	if id < 1 || id > len(parentInfo.Children) {
		fmt.Print("There are no files or folders with an id of " + strconv.Itoa(id) + "!")
		return nil
	}
	var files, directories []*pb.ItemBasicInfo
	for _, child := range parentInfo.Children {
		if child.Type == pb.ItemType_File {
			files = append(files, child)
		} else {
			directories = append(directories, child)
		}
	}
	directories = append(directories, files...)
	fmt.Print(pt.Base(directories[id-1].Name))
	return nil
}
