package libdto

type MultiForm struct {
	Ids    string `json:"ids" form:"ids"`
	IdList []int  `json:"id_list" form:"id_list"`
}

type Multi64Form struct {
	Ids    string  `json:"ids" form:"ids"`
	IdList []int64 `json:"id_list" form:"id_list"`
}
