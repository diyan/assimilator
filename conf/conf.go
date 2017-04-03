package conf

// Config struct holds all application settings
type Config struct {
	Port            int    `mapstructure:"port"`
	DatabaseURL     string `mapstructure:"db_url"`
	InitialTeam     string `mapstructure:"initial_team"`
	InitialProject  string `mapstructure:"initial_project"`
	InitialKey      string `mapstructure:"initial_key"`
	InitialPlatform string `mapstructure:"initial_platform"`
}
