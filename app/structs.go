package main

// CertResponse
// Represents a certificate response object
type CertResponse struct {
	IssuerCaID        int    `json:"number"`
	IssuerName        string `json:"string"`
	NameValue         string `json:"string"`
	MinCertID         int    `json:"number"`
	MinEntryTimeStamp string `json:"string"`
	NotBefore         string `json:"string"`
	NotAfter          string `json:"string"`
}

type CertResponsesList struct {
	certs []CertResponse
}
