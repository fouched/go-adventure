package main

import (
	"fmt"
	"github.com/fouched/go-adventure/internal/models"
)

func createWorld(currentGame *models.Game) (map[string]models.Room, int) {
	rooms := make(map[string]models.Room)
	monsters := 0

	for y := MAX_Y_AXIS; y >= MAX_Y_AXIS*-1; y-- {
		for x := MAX_X_AXIS * -1; x <= MAX_X_AXIS; x++ {
			location := fmt.Sprintf("%d,%d", x, y)
			rm := generateRoom(location)

			if rm.Monster != nil {
				monsters++
			}

			rooms[location] = rm
		}
	}

	return rooms, monsters
}
