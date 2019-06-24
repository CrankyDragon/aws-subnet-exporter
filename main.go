package main

import (
	"flag"

	"github.com/CrankyDragon/aws-subnet-exporter/cmd"
)

func main() {

	var (
		vpcID    = flag.String("vpc-id", "", "AWS VPC ID")
		vpcName  = flag.String("vpc-name", "", "AWS VPC Name")
		region   = flag.String("region", "us-west-2", "AWS Region")
		tagValue = flag.String("tag-val", "Type", "Tag value to group by")
	)

	flag.Parse()

	cmd := cmd.NewCommand(region, vpcName, vpcID, tagValue)
	cmd.Execute()
}
