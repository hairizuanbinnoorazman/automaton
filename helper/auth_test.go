package helper_test

import (
	"testing"

	"github.com/hairizuanbinnoorazman/automaton/helper"
)

func TestGetClient(t *testing.T) {
	_, err := helper.GetClient("petaccount")
	if err != nil {
		t.Fatalf("Unable to load and generate http client. %v", err.Error())
	}
}
