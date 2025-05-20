package collector

type status struct {
	Name             string    `json:"name"`
	Connector        connector `json:"connector"`
	Tasks            []task    `json:"tasks"`
	ValidationErrors []string  `json:"validation_errors"`
}

type connector struct {
	State    string `json:"state"`
	WorkerId string `json:"worker_id"`
}

type task struct {
	State    string  `json:"state"`
	Id       float64 `json:"id"`
	WorkerId string  `json:"worker_id"`
}

func (s status) ConsumerGroup() string {
	return "connect-" + s.Name
}
