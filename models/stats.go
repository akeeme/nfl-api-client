package models

type Stat struct {
	PLAYER_KEY			uint	`gorm:"foreignkey:PLAYER_KEY"`
	Pass_Yds			uint	`gorm:"column:PASS_YDS"`
	Pass_Yds_Att		float32	`gorm:"column:PASS_YDS/ATT"`
	Pass_Att			uint	`gorm:"column:PASS_ATT"`
	Pass_Cmp			uint	`gorm:"column:PASS_CMP"`
	Pass_Cmp_Percent	float32	`gorm:"column:PASS_CMP_PERC"`
	Pass_TD				uint	`gorm:"column:PASS_TD"`
	Pass_Int			uint	`gorm:"column:PASS_INT"`
	Pass_Rate			float32	`gorm:"column:PASS_RATE"`
	Pass_1st 			uint	`gorm:"column:PASS_1ST"`
	Pass_1st_Percent	float32	`gorm:"column:PASS_1ST_PERC"`
	Pass_20_Plus		uint	`gorm:"column:PASS_20+"`
	Pass_40_Plus		uint	`gorm:"column:PASS_40+"`
	Pass_Lng			uint	`gorm:"column:PASS_LNG"`
	Pass_Sck			uint	`gorm:"column:PASS_SCK"`
	Pass_Sck_Yds		uint	`gorm:"column:PASS_SCKY"`
	Rush_Yds			uint	`gorm:"column:RUSH_YDS"`
	Rush_Att			uint	`gorm:"column:RUSH_ATT"`
	Rush_TD				uint	`gorm:"column:RUSH_TD"`
	Rush_20_Plus		uint	`gorm:"column:RUSH_20+"`
	Rush_40_Plus		uint	`gorm:"column:RUSH_40+"`
	Rush_Lng			uint	`gorm:"column:RUSH_LNG"`
	Rush_1st			uint	`gorm:"column:RUSH_1ST"`
	Rush_1st_Percent	float32	`gorm:"column:RUSH_1ST_PERC"`
	Rush_FUM			uint	`gorm:"column:RUSH_FUM"`
}