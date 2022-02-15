package config

type CONFIG struct {
	ServeConfig
	MysqlConfig
}

var Configs *CONFIG

func init() {
	Configs = &CONFIG{
		ServeConfig{
			Port: ":8088",
			// Ip:   "http://localhost:8088",
			Ip: "http://yangyangcsy.cn",
		},
		MysqlConfig{
			"root",
			"CSY19961222",
			"101.34.66.232:3306",
			"blog1222",
		},
	}
}

func (c *CONFIG) GetDNS() string {
	return c.Name + ":" + c.Password + "@tcp(" + c.MysqlConfig.Ip + ")/" + c.Dbname
}
