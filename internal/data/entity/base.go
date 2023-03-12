package entity

import (
	"fmt"

	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
}

func Paginate(page int64, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		fmt.Println(page, pageSize)

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit((int(pageSize)))
	}
}
