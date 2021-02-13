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

type ResponseWartegDetail struct {
	Status     string           `json:"status"`
	StatusCode int              `json:"status_code"`
	Message    string           `json:"message"`
	Timestamp  time.Time        `json:"timestamp"`
	Data       DataWartegDetail `json:"data"`
}

type DataWartegDetail struct {
	WartegId          string `json:"warteg_id"`
	WartegName        string `json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddress     string `json:"warteg_addr"`
	WartegContactName string `json:"warteg_contact_name"`
	WartegPhone       string `json:"warteg_phone"`
}

type ResponseWartegList struct {
	Status     string           `json:"status"`
	StatusCode int              `json:"status_code"`
	Message    string           `json:"message"`
	Timestamp  time.Time        `json:"timestamp"`
	Data       []DataWartegList `json:"data"`
}

type DataWartegList struct {
	WartegId   string `json:"warteg_id"`
	WartegName string `json:"warteg_name"`
	WartegAddr string `json:"warteg_addr"`
}

type ResponseWartegAdd struct {
	Status     string        `json:"status"`
	StatusCode int           `json:"status_code"`
	Message    string        `json:"message"`
	Timestamp  time.Time     `json:"timestamp"`
	Data       DataWartegAdd `json:"data"`
}

type DataWartegAdd struct {
	WartegName        string `json:"warteg_name"`
	WartegDesc        string `json:"warteg_desc"`
	WartegAddress     string `json:"warteg_addr"`
	WartegContactName string `json:"warteg_contact_name"`
	WartegPhone       string `json:"warteg_phone"`
}

type ResponseMenu struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Timestamp  time.Time   `json:"timestamp"`
	Data       interface{} `json:"data"`
}
