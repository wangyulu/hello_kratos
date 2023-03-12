package entity

type User struct {
	Model
	Id     int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name   string `gorm:"column:name;NOT NULL" json:"name"`         // 客易达平台应用ID
	Age    int32  `gorm:"column:age;default:0;NOT NULL" json:"age"` // 租户id
	Mobile string `gorm:"column:mobile;NOT NULL" json:"mobile"`     // 客易达平台open_kid
}

func (e *User) TableName() string {
	return "user"
}
