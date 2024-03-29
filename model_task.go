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
	Keyword1            string          `json:"Keyword1"`
	Keyword2            string          `json:"Keyword2"`
	Keyword3            string          `json:"Keyword3"`
	UserName            string          `json:"UserName"`
	UserPass            string          `json:"UserPass"`
	FullPageDownload    bool            `json:"FullPageDownload"`
	DownloadHtml        bool            `json:"Download_Html"`
	DownloadFrames      bool            `json:"Download_Frames"`
	DownloadStyleSheets bool            `json:"Download_StyleSheets"`
	DownloadScripts     bool            `json:"Download_Scripts"`
	DownloadImages      bool            `json:"Download_Images"`
	DownloadObjects     bool            `json:"Download_Objects"`
	DownloadApplets     bool            `json:"Download_Applets"`
	DownloadAdditional  bool            `json:"Download_Additional"`
	GetParams           []TaskGetParams `json:"GetParams"`
	PostParams          []TaskGetParams `json:"PostParams"`
	HeaderParams        []TaskGetParams `json:"HeaderParams"`
	PrepareScript       string          `json:"PrepareScript"`
	DNSResolveMode      string          `json:"DNSResolveMode"`
	DNSserverIP         *string         `json:"DNSserverIP"`
	DeviceId            int32           `json:"Device_Id,omitempty"`
	TaskTypeId          int32           `json:"Task_Type_Id,omitempty"`
	Name                string          `json:"Name,omitempty"`
	Timeout             int32           `json:"Timeout,omitempty"`
}
