// go test -vet off -v
package main

import (
	"testing"
)

func TestLang(t *testing.T) {
	origLang := GetLang("hi")
	if origLang != "eng" {
		t.Errorf("GetLang was incorrect, got: %d, want: %d.", origLang, "eng")
	}
}
