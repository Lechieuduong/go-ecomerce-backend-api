package setting

type Config struct{
	Mysql MySQLSetting `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type LoggerSetting struct {
	Log_level string `maopstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups int `mapstructure:"max_backups"`
	Max_age int `mapstructure:"max_age"`
	Max_size int `mapstructure:"max_size"`
	Compress bool `mapstructure:"compress"`
}