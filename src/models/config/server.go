package config

type ServerConfig struct {
	JWT          JWT          `yaml:"jwt"`
	TokenRedis   Redis        `yaml:"token_redis"`
	QueueRedis   Redis        `yaml:"queue_redis"`
	LocalStorage LocalStorage `yaml:"local_storage"`
	Database     MySQL        `yaml:"database"`
	Zap          Zap          `yaml:"zap"`
}

type MySQL struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config_models" json:"config-models" yaml:"config_models"`
	Dbname       string `mapstructure:"db_name" json:"dbname" yaml:"db_name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max-idle-conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max-open-conns" yaml:"max_open_conns"`
	LogMode      bool   `mapstructure:"log_mode" json:"log-mode" yaml:"log_mode"`
	LogZap       string `mapstructure:"log_zap" json:"log-zap" yaml:"log_zap"`
}

func (m *MySQL) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type JWT struct {
	SigningKey  string `mapstructure:"signing_key" json:"signing-key" yaml:"signing_key"`
	ExpiresTime int64  `mapstructure:"expires_time" json:"expires-time" yaml:"expires_time"`
	BufferTime  int64  `mapstructure:"buffer_time" json:"buffer-time" yaml:"buffer_time"`
}

type LocalStorage struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link_name" json:"link-name" yaml:"link_name"`
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show_line"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode_level"`
	StacktraceKey string `mapstructure:"stacktrace_key" json:"stacktrace-key" yaml:"stacktrace_key"`
	LogInConsole  bool   `mapstructure:"log_in_console" json:"log-in-console" yaml:"log_in_console"`
}
