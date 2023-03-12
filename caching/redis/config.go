package redis

type Config struct {
	DB      int    `mapstructure:"db"`
	Passwd  string `mapstructure:"passwd"`
	Address string `mapstructure:"address"`
}
