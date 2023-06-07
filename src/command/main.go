package command

import (
	"fmt"
)

// Editor is the receiver object that executes the operations.
type Editor struct {
	Text string
}

// Command interface is the command that encapsulates an operation.
type Command interface {
	Execute() string
}

type CopyCommand struct {
	App *Application
}

func (c *CopyCommand) Execute() string {
	// Copy command logic...
	return fmt.Sprintf("Copying text: %s", c.App.Editor.Text)
}

type CutCommand struct {
	App *Application
}

func (c *CutCommand) Execute() string {
	// Cut command logic...
	text := c.App.Editor.Text
	c.App.Editor.Text = ""
	return fmt.Sprintf("Cutting text: %s", text)
}

type PasteCommand struct {
	App *Application
}

func (p *PasteCommand) Execute() string {
	// Paste command logic...
	p.App.Editor.Text = "Pasted text"
	return p.App.Editor.Text
}

type UndoCommand struct {
	App *Application
}

func (u *UndoCommand) Execute() string {
	// Undo command logic...
	if len(u.App.CommandHistory) > 0 {
		u.App.CommandHistory = u.App.CommandHistory[:len(u.App.CommandHistory)-1]
		return "Undo last command"
	}
	return "Nothing to undo"
}

// Application is the invoker object that invokes the command to perform an operation.
type Application struct {
	Editor         *Editor
	CommandHistory []Command
}

func (a *Application) ExecuteCommand(command Command) {
	a.CommandHistory = append(a.CommandHistory, command)
	fmt.Println(command.Execute())
}

func main() {
	editor := &Editor{Text: "Hello, World!"}
	app := &Application{Editor: editor}

	copyCmd := &CopyCommand{App: app}
	app.ExecuteCommand(copyCmd)

	cutCmd := &CutCommand{App: app}
	app.ExecuteCommand(cutCmd)

	pasteCmd := &PasteCommand{App: app}
	app.ExecuteCommand(pasteCmd)

	undoCmd := &UndoCommand{App: app}
	app.ExecuteCommand(undoCmd)
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "command pattern"
}

func (e Executer) Do() {
	main()
}
