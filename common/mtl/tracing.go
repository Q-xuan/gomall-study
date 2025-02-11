package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// provider.WithExportEndpoint("127.0.0.1:4317"),
		// provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)
	return p
}
