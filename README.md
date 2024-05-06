# otelcol-metadata-schema

![GitHub License](https://img.shields.io/github/license/Arthur1/otelcol-metadata-schema)
![go-version](https://img.shields.io/github/go-mod/go-version/Arthur1/otelcol-metadata-schema?filename=go.mod)

JSON Schema for OpenTelemetry Collector metadata.yaml and its generator

![The description of `aggregation_temporality` is shown on the VSCode.](docs/img/screenshot.png)

## How to Use

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/Arthur1/otelcol-metadata-schema/main/metadata.schema.json

type: sample
scope_name: go.opentelemetry.io/collector/internal/receiver/samplereceiver

...
```

## Generate JSON Schema File

```console
$ pwd
/path/to/otelcol-metadata-schema
$ go generate .
```

## License

### Source Code

MIT License

### Schema File

CC BY 4.0
