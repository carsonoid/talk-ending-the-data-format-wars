package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/types"
)

var tomlCode = `
// START TOML OMIT
AmiId = "ami-0ff8a91507f77f867"
InstanceType = "t2.micro"

[Placement]
AvailabilityZone = "us-east-1a"

[[BlockDeviceMappings]] # device 1
DeviceName = "/dev/sda1"

# you can use dot notation for nested tables
Ebs.DeleteOnTermination = true
Ebs.VolumeSize = 8

[[BlockDeviceMappings]] # device 2
DeviceName = "/dev/sdb"

# or you can use inline tables
Ebs = { VolumeSize = 20, DeleteOnTermination = true }
// END TOML OMIT
`

func main() {
	run.PrintAs(types.RunInstancesInput{}, run.TOMLExample(tomlCode))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
