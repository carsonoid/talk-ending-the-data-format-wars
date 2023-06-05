package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/types"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{"ImageId":"ami-12345678","InstanceType":"t2.micro","Placement":{"AvailabilityZone":"us-east-1a"},"BlockDeviceMappings":[{"DeviceName":"/dev/sda1","Ebs":{"DeleteOnTermination":true,"VolumeSize":8}},{"DeviceName":"/dev/sdb","Ebs":{"DeleteOnTermination":true,"VolumeSize":20}}]}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
# You can of course use the json as-is as yaml
{"ImageId":"ami-12345678","InstanceType":"t2.micro","Placement":{"AvailabilityZone":"us-east-1a"},"BlockDeviceMappings":[{"DeviceName":"/dev/sda1","Ebs":{"DeleteOnTermination":true,"VolumeSize":8}},{"DeviceName":"/dev/sdb","Ebs":{"DeleteOnTermination":true,"VolumeSize":20}}]}
// END YAML OMIT
`

var yamlCode2 = `
// START YAML2 OMIT
# Or you can drop the quotes but be warned:
# every single bit of whitespace here is required!
{ ImageId: "ami-12345678", InstanceType: "t2.micro", Placement: { AvailabilityZone: "us-east-1a" }, BlockDeviceMappings: [ { DeviceName: "/dev/sda1", Ebs: { DeleteOnTermination: true, VolumeSize: 8 } }, { DeviceName: "/dev/sdb", Ebs: { DeleteOnTermination: true, VolumeSize: 20 } } ] }
// END YAML2 OMIT
`

func main() {
	run.PrintAs(types.RunInstancesInput{}, run.JSONExample(jsonCode))
	run.PrintAs(types.RunInstancesInput{}, run.YAMLExample(yamlCode))
	run.PrintAs(types.RunInstancesInput{}, run.YAMLExample(yamlCode2))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
