package cassette

import (
	"testing"
)

func BenchmarkLoadingYAML(b *testing.B) {
	_, err := Load(`testdata/llama`)
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkLoadingJSON(b *testing.B) {
	_, err := LoadJSON(`testdata/llama`)
	if err != nil {
		b.Fatal(err)
	}
}
