// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule checks the pattern is valid
type AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule returns new rule with default attributes
func NewAwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule() *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule {
	return &AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule{
		resourceType:  "aws_globalaccelerator_endpoint_group",
		attributeName: "health_check_protocol",
		enum: []string{
			"TCP",
			"HTTP",
			"HTTPS",
		},
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Name() string {
	return "aws_globalaccelerator_endpoint_group_invalid_health_check_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as health_check_protocol`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}