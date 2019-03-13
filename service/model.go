package main

//Request structure for input
type Request struct {
	IP        string   `json:"ip"`
	Countries []string `json:"countries"`
}

//Body of http response
type Body struct {
	Action string `json:"action"`
}
