package ring

type State uint8

const (
	NonParticipant State = iota
	Participant
)

var (
	stateName = map[State]string{
		NonParticipant: "Non Participant",
		Participant:    "Participant",
	}
)

func (state State) String() string {
	return stateName[state]
}
