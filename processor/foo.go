package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/mfamador/go-opentelemetry/servicev1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	err := service.RegisterProcessor("foo", fooConfigSpec, newFooProcessor)
	if err != nil {
		panic(err)
	}
}

var fooConfigSpec = service.NewConfigSpec().
	Summary("This processor adds the field `foo` to the message")

type fooProcessor struct {
	conf   *service.ParsedConfig
	logger *service.Logger
	client servicev1.ServiceClient
}

func newFooProcessor(conf *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {
	conn, _ := grpc.Dial(fmt.Sprintf("localhost:%d", 8181),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	client := servicev1.NewServiceClient(conn)

	return &fooProcessor{
		logger: mgr.Logger(),
		conf:   conf,
		client: client,
	}, nil
}

//------------------------------------------------------------------------------

func (m *fooProcessor) Process(ctx context.Context, msg *service.Message) (service.MessageBatch, error) {
	newMsg := msg.Copy()
	// call grpc service
	resp, _ := m.client.Ping(ctx, &servicev1.PingRequest{Message: "foo"})

	msgs, _ := newMsg.AsStructured()
	unboxed, ok := msgs.(map[string]any)
	if ok {
		unboxed["foo"] = resp.Message
		bytes, _ := json.Marshal(unboxed)
		newMsg.SetBytes(bytes)
	}

	return []*service.Message{newMsg}, nil
}

func (m *fooProcessor) Close(ctx context.Context) error {
	return nil
}
