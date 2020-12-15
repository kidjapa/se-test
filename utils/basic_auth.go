package utils

import (
    "errors"
    "fmt"
    "github.com/labstack/echo/v4"
    "se-test/config"
)

func BasicAuthValidator(username string, password string, e echo.Context) (bool, error) {
    globalConfig := config.GetConfig()
    for _, user := range globalConfig.Users {
        if user.User == username && user.Pass == password {
            return true, nil
        }
    }
    fmt.Println("User try connection", username, password, e.RealIP())
    return false, errors.New("failed to authenticate")
}