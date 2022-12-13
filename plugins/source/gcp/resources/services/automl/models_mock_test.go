// Code generated by codegen; DO NOT EDIT.

package automl

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"testing"

	pb "cloud.google.com/go/automl/apiv1/automlpb"
)

func createModels(gsrv *grpc.Server) error {
	fakeServer := &fakeModelsServer{}
	pb.RegisterAutoMlServer(gsrv, fakeServer)
	return nil
}

type fakeModelsServer struct {
	pb.UnimplementedAutoMlServer
}

func (f *fakeModelsServer) ListModels(context.Context, *pb.ListModelsRequest) (*pb.ListModelsResponse, error) {
	resp := pb.ListModelsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestModels(t *testing.T) {
	client.MockTestGrpcHelper(t, Models(), createModels, client.TestOptions{})
}
