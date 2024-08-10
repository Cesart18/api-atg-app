package models

import "gorm.io/gorm"

type MatchPlayer struct {
	gorm.Model
	MatchID  uint
	PlayerID uint
	Winner   bool // Indica si el jugador ganó
}
