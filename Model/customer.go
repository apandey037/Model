package Model

//type Address struct {
//	State   string
//	City    string
//	PinCode int
//}

type Customer struct {
	AccountNo   int    `json:"account_no"`
	Name        string `json:"name"`
	PanNo       string `json:"pan_no"`
	Mobile      int    `json:"mobile"`
	AccountType string `json:"account_type"`
	//Address     Address
	AdharNo  string `json:"adhar_no"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
