package command

type CommandPipe struct {
	stagedCommands []Command
	undoneCommands []Command
}

func (cp *CommandPipe) Stage(command Command) {
	if hadEffect := command.Stage(); hadEffect { // TODO: It isn't very clear that Stage() has other effects than returning a bool
		cp.stagedCommands = append(cp.stagedCommands, command)
	}
}

func (cp *CommandPipe) Undo() {
	if len(cp.stagedCommands) > 0 {
		toUndo := popBack(&cp.stagedCommands)
		toUndo.Undo()
		cp.undoneCommands = append(cp.undoneCommands, toUndo)
	}
}

func (cp *CommandPipe) Redo() {
	if len(cp.undoneCommands) > 0 {
		toRedo := popBack(&cp.undoneCommands)
		cp.Stage(toRedo)
	}
}

func popBack(commands *[]Command) Command {
	popped := (*commands)[len(*commands)-1]
	*commands = (*commands)[:len(*commands)-1]
	return popped
}

func (cp *CommandPipe) Commit() {
	for _, stagedCommand := range cp.stagedCommands {
		stagedCommand.Execute()
	}
	cp.stagedCommands = make([]Command, 0, 0)
}

func (cp *CommandPipe) NumberOfStaged() int {
	return len(cp.stagedCommands)
}

func (cp *CommandPipe) NumberOfUndone() int {
	return len(cp.undoneCommands)
}
