package repositories

type Agent struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewAgent() *Agent {
	return &Agent{}
}

func GetAgent(id int64) *Agent {
	return NewAgent()
}
