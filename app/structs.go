package app

import "time"

// CertResponse represents a certificate response object
type CertResponse struct {
	IssuerCaID     int              `json:"issuer_ca_id"`
	IssuerName     string           `json:"issuer_name"`
	CommonName     string           `json:"common_name"`
	NameValue      string           `json:"name_value"`
	ID             int              `json:"id"`
	EntryTimestamp ISO8601LocalTime `json:"entry_timestamp"`
	NotBefore      ISO8601LocalTime `json:"not_before"`
	NotAfter       ISO8601LocalTime `json:"not_after"`
	SerialNumber   string           `json:"serial_number"`
}

// ToArray returns the cert information as an array
func (cr CertResponse) ToArray() []string {
	return []string{cr.CommonName, cr.NameValue, ISO8601LocalTime.String(cr.EntryTimestamp), cr.IssuerName, ISO8601LocalTime.String(cr.NotAfter), ISO8601LocalTime.String(cr.NotBefore)}
}

// ISO8601LocalTime struct for local time in ISO8601 format
type ISO8601LocalTime struct {
	time.Time
}

// UnmarshalJSON unmarshal datetime in ISO8601 format
func (lt *ISO8601LocalTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1]

	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {

		t, err = time.Parse("2006-01-02T15:04:05", s)
	}
	lt.Time = t
	return
}
