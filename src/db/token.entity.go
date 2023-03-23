package db

type Token struct {
	Base Base `gorm:"embedded"`
	Code string
}
