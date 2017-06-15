// Package config contains application config for daemon and bench
package config

import (
	"gopkg.in/yaml.v2"
)

type Config interface {
	yaml.Unmarshaler
	Apply() error
	Validate() error
	// TODO: we can't have original because YAML does not have []byte like in JSON for Unmarshaler
	// TODO: we call validate in apply, which trigger validate of all the children, then the apply of all the children
	// also trigger the validate of their own and their children, the lowest level would have its validate called n times
	// where n is their nested level
}
