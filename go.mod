module github.com/FrostyTheSouthernSnowman/Box

go 1.17

replace github.com/FrostyTheSouthernSnowman/Box/core => ./core

replace github.com/FrostyTheSouthernSnowman/Box/boxes => ./boxes

require (
	github.com/FrostyTheSouthernSnowman/Box/box v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
