package main

import (
	"context"
	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"github.com/benthosdev/benthos/v4/public/service"
	_ "github.com/mfamador/go-opentelemetry/processor"
)

func main() {
	service.RunCLI(context.Background())
}
