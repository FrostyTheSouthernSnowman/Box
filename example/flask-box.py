import flask
app = flask.Flask(__name__)
@app.route('/')
def home():
  return 'Micron will automatically import Flask and create an app variable because use_default was specified.'
if __name__=="__main__":
	app.run(port=80)
