package config

type EnvVar struct {
	Key         Key    `json:"key"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}
