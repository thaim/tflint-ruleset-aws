// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmMaintenanceWindowTargetInvalidDescriptionRule checks the pattern is valid
type AwsSsmMaintenanceWindowTargetInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSsmMaintenanceWindowTargetInvalidDescriptionRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowTargetInvalidDescriptionRule() *AwsSsmMaintenanceWindowTargetInvalidDescriptionRule {
	return &AwsSsmMaintenanceWindowTargetInvalidDescriptionRule{
		resourceType:  "aws_ssm_maintenance_window_target",
		attributeName: "description",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowTargetInvalidDescriptionRule) Name() string {
	return "aws_ssm_maintenance_window_target_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowTargetInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowTargetInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowTargetInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowTargetInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 128 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"description must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
