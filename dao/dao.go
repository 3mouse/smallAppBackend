package dao

type User struct {
	Id 			string `json:"id"`
	UserName 	string `json:"user_name"`
	Password 	string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Token 		string `json:"token"`
}

type Good struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Price 	string `json:"price"`
	ProductionDate string `json:"production_date"`
	Size 	string `json:"size"`
}

type Order struct {
	Id string	`json:"id"`
	Time string `json:"time"`
	Owner string `json:"owner"`
	Content string `json:"content"`
}
