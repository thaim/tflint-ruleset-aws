// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventPermissionInvalidActionRule checks the pattern is valid
type AwsCloudwatchEventPermissionInvalidActionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventPermissionInvalidActionRule returns new rule with default attributes
func NewAwsCloudwatchEventPermissionInvalidActionRule() *AwsCloudwatchEventPermissionInvalidActionRule {
	return &AwsCloudwatchEventPermissionInvalidActionRule{
		resourceType:  "aws_cloudwatch_event_permission",
		attributeName: "action",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^events:[a-zA-Z]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventPermissionInvalidActionRule) Name() string {
	return "aws_cloudwatch_event_permission_invalid_action"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventPermissionInvalidActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventPermissionInvalidActionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventPermissionInvalidActionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventPermissionInvalidActionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"action must be 64 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"action must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^events:[a-zA-Z]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
