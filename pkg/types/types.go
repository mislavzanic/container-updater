package types

type Payload struct {
	Image string
	PortBindings []PortBinding
}

type PortBinding struct {
	HostIP string
	HostPort string
	ContainerPort string
}
