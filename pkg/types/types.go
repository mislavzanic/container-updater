package types

const (
	CreateUpdate string = "CreateUpdate"
	Update       string = "Update"
)

type Payload struct {
	Image string
	Name string
	PortBindings []PortBinding
	RequestType string
	Secret string
}

type PortBinding struct {
	HostIP string
	HostPort string
	ContainerPort string
}
