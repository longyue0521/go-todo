package main

import "errors"

var (
	_                    Todo = &todo{}
	ErrItemHasMaskedDone      = errors.New("item has done")
	ErrNoSuchItem             = errors.New("no souch item")
)

type Todo interface {
	Add(string)
	List() []Item
	Done(id uint64) error
}

type Item struct {
	ID    uint64
	Value string
	Done  bool
}

type todo struct {
	items         []Item
	undoneItemsID []uint64
	doneItemsID   []uint64
}

func NewTodo() Todo {
	return &todo{
		items:       make([]Item, 0),
		doneItemsID: make([]uint64, 0),
	}
}

func (t *todo) Add(value string) {
	id := uint64(len(t.items))
	t.items = append(t.items, Item{ID: id, Value: value, Done: false})
}

func (t *todo) List() []Item {
	list := []Item{}
	for _, v := range t.items {
		if !v.Done {
			list = append(list, v)
		}
	}
	for i := len(t.doneItemsID) - 1; i >= 0; i-- {
		list = append(list, t.items[t.doneItemsID[i]])
	}
	return list
}

func (t *todo) Done(id uint64) error {

	if uint64(len(t.items)) <= id {
		return ErrNoSuchItem
	}

	if t.items[id].Done == true {
		return ErrItemHasMaskedDone
	}

	t.items[id].Done = true
	t.doneItemsID = append(t.doneItemsID, id)
	return nil
}
