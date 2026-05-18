package setting

type Config struct {
	Postgre PostgreSetting `mapstructure:"postgre"`
	Logger  LoggerSetting  `mapstructure:"logger"`
}

type PostgreSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbName"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	FileName    string `mapstructure:"fileName"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxSize     int    `mapstructure:"log_level"`
	MaxAge      int    `mapstructure:"log_level"`
	Compress    bool   `mapstructure:"log_level"`
}
