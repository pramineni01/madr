module server

go 1.14

// replace github.com/pramineni01/go-workspace/src/madr/vehicle_logs_service_mock => ../vehicle_logs_service_mock
// replace /usr/local/go/src/vehicle_logs_service_mock => ./vehicle_logs_service_mock

require (
	// github.com/pramineni01/go-workspace/src/madr/vehicle_logs_service_mock v1.0.0 // indirect
	google.golang.org/genproto v0.0.0-20200420144010-e5e8543f8aeb
	google.golang.org/grpc v1.29.0
)
