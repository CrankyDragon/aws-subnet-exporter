package cmd

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Command ...
type Command struct {
	region   *string
	vpcName  *string
	vpcID    *string
	tagValue *string
	svc      *ec2.EC2
}

// NewCommand ...
func NewCommand(region, vpcName, vpcID, tagValue *string) *Command {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: region,
	}))

	return &Command{
		svc:      ec2.New(sess),
		region:   region,
		vpcName:  vpcName,
		vpcID:    vpcID,
		tagValue: tagValue,
	}

}

// Execute ..
func (c *Command) Execute() {
	vpcID := c.getVpcID()

	describeSubnetsInput := &ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []*string{vpcID},
			},
		},
	}

	describeSubnetsOutput, err := c.svc.DescribeSubnets(describeSubnetsInput)
	if err != nil {
		log.Fatal(err.Error())
	}

	output := NewOutput()
	for _, subnet := range describeSubnetsOutput.Subnets {

		value := ""
		for _, tag := range subnet.Tags {
			if *tag.Key == *c.tagValue {
				value = *tag.Value
				break
			}
		}

		if value != "" {
			output.Add(value, *subnet.SubnetId)
		}

	}

	output.Render()
}

func (c *Command) getVpcID() *string {

	if !c.issetVpcName() {
		return c.vpcID
	}

	descrubVpcInput := &ec2.DescribeVpcsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{c.vpcName},
			},
		},
	}

	describeVpcOutput, err := c.svc.DescribeVpcs(descrubVpcInput)
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(describeVpcOutput.Vpcs) == 0 {
		log.Fatalf("Could not find VPC %s\n", *c.vpcName)
	}

	return describeVpcOutput.Vpcs[0].VpcId
}

func (c *Command) issetVpcName() bool {
	return *c.vpcID == ""
}
