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

type InputMenu struct {
	Name       string `json:"name"`
	Detail     string `json:"detail"`
	Picture    string `json:"picture"`
	Price      int    `json:"price"`
	MenuTypeId int    `json:"menu_type_id"`
	WartegId   string `json:"warteg_id"`
}

type AddMenu struct {
	MenuName    string `json:"menu_name"`
	MenuDetail  string `json:"menu_detail"`
	MenuPicture string `json:"menu_picture"`
	MenuPrice   int    `json:"menu_price"`
	MenuTypeId  int    `json:"menu_type_id"`
	WartegId    string `json:"warteg_id"`
}
