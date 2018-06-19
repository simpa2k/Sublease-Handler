package mockCommand

type MockCommand struct {
	Staged bool
	Undone bool
}

func (mc *MockCommand) Stage() bool {
	mc.Staged = true
	mc.Undone = false
	return true
}

func (mc *MockCommand) Execute() {}

func (mc *MockCommand) Undo() {
	mc.Staged = false
	mc.Undone = true
}
