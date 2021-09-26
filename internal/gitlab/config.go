package gitlab

// Config is a Gitlab client configuration structure that embedded in common
// Config struct.
// nolint:tagliatelle
type Config struct {
	// BaseURL is a base part of URL for Gitlab.
	BaseURL string `yaml:"base_url"`
	// Token is a personal token to use for interacting with Gitlab API.
	Token string `yaml:"token"`
}
