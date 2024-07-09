// Package po

package po

type User struct {
	Id               int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Guid             string `gorm:"column:guid" json:"guid"`
	Username         string `gorm:"column:username" json:"username"`
	Password         string `gorm:"column:password" json:"password"`
	NickName         string `gorm:"column:nick_name" json:"nick_name"`
	MerchantId       int64  `gorm:"column:merchant_id" json:"merchant_id"`
	Avatar           string `gorm:"column:avatar" json:"avatar"`
	Email            string `gorm:"column:email" json:"email"`
	Phone            string `gorm:"column:phone" json:"phone"`
	CreateTime       int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime       int64  `gorm:"column:update_time" json:"update_time"`
	LastLoginTime    int64  `gorm:"column:last_login_time" json:"last_login_time"`
	LastLoginOutTime int64  `gorm:"column:last_login_out_time" json:"last_login_out_time"`
}

func (s *User) TableName() string {
	return "model"
}
