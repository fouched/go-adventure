package models

import "math/rand/v2"

type Room struct {
	Description string
	Sound       string
	Smell       string
	Items       map[string]ArmoryItem
	Monster     *Monster
	Location    string
}

var descriptions = []string{
	"Whatever furniture used to adorn this chamber has long since been carried away. \nIt is completely empty, save for a few splinters of rotted wood and some lichen. ",
	"This room appears to have once been a barracks of some sort, if the remnants of the cots lining each wall are any indication. ",
	"Based on the rusted implements littering the floor, this room must have been a torture chamber of some sort in the distant past. ",
	"The flickering light from your torch does little to lighten the dim hallway you move through. ",
	"It is impossible to tell just what this location might have been in the past, but bones are strewn everywhere.",
	"Scattered bits of bone, armor, and rusted weaponry suggest that a major battle took place here long ago.",
	"From the scorch marks on the walls and floor, it appears that something burned here not too long ago. \nGiven the amount of scorching, it must have been a huge blaze.",
}

var sounds = []string{
	"The faint sound of dripping water echoes off the walls.",
	"Off in the darkness you hear something scraping against the stone floor.",
	"You think you hear the clicking of talons against the stone floor somewhere in the distance. ",
	"The faint echo of something suspiciously like a growl comes from further out in the darkness.",
	"You hear a faint grunting noise from somewhere in the distance.",
	"What sounds like faint laughter wends its way into the room.",
	"You're not sure, but something like a faint hiss seems to sound off in the distance.",
}

var smells = []string{
	"There is a faint acrid smell in the air. ",
	"Given the stench in the air, something must have died nearby recently. ",
	"The faint smell of wood smoke suggests that someone, or something, has been through here recently. ",
	"The coppery smell of blood hangs in the air. ",
	"You are not certain what is causing the foul stench in the air, but you have your suspicions...",
	"The smell of rancid meat suggests that something was rotting here not too long ago.",
}

func NewRoom(location string) Room {
	return Room{
		Description: descriptions[rand.IntN(len(descriptions))],
		Sound:       sounds[rand.IntN(len(sounds))],
		Smell:       smells[rand.IntN(len(smells))],
		Items:       make(map[string]ArmoryItem),
		Location:    location,
	}
}

func (r *Room) PrintDescription() {
	green.Println(r.Description)
	green.Println(r.Sound)
	green.Println(r.Smell)
}
