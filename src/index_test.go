// go test -vet off -v
package main

import (
	"testing"
)

func TestLang(t *testing.T) {
	var input string = "How are you?"
	origLang := GetLang(input)
	if origLang != "en" {
		t.Errorf("GetLang was incorrect")
	}
}
