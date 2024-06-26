// Data structure for the OAI-PMH protocol responses:
package oai

type Header struct {
	Status     string   `xml:"status,attr"`
	Identifier string   `xml:"identifier"`
	DateStamp  string   `xml:"datestamp"`
	SetSpec    []string `xml:"setSpec"`
}

type Metadata struct {
	Body []byte `xml:",innerxml"`
}

// Formatter for Metadata content
func (md Metadata) GoString() string { return string(md.Body) }

type About struct {
	Body []byte `xml:",innerxml"`
}

// Formatter for About content
func (ab About) GoString() string { return string(ab.Body) }

type Record struct {
	Header   Header   `xml:"header"`
	Metadata Metadata `xml:"metadata"`
	About    About    `xml:"about"`
	Raw      []byte   `xml:",innerxml"`
}

func (r *Record) GoString() string {
	return "<record>\n" + string(r.Raw) + "</record>\n"
}

type ListIdentifiers struct {
	Headers         []Header        `xml:"header"`
	ResumptionToken ResumptionToken `xml:"resumptionToken"`
}

type ResumptionToken struct {
	Token            string `xml:",chardata"`
	CompleteListSize int    `xml:"completeListSize,attr"`
}

type ListRecords struct {
	Records         []Record        `xml:"record"`
	ResumptionToken ResumptionToken `xml:"resumptionToken"`
}

type GetRecord struct {
	Record Record `xml:"record"`
}

type RequestNode struct {
	Verb           string `xml:"verb,attr"`
	Set            string `xml:"set,attr"`
	MetadataPrefix string `xml:"metadataPrefix,attr"`
}

type OAIError struct {
	Code    string `xml:"code,attr"`
	Message string `xml:",chardata"`
}

type MetadataFormat struct {
	MetadataPrefix    string `xml:"metadataPrefix"`
	Schema            string `xml:"schema"`
	MetadataNamespace string `xml:"metadataNamespace"`
}

type ListMetadataFormats struct {
	MetadataFormat []MetadataFormat `xml:"metadataFormat"`
}

type Description struct {
	Body []byte `xml:",innerxml"`
}

// Formatter for Description content
func (desc Description) GoString() string { return string(desc.Body) }

type Set struct {
	SetSpec        string      `xml:"setSpec"`
	SetName        string      `xml:"setName"`
	SetDescription Description `xml:"setDescription"`
}

type ListSets struct {
	Set []Set `xml:"set"`
}

type Identify struct {
	RepositoryName    string        `xml:"repositoryName"`
	BaseURL           string        `xml:"baseURL"`
	ProtocolVersion   string        `xml:"protocolVersion"`
	AdminEmail        []string      `xml:"adminEmail"`
	EarliestDatestamp string        `xml:"earliestDatestamp"`
	DeletedRecord     string        `xml:"deletedRecord"`
	Granularity       string        `xml:"granularity"`
	Description       []Description `xml:"description"`
}

// The struct representation of an OAI-PMH XML response
type Response struct {
	ResponseDate string      `xml:"responseDate"`
	Request      RequestNode `xml:"request"`
	Error        OAIError    `xml:"error"`

	Identify            Identify            `xml:"Identify"`
	ListMetadataFormats ListMetadataFormats `xml:"ListMetadataFormats"`
	ListSets            ListSets            `xml:"ListSets"`
	GetRecord           GetRecord           `xml:"GetRecord"`
	ListIdentifiers     ListIdentifiers     `xml:"ListIdentifiers"`
	ListRecords         ListRecords         `xml:"ListRecords"`
}
