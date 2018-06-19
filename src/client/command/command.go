package command

type Command interface {
	Stage() bool
	Execute()
	Undo()
}
