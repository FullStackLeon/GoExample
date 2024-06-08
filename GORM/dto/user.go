package dto

type UserInfo struct {
	Id    int64  `json:"id,omitempty" gorm:"column:user_id"`
	Name  string `json:"name,omitempty" gorm:"column:username"`
	Email string `json:"email,omitempty" gorm:"column:email"`
}
