package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"hello/api"
)

var (
	// MarshalOptions is a configurable JSON format marshaller.
	MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	// UnmarshalOptions is a configurable JSON format parser.
	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

func marshal(v interface{}) ([]byte, error) {
	switch m := v.(type) {
	case json.Marshaler:
		return m.MarshalJSON()
	case proto.Message:
		return MarshalOptions.Marshal(m)
	default:
		return json.Marshal(m)
	}
}
func DefEncodeResponseFunc() func(http.ResponseWriter, *http.Request, interface{}) error {
	return func(writer http.ResponseWriter, request *http.Request, v interface{}) error {
		byts, _ := marshal(v)
		tmpStruct := &structpb.Struct{}
		_ = protojson.Unmarshal(byts, tmpStruct)

		response := api.Response{
			Code:    int32(0),
			Reason:  "SUCCESS",
			Message: "success",
			Data:    tmpStruct,
		}

		respp, err := marshal(&response)
		writer.WriteHeader(int(200))
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(respp)

		return err
	}
}

func DefEncodeErrorFunc() func(http.ResponseWriter, *http.Request, error) {
	return func(writer http.ResponseWriter, request *http.Request, err error) {
		byts := []byte("{}")
		tmpStruct := &structpb.Struct{}
		_ = protojson.Unmarshal(byts, tmpStruct)

		response := api.Response{
			Code:    1,
			Reason:  "UNKNOWN ERROR",
			Message: "unknown error",
			Data:    tmpStruct,
		}

		if se := new(errors.Error); errors.As(err, &se) {
			code, _ := MappingReasonCode[se.Reason]
			response.Code = code
			response.Reason = se.Reason
			response.Message = se.Message
		}

		respp, err := marshal(&response)
		writer.WriteHeader(int(200))
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(respp)

		return
	}
}
