package unit_tests

import (
	"testing"

	"github.com/FrostyTheSouthernSnowman/Box/boxes"
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

func TestFlaskWebServer(t *testing.T) {
	expecteBegining, expectedEnd := `import flask
app = flask.Flask(__name__)`, `if __name__=="__main__":
	app.run(port=5050)`
	begining, end := boxes.FlaskWebServer("5050")
	if expecteBegining != begining || expectedEnd != end {
		t.Fatalf("FlaskWebServer returned\n\"%s,%s\", \nbut expected\"\n%s,%s\"!", begining, end, expecteBegining, expectedEnd)
	}
}

func TestGetStringInBetween(t *testing.T) {
	stringToTest := "<tag>This data should be returned.</tag>"
	expect := "This data should be returned."
	stringInBetween, found := boxes.GetStringInBetween(stringToTest, "<tag>", "</tag>")
	if stringInBetween != expect || !found {
		t.Fatalf("string in between should return \"%s\", but instead returned \"%s\"!", expect, stringInBetween)
	}
}

func TestBoxConfigConstructor(t *testing.T) {
	Config := boxes.CreateBoxConfig("test config", "test base", []string{"test code"})
	if Config.Name != "test config" || Config.Base != "test base" || !ArrayEqual(Config.Code, []string{"test code"}) {
		t.Fatalf("Expected: Returned, %s:%s, %s:%s, {%s}:{%s}", Config.Name, "test config", Config.Base, "test base", Config.Code[0], "test code")
	}
}
