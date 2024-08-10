package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name              string        `json:"name" gorm:"unique"`
	SinglePoints      int           `json:"singlePoints" gorm:"default 0"`
	DoublePoints      int           `json:"doublePoints" gorm:"default 0"`
	IsMembershipValid bool          `json:"isMembershipValid" gorm:"default false"`
	IsPayedBalls      bool          `json:"isPayedBalls" gorm:"default false"`
	MatchPlayers      []MatchPlayer `gorm:"foreignKey:PlayerID;references:ID"`
}
