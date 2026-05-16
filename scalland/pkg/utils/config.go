package utils

// Config holds the full application configuration tree.
type Config struct {
	App           AppConfig
	Databases     map[string]DatabaseConfig
	SMTPs         SMTPConfigs
	Upgrade       UpgradeConfig
	OTP           OTPConfig
	ObjectStorage map[string]ObjectStorageConfig
}

// AppConfig holds main application settings.
type AppConfig struct {
	ServerPort string `mapstructure:"server_port"`
	JWTSecret  string `mapstructure:"jwt_secret"`
	Timezone   string
	Debug      bool
}

// DatabaseConfig holds DB connection settings.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
	Timezone string
}

// SMTPConfig holds email settings.
type SMTPConfig struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

// SMTPConfigs is a slice of SMTP configs with helper methods.
type SMTPConfigs []SMTPConfig

// Default returns the first SMTP config, or nil if none configured.
func (s SMTPConfigs) Default() *SMTPConfig {
	if len(s) == 0 {
		return nil
	}
	return &s[0]
}

// Config returns the SMTP config with the given name.
func (s SMTPConfigs) Config(name string) *SMTPConfig {
	for i := range s {
		if s[i].Name == name {
			return &s[i]
		}
	}
	return nil
}

// UpgradeConfig holds self-upgrade settings.
type UpgradeConfig struct {
	Enabled          bool
	DefaultChannel   string   `mapstructure:"default_channel"`
	BinDirectory     string   `mapstructure:"bin_directory"`
	PreserveVersions int      `mapstructure:"preserve_versions"`
	Channels         UpgradeChannels
}

// UpgradeChannels holds config for each upgrade channel.
type UpgradeChannels struct {
	GitHub GitHubChannelConfig
	S3     S3ChannelConfig
	HTTP   HTTPChannelConfig
	Local  LocalChannelConfig
}

type GitHubChannelConfig struct {
	Enabled      bool
	Owner        string
	Repo         string
	Token        string
	AssetPattern string `mapstructure:"asset_pattern"`
}

type S3ChannelConfig struct {
	Enabled   bool
	Bucket    string
	Region    string
	Prefix    string
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

type HTTPChannelConfig struct {
	Enabled     bool
	BaseURL     string `mapstructure:"base_url"`
	ManifestURL string `mapstructure:"manifest_url"`
}

type LocalChannelConfig struct {
	Enabled bool
	Path    string
}

// OTPConfig holds one-time-password settings.
type OTPConfig struct {
	Length    int
	OTPType  string `mapstructure:"otp_type"`
	ValidTill int   `mapstructure:"valid_till"`
}

// ObjectStorageConfig holds S3/B2 storage settings.
type ObjectStorageConfig struct {
	Enabled   bool
	Provider  string
	Bucket    string
	Region    string
	Endpoint  string
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}
