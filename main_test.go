package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
)

// https://github.com/aws/aws-sdk-go/blob/master/aws/ec2metadata/api_test.go#L21
const instanceIdentityDocumentJSON = `{
  "devpayProductCodes" : null,
  "availabilityZone" : "us-east-1d",
  "privateIp" : "10.158.112.84",
  "version" : "2010-08-31",
  "region" : "us-east-1",
  "instanceId" : "i-1234567890abcdef0",
  "billingProducts" : null,
  "instanceType" : "t1.micro",
  "accountId" : "123456789012",
  "pendingTime" : "2015-11-19T16:32:11Z",
  "imageId" : "ami-5fb8c835",
  "kernelId" : "aki-919dcaf8",
  "ramdiskId" : null,
  "architecture" : "x86_64"
}`

func TestGenerateHostname(t *testing.T) {
	instanceIdentityDocument := ec2metadata.EC2InstanceIdentityDocument{}
	err := json.Unmarshal([]byte(instanceIdentityDocumentJSON), &instanceIdentityDocument)
	if err != nil {
		t.Fatal(err)
	}
	hostname, err := GenerateHostname(instanceIdentityDocument)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("hostname is: %s\n", *hostname)
}
