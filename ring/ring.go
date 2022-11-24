package ring

type Ring interface {
	GetPid() uint32
	GetState() State
	SetState(state State)
	SetElected(pid uint32)
	SendElection(pid uint32)
	SendElected(pid uint32)
}

func Init(ring Ring) {
	ring.SetState(NonParticipant)
	ring.SetElected(0)
}

func CallElection(ring Ring) {
	ring.SendElection(ring.GetPid())
}

func ReceiveElected(ring Ring, otherPid uint32) {
	if ring.GetPid() == otherPid {
		return
	}

	ring.SetElected(otherPid)
	ring.SendElected(otherPid)
}

func ReceiveElection(ring Ring, otherPid uint32) {
	pid := ring.GetPid()

	if otherPid == pid {
		ring.SetElected(pid)
		ring.SendElected(pid)
	} else if otherPid > pid {
		ring.SendElection(otherPid)
	} else if ring.GetState() == NonParticipant {
		ring.SendElection(pid)
	}

	ring.SetState(Participant)
}
