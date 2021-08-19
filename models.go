package main

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Name  string
	Views int
}
