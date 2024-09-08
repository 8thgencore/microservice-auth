package swagger

import (
	"embed"
)

// SwaggerFiles contains the embedded Swagger UI and JSON files.
//
//go:embed index.html api.swagger.json
var SwaggerFiles embed.FS
