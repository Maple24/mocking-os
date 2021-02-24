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
	Time   uint32
}

func (this *Process) ChangStatus(Status State) {
	this.Status = Status
}

func (this *Process) StopProcess() {
	this.Status = Blocked
	this.Time = 0
}
