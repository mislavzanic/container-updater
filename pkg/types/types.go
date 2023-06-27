package types

const (
	CreateUpdate string = "CreateUpdate"
	Update       string = "Update"
)

type Payload struct {
	Image string
	PortBindings []PortBinding
	RequestType string
}

type PortBinding struct {
	HostIP string
	HostPort string
	ContainerPort string
}
