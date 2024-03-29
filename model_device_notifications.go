/*
 * dotcom-monitor
 *
 * dotcom-monitor API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package dotcommonitor

type DeviceNotifications struct {
	EMailFlag               bool                                    `json:"E_Mail_Flag"`
	EMailAddress            string                                  `json:"E_Mail_Address,omitempty"`
	EMailTimeIntervalMin    int32                                   `json:"E_Mail_TimeInterval_Min,omitempty"`
	WLDeviceFlag            bool                                    `json:"WL_Device_Flag"`
	WLDeviceEmailAddress    string                                  `json:"WL_Device_Email_Address,omitempty"`
	WLDeviceTimeIntervalMin int32                                   `json:"WL_Device_TimeInterval_Min,omitempty"`
	PagerFlag               bool                                    `json:"Pager_Flag"`
	PagerAreaCode           string                                  `json:"Pager_Area_Code,omitempty"`
	PagerPhone              string                                  `json:"Pager_Phone,omitempty"`
	PagerNumCode            string                                  `json:"Pager_Num_Code,omitempty"`
	PagerTimeIntervalMin    int32                                   `json:"Pager_TimeInterval_Min,omitempty"`
	PhoneFlag               bool                                    `json:"Phone_Flag"`
	PhoneAreaCode           string                                  `json:"Phone_Area_Code,omitempty"`
	PhonePhone              string                                  `json:"Phone_Phone,omitempty"`
	PhoneTimeIntervalMin    int32                                   `json:"Phone_TimeInterval_Min,omitempty"`
	SMSFlag                 bool                                    `json:"SMS_Flag"`
	SMSPhone                string                                  `json:"SMS_Phone,omitempty"`
	SMSTimeIntervalMin      int32                                   `json:"SMS_TimeInterval_Min,omitempty"`
	ScriptFlag              bool                                    `json:"Script_Flag"`
	ScriptBatchFileName     string                                  `json:"Script_Batch_File_Name,omitempty"`
	ScriptTimeIntervalMin   int32                                   `json:"Script_TimeInterval_Min,omitempty"`
	SNMPTimeIntervalMin     int32                                   `json:"SNMP_TimeInterval_Min,omitempty"`
	NotificationGroups      []DeviceNotificationsNotificationGroups `json:"Notification_Groups"`
}
