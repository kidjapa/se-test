package tests

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "os"
    "se-test/config"
    "se-test/db"
)

func getGoEnv() *string {
    goEnv := "local"
    if os.Getenv("GO_ENV") != "" {
        goEnv = os.Getenv("GO_ENV")
    }
    return &goEnv
}

func NewTestDBConnectionsHandlerLocal() *db.Handler {
    testConfig := NewConfigurationTestSchema()
    return &db.Handler{
        Config: testConfig.DbConfig,
        Conn:   testConfig.getConnection(),
    }
}

type HelperTestConfiguration struct {
    DbConfig config.Configuration
}

/**
Get basic configuration for get ConnectionDBHandler for test scripts
*/
func NewConfigurationTestSchema() *HelperTestConfiguration {
    goEnv := getGoEnv()

    if *goEnv == "local" {
        return &HelperTestConfiguration{
            DbConfig: config.Configuration{
                DbUsername: "root",
                DbPassword: "root",
                DbPort:     "3306",
                DbHost:     "127.0.0.1",
                DbName:     "se_test",
                Users: []config.AccessUser{
                    {
                        User: "bruno",
                        Pass: "zShM39hkPqcyjQu",
                    },
                },
            },
        }
    } else {
        panic("goenv not ok, expected loca, " + *goEnv + " given.")
    }
}

/**
Return a new connection for Bemtevi
*/
func (h *HelperTestConfiguration) getConnection() *gorm.DB {
    connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&parseTime=true&loc=Local",
        h.DbConfig.DbUsername,
        h.DbConfig.DbPassword,
        h.DbConfig.DbHost,
        h.DbConfig.DbPort,
        h.DbConfig.DbName,
    )
    d, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        panic("\nFalha ao conectar ao banco de dados")
    } else {
        fmt.Printf(`==========================================
Banco de dados conectado com sucesso: %v - %v
==========================================`, h.DbConfig.DbName, h.DbConfig.DbPort)
    }
    return d
}
