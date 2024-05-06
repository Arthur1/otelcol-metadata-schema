package main

import (
	"encoding/json"
	"log"
	"os"

	otelcolmetadataschema "github.com/Arthur1/otelcol-metadata-schema"
	"github.com/invopop/jsonschema"
)

func main() {
	file := "metadata.schema.json"
	mdata := (*otelcolmetadataschema.Metadata)(nil)
	schema := jsonschema.Reflect(mdata)
	b, err := schema.MarshalJSON()
	if err != nil {
		log.Fatalln(err)
	}

	// workaround: *jsonschema.Schema does not have MarshalIndent()
	var a any
	if err := json.Unmarshal(b, &a); err != nil {
		log.Fatalln(err)
	}
	pb, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	pb = append(pb, byte('\n'))

	if err := os.WriteFile(file, pb, 0666); err != nil {
		log.Fatalln(err)
	}
}
