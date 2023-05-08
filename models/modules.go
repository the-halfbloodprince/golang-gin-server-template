package models

// Module is the structure of a module
type Module struct {
	Name      string
	Prefix    string
	Endpoints []Endpoint
}
