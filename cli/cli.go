package cli

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/longyue0521/go-todo/todo"
)

var (
	ErrInvalidItemIndex = errors.New("invalid item index")
)

type CLI struct {
	t              todo.Todo
	indexItemIDMap map[int]uint64
}

func NewCLI() *CLI {
	var cli CLI
	cli.t = todo.NewTodo()
	cli.indexItemIDMap = make(map[int]uint64, 0)
	return &cli
}

func (c *CLI) Add(item string) string {
	id := c.t.Add(item)
	_, _, str := c.Items(true)
	buf := bytes.NewBufferString(str)
	for index, i := range c.indexItemIDMap {
		if i == id {
			buf.WriteString(fmt.Sprintf("Item %d added", index))
			break
		}
	}
	return buf.String()
}

func (c *CLI) Items(undoneOnly bool) (int, int, string) {
	undoneN, doneN, buf := 0, 0, new(bytes.Buffer)
	for k, v := range c.t.List() {
		index := k + 1
		c.indexItemIDMap[index] = v.ID

		if undoneOnly {
			if v.Done {
				continue
			}
		}
		//fmt.Println(v)
		buf.WriteString(fmt.Sprintf("%d. %s\n", index, v))
		if v.Done {
			doneN += 1
		} else {
			undoneN += 1
		}
	}
	return undoneN, doneN, buf.String()
}

func (c *CLI) List(isAll bool) string {
	undoneOnly := !isAll
	fmt.Println(undoneOnly)
	undone, done, str := c.Items(undoneOnly)
	//fmt.Printf("---%s", str)
	buf := bytes.NewBufferString(str)
	if undoneOnly {
		buf.WriteString(fmt.Sprintf("Total: %d items", undone))
	} else {
		buf.WriteString(fmt.Sprintf("Total: %d items, %d items done", undone, done))
	}
	return buf.String()
}

func (c *CLI) Done(index int) string {
	buf := new(bytes.Buffer)
	id := c.indexItemIDMap[index]
	if err := c.t.Done(id); err != nil {
		buf.WriteString(fmt.Sprintf("Error: %s\n", ErrInvalidItemIndex))
	} else {
		buf.WriteString(fmt.Sprintf("Item %d done.", index))
	}
	return buf.String()
}
