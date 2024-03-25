package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	return errors.New("error fake")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}