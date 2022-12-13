package models

type Player struct {
	PLAYER_KEY		uint	`gorm:"primaryKey"`
	Name			string	`gorm:"column:NAME"`
	Jersey_Number	uint	`gorm:"column:JERSEY_NUMBER"`
	Position		string	`gorm:"column:POSITION"`
	Status			string	`gorm:"column:STATUS"`
	Height			uint	`gorm:"column:HEIGHT"`
	Weight			uint	`gorm:"column:WEIGHT"`
	Experience		uint	`gorm:"column:EXPERIENCE"`
	College			string	`gorm:"column:COLLEGE"`
	TEAM_KEY		uint	`gorm:"foreignKey:TEAM_KEY"`
}