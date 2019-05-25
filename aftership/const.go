package aftership

import (
	"errors"
)

const (
	// APIUri holds the AfterShip API uri.
	APIUri = "https://api.aftership.com"
	// APIVersion holds the AfterShip API version.
	APIVersion = "v4"
	// APIKeyHeader is the header key for the API.
	APIKeyHeader = "aftership-api-key"

	// TrackingsEndpoint is the API endoint for trackings.
	TrackingsEndpoint = "/trackings"

	// CouriersEndpoint is the API endpoint for couriers.
	CouriersEndpoint = "/couriers"
	// CouriersAllEndpoint is the API endpoint for all couriers.
	CouriersAllEndpoint = "/couriers/all"
	// CouriersDetectEndpoint is the API endpoint for detecting couriers.
	CouriersDetectEndpoint = "/couriers/detect"

	// LastCheckpointEndpoint is the API endpoint for the last checkpoint of a tracking.
	LastCheckpointEndpoint = "/last_checkpoint"

	// NotificationsEndpoint is the API endpoint for interacting with notifications.
	NotificationsEndpoint = "/notifications"
)

var (
	// ErrorEmptyResult defines the error when the result is empty.
	ErrorEmptyResult = errors.New("empty result")
)
