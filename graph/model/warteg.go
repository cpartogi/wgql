package model

type WartegDetail struct {
	WartegId    string `json:"warteg_id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Address     string `json:"address"`
	ContactName string `json:"contact_name"`
	Phone       string `json:"phone"`
}

type WartegList struct {
	WartegId string `json:"warteg_id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
}

type InputWarteg struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	ContactName string `json:"contact_name"`
	Phone       string `json:"phone"`
}

type Resdata struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type AddWarteg struct {
	WartegName        string `json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddr        string `json:"warteg_addr"`
	WartegContactName string `json:"warteg_contact_name"`
	WartegPhone       string `json:"warteg_phone"`
}
