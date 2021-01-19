package partialresponse

// PartialResponse is a response that only contains the value and wether or not it's displayed
type PartialResponse struct {
	Response string
	Hidden   bool
}

// New returns an initialized OperationResult
func New() *PartialResponse {
	partialResponse := new(PartialResponse)
	partialResponse.Response = ""
	partialResponse.Hidden = false
	return partialResponse
}
