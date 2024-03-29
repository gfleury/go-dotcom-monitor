/*
 * dotcom-monitor
 *
 * dotcom-monitor API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package dotcommonitor

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestLoginApiService_Login(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx        context.Context
		credential BasicAuth
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Result
		want1   int
		wantErr bool
	}{
		{"FailedLogin", fields{client: generateConfigFake()}, args{ctx: generateContextFakeAuth()}, Result{Success: false}, http.StatusForbidden, true},
		{"NoConnectivity", fields{client: generateConfigFake()}, args{ctx: generateContextCanceled()}, Result{Success: false}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LoginApiService{
				client: tt.fields.client,
			}
			got, got1, err := a.Login(tt.args.ctx, tt.args.credential)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginApiService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginApiService.Login() got = %v, want %v", got, tt.want)
			}
			if got1 != nil && !reflect.DeepEqual(got1.StatusCode, tt.want1) {
				t.Errorf("LoginApiService.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
