package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsInstancePreviousType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "t1.micro is previous type",
			Content: `
resource "aws_instance" "web" {
    instance_type = "t1.micro"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsInstancePreviousTypeRule(),
					Message: "\"t1.micro\" is previous generation instance type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 21},
						End:      hcl.Pos{Line: 3, Column: 31},
					},
				},
			},
		},
		{
			Name: "t2.micro is not previous type",
			Content: `
resource "aws_instance" "web" {
    instance_type = "t2.micro"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsInstancePreviousTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
