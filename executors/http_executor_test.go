package executors

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestHttpVal_DoExecute(t *testing.T) {
	type args struct {
		requestBody interface{}
	}
	tests := []struct {
		name       string
		fields     HttpVal
		args       args
		want       interface{}
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "TestShouldExecuteHTTPStep",
			fields: HttpVal{
				Method:  "GET",
				Url:     "http://54.190.25.178:3333/api/v1/user",
				Headers: "",
			},
			args: args{
				requestBody: map[string]interface{}{"k": "v"},
			},
			want:       map[string]interface{}{"id": "1234", "name": "ABC", "email": "abc@sahaj.com", "org": "sahaj"},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "TestShouldExecuteHTTPStepWithHeaders",
			fields: HttpVal{
				Method:  "GET",
				Url:     "http://54.190.25.178:3333/api/v1/user",
				Headers: "Content-Type:application/json;token:abc",
			},
			args: args{
				requestBody: map[string]interface{}{"k": "v"},
			},
			want:       map[string]interface{}{"id": "1234", "name": "ABC", "email": "abc@sahaj.com", "org": "sahaj"},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "TestShouldThrowErrorWhileExecutingStep",
			fields: HttpVal{
				Method:  "GET",
				Url:     "http://54.190.25.178:3333/api/v1/asd",
				Headers: "",
			},
			args: args{
				requestBody: map[string]interface{}{"k": "v"},
			},
			want:       nil,
			wantErr:    true,
			wantErrMsg: "404 page not found",
		},
		//{
		//	name: "TestShouldThrowErrorForHTTPStep",
		//	fields: HttpVal{
		//		Method:  "GET",
		//		Url:     "http://localhost:3333/api/v1/user",
		//		Headers: "",
		//	},
		//	args: args{
		//		requestBody: map[string]interface{}{"k": "v"},
		//	},
		//	want:       map[string]interface{}{"id": "1234", "name": "ABC", "email": "abc@sahaj.com", "org": "sahaj"},
		//	wantErr:    true,
		//	wantErrMsg: "Get http://localhost:3333/api/v1/user: dial tcp [::1]:3333: connect: connection refused",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpVal := HttpVal{
				Method:  tt.fields.Method,
				Url:     tt.fields.Url,
				Headers: tt.fields.Headers,
			}
			got, err := httpVal.DoExecute(tt.args.requestBody, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("DoExecute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
			} else {
				var responsePayload map[string]interface{}
				json.Unmarshal([]byte(got.(string)), &responsePayload)
				if !reflect.DeepEqual(responsePayload, tt.want) {
					t.Errorf("DoExecute() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
