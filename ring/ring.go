package ring

type Ring interface {
	GetPid() uint32
	GetState() State
	SetState(state State)
	SetElected(pid uint32)
	Election(pid uint32)
	Elected(pid uint32)
}

func Init(ring Ring) {
	ring.SetState(NonParticipant)
	ring.SetElected(0)
}

func CallElection(ring Ring) {
	ring.Election(ring.GetPid())
}

func ReceiveElected(ring Ring, otherPid uint32) {
	if ring.GetPid() == otherPid {
		return
	}

	ring.SetElected(otherPid)
	ring.Elected(otherPid)
}

func ReceiveElection(ring Ring, otherPid uint32) {
	pid := ring.GetPid()

	if otherPid == pid {
		ring.SetElected(pid)
		ring.Elected(pid)
	} else if otherPid > pid {
		ring.Election(otherPid)
	} else if ring.GetState() == NonParticipant {
		ring.Election(pid)
	}

	ring.SetState(Participant)
}
