package client

import (
	"testing"
	mc "subLease/test/utils/mockCommand"
	"subLease/src/client/command"
)

func TestStage(t *testing.T) {
	mockCommand := mc.MockCommand{}
	commandPipe := command.CommandPipe{}

	commandPipe.Stage(&mockCommand)

	if !mockCommand.Staged {
		t.Error("Command was not staged")
	}

	if commandPipe.NumberOfStaged() > 1 {
		t.Error("Command was not added to staged commands")
	}
}

func TestUndo(t *testing.T) {
	mockCommand := mc.MockCommand{}
	commandPipe := command.CommandPipe{}

	commandPipe.Stage(&mockCommand)
	commandPipe.Undo()

	if !mockCommand.Undone {
		t.Error("Command was not undone")
	}

	if commandPipe.NumberOfUndone() > 1 {
		t.Error("Command was not added to undone commands")
	}
}

func TestRedo(t *testing.T) {
	mockCommand := mc.MockCommand{}
	commandPipe := command.CommandPipe{}

	commandPipe.Stage(&mockCommand)
	commandPipe.Undo()
	commandPipe.Redo()

	if !mockCommand.Staged {
		t.Error("Command was not restaged")
	}

	if mockCommand.Undone {
		t.Error("Command was not redone")
	}

	if commandPipe.NumberOfUndone() > 0 {
		t.Error("Command was not removed from undone commands")
	}

	if commandPipe.NumberOfStaged() < 1 {
		t.Error("Command was not readded to staged commands")
	}
}
