package models

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	Score        string        `json:"score"`
	Date         string        `json:"date" gorm:"not null"`      // Fecha del partido
	MatchType    string        `json:"matchType" gorm:"not null"` // Tipo de partido: "individual" o "dobles"
	MatchPlayers []MatchPlayer `gorm:"foreignKey:MatchID;references:ID" json:"matchPlayers"`
}
