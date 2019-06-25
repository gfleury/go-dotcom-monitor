/*
 * dotcom-monitor
 *
 * dotcom-monitor API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package dotcommonitor

type Task struct {
	RequestType         string          `json:"RequestType,omitempty"`
	Url                 string          `json:"Url,omitempty"`
	Keyword1            string          `json:"Keyword1,omitempty"`
	FullPageDownload    bool            `json:"FullPageDownload,omitempty"`
	DownloadHtml        bool            `json:"Download_Html,omitempty"`
	DownloadFrames      bool            `json:"Download_Frames,omitempty"`
	DownloadStyleSheets bool            `json:"Download_StyleSheets,omitempty"`
	GetParams           []TaskGetParams `json:"GetParams,omitempty"`
	PostParams          []TaskGetParams `json:"PostParams,omitempty"`
	HeaderParams        []TaskGetParams `json:"HeaderParams,omitempty"`
	PrepareScript       string          `json:"PrepareScript,omitempty"`
	DeviceId            int32           `json:"Device_Id,omitempty"`
	TaskTypeId          int32           `json:"Task_Type_Id,omitempty"`
	Name                string          `json:"Name,omitempty"`
	Timeout             int32           `json:"Timeout,omitempty"`
}