package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	_, err := GetConfig("config.yaml")

	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = GetConfig("config2.yaml")

	if err == nil {
		t.Error("Expected err, got ", err)
	}
}
