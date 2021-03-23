package activecampaign

import (
	go_types "github.com/leapforce-libraries/go_types"
)

type Meta struct {
	Total     go_types.Int64String `json:"total"`
	Sortable  bool                 `json:"sortable"`
	PageInput struct {
		SegmentID  int64            `json:"segmentid"`
		FormID     int64            `json:"formid"`
		ListID     int64            `json:"listid"`
		TagID      int64            `json:"tagid"`
		Limit      int64            `json:"limit"`
		Offset     int64            `json:"offset"`
		Search     *string          `json:"search"`
		Sort       *string          `json:"sort"`
		SeriesID   int64            `json:"seriesid"`
		WaitID     int64            `json:"waitid"`
		Status     int64            `json:"status"`
		ForceQuery go_types.BoolInt `json:"forceQuery"`
		CacheID    *string          `json:"cacheid"`
	} `json:"page_input"`
}
