package todo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TodoTestSuite struct {
	suite.Suite
	items []Item
	todo  Todo
}

func (suite *TodoTestSuite) SetupTest() {
	suite.todo = NewTodo()
	suite.items = []Item{
		{ID: uint64(1), Value: "go", Done: false},
		{ID: uint64(2), Value: "java", Done: false},
		{ID: uint64(3), Value: "python", Done: false},
		{ID: uint64(4), Value: "rust", Done: false},
		{ID: uint64(5), Value: "cpp", Done: false},
	}
	for k, v := range suite.items {
		id := suite.todo.Add(v.Value)
		assert.Equal(suite.T(), v.ID, id)
		assert.Equal(suite.T(), uint64(k+1), id)
	}
}

func (suite *TodoTestSuite) TestItemString() {

	suite.items[0].Done = true
	suite.items[2].Done = true
	suite.items[4].Done = true

	buf := new(bytes.Buffer)
	for _, item := range suite.items {
		buf.WriteString(item.String())
	}

	expected := `[Done] gojava[Done] pythonrust[Done] cpp`

	assert.Equal(suite.T(), expected, buf.String())
}

func (suite *TodoTestSuite) TestEmptyList() {
	assert.Equal(suite.T(), []Item{}, NewTodo().List())
}

func (suite *TodoTestSuite) TestAddAndList() {
	assert.Equal(suite.T(), suite.items, suite.todo.List())
}

func (suite *TodoTestSuite) TestDoneError() {

	err := suite.todo.Done(0)
	assert.ErrorIs(suite.T(), err, ErrNoSuchItem)

	err = suite.todo.Done(uint64(len(suite.items) + 1))
	assert.ErrorIs(suite.T(), err, ErrNoSuchItem)

	err = suite.todo.Done(1)
	assert.NoError(suite.T(), err)

	err = suite.todo.Done(1)
	assert.ErrorIs(suite.T(), err, ErrItemHasMaskedDone)

}

func (suite *TodoTestSuite) TestDoneAndList() {

	err := suite.todo.Done(1)
	assert.NoError(suite.T(), err)

	expectd := []Item{
		{ID: uint64(2), Value: "java", Done: false},
		{ID: uint64(3), Value: "python", Done: false},
		{ID: uint64(4), Value: "rust", Done: false},
		{ID: uint64(5), Value: "cpp", Done: false},
		{ID: uint64(1), Value: "go", Done: true},
	}

	assert.Equal(suite.T(), expectd, suite.todo.List())
}

func (suite *TodoTestSuite) TestMultipleDoneAndMultipleList() {

	err := suite.todo.Done(1)
	assert.NoError(suite.T(), err)

	err = suite.todo.Done(3)
	assert.NoError(suite.T(), err)

	expectd := []Item{
		{ID: uint64(2), Value: "java", Done: false},
		{ID: uint64(4), Value: "rust", Done: false},
		{ID: uint64(5), Value: "cpp", Done: false},
		{ID: uint64(3), Value: "python", Done: true},
		{ID: uint64(1), Value: "go", Done: true},
	}

	assert.Equal(suite.T(), expectd, suite.todo.List())
	// 多次调用返回相同结果
	assert.Equal(suite.T(), expectd, suite.todo.List())
}

func TestTodoTestSuite(t *testing.T) {
	suite.Run(t, new(TodoTestSuite))
}

func TestTodo(t *testing.T) {

	items := []Item{
		{ID: uint64(1), Value: "go", Done: false},
		{ID: uint64(2), Value: "java", Done: false},
		{ID: uint64(3), Value: "python", Done: false},
		{ID: uint64(4), Value: "rust", Done: false},
		{ID: uint64(5), Value: "cpp", Done: false},
	}

	todo := NewTodo()
	assert.Equal(t, []Item{}, todo.List())

	for _, v := range items {
		todo.Add(v.Value)
	}

	assert.Equal(t, items, todo.List())

	err := todo.Done(0)
	assert.ErrorIs(t, err, ErrNoSuchItem)

	err = todo.Done(uint64(len(items) + 1))
	assert.ErrorIs(t, err, ErrNoSuchItem)

	err = todo.Done(1)
	assert.NoError(t, err)

	err = todo.Done(1)
	assert.ErrorIs(t, err, ErrItemHasMaskedDone)

	expectedListResult := []Item{
		{ID: uint64(2), Value: "java", Done: false},
		{ID: uint64(3), Value: "python", Done: false},
		{ID: uint64(4), Value: "rust", Done: false},
		{ID: uint64(5), Value: "cpp", Done: false},
		// 移动到最后，Done为true
		{ID: uint64(1), Value: "go", Done: true},
	}

	assert.Equal(t, expectedListResult, todo.List())

	err = todo.Done(3)
	assert.NoError(t, err)

	expectedListResult = []Item{
		{ID: uint64(2), Value: "java", Done: false},
		{ID: uint64(4), Value: "rust", Done: false},
		{ID: uint64(5), Value: "cpp", Done: false},
		// 移动到倒数第二，Done为true
		{ID: uint64(3), Value: "python", Done: true},
		// 移动到最后，Done为true
		{ID: uint64(1), Value: "go", Done: true},
	}

	assert.Equal(t, expectedListResult, todo.List())
	// 多次调用返回相同结果
	assert.Equal(t, expectedListResult, todo.List())

}
