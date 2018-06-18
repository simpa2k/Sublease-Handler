package commands

type Command interface {
	Stage() bool
	Execute()
	Undo()
}
