package constants

// Defines the maximum size for requests
const (
	MB = 1048576
	MaxRequestSize = 1 * MB
	MaxFileUploadSize = 16 * MB
)

// Timeout is the default request timeout
var Timeout int64 = 10

// ClientAuthorizationHeaderField is the header to use for token authorization
var ClientAuthorizationHeaderField = "Authorization"

// AccessControlAllowHeaders is the headers allowed by CORS
var AccessControlAllowHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
