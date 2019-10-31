package pkg

// HTTPError -
type HTTPError string

func (h HTTPError) String() string {
	return string(h)
}

const (
	// ProcessingError -
	ProcessingError HTTPError = "processing_error"
	// RequestError -
	RequestError HTTPError = "request_error"
)
