package unit_tests

import (
	"testing"

	"github.com/FrostyTheSouthernSnowman/Box/boxes"
	"github.com/FrostyTheSouthernSnowman/Box/core"
)

func ArrayEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestGetStringInBetween(t *testing.T) {
	stringToTest := "<tag>This data should be returned.</tag>"
	expect := "This data should be returned."
	stringInBetween, found := core.GetStringInBetween(stringToTest, "<tag>", "</tag>")
	if stringInBetween != expect || !found {
		t.Fatalf("string in between should return \"%s\", but instead returned \"%s\"!", expect, stringInBetween)
	}
}

func TestBoxConfigConstructor(t *testing.T) {
	Config := core.CreateBoxConfig("test config", boxes.Box{"test base", "foo", "bar"}, []string{"test code"})
	if Config.Name != "test config" || Config.Base.Name != "test base" || !ArrayEqual(Config.Code, []string{"test code"}) {
		t.Fatalf("Expected: Returned, %s:%s, %s:%s, {%s}:{%s}", Config.Name, "test config", Config.Base, "test base", Config.Code[0], "test code")
	}
}
