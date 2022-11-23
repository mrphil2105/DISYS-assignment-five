package bully

import "errors"

type Bully interface {
	GetPid() uint32
	GetAllPids() []uint32
	GetCoordinator() uint32
	SetCoordinator(uint32)
	SendElection(targetPid uint32)
	MulticastCoordinator()
	SendAnswer(targetPid uint32)
}

func Start(bully Bully) {
	bully.SetCoordinator(0)

	pid := bully.GetPid()
	allPids := bully.GetAllPids()
	maxPid := GetMaxValue(allPids)

	if pid == maxPid {
		bully.MulticastCoordinator()
	} else {
		for i := 0; i < len(allPids); i++ {
			otherPid := allPids[i]

			if otherPid > pid {
				bully.SendElection(pid)
			}
		}
	}
}

func ElectionReceived(bully Bully, otherPid uint32) error {
	if bully.GetPid() <= otherPid {
		return errors.New("the specified pid is not lower than our pid")
	}

	bully.SendAnswer(otherPid)
	Start(bully)
}

func GetMaxValue(values []uint32) uint32 {
	max := values[0]

	for i := 1; i < len(values); i++ {
		value := values[i]

		if value > max {
			max = value
		}
	}

	return max
}
