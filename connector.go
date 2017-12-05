package goco

// Connector connects an observed to an actual dom component
// There are connectors for each type of connection:
// value
// text
// html
// etc.

type Connector interface {
	Attach(selector string)
	Detach()
}

type ValueConnector interface {
	Set(s string)
	Get() string
}
