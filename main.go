package main

import (
	"Monitoring/prom-json-exporter-mod/jsonexporter"

	"github.com/kawamuray/prometheus-exporter-harness/harness"
)

func main() {

	opts := harness.NewExporterOpts("json_exporter", jsonexporter.Version)
	opts.Usage = "[OPTIONS] HTTP_ENDPOINT CONFIG_PATH"
	opts.Init = jsonexporter.Init
	harness.Main(opts)

}
