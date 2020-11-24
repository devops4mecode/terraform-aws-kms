// Managed By : DevOps4Me
// Description : This Terratest is used to test the Terraform VPC module.
// Copyright @ DevOps4Me. All Right Reserved.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// Source path of Terraform directory.
		TerraformDir: "../gorun",
	}

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// To clean up any resources that have been created, run 'terraform destroy' towards the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// To get the value of an output variable, run 'terraform output'
	keyArn := terraform.Output(t, terraformOptions, "key_arn")
	Tags := terraform.OutputMap(t, terraformOptions, "tags")

	// Check that we get back the outputs that we expect
	assert.Contains(t, keyArn, "arn:aws:kms")
	assert.Equal(t, "test-devops4me-kms", Tags["Name"])
