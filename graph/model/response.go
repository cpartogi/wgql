package model

import "time"

type ResponseMenuType struct {
	Status     string         `json:"status"`
	StatusCode int            `json:"status_code"`
	Message    string         `json:"message"`
	Timestamp  time.Time      `json:"timestamp"`
	Data       []DataMenuType `json:"data"`
}

type DataMenuType struct {
	MenuTypeId   int    `json:"menu_type_id"`
	MenuTypeName string `json:"menu_type_name"`
}

type ResponseMenuDetail struct {
	Status     string         `json:"status"`
	StatusCode int            `json:"status_code"`
	Message    string         `json:"message"`
	Timestamp  time.Time      `json:"timestamp"`
	Data       DataMenuDetail `json:"data"`
}

type DataMenuDetail struct {
	MenuId       string `json:"menu_id"`
	MenuTypeName string `json:"menu_type_name"`
	WartegId     string `json:"warteg_id"`
	MenuName     string `json:"menu_name"`
	MenuDetail   string `json:"menu_detail"`
	MenuPicture  string `json:"menu_picture"`
	MenuPrice    int    `json:"menu_price"`
}

type ResponseMenuList struct {
	Status     string         `json:"status"`
	StatusCode int            `json:"status_code"`
	Message    string         `json:"message"`
	Timestamp  time.Time      `json:"timestamp"`
	Data       []DataMenuList `json:"data"`
}

type DataMenuList struct {
	MenuId       string `json:"menu_id"`
	MenuTypeName string `json:"menu_type_name"`
	WartegId     string `json:"warteg_id"`
	MenuName     string `json:"menu_name"`
	MenuPrice    int    `json:"menu_price"`
}
