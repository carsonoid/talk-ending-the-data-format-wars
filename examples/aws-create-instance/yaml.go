package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/types"
)

// //go:embed example.json
// var jsonBytes []byte

var yamlCode = `
// START YAML OMIT
{
   "ImageId": "ami-12345678",
   "InstanceType": "t2.micro",
   "Placement": {
      "AvailabilityZone": "us-east-1a"
   },
   "BlockDeviceMappings": [
      {
         "DeviceName": "/dev/sda1",
         "Ebs": {
            "DeleteOnTermination": true,
            "VolumeSize": 8
         }
      },
      {
         "DeviceName": "/dev/sdb",
         "Ebs": { "DeleteOnTermination": true, "VolumeSize": 20 }
      }
   ]
}
// END YAML OMIT
`

func main() {
	run.PrintAs(types.RunInstancesInput{}, run.YAMLExample(yamlCode))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
