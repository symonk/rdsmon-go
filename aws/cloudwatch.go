package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type CloudWatchMonitor struct {
	client *cloudwatch.Client
}

// Initialize a new cloud watch client from aws credentials
func NewCloudWatchMonitor(config aws.Config) *CloudWatchMonitor {
	return &CloudWatchMonitor{
		client: cloudwatch.NewFromConfig(config),
	}
}

// Reports the available metrics for the instance
func (c *CloudWatchMonitor) ReportAvailableMetrics() {
}
