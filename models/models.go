package models

import (
	"time"
)

type User struct {
	UserID       uint      `gorm:"primaryKey"`
	Username     string    `gorm:"size:100;unique;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	Email        string    `gorm:"size:255;unique;not null"`
	FullName     string    `gorm:"size:255"`
	JoinDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Member struct {
	MemberID        uint      `gorm:"primaryKey"`
	UserID          uint      `gorm:"unique;not null"`
	User            User      `gorm:"foreignKey:UserID"`
	RegNum          string    `gorm:"size:225;unique;not null"`
	MembershipLevel string    `gorm:"size:50"`
	JoinDate        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type CoreMember struct {
	CoreMemberID uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"unique;not null"`
	User         User      `gorm:"foreignKey:UserID"`
	Role         string    `gorm:"size:100"`
	JoinDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Staff struct {
	StaffID    uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"unique;not null"`
	User       User      `gorm:"foreignKey:UserID"`
	Role       string    `gorm:"size:100"`
	Department string    `gorm:"size:100"`
	HireDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Event struct {
	EventID     uint   `gorm:"primaryKey"`
	EventName   string `gorm:"size:255;not null"`
	EventDate   time.Time
	Description string    `gorm:"type:text"`
	Location    string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Attendee struct {
	UserID  uint `gorm:"primaryKey"`
	EventID uint `gorm:"primaryKey"`
}

type Gallery struct {
	ImageID     uint      `gorm:"primaryKey"`
	EventID     uint      `gorm:"foreignKey:EventID"`
	UserID      uint      `gorm:"foreignKey:UserID"`
	ImageURL    string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	UploadedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Comment struct {
	CommentID   uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"foreignKey:UserID"`
	EventID     uint      `gorm:"foreignKey:EventID"`
	CommentText string    `gorm:"type:text"`
	CommentedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Project struct {
	ProjectID   uint      `gorm:"primaryKey"`
	ProjectName string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"size:50;default:'Pending'"`
	CreatedByID uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ProjectRequest struct {
	RequestID   uint      `gorm:"primaryKey"`
	ProjectID   uint      `gorm:"foreignKey:ProjectID"`
	UserID      uint      `gorm:"foreignKey:UserID"`
	Status      string    `gorm:"size:50;default:'Pending'"`
	RequestedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ProjectMember struct {
	ProjectID uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
}

type Notice struct {
	NoticeID    uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:255;not null"`
	Content     string    `gorm:"type:text"`
	CreatedByID uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
