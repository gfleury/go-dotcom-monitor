/*
 * dotcom-monitor
 *
 * dotcom-monitor API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package dotcommonitor

type Platform struct {
	Id        int32              `json:"Id,omitempty"`
	Name      string             `json:"Name,omitempty"`
	Available bool               `json:"Available,omitempty"`
	Packages  []PlatformPackages `json:"Packages,omitempty"`
}
