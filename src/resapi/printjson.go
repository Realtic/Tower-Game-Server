package resapi

// PrintJSON ...
func (r *ResponseAPI) PrintJSON(message interface{}) {
	var blob = new(ResponseBlob)
	blob.Status = statusOk
	blob.Data = message

	r.writeBlob(blob)
}
