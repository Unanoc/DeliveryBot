package config

// Config is struct for parsing of configuration file.
type Config struct {
	Connection  string `json:"connection"`
	AccessToken string `json:"access_token"`
}
