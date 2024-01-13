package techcast

import (
	_ "embed"
	"fmt"
	"strings"
)

var version string
var Version = fmt.Sprintf("v%s", strings.TrimSpace(version))

const Name = "techcast"
const NameUpper = "TECHCAST"