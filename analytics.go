package saviynt

import (
	"fmt"
	"net/http"

	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/grokify/mogo/net/urlutil"
)

const (
	AnalyticsLimitDefault = 50
	AnalyticsLimitMax     = 500
)

// FetchRuntimeControlsDataV2 returns data from the `fetchRuntimeControlsDataV2` API endpoint.
// `analyticsName` is required. `requestor`, and `analyticsID` are optional.
func (c Client) FetchRuntimeControlsDataV2(analyticsName, requestor, analyticsID string, attrs map[string]any, limit, offset uint) (*httpsimple.Request, *http.Response, error) {
	// func (c Client) GetAuditLogRuntimeControlsData(name string, minutes, limit, offset uint) (*http.Response, error) {
	if limit == 0 {
		limit = AnalyticsLimitDefault
	}
	// attrs = map[string]any{"timeFrame": "3000"}
	apiURL := urlutil.JoinAbsolute(c.BaseURL, RelURLECM, RelURLAPI, RelURLLoginRuntimeControlsData)
	sreq := httpsimple.Request{
		Method:   http.MethodPost,
		URL:      apiURL,
		BodyType: httpsimple.BodyTypeJSON,
		Body: AnalyticsRequest{
			AnalyticsName: analyticsName,
			Requestor:     requestor,
			AnalyticsID:   analyticsID,
			Attributes:    attrs,
			Max:           fmt.Sprintf("%d", limit),
			Offset:        fmt.Sprintf("%d", offset),
		},
	}
	sclient := httpsimple.Client{
		BaseURL:    c.BaseURL,
		HTTPClient: c.HTTPClient}
	resp, err := sclient.Do(sreq)
	return &sreq, resp, err
}

type AnalyticsRequest struct {
	AnalyticsName string `json:"analyticsname"`
	AnalyticsID   string `json:"analyticsid"`
	Requestor     string `json:"requestor"`
	Attributes    any    `json:"attributes,omitempty"`
	Max           string `json:"max,omitempty"`
	Offset        string `json:"offset,omitempty"`
}

type AnalyticsRequestAttributes struct {
	TimeFrame string `json:"timeFrame"`
}

type AnalyticsResponse struct {
	DisplayCount int                 `json:"displaycount"`
	Msg          string              `json:"msg"`
	TotalCount   int                 `json:"totalcount"`
	Results      []map[string]string `json:"results"`
}

type UserLoginAccessDetail struct {
	Data       string `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	ObjectName string `json:"objectName,omitempty"`
}
