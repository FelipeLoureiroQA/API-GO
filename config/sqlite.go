package config

import (
	"os"

	"github.com/FelipeLoureiroQA/API-GO/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	//Verifica se o arquivo do banco de dados existe
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating it")
	}
	//Criação de diretorio e do arquivo do banco de dados
	err = os.MkdirAll("./db", os.ModePerm)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(dbPath)
	if err != nil {
		return nil, err
	}
	// Em sistema operacional é importante sempre fechar o arquivo
	file.Close()

	//Conexão com o banco de dados
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
	if err != nil {
		logger.Errof("Error initializing SQLite: %v", err.Error())
		return nil, err

	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errof("Error migrating schema: %v", err.Error())
		return nil, err
	}

	return db, nil
}
