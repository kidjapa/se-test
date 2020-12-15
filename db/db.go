package db

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "se-test/config"
)

type Handler struct {
    Conn   *gorm.DB
    Config config.Configuration
}

func NewDbHandler() *Handler {
    return &Handler{
        Conn:   getNewConnection(),
        Config: config.GetConfig(),
    }
}

/**
Return a new database connection
*/
func getNewConnection() *gorm.DB {
    configuration := config.GetConfig()
    connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&parseTime=true&loc=Local",
        configuration.DbUsername,
        configuration.DbPassword,
        configuration.DbHost,
        configuration.DbPort,
        configuration.DbName,
    )
    db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        panic("\nFalha ao conectar ao banco de dados")
    } else {
        fmt.Printf(`==========================================
Banco de dados conectado com sucesso: %v - %v
==========================================`, configuration.DbName, configuration.DbPort)
    }
    return db
}
