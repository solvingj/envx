package config

type EnvxConfigV0 struct {
	Version   string `json:"version,omitempty"`
	IsDefault bool   `json:"isDefault,omitempty"`
}
