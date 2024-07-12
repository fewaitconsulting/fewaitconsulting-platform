package models

import "gorm.io/gorm"

type Course struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}
