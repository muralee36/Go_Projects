package notes

import (
	"fmt"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"createdTime"`
}

func (n *Note) Show() {
	fmt.Println("Title : ", n.Title)
	fmt.Println("Dated: ", n.CreatedTime.Format("02-01-2006"))
	fmt.Println(n.Content)
}

func New(title, note string) *Note {
	return &Note{
		Title:       title,
		Content:     note,
		CreatedTime: time.Now(),
	}
}
