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

type trackingRequest struct {
	Tracking Tracking `json:"tracking"`
}

// GetTracking returns a specific tracking.
// From: https://docs.aftership.com/api/4/trackings/get-trackings-slug-tracking_number
func (c *Client) GetTracking(tracking Tracking) (Tracking, error) {
	data, err := c.doRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s/%s", TrackingsEndpoint, tracking.Slug, tracking.TrackingNumber),
		nil,
	)
	if err != nil {
		return Tracking{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if data == nil {
		return Tracking{}, nil
	}

	// Return the first tracking as that should be ours.
	return data.Tracking, nil
}

// PostTracking creates a new tracking.
// From: https://docs.aftership.com/api/4/trackings/get-trackings
func (c *Client) PostTracking(tracking Tracking) (Tracking, error) {
	data, err := c.doRequest(
		http.MethodPost,
		TrackingsEndpoint,
		trackingRequest{Tracking: tracking},
	)
	if err != nil {
		return Tracking{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if data == nil || len(data.Trackings) <= 0 {
		return Tracking{}, nil
	}

	// We should not receive more than one tracking.
	if len(data.Trackings) > 1 {
		return Tracking{}, fmt.Errorf("got %d results, only expected 1", len(data.Trackings))
	}

	// Return the first tracking as that should be ours.
	return data.Trackings[0], nil
}

// DeleteTracking deletes a specific tracking.
// From: https://docs.aftership.com/api/4/trackings/delete-trackings
func (c *Client) DeleteTracking(tracking Tracking) error {
	_, err := c.doRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/%s/%s", TrackingsEndpoint, tracking.Slug, tracking.TrackingNumber),
		nil,
	)

	return err
}
