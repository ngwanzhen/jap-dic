package main

import (
	"testing"
)

func TestConvertToVocab(t *testing.T) {
	v := convertToMap("apple,りんご")
	a := map[string]string{
		"りんご": "apple",
	}

	if v["りんご"] != a["りんご"] {
		t.Errorf("Expected %s Got %s", a["りんご"], v["りんご"])
	}
}
