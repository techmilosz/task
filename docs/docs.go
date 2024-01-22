package docs

import _ "embed"

//go:embed oapi.yaml
var OAPI []byte

//go:embed swagger.html
var Swagger []byte
