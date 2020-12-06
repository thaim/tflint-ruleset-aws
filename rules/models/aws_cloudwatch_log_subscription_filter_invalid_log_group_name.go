// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule checks the pattern is valid
type AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule returns new rule with default attributes
func NewAwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule() *AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule {
	return &AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule{
		resourceType:  "aws_cloudwatch_log_subscription_filter",
		attributeName: "log_group_name",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^[\.\-_/#A-Za-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule) Name() string {
	return "aws_cloudwatch_log_subscription_filter_invalid_log_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"log_group_name must be 512 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"log_group_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\.\-_/#A-Za-z0-9]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}