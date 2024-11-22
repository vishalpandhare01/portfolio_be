package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	UserName  string    `gorm:"type:varchar(36);not null"`
	Password  string    `gorm:"type:text;not null"`
	Phone     string    `gorm:"type:varchar(36);not null"`
	Role      string    `gorm:"type:Enum('admin','user');not null"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

type Contacts struct {
	ID        string     `gorm:"type:char(36);primaryKey"`
	UserID    string     `gorm:"type:primaryKey;char(36)"`
	Name      string     `gorm:"type:varchar(36);not null"`
	Phone     string     `gorm:"type:varchar(10);not null"`
	Email     string     `gorm:"type:varchar(36);not null"`
	Message   string     `gorm:"type:text;not null"`
	IpAdress  string     `gorm:"type:varchar(36);not null"`
	Location  string     `gorm:"type:varchar(36);not null"`
	User      *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

type UserProfile struct {
	ID          string     `gorm:"type:char(36);primaryKey"`
	UserID      string     `gorm:"type:primaryKey;char(36)"`
	Name        string     `gorm:"type:varchar(36);not null"`
	Phone       string     `gorm:"type:varchar(10);not null"`
	Email       string     `gorm:"type:varchar(36);not null"`
	ProfilePic  string     `gorm:"type:text;not null"`
	Banner      string     `gorm:"type:text;not null"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Discription string     `gorm:"type:text;not null"`
	Skills      string     `gorm:"type:text;not null"`
	Projects    string     `gorm:"type:text;not null"`
	Services    string     `gorm:"type:text;not null"`
	About       string     `gorm:"type:text;not null"`
	SocialMedia string     `gorm:"type:text;not null"`
	User        *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (U *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(U.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	U.ID = uuid.New().String()
	U.Password = string(hash)
	return nil
}
func (U *UserProfile) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return nil
}
func (U *Contacts) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return nil
}
