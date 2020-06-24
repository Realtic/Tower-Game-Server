package resapi

// Error ...
func (r *ResponseAPI) Error(message string) {
	var blob = new(ResponseBlob)
	blob.Status = statusError
	blob.Data = message

	r.writeBlob(blob)
}
