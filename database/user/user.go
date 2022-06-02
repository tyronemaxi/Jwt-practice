package user

import (
	"gorm.io/gorm"
	"time"
	"website/database"
)

const (
	UserTableName = "user"
)

type User struct {
	PK             int64     `json:"pk" gorm:"column:pk; primary_key;AUTO_INCREMENT"`
	UserID         string    `json:"user_id" gorm:"column:user_id"`
	DomainID       string    `json:"domain_id" gorm:"column:domain_id"`
	UserName       string    `json:"user_name" gorm:"column:username"`
	Password       string    `json:"password" gorm:"column:password"`
	DisplayName    string    `json:"display_name" gorm:"column:display_name"`
	SourceType     string    `json:"source_type" gorm:"column:source_type"`
	MobilePhone    string    `json:"mobile_phone" gorm:"column:mobile_phone"`
	MobileVerified int       `json:"mobile_verified" gorm:"column:mobile_verified"`
	Email          string    `json:"email" gorm:"column:email"`
	EmailVerified  int       `json:"email_verified" gorm:"column:email_verified"`
	Extra          string    `json:"extra" gorm:"extra"`
	CreateTime     time.Time `json:"create_time" gorm:"create_time;default:null"`
	UpdateTime     time.Time `json:"update_time" gorm:"update_time;default:null"`
	DeletedTime    time.Time `json:"deleted_time" gorm:"deleted_time;default:null"`
	ISDeleted      int       `json:"is_deleted" gorm:"is_deleted"`
}

type UserDao struct {
	database.BaseDao
}

type UsernameFilter struct {
	UserName []string
	IsDelete []int
}

func (u *User) TableName() string {
	return UserTableName
}

func (p *UserDao) getListQuery(q *gorm.DB, filter UsernameFilter) *gorm.DB {
	if len(filter.UserName) > 0 {
		q = q.Where("username in (?)", filter.UserName)
	}

	return q
}

func (p *UserDao) CreateUser(u User) error {
	db := p.GetDB()
	db = db.Create(&u)

	return db.Error
}

func (p *UserDao) GetUserCount(filter UsernameFilter) (int, error) {
	// select user from user where username='xxx';
	var num int64
	q := p.GetDB().Table("user")
	q = p.getListQuery(q, filter)

	if len(filter.IsDelete) > 0 {
		q = q.Where("is_deleted in (?)", filter.IsDelete)
		q = q.Unscoped().Count(&num)
	} else {
		q = q.Where("is_deleted = ?", database.NotDelete)
		q.Count(&num)
	}

	return int(num), q.Error
}

func (p *UserDao) GetUser(username string) (*User, error) {
	q := p.GetDB()
	res := &User{}

	q = q.Where("username", username).First(&res)

	return res, q.Error
}

func NewUserDao() *UserDao {
	return &UserDao{
		database.BaseDao{
			Engine: database.DB(),
		},
	}
}
