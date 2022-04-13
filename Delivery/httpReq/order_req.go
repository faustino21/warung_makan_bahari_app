package httpReq

type OrderReq struct {
	Customer string     `json:"customer"`
	Table    int        `json:"table_req"`
	ListMenu []ListMenu `json:"list_order"`
}

type ListMenu struct {
	IdMenu   int `json:"id_menu"`
	Quantity int `json:"quantity"`
}
