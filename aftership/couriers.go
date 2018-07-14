package aftership

import (
	"fmt"
	"net/http"
)

// DetectCourier returns a matched courier based on tracking number.
// From: https://docs.aftership.com/api/4/couriers/post-couriers-detect
func (c *Client) DetectCourier(tracking Tracking) (Courier, error) {
	data, err := c.doRequest(
		http.MethodPost,
		CouriersDetectEndpoint,
		trackingRequest{Tracking: tracking},
	)
	if err != nil {
		return Courier{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if data == nil || len(data.Couriers) <= 0 {
		return Courier{}, nil
	}

	// We should not receive more than one courier.
	if len(data.Couriers) > 1 {
		return Courier{}, fmt.Errorf("got %d results, only expected 1", len(data.Couriers))
	}

	// Return the first courier as that should be ours.
	return data.Couriers[0], nil
}
