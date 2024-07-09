package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
    Body        string      `json:"body"`
}

func New(body string) (Todo, error) {
    if body == "" {
        return Todo{}, errors.New("invalid input")
    }
    return Todo{
        body,
    }, nil
}

func (todo Todo) Display() {
    fmt.Println("\n***** DISPLAY *****")
    fmt.Printf("Your todo contains the following:\n%v\n", todo.Body)
    fmt.Println("***** END *****")
}

func (todo Todo) Save() error {
    fileName := "todos.json"
    
    json, err := json.Marshal(todo)
    if err != nil {
        return err
    }

    return os.WriteFile(fileName, json, 0644)
}
