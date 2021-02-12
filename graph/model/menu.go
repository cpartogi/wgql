package model

type MenuType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MenuDetail struct {
	MenuId   string `json:"menu_id"`
	Name     string `json:"name"`
	Detail   string `json:"detail"`
	Picture  string `json:"picture"`
	Price    int    `json:"price"`
	TypeName string `json:"type_name"`
	WartegId string `json:"warteg_id"`
}
