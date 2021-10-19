module github.com/FrostyTheSouthernSnowman/Micron

go 1.17

replace src => ./src

replace FlaskWebServer => ./src/FlaskWebServer

require FlaskWebServer v0.0.0-00010101000000-000000000000
