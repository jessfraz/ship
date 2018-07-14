package aftership

import (
	"fmt"
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

// PostTracking creates a new tracking.
// From: https://docs.aftership.com/api/4/trackings/get-trackings
func (c *Client) PostTracking(tracking Tracking) (Tracking, error) {
	data, err := c.doRequest(http.MethodPost, TrackingsEndpoint, tracking)
	if err != nil {
		return Tracking{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if data == nil || len(data.Trackings) <= 0 {
		return Tracking{}, ErrorEmptyResult
	}

	// We should not receive more than one tracking.
	if len(data.Trackings) > 1 {
		return Tracking{}, fmt.Errorf("got %d results, only expected 1", len(data.Trackings))
	}

	// Return the first tracking as that should be ours.
	return data.Trackings[0], nil
}
