package models

import "gorm.io/gorm"

type MatchPlayer struct {
	gorm.Model
	PlayerID uint  `json:"playerId"`
	MatchID  uint  `json:"matchId"`
	Winner   bool  `json:"winner"`
	Match    Match `gorm:"foreignKey:MatchID"`
}
