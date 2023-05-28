package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/types"
)

var tomlCode2 = `
// START TOML OMIT
Document = {ImageId="ami-12345678",InstanceType="t2.micro",Placement={AvailabilityZone="us-east-1a"},BlockDeviceMappings=[{DeviceName="/dev/sda1",Ebs={DeleteOnTermination=true,VolumeSize=8}},{DeviceName="/dev/sdb",Ebs={DeleteOnTermination=true,VolumeSize=20}}]}
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
Document = {ImageId="ami-12345678",InstanceType="t2.micro",Placement={AvailabilityZone="us-east-1a"},BlockDeviceMappings=[{DeviceName="/dev/sda1",Ebs={DeleteOnTermination=true,VolumeSize=8}},{DeviceName="/dev/sdb",Ebs={DeleteOnTermination=true,VolumeSize=20}}]}
// END HCL OMIT
`

// START WRAPPER OMIT
type wrapper struct {
	Document types.RunInstancesInput `toml:"Document" hcl:"Document"`
}

// END WRAPPER OMIT

func main() {
	run.PrintAs(wrapper{}, run.TOMLExample(tomlCode2))
	run.PrintAs(wrapper{}, run.HCLExample(hclCode))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
