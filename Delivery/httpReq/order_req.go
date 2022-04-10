package httpReq

type OrderReq struct {
	Customer  string     `json:"customer"`
	ListOrder []ListMenu `json:"list_order"`
}

type ListMenu struct {
	IdMenu   int `json:"id_menu"`
	Quantity int `json:"quantity"`
}
