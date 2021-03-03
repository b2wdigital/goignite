package recoverer

import (
	gichi "github.com/b2wdigital/goignite/chi/v1"
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	enabled = gichi.ExtRoot + ".recover.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable recover middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
