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

func TestLocationApiService_GetLocations(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx          context.Context
		platformName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Location
		want1   *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LocationApiService{
				client: tt.fields.client,
			}
			got, got1, err := a.GetLocations(tt.args.ctx, tt.args.platformName)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocationApiService.GetLocations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocationApiService.GetLocations() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LocationApiService.GetLocations() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}