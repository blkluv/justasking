package operationresult

// OperationResult is an object for passing information about an operation between functions
type OperationResult struct {
	Status  string
	Message string
	Error   error
}

// Success is a constant used as a status
const Success string = "SUCCESS"

// Error is a constant used as a status
const Error string = "ERROR"

// NotFound is a constant used as a status
const NotFound string = "NOTFOUND"

// Unauthorized is a constant used as a status
const Unauthorized string = "UNAUTHORIZED"

// Conflict is a constant used as a status
const Conflict string = "CONFLICT"

// Forbidden is a constant used as a status
const Forbidden string = "FORBIDDEN"

// PaymentRequired is a constant used as a status
const PaymentRequired string = "PAYMENTREQUIRED"

// UnprocessableEntity is a contsant used as a status
const UnprocessableEntity string = "UNPROCESSABLE"

// Gone is a contsant used as a status
const Gone string = "GONE"

// New returns an initialized OperationResult
func New() *OperationResult {
	operationResult := new(OperationResult)
	operationResult.Status = Success
	operationResult.Message = ""
	return operationResult
}

// IsSuccess returns information on whether the OperationResult was a success or not
func (operationResult OperationResult) IsSuccess() bool {
	return operationResult.Status == Success
}

// CreateErrorResult creates an operationResult with Error as the status so that we don't use three lines in the domains
func CreateErrorResult(message string, err error) *OperationResult {
	return &OperationResult{Status: Error, Message: message, Error: err}
}
