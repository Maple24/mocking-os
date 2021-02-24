package model

type State uint32

const (
	Ready State = iota
	Running
	Blocked
)

type Process struct {
	Pid    uint32
	Status State
}

func (this *Process) ChangStatus(Status State) {
	this.Status = Status
}
