package boxes

import (
	"strconv"

	utils "github.com/FrostyTheSouthernSnowman/Box/utils"
)

var _ bool = AddBox(FlaskWebServer())

func FlaskWebServer() Box {
	return Box{"flask-web-server", "python", `import flask
app = flask.Flask(__name__)`, `if __name__=="__main__":
	app.run(port=` + strconv.Itoa(utils.Config.Box.Port) + `)`, "docker"}
}
