package FlaskWebServer

func FlaskWebServer(port string) (string, string) {
	return `import flask
app = flask.Flask(__name__)`, `if __name__=="__main__":
	app.run(port=` + port + `)`
}
