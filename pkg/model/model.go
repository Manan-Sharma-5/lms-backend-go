package models

import (
	"time"
)

type User struct {
    ID                    string                 `gorm:"primaryKey;default:gen_random_uuid()"` // Using PostgreSQL's gen_random_uuid() for UUIDs
    Name                  string
    Password              string
    Email                 string                 `gorm:"uniqueIndex"`
    Role                  string
    DateJoined           time.Time
    Notes                []Note                  `gorm:"foreignKey:UserID"`
    Books                []Book                  `gorm:"foreignKey:UserID"`
    PreviousYearQuestions []PreviousYearQuestion  `gorm:"foreignKey:UserID"`
    Classrooms           []Classroom             `gorm:"foreignKey:UserID"`
}

type Note struct {
    ID          string     `gorm:"primaryKey;default:gen_random_uuid()"`
    Title       string
    Description string
    Content     string
    Status      string
    Stream      string
    Subject     string
    Year        int
    UploadDate  time.Time
    UserID      string     `gorm:"index"`
    User        User       `gorm:"foreignKey:UserID;references:ID"`
}

type PreviousYearQuestion struct {
    ID         string     `gorm:"primaryKey;default:gen_random_uuid()"`
    Subject    string
    Stream     string
    Year       int
    DatePosted time.Time
    UserID     string     `gorm:"index"`
    Content    string
    User       User       `gorm:"foreignKey:UserID;references:ID"`
}

type Book struct {
    ID          string     `gorm:"primaryKey;default:gen_random_uuid()"`
    Title       string
    Author      string
    Description string
    Price       float64
    Image       string
    Status      string
    UserID      string     `gorm:"index"`
    User        User       `gorm:"foreignKey:UserID;references:ID"`
}

type Classroom struct {
    ID          string     `gorm:"primaryKey;default:gen_random_uuid()"`
    Name        string
    Stream      string
    Subject     string
    Year        int
    URL         string
    CreatedDate time.Time
    UserID      string     `gorm:"index"`
    User        User       `gorm:"foreignKey:UserID;references:ID"`
}
