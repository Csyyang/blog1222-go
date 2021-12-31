package config

type CONFIG struct {
	ServeConfig
}

var Configs CONFIG

func init() {
	Configs = CONFIG{
		ServeConfig{
			Port: ":8088",
		},
	}
}
