package topgun

type Instance interface {
	GetAddress() string
}

type InstanceManager interface {
	GetWebInstances() []Instance
}
