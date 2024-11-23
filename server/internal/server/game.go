package server

import (
	"chessing/internal/database"
)

type Game struct {
	White *database.User
	Black *database.User
	Board Board
}
