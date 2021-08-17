package api

type MsbContext interface {
	GetTracingContext() (string, string)
}
