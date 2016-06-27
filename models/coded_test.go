package models

import "testing"

func TestIntersection(t *testing.T) {
	if len(computeIntersection([]string{"a", "b"}, []string{"a"})) != 1 {
		t.Error("Incorrect number of intersecting elements")
	}
	if len(computeIntersection([]string{"a", "b"}, []string{"a", "b"})) != 2 {
		t.Error("Incorrect number of intersecting elements")
	}
	if len(computeIntersection([]string{"a", "b"}, []string{"c", "d"})) != 0 {
		t.Error("Incorrect number of intersecting elements")
	}
}

func TestPreferredCode(t *testing.T) {
	codes := make(map[string][]string, 2)
	codes["a"] = []string{"aa", "ab"}
	codes["b"] = []string{"ba", "bb"}
	coded := Coded{Codes: codes}
	prefCode := coded.PreferredCode([]string{"b"})
	if prefCode.Code != "ba" {
		t.Error("Returned incorrect code, expected", "ba", "got", prefCode.Code)
	}
}