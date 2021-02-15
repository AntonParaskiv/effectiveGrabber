package ApiRepository

import (
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/interafaces/HttpClientInterface"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/interafaces/mock_HttpClientInterface"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestRepository_getAny(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	httpClientMockSuccess := mock_HttpClientInterface.NewMockClient(ctrl)
	httpClientMockSuccess.
		EXPECT().
		Get(endpointOne).Return([]byte(`{ "result": 1 }`), http.StatusOK, nil)

	httpClientMockRequestFailed := mock_HttpClientInterface.NewMockClient(ctrl)
	httpClientMockRequestFailed.
		EXPECT().
		Get(endpointOne).Return(nil, http.StatusForbidden, nil)

	httpClientMockBadResponse := mock_HttpClientInterface.NewMockClient(ctrl)
	httpClientMockBadResponse.
		EXPECT().
		Get(endpointOne).Return([]byte(`{ "result": 0 }`), http.StatusOK, nil)

	type fields struct {
		httpClient HttpClientInterface.Client
	}
	type args struct {
		endpoint string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				httpClient: httpClientMockSuccess,
			},
			args: args{
				endpoint: endpointOne,
			},
			wantErr: false,
		},
		{
			name: "Request error",
			fields: fields{
				httpClient: httpClientMockRequestFailed,
			},
			args: args{
				endpoint: endpointOne,
			},
			wantErr: true,
		},
		{
			name: "Bad response",
			fields: fields{
				httpClient: httpClientMockBadResponse,
			},
			args: args{
				endpoint: endpointOne,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				httpClient: tt.fields.httpClient,
			}
			if err := r.getAny(tt.args.endpoint); (err != nil) != tt.wantErr {
				t.Errorf("getAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
