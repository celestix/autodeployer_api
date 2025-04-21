package db

type User struct {
	Id       uint `gorm:"primary_key;autoIncrement:true" json:"id"`
	Name     string
	Username string
	GhOToken string `json:"gho"`
}

func AddUser(gho string) error {
	var w = User{GhOToken: gho}
	tx := SESSION.Begin()
	tx.Create(&w)
	tx.Commit()
	return tx.Error
}

func GetuserById(userid uint) *User {
	var w User
	if err := SESSION.Model(&User{}).First(&w, "id = ?", userid).Error; err != nil {
		return nil
	}
	return &w
}

func UpdateUser(user *User) error {
	tx := SESSION.Begin()
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
