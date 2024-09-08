//go:build test

package swagger

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

// Replacing SwaggerFiles with a stub for tests
var SwaggerFiles = fstest.MapFS{
	"index.html":       &fstest.MapFile{Data: []byte("<html>Test Index</html>")},
	"api.swagger.json": &fstest.MapFile{Data: []byte(`{"swagger": "2.0", "info": {"title": "Test API"}}`)},
}

func TestSomething(t *testing.T) {
	// Try to read the index.html file from the SwaggerFiles stub
	indexFile, err := SwaggerFiles.Open("index.html")
	if err != nil {
		t.Fatalf("failed to open index.html file: %v", err)
	}

	// Check its contents
	indexData := make([]byte, 1024)
	n, err := indexFile.Read(indexData)
	if err != nil && err != fs.ErrClosed {
		t.Fatalf("error reading index.html file: %v", err)
	}

	expectedIndexContent := "<html>Test Index</html>"
	if string(indexData[:n]) != expectedIndexContent {
		t.Errorf("unexpected content of index.html file: %s, expected: %s", string(indexData[:n]), expectedIndexContent)
	}

	// Repeat the test with the api.swagger.json file
	swaggerFile, err := SwaggerFiles.Open("api.swagger.json")
	if err != nil {
		t.Fatalf("failed to open api.swagger.json file: %v", err)
	}

	swaggerData := make([]byte, 1024)
	n, err = swaggerFile.Read(swaggerData)
	if err != nil && err != fs.ErrClosed {
		t.Fatalf("error reading api.swagger.json file: %v", err)
	}

	expectedSwaggerContent := `{"swagger": "2.0", "info": {"title": "Test API"}}`
	if string(swaggerData[:n]) != expectedSwaggerContent {
		t.Errorf("unexpected content of api.swagger.json file: %s, expected: %s", string(swaggerData[:n]), expectedSwaggerContent)
	}
}
