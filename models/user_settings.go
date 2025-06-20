package models

type UserSettings struct {
	UserID            uint   `json:"user_id" gorm:"primaryKey"` // 1:1 with User
	Language          string `json:"language" gorm:"size:10;default:'en'"`
	Theme             string `json:"theme" gorm:"size:10;default:'light'"`
	ConcentrationTime int    `json:"concentration_time" gorm:"default:25"`
	RelaxTime         int    `json:"relax_time" gorm:"default:5"`
}
