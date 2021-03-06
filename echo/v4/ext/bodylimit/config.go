package giechobodylimit

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
)

const (
	enabled = giecho.ExtRoot + ".bodylimit.enabled"
	size    = giecho.ExtRoot + ".bodylimit.size"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable body limit middleware")
	giconfig.Add(size, "8M", "body limit size")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetSize() string {
	return giconfig.String(size)
}
