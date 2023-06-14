package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/types"
)

var hclCode = `
// START HCL OMIT
{"ImageId":"ami-12345678","InstanceType":"t2.micro","Placement":{"AvailabilityZone":"us-east-1a"},"BlockDeviceMappings":[{"DeviceName":"/dev/sda1","Ebs":{"DeleteOnTermination":true,"VolumeSize":8}},{"DeviceName":"/dev/sdb","Ebs":{"DeleteOnTermination":true,"VolumeSize":20}}]}
// END HCL OMIT
`

var hclCode2 = `
// START HCL2 OMIT
Document = {ImageId="ami-12345678",InstanceType="t2.micro",Placement={AvailabilityZone="us-east-1a"},BlockDeviceMappings=[{DeviceName="/dev/sda1",Ebs={DeleteOnTermination=true,VolumeSize=8}},{DeviceName="/dev/sdb",Ebs={DeleteOnTermination=true,VolumeSize=20}}]}
// END HCL2 OMIT
`

type wrapper struct {
	Document types.RunInstancesInput `toml:"Document" hcl:"Document"`
}

func main() {
	run.PrintAs(types.RunInstancesInput{}, run.HCLExample(hclCode))
	run.PrintAs(wrapper{}, run.HCLExample(hclCode2))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
