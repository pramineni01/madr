package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"

	vls "github.com/pramineni01/madr/vehicle_logs_service_mock"
)



// Server mock
type Server struct {
	srv *vls.UnimplementedLogServiceServer

	// config store holding key/value pairs
	// key= vin-API, value= ErrorCode
	cfgs map[string]uint32
}

// GetLogSettings mock api implementation
func (s *Server) GetLogSettings(context.Context, req *vls.GetLogSettingsRequest) (*vls.GetLogSettingsResponse, error) {
	return handleAPICall(req.GetVin(), req.GetAPI())
}

// UploadLogMessages mock api implementation
func (s *Server) UploadLogMessages(stream vls.LogService_UploadLogMessagesServer) error {
	// Read requests and send responses.
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return handleAPICall(req.GetVin(), req.GetAPI())
		}

		if err != nil {
			s:= status.New(c.InvalidArgument, "Invalid request")
			return nil, s.Err()
		}
	}
}

// GetAttachmentParameters mock api implementation
func (s *Server) GetAttachmentParameters(context.Context, *vls.GetAttachmentParametersRequest) (*vls.GetAttachmentParametersResponse, error) {
	return handleAPICall(req.GetVin(), req.GetAPI())
}

// SetMockConfig mock api implementation
func (s *Server) SetMockConfig(context.Context, req *vls.SetMockConfigRequest) (*vls.SetMockConfigResponse, error) {
	mcfgs := req.GetConfigs()
	
	for cfg := range mcfgs {
		var c codes.Code
		ecs := cfg.GetErrorCodeStr()
		if len(ecs) == 0 {
			ecs = strconv.Itoa(cfg.GetErrorCode())
		}

		if err := c.UnmarshalJSON([]byte(ecs)); err != nil {
			e := buildGrpcError(c.InvalidArgument, fmt.Sprintf( "Invalid code provided. Vin:%s. API: %s", cfg.GetVin(), cfg.GetAPI())
			return nil, e.Err()
		}

		s.cfgs[fmt.Sprintf ("%s-%s", cfg.GetVin(), cfg.GetAPI())] = c
	}

	if len(s.cfgs) == 0 {
		e := buildGrpcError(c.InvalidArgument, fmt.Sprintf( "Configurations missing"))
		return nil, e.Err()
	}

	return &GetLogSettingsResponse{}, nil
}

func (s *Server) buildGrpcError(c codes.Code, msg string) error {
	st := status.New(c, msg)
	ds, err := st.WithDetails(
		&epb.ErrorInfo{
			Reason: msg,
		},
	)
	return err != nil ? st.Err() : ds.Err()
}

func (s *Server) handleAPICall(vin, api string) (*proto.Message, error) {
	var e error
	if errorCode, ok := s.cfgs[fmt.Sprintf("%s-%s", vin, api]; ok == false {
		e := buildGrpcError(c.InvalidArgument, fmt.Sprintf( "Configuration missing"))
	} else {
		e := buildGrpcError(errorCode, "API: %s, VIN:%s, Error Code: %d", vin, api, errorCode)
	}
	return nil, e

}


		
