package resapi

// Print ...
func (r *ResponseAPI) Print(message string) {
	var blob = new(ResponseBlob)
	blob.Status = statusOk
	blob.Data = message

	r.writeBlob(blob)
}
