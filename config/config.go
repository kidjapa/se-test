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

// Get all configurations from .json config files
type Configuration struct {
    DbUsername   string       `json:"db_username"`
    DbPassword   string       `json:"db_password"`
    DbPort       string       `json:"db_port"`
    DbHost       string       `json:"db_host"`
    DbName       string       `json:"db_name"`
    Users        []AccessUser `json:"access_users"`
    LegacyAccess LegacyAccess `json:"legacy_access"`
}

// Get enabled users access list
type AccessUser struct {
    User string `json:"user"`
    Pass string `json:"pass"`
}

// get the legacy system configurations
type LegacyAccess struct {
    Url string `json:"url"`
}

func GetConfig() Configuration {
    configuration := Configuration{}

    goEnv, err := getenvStr("GO_ENV")
    if err != nil {
        goEnv = "local"
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
            return 9090
        }
    }
    return 9090
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
