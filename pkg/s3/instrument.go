package s3

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("github.com/arfan21/project-sprint-banking-api/pkg/s3")
