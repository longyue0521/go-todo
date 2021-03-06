package todo

import (
	"errors"
	"fmt"
)

var (
	_                    Todo = &todo{}
	ErrEmptyItem              = errors.New("empty item")
	ErrItemHasMaskedDone      = errors.New("item has done")
	ErrNoSuchItem             = errors.New("no souch item")
)

type Todo interface {
	Add(item string) (itemID uint64, err error)
	List() []Item
	Done(itemID uint64) error
}

type Item struct {
	ID    uint64
	Value string
	Done  bool
}

func (i Item) String() string {
	if i.Done {
		return fmt.Sprintf("[Done] %s", i.Value)
	}
	return fmt.Sprintf("%s", i.Value)
}

type todo struct {
	items          []Item
	doneItemsIndex []int
}

func NewTodo() Todo {
	return &todo{
		items:          make([]Item, 0),
		doneItemsIndex: make([]int, 0),
	}
}

func (t *todo) Add(item string) (id uint64, err error) {
	if item == "" {
		return 0, ErrEmptyItem
	}
	id = uint64(len(t.items) + 1)
	t.items = append(t.items, Item{ID: id, Value: item, Done: false})
	return
}

func (t *todo) List() []Item {
	list := []Item{}
	for _, v := range t.items {
		if !v.Done {
			list = append(list, v)
		}
	}
	for i := len(t.doneItemsIndex) - 1; i >= 0; i-- {
		list = append(list, t.items[t.doneItemsIndex[i]])
	}
	return list
}

func (t *todo) Done(id uint64) error {

	if id < 1 || uint64(len(t.items)) < id {
		return ErrNoSuchItem
	}

	index := int(id - 1)
	if t.items[index].Done == true {
		return ErrItemHasMaskedDone
	}

	t.items[index].Done = true
	t.doneItemsIndex = append(t.doneItemsIndex, index)
	return nil

}
