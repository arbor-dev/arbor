module github.com/arbor-dev/arbor

go 1.13

replace github.com/koding/websocketproxy => github.com/Nydauron/websocketproxy v0.0.0-20220622003505-9c247a147a15

require (
	github.com/gorilla/mux v0.0.0-20170922205414-3f19343c7d9c
	github.com/gorilla/websocket v1.5.0
	github.com/kennygrant/sanitize v1.2.3
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c
	github.com/syndtr/goleveldb v0.0.0-20170725064836-b89cc31ef797
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-20180719183105-8007e27cdb32
)

require (
	github.com/golang/snappy v0.0.0-20170215233205-553a64147049 // indirect
	github.com/gorilla/context v0.0.0-20160226214623-1ea25387ff6f // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
)
