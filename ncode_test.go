package narou

import (
	"testing"
)

func TestNcode2Serial(t *testing.T) {
	s, e := Ncode2Serial("n9999cx")
	if e != nil {
		t.Error("代入することができません")
	}

	if s != 759924 {
		t.Error("結果が不正確")
	}
}
