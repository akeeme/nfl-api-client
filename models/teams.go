package models

type Team struct {
	TEAM_KEY 	uint 	`gorm:"primaryKey"`
	Name		string	`gorm:"column:NAME"`
	City		string	`gorm:"column:CITY"`
	Website		string	`gorm:"column:WEBSITE"`
	Wins		uint		`gorm:"column:WINS"`
	Losses		uint		`gorm:"column:LOSSES"`
	Ties		uint		`gorm:"column:TIES"`
	Position	string	`gorm:"column:POSITION"`
	Division	string	`gorm:"column:DIVISION"`
	Coach		string	`gorm:"column:COACH"`
	Stadium		string	`gorm:"column:STADIUM"`
	Owners		string	`gorm:"column:OWNERS"`
	Established uint	`gorm:"column:ESTABLISHED"`
}