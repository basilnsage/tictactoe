package board

import (
	"github.com/google/uuid"
)

// Player structure for use in main package
type Player struct {
	Mark int8
	id   [16]byte
}

// InitPlayer create new player struct with uuid as a unique id
func InitPlayer(mark int8) Player {
	return Player{
		Mark: mark,
		id:   uuid.New(),
	}
}
