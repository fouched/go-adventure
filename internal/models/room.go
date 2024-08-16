package models

type Room struct {
	Description string
	Sound       string
	Smell       string
}

func NewRoom(description string, sound string, smell string) *Room {
	return &Room{
		description,
		sound,
		smell,
	}
}

func (r *Room) PrintDescription() {
	green.Println(r.Description)
	green.Println(r.Sound)
	green.Println(r.Smell)
}
