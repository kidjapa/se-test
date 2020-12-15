package config

import (
    "errors"
    "fmt"
    "github.com/tkanos/gonfig"
    "os"
    "strconv"
)

var (
    ErrEnvVarEmpty = errors.New("getenv: environment variable empty")
)

// Representa todas as configurações da aplicação
type Configuration struct {
    DbUsername string       `json:"db_username"`
    DbPassword string       `json:"db_password"`
    DbPort     string       `json:"db_port"`
    DbHost     string       `json:"db_host"`
    DbName     string       `json:"db_name"`
    Users      []AccessUser `json:"access_users"`
}

type AccessUser struct {
    User string `json:"user"`
    Pass string `json:"pass"`
}

func GetConfig() Configuration {
    configuration := Configuration{}

    goEnv := "local"
    if os.Getenv("GO_ENV") == "" {
        goEnv = os.Getenv("GO_ENV")
    }

    fmt.Println("======================")
    fmt.Println("GO_ENV: " + goEnv)
    fmt.Println("======================")
    if err := gonfig.GetConf("config/config_files/config."+goEnv+".json", &configuration); err != nil {
        panic("\nErro ao recuperar configurações do serviço - se-test")
    }
    return configuration
}

/**
Return the Environment variable "PORT"
*/
func GetPortEnv() int64 {
    if os.Getenv("PORT") != "" {
        if port, err := getEnvInt64("PORT"); err == nil {
            return port
        } else {
            return 8085
        }
    }
    return 8085
}

func getenvStr(key string) (string, error) {
    v := os.Getenv(key)
    if v == "" {
        return v, ErrEnvVarEmpty
    }
    return v, nil
}

func getEnvInt64(key string) (int64, error) {
    v := os.Getenv(key)
    if v == "" {
        return -1, ErrEnvVarEmpty
    }
    if i, e := strconv.ParseInt(v, 10, 64); e == nil {
        return i, nil
    } else {
        return -1, e
    }
}

func getenvBool(key string) (bool, error) {
    s, err := getenvStr(key)
    if err != nil {
        return false, err
    }
    v, err := strconv.ParseBool(s)
    if err != nil {
        return false, err
    }
    return v, nil
}
