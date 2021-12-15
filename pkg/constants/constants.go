package constants

import (
	_ "embed"
)

//go:embed service-constants.json
var constantsConfig []byte
