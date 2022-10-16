package config

type ServerConfig struct {
	Name        string           `mapstructure:"name"`
	Port        int              `mapstructure:"port"`
	MysqlInfo   PostgresqlConfig `mapstructure:"postgresql"`
	RedisInfo   RedisConfig      `mapstructure:"redis"`
	LogsAddress string           `mapstructure:"logsAddress"`
}

type PostgresqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
