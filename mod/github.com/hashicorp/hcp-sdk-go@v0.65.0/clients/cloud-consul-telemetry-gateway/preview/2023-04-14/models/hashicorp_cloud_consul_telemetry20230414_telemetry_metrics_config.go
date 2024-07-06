// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulTelemetry20230414TelemetryMetricsConfig hashicorp cloud consul telemetry 20230414 telemetry metrics config
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.TelemetryMetricsConfig
type HashicorpCloudConsulTelemetry20230414TelemetryMetricsConfig struct {

	// Disabled is a bool used to indicate whether metrics pushing from the server or proxy disabled
	Disabled bool `json:"disabled,omitempty"`

	// Endpoint is the URL to use when exporting metrics via OTLP. If set will override the global endpoint in TelemetryConfig
	Endpoint string `json:"endpoint,omitempty"`

	// IncludeList is a list of regular expressions used to configure the metrics pipeline filter
	IncludeList []string `json:"include_list"`
}

// Validate validates this hashicorp cloud consul telemetry 20230414 telemetry metrics config
func (m *HashicorpCloudConsulTelemetry20230414TelemetryMetricsConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud consul telemetry 20230414 telemetry metrics config based on context it is used
func (m *HashicorpCloudConsulTelemetry20230414TelemetryMetricsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414TelemetryMetricsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414TelemetryMetricsConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulTelemetry20230414TelemetryMetricsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}