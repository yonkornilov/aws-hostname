package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// https://github.com/aws/aws-sdk-go/blob/master/aws/ec2metadata/api_test.go#L21
const instanceIdentityDocument = `{
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

// https://github.com/aws/aws-sdk-go/blob/master/aws/ec2metadata/api_test.go#L52
func initTestServer(path string, resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != path {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		w.Write([]byte(resp))
	}))
}

func TestGetInstance(t *testing.T) {
	server := initTestServer(
		"/latest/dynamic/instance-identity/document",
		instanceIdentityDocument,
	)
	defer server.Close()
	meta := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	svc := ec2.New(unit.Session)

	instance, err := getInstance(meta, svc)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("hostname is: %s\n", *instance)
}
