package constants

// Defines the maximum size for requests
const (
	MB = 1048576
	MaxRequestSize = 1 * MB
	MaxFileUploadSize = 16 * MB
)

// Timeout is the default request timeout
var Timeout int64 = 10

// Client Authorization Token Field
var ClientAuthorizationHeaderField = "Authorization"

// The headers allowed by CORS
var AccessControlAllowHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
