// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferServerInvalidURLRule checks the pattern is valid
type AwsTransferServerInvalidURLRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsTransferServerInvalidURLRule returns new rule with default attributes
func NewAwsTransferServerInvalidURLRule() *AwsTransferServerInvalidURLRule {
	return &AwsTransferServerInvalidURLRule{
		resourceType:  "aws_transfer_server",
		attributeName: "url",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidURLRule) Name() string {
	return "aws_transfer_server_invalid_url"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidURLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidURLRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidURLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidURLRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"url must be 255 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}