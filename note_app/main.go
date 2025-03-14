package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	fileop "note_app/fileOp"
	"note_app/notes"
	"os"
	"strings"
)

const NotesFile string = "notes.json"

func main() {
	title, content := getNote()
	newNote := notes.New(title, content)
	newNote.Show()
	data, err := fileop.ReadJsonFile(NotesFile)
	if err != nil {
		fmt.Print(err)
	}
	newNoteByte, err := json.Marshal(newNote)
	if err != nil {
		fmt.Print(err)
	}
	var newNoteJson map[string]interface{}
	err = json.Unmarshal(newNoteByte, &newNoteJson)
	if err != nil {
		fmt.Print(err)
	}
	data = append(data, newNoteJson)
	err = fileop.WriteJsonFile(NotesFile, data)
	if err != nil {
		fmt.Print(err)
	}
}

func getNote() (string, string) {
	title, err1 := getInput("Enter Note Title : ")
	if err1 != nil {
		fmt.Println(err1)
		return "", ""
	}
	content, err2 := getInput("Enter Your Note : ")
	if err2 != nil {
		fmt.Println(err1)
		return "", ""
	}
	return title, content
}

func getInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	if err != nil {
		return "", errors.New("something went wrong")
	}
	return text, nil
}
