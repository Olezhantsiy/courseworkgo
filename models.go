package main

type client struct {
	Id        int    `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

type service struct {
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Price int    `json:"Price"`
}

type payment struct {
	Id        int    `json:"Id"`
	ClientId  int    `json:"ClientId"`
	ServiceId int    `json:"ServiceId"`
	Date      string `json:"Date"`
	Quantity  int    `json:"Quantity"`
	Amount    int    `json:"Amount"`
}

type InfoDemo struct {
	FirstName   string  `json:"FirstName"`
	LastName    string  `json:"LastName"`
	Phone       string  `json:"Phone"`
	ServiceName string  `json:"ServiceName"`
	Quantity    int     `json:"Quantity"`
	Amount      float64 `json:"Amount"`
}
