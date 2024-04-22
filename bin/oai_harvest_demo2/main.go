package main

import (
	"fmt"
	"github.com/kiivihal/goharvest/oai"
)

// Dump a snippet of the Record metadata
func dump(record *oai.Record) {
	fmt.Printf("%s\n\n", record.Metadata.Body[0:500])
}

// Demonstrates harvesting using the ListRecords verb with HarvestRecords
func main() {
	req := &oai.Request{
		BaseUrl:        "http://services.kb.nl/mdo/oai",
		Set:            "DTS",
		MetadataPrefix: "dcx",
		Verb:           "ListRecords",
		From:           "2012-09-06T014:00:00.000Z",
	}
	// HarvestRecords passes each individual metadata record to the dump
	// function as a Record object
	req.HarvestRecords(dump)
}
