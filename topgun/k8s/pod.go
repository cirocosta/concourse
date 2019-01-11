package k8s

type Pod struct {
	Status struct {
		ContainerStatuses []struct {
			Name  string `json:"name"`
			Ready bool   `json:"ready"`
		} `json:"containerStatuses"`
		Phase  string `json:"phase"`
		HostIp string `json:"hostIP"`
		Ip     string `json:"podIP"`
	} `json:"status"`
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
}

func (p *Pod) GetAddress() string {
	return p.Status.Ip
}
