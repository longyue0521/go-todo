package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CLITestSuite struct {
	suite.Suite
	cli *CLI
}

func (suite *CLITestSuite) SetupTest() {
	suite.cli = NewCLI()
}

func (suite *CLITestSuite) TestItems() {

	suite.cli.Add("golang")
	suite.cli.Add("python")
	suite.cli.Add("java")
	suite.cli.Add("rust")
	undone, done, actual := suite.cli.Items(true)
	expected := `1. golang
2. python
3. java
4. rust
`
	assert.Equal(suite.T(), expected, actual)
	assert.Equal(suite.T(), 4, undone)
	assert.Equal(suite.T(), 0, done)

	suite.cli.Done(1)
	suite.cli.Done(3)

	undone, done, actual = suite.cli.Items(true)
	expected = `1. python
2. rust
`
	assert.Equal(suite.T(), expected, actual)
	assert.Equal(suite.T(), 2, undone)
	assert.Equal(suite.T(), 0, done)

	undone, done, actual = suite.cli.Items(false)
	expected = `1. python
2. rust
3. [Done] java
4. [Done] golang
`
	assert.Equal(suite.T(), expected, actual)
	assert.Equal(suite.T(), 2, undone)
	assert.Equal(suite.T(), 2, done)
}

func (suite *CLITestSuite) TestEmptyList() {

	expected := "Total: 0 items"
	assert.Equal(suite.T(), expected, suite.cli.List(false))
	assert.Equal(suite.T(), expected, suite.cli.List(false))

	expected = "Total: 0 items, 0 items done"
	assert.Equal(suite.T(), expected, suite.cli.List(true))
	assert.Equal(suite.T(), expected, suite.cli.List(true))
}

func (suite *CLITestSuite) TestAddAndList() {

	suite.cli.Add("golang")
	suite.cli.Add("python")
	actual := suite.cli.Add("java")
	expected := `1. golang
2. python
3. java
Item 3 added`
	assert.Equal(suite.T(), expected, actual)

	actual = suite.cli.List(false)
	expected = `1. golang
2. python
3. java
Total: 3 items`
	assert.Equal(suite.T(), expected, actual)

	actual = suite.cli.List(true)
	expected = `1. golang
2. python
3. java
Total: 3 items, 0 items done`
	assert.Equal(suite.T(), expected, actual)

}

func (suite *CLITestSuite) TestDoneError() {
	actual := suite.cli.Done(0)
	expected := "Error: invalid item index\n"
	assert.Equal(suite.T(), expected, actual)

	actual = suite.cli.Done(1)
	expected = "Error: invalid item index\n"
	assert.Equal(suite.T(), expected, actual)
}

func (suite *CLITestSuite) TestDoneAndList() {

	suite.cli.Add("golang")
	suite.cli.Add("python")
	actual := suite.cli.Add("java")
	expected := `1. golang
2. python
3. java
Item 3 added`
	assert.Equal(suite.T(), expected, actual)

	actual = suite.cli.Done(1)
	expected = `Item 1 done.`
	assert.Equal(suite.T(), expected, actual)

	actual = suite.cli.List(false)
	expected = `1. python
2. java
Total: 2 items`
	assert.Equal(suite.T(), expected, actual)

	actual = suite.cli.List(true)
	expected = `1. python
2. java
3. [Done] golang
Total: 2 items, 1 items done`
	assert.Equal(suite.T(), expected, actual)
}

func TestCmdTestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}
