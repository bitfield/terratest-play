package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestPlay(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "testdata/terratest-play-example",
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Validate your code works as expected
	want := []byte("foo!")
	got, err := ioutil.ReadFile("testdata/terratest-play-example/foo.bar")
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("want %q, got %q", want, got)
	}
}
