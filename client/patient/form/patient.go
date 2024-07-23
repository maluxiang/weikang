package form

type Patient struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Sex        int    `json:"sex"`
	IdCard     string `json:"id_card"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Doctor     string `json:"doctor"`
	Status     int    `json:"status"`
}

type EsSearch struct {
	Index     string `json:"index"`
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
}
