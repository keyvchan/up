package config

var HOME string
var ConfigPath string

type Item struct {
	Kind    string
	Name    string
	Command string
	Args    []string
}
