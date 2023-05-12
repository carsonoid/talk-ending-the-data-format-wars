package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/types"
)

var tomlCode = `
// START TOML OMIT
ImageId = "ami-12345678"
InstanceType = "t2.micro"
Placement = { AvailabilityZone = "us-east-1a" }
BlockDeviceMappings = [
   # each inline table must be on a single line to be valid TOML
   {  DeviceName = "/dev/sda", Ebs = { DeleteOnTermination = true, VolumeSize = 8 } },
   {  DeviceName = "/dev/sdb", Ebs = { DeleteOnTermination = true, VolumeSize = 20 } },
]
// END TOML OMIT
`

func main() {
	run.PrintAs(types.RunInstancesInput{}, run.TOMLExample(tomlCode))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
