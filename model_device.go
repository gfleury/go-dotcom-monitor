/*
 * dotcom-monitor
 *
 * dotcom-monitor API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package dotcommonitor

type Device struct {
	AvoidSimultaneousChecks bool                 `json:"Avoid_Simultaneous_Checks,omitempty"`
	AlertSilenceMin         int32                `json:"Alert_Silence_Min,omitempty"`
	FalsePositiveCheck      bool                 `json:"False_Positive_Check,omitempty"`
	Locations               []int32              `json:"Locations,omitempty"`
	SendUptimeAlert         bool                 `json:"Send_Uptime_Alert,omitempty"`
	StatusDescription       string               `json:"Status_Description,omitempty"`
	Postpone                bool                 `json:"Postpone"`
	OwnerDeviceId           int32                `json:"Owner_Device_Id,omitempty"`
	Frequency               int32                `json:"Frequency,omitempty"`
	FilterId                int32                `json:"Filter_Id,omitempty"`
	SchedulerId             int32                `json:"Scheduler_Id,omitempty"`
	NumberOfTasks           int32                `json:"Number_Of_Tasks,omitempty"`
	Id                      int32                `json:"Id,omitempty"`
	PlatformId              int32                `json:"Platform_Id,omitempty"`
	PackageId               int32                `json:"Package_Id,omitempty"`
	Name                    string               `json:"Name,omitempty"`
	Notifications           *DeviceNotifications `json:"Notifications,omitempty"`
}
