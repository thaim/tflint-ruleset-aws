// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventTargetInvalidRuleRule checks the pattern is valid
type AwsCloudwatchEventTargetInvalidRuleRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventTargetInvalidRuleRule returns new rule with default attributes
func NewAwsCloudwatchEventTargetInvalidRuleRule() *AwsCloudwatchEventTargetInvalidRuleRule {
	return &AwsCloudwatchEventTargetInvalidRuleRule{
		resourceType:  "aws_cloudwatch_event_target",
		attributeName: "rule",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\.\-_A-Za-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventTargetInvalidRuleRule) Name() string {
	return "aws_cloudwatch_event_target_invalid_rule"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventTargetInvalidRuleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventTargetInvalidRuleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventTargetInvalidRuleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventTargetInvalidRuleRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"rule must be 64 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"rule must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\.\-_A-Za-z0-9]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
