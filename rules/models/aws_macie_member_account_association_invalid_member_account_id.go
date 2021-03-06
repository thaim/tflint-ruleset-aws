// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule checks the pattern is valid
type AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsMacieMemberAccountAssociationInvalidMemberAccountIDRule returns new rule with default attributes
func NewAwsMacieMemberAccountAssociationInvalidMemberAccountIDRule() *AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule {
	return &AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule{
		resourceType:  "aws_macie_member_account_association",
		attributeName: "member_account_id",
		pattern:       regexp.MustCompile(`^[0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule) Name() string {
	return "aws_macie_member_account_association_invalid_member_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacieMemberAccountAssociationInvalidMemberAccountIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]{12}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
