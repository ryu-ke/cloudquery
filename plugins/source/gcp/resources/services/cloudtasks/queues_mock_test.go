// Code generated by codegen; DO NOT EDIT.

package cloudtasks

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"testing"

	pb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
)

func createQueues(gsrv *grpc.Server) error {
	fakeServer := &fakeQueuesServer{}
	pb.RegisterCloudTasksServer(gsrv, fakeServer)
	return nil
}

type fakeQueuesServer struct {
	pb.UnimplementedCloudTasksServer
}

func (f *fakeQueuesServer) ListQueues(context.Context, *pb.ListQueuesRequest) (*pb.ListQueuesResponse, error) {
	resp := pb.ListQueuesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestQueues(t *testing.T) {
	client.MockTestGrpcHelper(t, Queues(), createQueues, client.TestOptions{})
}
