package ApiRepository

import (
	"fmt"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/interafaces/HttpClientInterface"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/interafaces/mock_HttpClientInterface"
	"github.com/golang/mock/gomock"
	"net/http"
	"reflect"
	"testing"
)

var SimulatedError = fmt.Errorf("simulated error")

func TestRepository_requestJsonGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	httpClientMockSuccess := mock_HttpClientInterface.NewMockClient(ctrl)
	httpClientMockSuccess.
		EXPECT().
		Get(endpointOne).Return([]byte(`{ "result": 1 }`), http.StatusOK, nil)

	httpClientMockClientError := mock_HttpClientInterface.NewMockClient(ctrl)
	httpClientMockClientError.
		EXPECT().
		Get(endpointOne).Return(nil, 0, SimulatedError)

	httpClientMockBadStatus := mock_HttpClientInterface.NewMockClient(ctrl)
	httpClientMockBadStatus.
		EXPECT().
		Get(endpointOne).Return(nil, http.StatusForbidden, nil)

	httpClientMockInvalidJson := mock_HttpClientInterface.NewMockClient(ctrl)
	httpClientMockInvalidJson.
		EXPECT().
		Get(endpointOne).Return([]byte(`{ invalid json }`), http.StatusOK, nil)

	type fields struct {
		httpClient HttpClientInterface.Client
	}
	type args struct {
		url            string
		response       interface{}
		wantStatusCode int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		wantResponse interface{}
	}{
		{
			name: "Success",
			fields: fields{
				httpClient: httpClientMockSuccess,
			},
			args: args{
				url:            endpointOne,
				response:       &Response{},
				wantStatusCode: http.StatusOK,
			},
			wantErr: false,
			wantResponse: &Response{
				Result: 1,
			},
		},
		{
			name: "Client error",
			fields: fields{
				httpClient: httpClientMockClientError,
			},
			args: args{
				url:            endpointOne,
				response:       &Response{},
				wantStatusCode: http.StatusOK,
			},
			wantErr:      true,
			wantResponse: &Response{},
		},
		{
			name: "Bad status code",
			fields: fields{
				httpClient: httpClientMockBadStatus,
			},
			args: args{
				url:            endpointOne,
				response:       &Response{},
				wantStatusCode: http.StatusOK,
			},
			wantErr:      true,
			wantResponse: &Response{},
		},
		{
			name: "Invalid json",
			fields: fields{
				httpClient: httpClientMockInvalidJson,
			},
			args: args{
				url:            endpointOne,
				response:       &Response{},
				wantStatusCode: http.StatusOK,
			},
			wantErr:      true,
			wantResponse: &Response{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				httpClient: tt.fields.httpClient,
			}
			if err := r.requestJsonGet(tt.args.url, tt.args.response, tt.args.wantStatusCode); (err != nil) != tt.wantErr {
				t.Errorf("requestJsonGet() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.response, tt.wantResponse) {
				t.Errorf("response = %v, want %v", tt.args.response, tt.wantResponse)
			}
		})
	}
}
