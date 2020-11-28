package app

import "fmt"

// CertResponse represents a certificate response object
type CertResponse struct {
	IssuerCaID     int    `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	ID             int    `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
	SerialNumber   string `json:"serial_number"`
}

// PrintCertText prints the cert info in a text format
func (cr CertResponse) PrintCertText() {
	fmt.Printf("%s %s %s '%s' %s %s\n", cr.CommonName, cr.NameValue, cr.EntryTimestamp, cr.IssuerName, cr.NotAfter, cr.NotBefore)
}

// ToArray returns the cert information as an array
func (cr CertResponse) ToArray() []string {
	return []string{cr.CommonName, cr.NameValue, cr.EntryTimestamp, cr.IssuerName, cr.NotAfter, cr.NotBefore}
}
