package config

import (
	"os"
	"strconv"
	"testing"
)

func TestGetPortEnv(t *testing.T) {
	_ = os.Setenv("PORT", "8080")

	if p := GetPortEnv(); p != 8080 {
		t.Error("port need to be 8080, \"" + strconv.FormatInt(p, 10) + "\" given.")
	}

	if p := GetPortEnv(); p == 8085 {
		t.Error("port need to be 8080, \"" + strconv.FormatInt(p, 10) + "\" given.")
	}

	_ = os.Setenv("PORT", "")
	if p := GetPortEnv(); p != 8085 {
		t.Error("port need to be 8085, \"" + strconv.FormatInt(p, 10) + "\" given.")
	}

	_ = os.Setenv("PORT", "asdf")
	if p := GetPortEnv(); p != 8085 {
		t.Error("port need to be 8085, \"" + strconv.FormatInt(p, 10) + "\" given.")
	}

}