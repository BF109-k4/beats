package clustersummary


type SpoutInfo struct{
	SpoutId string `json:"spoutId"`
	Emitted float64 `json:"emitted"`
	Tasks float64 `json:"tasks"`
	Failed float64 `json:"failed"`
	Acked float64 `json:"acked"`
}


type BoltInfo struct{
	BoltId string `json:"boltId"`
	Emitted float64 `json:"emitted"`
	Tasks float64 `json:"tasks"`
	Failed float64 `json:"failed"`
	Acked float64 `json:"acked"`

}

type TopoInfo struct {
	Name string `json:"name"`
	Status string `json:"status"`
	Uptime string `json:"uptime"`
	WorkersTotal float64 `json:"workersTotal"`
	TasksTotal float64 `json:"tasksTotal"`
	ExecutorsTotal float64 `json:"executorsTotal"`
	ReplicationCount float64 `json:"replicationCount"`
	AssignedMemOnHeap float64 `json:"assignedMemOnHeap"`
	Spouts []SpoutInfo `json:"spouts"`
	Bolts []BoltInfo `json:"bolts"`
}
