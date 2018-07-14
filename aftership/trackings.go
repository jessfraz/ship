package aftership

import (
	"net/http"
)

// GetTrackings returns the trackings.
// From: https://docs.aftership.com/api/4/trackings/get-trackings
func (c *Client) GetTrackings() ([]Tracking, error) {
	data, err := c.doRequest(http.MethodGet, TrackingsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	// Check if we didn't get a result and return an error if true.
	if data == nil || len(data.Trackings) <= 0 {
		return nil, ErrorEmptyResult
	}

	return data.Trackings, nil
}
