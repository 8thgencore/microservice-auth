package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"

	accessAPI "github.com/8thgencore/microservice-auth/internal/delivery/access"
	"github.com/8thgencore/microservice-auth/internal/service"
	serviceMocks "github.com/8thgencore/microservice-auth/internal/service/mocks"
	accessv1 "github.com/8thgencore/microservice-auth/pkg/pb/access/v1"
)

func TestCheck(t *testing.T) {
	t.Parallel()

	type accessServiceMockFunc func(mc *minimock.Controller) service.AccessService

	type args struct {
		ctx context.Context
		req *accessv1.CheckRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		endpoint = "/chat_v1.ChatV1/Create"

		serviceErr = fmt.Errorf("service error")

		req = &accessv1.CheckRequest{
			Endpoint: endpoint,
		}

		res = &empty.Empty{}
	)

	tests := []struct {
		name              string
		args              args
		want              *empty.Empty
		err               error
		accessServiceMock accessServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			accessServiceMock: func(mc *minimock.Controller) service.AccessService {
				mock := serviceMocks.NewAccessServiceMock(mc)
				mock.CheckMock.Expect(minimock.AnyContext, endpoint).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			accessServiceMock: func(mc *minimock.Controller) service.AccessService {
				mock := serviceMocks.NewAccessServiceMock(mc)
				mock.CheckMock.Expect(minimock.AnyContext, endpoint).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			accessServiceMock := tt.accessServiceMock(mc)
			api := accessAPI.NewImplementation(accessServiceMock)

			res, err := api.Check(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
