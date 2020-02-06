package logger

import (
	"testing"
)

func TestGetConfig(t *testing.T) {

	details := []string{"prod", "dev", "error"}

	_, err := GetLogger(details[0], "")
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = GetLogger(details[1], "")
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = GetLogger(details[2], "")
	if err == nil {
		t.Error("Expected err, got ", err)
	}
	_, err = GetLogger(details[1], "mylog")
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	_, err = GetLogger(details[2], "mylog")
	if err == nil {
		t.Error("Expected err, got ", err)
	}
}
