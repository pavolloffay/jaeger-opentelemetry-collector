package main

import (
	"log"

	"github.com/jaegertracing/jaeger-opentelemetry-collector/pkg/exporter/elasticsearch"

	"github.com/open-telemetry/opentelemetry-collector/defaults"
	"github.com/open-telemetry/opentelemetry-collector/service"
)

func main() {
	handleErr := func(err error) {
		if err != nil {
			log.Fatalf("Failed to run the service: %v", err)
		}
	}

	info := service.ApplicationStartInfo{
		ExeName:  "jaeger-opentelemetry-collector",
		LongName: "Jaeger OpenTelemetry Collector",
		// TODO
		//Version:  version.Version,
		//GitHash:  version.GitHash,
	}

	cmpts, err := defaults.Components()
	handleErr(err)

	es := elasticsearch.Factory{}
	cmpts.Exporters[es.Type()] = es

	svc, err := service.New(cmpts, info)
	handleErr(err)

	err = svc.Start()
	handleErr(err)
}