package boxes

var _ bool = AddBox("flask-web-server")

func FlaskWebServer(port string) Box {
	return Box{`import flask
app = flask.Flask(__name__)`, `if __name__=="__main__":
	app.run(port=` + port + `)`}
}
