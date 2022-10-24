package main

import (
	"C-NetDisk/CLI/operation"
	pb "C-NetDisk/CLI/proto"
	"bufio"
	"fmt"
	"os"
	pt "path"
	"strconv"
	"strings"
)

type CommandHandler struct {
	Header string
}

func NewCommandHandler(header string) *CommandHandler {
	return &CommandHandler{Header: header}
}

func (handler *CommandHandler) WaitForCommand() {
	fmt.Print(handler.Header + "> ")
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	params := strings.Split(string(line), " ")
	if len(params) < 1 || params[0] == "" {
		handler.WaitForCommand()
	}
	handler.HandleCommand(params[0], params[1:]...)
	fmt.Println()
	handler.WaitForCommand()
}

func (handler *CommandHandler) HandleCommand(command string, params ...string) {
	if command == "list" {
		err := ShowChildren(handler.Header)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		return
	}
	if command == "fullname" {
		if len(params) < 1 {
			fmt.Println("You didn't provide the id of the file or folder!")
			fmt.Println("Usage: fullname [id]")
			return
		}
		id, err := strconv.Atoi(params[0])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		err = ShowFullName(handler.Header, id)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		return
	}
	if command == "enter" {
		if len(params) < 1 {
			fmt.Println("You didn't provide the folder you want to enter!")
			fmt.Println("Usage: enter [folder]")
			return
		}
		name := handler.Header + "/" + params[0]
		if !operation.IsItemExist(name) {
			fmt.Print("Folder \"" + name + "\" doesn't exist!")
			return
		}
		info, err := operation.GetItemInfo(name)
		if err != nil {
			fmt.Print("Error: ", err)
			return
		}
		if info.Type != pb.ItemType_Directory {
			fmt.Print("\"" + name + "\" is not a folder!")
			return
		}
		handler.Header = name
		fmt.Print("Enter into " + name)
		return
	}
	if command == "back" {
		if pt.Dir(handler.Header) != "." {
			handler.Header = pt.Dir(handler.Header)
		}
		fmt.Print("Get back to " + handler.Header)
		return
	}
	if command == "upload" {
		if len(params) < 1 {
			fmt.Println("You didn't provide the file or folder to be uploaded!")
			fmt.Println("Usage: upload [file|folder]")
		} else {
			err := operation.UploadItem(params[0], handler.Header)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		}
		return
	}
	if command == "download" {
		if len(params) < 1 {
			fmt.Println("You didn't provide the file or folder to be downloaded!")
			fmt.Println("Usage: download [file|folder] [save_path(default:\".\")]")
			return
		}
		name := handler.Header + "/" + params[0]
		savePath := "."
		if len(params) >= 2 {
			savePath = params[1]
		}
		err := operation.DownloadItem(name, savePath)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		return
	}
	if command == "delete" {
		if len(params) < 1 {
			fmt.Println("You didn't provide the file or folder to be deleted!")
			fmt.Println("Usage: delete [file|folder]")
			return
		}
		name := handler.Header + "/" + params[0]
		err := operation.DeleteItem(name)
		if err != nil {
			fmt.Print("Error: ", err)
		} else {
			fmt.Print("Delete Success")
		}
		return
	}
	if command == "rename" {
		if len(params) < 1 {
			fmt.Println("You didn't provide the file or folder to be renamed!")
			fmt.Println("Usage: rename [file|folder] [new_name]")
			return
		}
		if len(params) < 2 {
			fmt.Println("You didn't provide the new name of the file or folder!")
			fmt.Println("Usage: rename [file|folder] [new name]")
			return
		}
		name := handler.Header + "/" + params[0]
		newItemName := params[1]
		err := operation.RenameItem(name, newItemName)
		if err != nil {
			fmt.Print("Error: ", err)
		} else {
			fmt.Print("Rename Success")
		}
		return
	}
	if command == "commands" {
		ShowOptionalCommands()
		return
	}
	if command == "exit" {
		fmt.Println("Thanks for using C-NetDisk!")
		os.Exit(0)
	}
	fmt.Println("Unknown Command: " + command)
	ShowOptionalCommands()
}
