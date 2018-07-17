package aftership

import (
	"strings"
	"time"
)

// Response defines the data struct for an API response object.
// From: https://docs.aftership.com/api/4/overview
type Response struct {
	Meta `json:"meta,omitempty"`
	Data Data `json:"data,omitempty"`
}

// Data holds the data in a response object,
// From: https://docs.aftership.com/api/4/overview
type Data struct {
	Checkpoint   CheckPoint   `json:"checkpoint,omitempty"`
	Couriers     []Courier    `json:"couriers,omitempty"`
	Notification Notification `json:"notification,omitempty"`
	Trackings    []Tracking   `json:"trackings,omitempty"`
	Tracking     Tracking     `json:"tracking,omitempty"`
}

// Meta defines the data struct for the metadata.
// From: https://docs.aftership.com/api/4/overview
type Meta struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

// Tracking define the data struct for a tracking object.
// From: https://docs.aftership.com/api/4/trackings/
type Tracking struct {
	ID             string `json:"id,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`

	Active bool `json:"active,omitempty"`

	Slug  string `json:"slug,omitempty"`
	Title string `json:"title,omitempty"`
	Tag   string `json:"tag,omitempty"`

	PhoneNumbers []string `json:"smses,omitempty"`
	Emails       []string `json:"emails,omitempty"`

	OrderID     string `json:"order_id,omitempty"`
	OrderIDPath string `json:"order_id_path,omitempty"`

	CustomFields map[string]string `json:"custom_fields,omitempty"`

	Language string `json:"language,omitempty"`

	CreatedAt TimeJSON `json:"created_at,omitempty"`
	UpdatedAt TimeJSON `json:"updated_at,omitempty"`

	TrackingAccountNumber string `json:"tracking_account_number,omitempty"`
	TrackingPostalCode    string `json:"tracking_postal_code,omitempty"`
	TrackingShipDate      string `json:"tracking_ship_date,omitempty"`

	ExpectedDelivery string `json:"expected_delivery,omitempty"`

	Note string `json:"note,omitempty"`

	OriginCountryISO string `json:"origin_country_iso3,omitempty"`

	ShipmentPackageCount int    `json:"shipment_package_count,omitempty"`
	ShipmentType         string `json:"shipment_type,omitempty"`

	SignedBy string `json:"signed_by,omitempty"`

	Source string `json:"source,omitempty"`

	TrackedCount int `json:"tracked_count,omitempty"`

	UniqueToken string `json:"unique_token,omitempty"`

	Checkpoints []CheckPoint `json:"checkpoints,omitempty"`
}

// CheckPoint defines the data struct for a checkpoint object.
type CheckPoint struct {
	Slug    string `json:"slug,omitempty"`
	Tag     string `json:"tag,omitempty"`
	Message string `json:"message,omitempty"`

	CreatedAt      TimeJSON `json:"created_at,omitempty"`
	CheckPointTime TimeJSON `json:"checkpoint_time,omitempty"`

	City        string   `json:"city,omitempty"`
	State       string   `json:"state,omitempty"`
	Zip         string   `json:"zip,omitempty"`
	CountryName string   `json:"country_name,omitempty"`
	CountryISO  string   `json:"country_iso3,omitempty"`
	Coordinates []string `json:"coordinates,omitempty"`
	Location    string   `json:"location,omitempty"`
}

// Notification defines the data struct for a notification object.
// From: https://docs.aftership.com/api/4/notifications/
type Notification struct {
	AndroidPhones []string `json:"android,omitempty"`
	PhoneNumbers  []string `json:"smses,omitempty"`
	Emails        []string `json:"emails,omitempty"`
	ApplePhones   []string `json:"ios,omitempty"`
}

// Courier defines the data struct for a courier object.
// From: https://docs.aftership.com/api/4/couriers/
type Courier struct {
	Name      string `json:"name,omitempty"`
	OtherName string `json:"other_name,omitempty"`
	Slug      string `json:"slug,omitempty"`

	Phone string `json:"phone,omitempty"`
	URL   string `json:"web_url,omitempty"`

	RequiredFields []string `json:"required_fields,omitempty"`
	OptionalFields []string `json:"optional_fields,omitempty"`
}

// TimeJSON is the time format returned by the AfterShip API.
type TimeJSON struct {
	time.Time
}

// UnmarshalJSON sets the TimeJSON correctly from a string.
func (t *TimeJSON) UnmarshalJSON(b []byte) error {
	s := strings.Trim(strings.TrimSpace(string(b)), `"`)

	// The time format is not universal for all responses... oye.
	format := "2006-01-02T15:04:05-07:00"
	if !strings.Contains(s, "+") && strings.Count(s, "-") < 3 {
		if strings.HasSuffix(s, "Z") {
			format = "2006-01-02T15:04:05Z"
		} else {
			format = "2006-01-02T15:04:05"
		}
	}

	i, err := time.Parse(format, s)
	if err != nil {
		return err
	}

	*t = TimeJSON{i}

	return nil
}
