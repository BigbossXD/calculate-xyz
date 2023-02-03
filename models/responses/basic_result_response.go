package responses

type BasicResultResponse struct {
	X       interface{} `json:"x"`
	Y       interface{} `json:"y"`
	Z       interface{} `json:"z"`
	DataSet []int       `json:"dataSet"`
}
