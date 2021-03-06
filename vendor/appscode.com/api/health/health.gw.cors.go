// Code generated by protoc-gen-grpc-gateway-cors
// source: health.proto
// DO NOT EDIT!

/*
Package health is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package health

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportHealthCorsPatterns returns an array of grpc gatway mux patterns for
// Health service to enable CORS.
func ExportHealthCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Health_Status_0,
	}
}
