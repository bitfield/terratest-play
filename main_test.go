package main

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestPlay(t *testing.T) {
	t.Parallel()
	awsRegion := "ap-northeast-1"
	terraformOptions := &terraform.Options{
		TerraformDir: "testdata/terratest-play-example",
	}
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	instanceID := terraform.Output(t, terraformOptions, "instance_id")
	aws.AddTagsToResource(t, awsRegion, instanceID, map[string]string{"terratest": "true"})
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)
	testingTag, ok := instanceTags["terratest"]
	if !ok || testingTag != "true" {
		t.Error("tag not found")
	}
}
