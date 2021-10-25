import flask
app = flask.Flask(__name__)
@app.route('/')
def home():
  return '<h1>Box will automatically import Flask and create an app variable because use_default was specified. It will also run the server!</h1>'
if __name__=="__main__":
	app.run(port=80)
