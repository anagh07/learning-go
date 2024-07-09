package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
    Title       string      `json:"title"`
    Body        string      `json:"body"`
    CreatedAt   time.Time   `json:"createdAt"`
}

func New(title, body string) (Note, error) {
    if title == "" || body == "" {
        return Note{}, errors.New("invalid input")
    }
    return Note{
        Title: title,
        Body: body,
        CreatedAt: time.Now(),
    }, nil
}

func (note Note) Display() {
    fmt.Println("\n***** DISPLAY *****")
    fmt.Printf("Your note Titled '%v' created on %v contains the following text:\n\n%v\n", note.Title, note.CreatedAt, note.Body)
    fmt.Println("***** END *****")
}

func (note Note) Save() error {
    fileName := strings.ReplaceAll(note.Title, " ", "_")
    fileName = strings.ToLower(fileName) + ".json"
    
    json, err := json.Marshal(note)
    if err != nil {
        return err
    }

    return os.WriteFile(fileName, json, 0644)
}

func (n *Note) GetTitle() string {
    return n.Title 
}