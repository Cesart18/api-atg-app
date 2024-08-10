package models

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	Date        string        `json:"date" gorm:"not null"`      // Fecha del partido
	MatchType   string        `json:"matchType" gorm:"not null"` // Tipo de partido: "individual" o "dobles"
	GamePlayers []MatchPlayer `json:"gamePlayers"  gorm:"foreignKey:MatchID;references:ID"`
}
