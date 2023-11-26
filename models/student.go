package models

type GormStudent struct {
	Stud_id       uint64 `json:"stud_id" gorm:"AUTO_INCREMENT"`
	Stud_name     string `json:"stud_name" binding:"required"`
	Stud_age      uint64 `json:"stud_age" binding:"required"`
	Stud_address  string `json:"stud_address" binding:"required"`
	Stud_phonenum string `json:"stud_phonenum" binding:"required"`
}
