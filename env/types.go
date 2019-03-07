package env

type Env struct {
	EnvironmentId string            `json:"environmentId,omitempty"`
	IsDefault     bool              `json:"isDefault,omitempty"`
	Path          string            `json:"path,omitempty"`
	Vars          map[string]string `json:"vars,omitempty"`
}
