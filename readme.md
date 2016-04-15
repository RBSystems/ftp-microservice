# FTP Microservice

A microservice to send files to clients over FTP asyncronously. Specifically designed for Crestron systems.

## Endpoints

* `POST /send` Send a file.

### Send
  This endpoint will send a file based on the parameters in the JSON Payload. Required fields are in the form of

  ```
{
	"IPAddressHostname": "string",
	"CallbackAddress":"",
	"Path": "string",
	"File": "./test.txt"
}
```

* IPAddressHostname: IPAddress or hostname of the device.
* CallbackAddress: Complete address of the server to contact with POST request when complete.
* Path: The **directory** on the device, relative to root, to store the file transferred.
* File: The path to the file to send. Must be accessible from the server running the service. Must be readable by the service.

There are additional, optional fields that can be passed in.

```
{
  "Identifier": "string",
  "Timeout": int,
  "Username": "string",
  "Password": "string"
}
```
* Identifier: Value that will be passed to the callback address to aid in identification of the entity completed.
* Timeout: Amount of time, in seconds, to wait for the device to respond to a request to open an FTP connection. Defaults to 30 if no value or 0 passed in.
* Username: Username to use in authentication. Defaults to anonymous.
* Password: Password to use in authentication. Defaults to anonymous.

Thus the full payload with all optional and required fields will appear in the form of

```
{
	"IPAddressHostname": "string",
	"CallbackAddress":"",
	"Path": "string",
	"File": "./test.txt",
  "Identifier": "string",
  "Timeout": int,
  "Username": "string",
  "Password": "string"
}
```

## Response

Upon completion or error the service will send a post request to the value passed in as the `Callback Address` field.

The response will come in the form of

```
{
  "IPAddressHostname": "string",
  "Path": "string",
  "File": "./test.txt",
  "CallbackAddress": "string",
  "Identifier": "string",
  "Timeout": int,
  "Username": "string",
  "Password": "string",
  "SubmissionTime": dateTime (RFC 3339),
  "CompletionTime": dateTime (RFC 3339),
  "Status": "string",
  "Error": "string"
}
```

Most of the values are the same as were passed in, however there are some additional values.

* SubmissionTime: The time the ftp transfer was submitted to the service
* CompletionTime: The time the ftp transfer was completed or the time an error occurred.
* Status: Success or error, depending on the result of the operation.
* Error: Will be empty on success, on error will contain error information.
