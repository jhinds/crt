package app

// CertResponse
// Represents a certificate response object
type CertResponse struct {
	IssuerCaID        int
	IssuerName        string
	NameValue         string
	MinCertID         int
	MinEntryTimestamp string
	NotBefore         string
	NotAfter          string
}
