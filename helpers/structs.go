package helpers

import "time"

// Request represents an incoming request body passed via a RESTful POST
type Request struct {
	// Required fields
	IPAddressHostname string // Hostname to be sent the file
	Path              string // Path indicates where to store the file on the server
	File              string // Local location of the file to be sent (must be accessible from the server running the service)
	CallbackAddress   string // Complete address to send the notification that the file transfer is complete

	// Optional Fields
	Identifier string // Optional value to be passed in so the requester can identify the host when it's sent back
	Timeout    int    // Time in seconds to wait for timeout when trying to open an FTP connection (defaults to 30)
	Username   string // Username to authenticate with the device (defaults to anonymous)
	Password   string // Password to authenticate with the device (defaults to anonymous)

	// Fields not expected in request, will be filled by the service
	SubmissionTime time.Time // Will be filled by the service to indicate when the file transfer began
	CompletionTime time.Time // Will be filled by the service to indicate when the file transfer ended or timed out
	Status         string    // "Timeout" or "Success"
	Error          string
}
