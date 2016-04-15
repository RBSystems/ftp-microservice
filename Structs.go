package main

import "time"

type request struct {
	//Required fields.
	IPAddressHostname string //hostname to be sent the file
	Path              string //path to store the file on the server
	File              string //Location of the file to be sent - must be accessible from the server running the service
	CallbackAddress   string //complete address to send the notification that the file transfer is complete

	//Optional Fields
	Identifier string //Optional value to be passed in so the requester can identify the host when it's sent back.
	Timeout    int    //Time in seconds to wait for timeout when trying to open FTP connection. Defaults to 30.
	Username   string //Username and password to authenticate with the device. Defaults to anonymous/anonymous
	Password   string

	//Fields not expected in request, will be filled out by server
	SubmissionTime time.Time //Will be filled by the server as the time the file transfer began
	CompletionTime time.Time //Will be filled by the service as the time that a) the file transfer began or b) timed out
	Status         string    //Timeout or Success
}
