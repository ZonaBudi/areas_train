package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Area struct {
		ID        int64   `gorm:"column:id;primaryKey;"`
		AreaValue float64 `gorm:"column:area_value"`
		AreaType  string  `gorm:"column:type"`
	}
	AreaRepository struct {
		DB *gorm.DB
	}
)

func (_r *AreaRepository) InsertArea(param1 int64, param2 int64, typeArea string) (err error) {
	area := new(Area)
	switch typeArea {
	case "persegi panjang", "persegi":
		formula := param1 * param2
		area.AreaValue = float64(formula)
		area.AreaType = typeArea
	case "segitiga":
		formula := float64(0.5) * float64((param1 * param2))
		area.AreaValue = formula
		area.AreaType = typeArea
	default:
		area.AreaValue = 0
		area.AreaType = "undefined data"
	}
	err = _r.DB.Create(&area).Error
	if err != nil {
		return err
	}
	return nil
}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:password@(localhost:3306)/area?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type service struct {
	repository AreaRepository
}

func (_u service) Service() error {
	err := _u.repository.InsertArea(1, 2, "persegi panjang")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	gorm, err := Connect()
	if err != nil {
		panic(err)
	}
	gorm.AutoMigrate(&Area{})
	ar := AreaRepository{
		DB: gorm,
	}
	service := service{
		repository: ar,
	}
	service.Service()
}
