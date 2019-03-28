package schema

// MicroserviceDetail is the structure of the collection that stores
// the details while registering a microservice
type MicroserviceDetail struct {
	MicroserviceName             string
	MicroserviceDockerRepository string
}
