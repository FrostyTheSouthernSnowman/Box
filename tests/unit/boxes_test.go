package unit_tests

import (
	"testing"

	"github.com/FrostyTheSouthernSnowman/Box/boxes"
)

func TestFlaskWebServer(t *testing.T) {
	expecteBegining, expectedEnd := `import flask
app = flask.Flask(__name__)`, `if __name__=="__main__":
	app.run(port=5050)`
	box := boxes.FlaskWebServer("5050")
	if expecteBegining != box.Begining || expectedEnd != box.End {
		t.Fatalf("FlaskWebServer returned\n\"%s,%s\", \nbut expected\"\n%s,%s\"!", box.Begining, box.End, expecteBegining, expectedEnd)
	}
}
