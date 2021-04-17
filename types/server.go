package types

type Body struct {
	Sender string `json:"sender"`
	Reveiver string `json:"receiver"`
}

type GetResp struct {
	Data map[string]int `json:"data"`
}

func GetRespInit(data map[string]int) *GetResp {
	return &GetResp{ Data: data }
}