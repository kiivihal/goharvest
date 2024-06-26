// Package oai Data structure for the OAI-PMH protocol request:
package oai

import (
	"encoding/xml"
	"os"
)

// FromFile reads OAI PMH response XML from a file
func FromFile(filename string) (oaiResponse *Response) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Unmarshall all the data
	err = xml.Unmarshal(bytes, &oaiResponse)
	if err != nil {
		panic(err)
	}

	return
}
