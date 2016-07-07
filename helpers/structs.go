package helpers

import "time"

// Request represents an incoming request body passed via a RESTful POST
type Request struct {
	// Required fields
	DestinationAddress   string `json:"destinationAddress"`   // Hostname to be sent the file
	DestinationDirectory string `json:"destinationDirectory"` // Path indicates where to store the file on the server
	FileLocation         string `json:"fileLocation"`         // The remote URL of the file to be downloaded and sent
	CallbackAddress      string `json:"callbackAddress"`      // Complete address to send the notification that the file transfer is complete

	// Optional Fields
	CallbackIdentifier string `json:"callbackIdentifier,omitempty"` // Optional value to be passed in so the requester can identify the host when it's sent back
	Timeout            int    `json:"timeout,omitempty"`            // Time in seconds to wait for timeout when trying to open an FTP connection (defaults to 30)
	UsernameFTP        string `json:"usernameFTP,omitempty"`        // Username to authenticate with the device (defaults to anonymous)
	PasswordFTP        string `json:"passwordFTP,omitempty"`        // Password to authenticate with the device (defaults to anonymous)

	// Fields not expected in request, will be filled by the service
	Filename       string
	SubmissionTime time.Time // Will be filled by the service to indicate when the file transfer began
	CompletionTime time.Time // Will be filled by the service to indicate when the file transfer ended or timed out
	Status         string    // "Timeout" or "Success"
	Error          string
}
