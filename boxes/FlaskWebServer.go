package boxes

var _ bool = AddBox(Box{"flask-web-server", `import flask
app = flask.Flask(__name__)`, `if __name__=="__main__":
	app.run(port=)`})

func FlaskWebServer(port string) Box {
	return Box{"flask-web-server", `import flask
app = flask.Flask(__name__)`, `if __name__=="__main__":
	app.run(port=` + port + `)`}
}
