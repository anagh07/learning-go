package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"anagh.xyz/notes/note"
	"anagh.xyz/notes/todo"
)

type saver interface {
    Save() error
}

type outputtable interface {
    saver
    Display()
}

func getUserInput(prompt string) (string) {
    fmt.Printf("%v ", prompt)

    reader := bufio.NewReader(os.Stdin)
    data, err := reader.ReadString('\n')

    if err != nil {
        return ""
    }

    data = strings.TrimSuffix(data, "\n")
    data = strings.TrimSuffix(data, "\r")
    return data
}

func getNoteData() (string, string) {
    title := getUserInput("Enter title:")
    body := getUserInput("Enter body:")

    return title, body
}

func getTodoData() string {
    text := getUserInput("Enter todo:")
    return text
}

func saveData(obj saver) error {
    err := obj.Save()
    if err != nil {
        fmt.Println("failed to save")
        return err
    }
    fmt.Println("successfully saved to file")
    return err
}

func outPutData(obj outputtable) error {
    obj.Display()
    err := saveData(obj)
    if err != nil {
        return err
    }
    return nil
}

func printAny(obj any) {
    todoVal, isTodo := obj.(todo.Todo)
    if isTodo {
        todoVal.Display()
        return
    }

    switch obj.(type) {
    case int:
        fmt.Println("Integer: ", obj)
    case float64:
        fmt.Println("Float64: ", obj)
    case string:
        fmt.Println("Float64: ", obj)
    default:
        fmt.Println("Print not defined for datatype")
    }
}

func main() {
    printAny(7)    
    printAny(2.1234)

    title, text := getNoteData()

    userNote, err := note.New(title, text)
    if err != nil {
        fmt.Println(err)
        return
    } 
    err = outPutData(userNote)
    if err != nil {
        return
    }
    printAny(userNote)

    todoText := getTodoData()
    userTodo, err := todo.New(todoText)
    if err != nil {
        fmt.Println(err)
        return
    }
    outPutData(userTodo)
    printAny(userTodo)

    // ********** Generics
    sum := add(2.1, 5.81)
    fmt.Println(sum)
    
}
