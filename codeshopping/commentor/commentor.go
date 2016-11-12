package commentor

import (
	"io/ioutil"
)

type Comment struct {
	Item  string
	Title string
	Body  []byte // "byte slice", expected by the io package of go
}

func (c *Comment) Save() error {
	filename := c.Item + c.Title + ".txt"
	return ioutil.WriteFile(filename, c.Body, 0600) // 0600 unix permission of the file
}

func LoadComment(item string, title string) (*Comment, error) {
	filename := item + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Comment{Title: title, Body: body}, nil
}
