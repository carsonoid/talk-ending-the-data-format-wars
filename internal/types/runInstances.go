package types

// START OMIT
type RunInstancesInput struct {
	ImageId             string               `json:"ImageId" yaml:"ImageId", toml:"ImageId", hcl:"ImageId"`
	InstanceType        string               `json:"InstanceType" yaml:"InstanceType", toml:"InstanceType", hcl:"InstanceType"`
	Placement           Placement            `json:"Placement" yaml:"Placement", toml:"Placement", hcl:"Placement"`
	BlockDeviceMappings []BlockDeviceMapping `json:"BlockDeviceMappings" yaml:"BlockDeviceMappings", toml:"BlockDeviceMappings", hcl:"BlockDeviceMappings"`
}

type Placement struct {
	AvailabilityZone string `json:"AvailabilityZone" yaml:"AvailabilityZone", toml:"AvailabilityZone", hcl:"AvailabilityZone"`
}

type BlockDeviceMapping struct {
	DeviceName string `json:"DeviceName" yaml:"DeviceName", toml:"DeviceName", hcl:"DeviceName"`
	Ebs        Ebs    `json:"Ebs" yaml:"Ebs", toml:"Ebs", hcl:"Ebs"`
}

type Ebs struct {
	DeleteOnTermination bool `json:"DeleteOnTermination" yaml:"DeleteOnTermination", toml:"DeleteOnTermination", hcl:"DeleteOnTermination"`
	VolumeSize          int  `json:"VolumeSize" yaml:"VolumeSize", toml:"VolumeSize", hcl:"VolumeSize"`
}

// END OMIT
