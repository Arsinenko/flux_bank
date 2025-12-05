package customerAdress

type CustomerAddress struct {
	Id          *int32
	Customer_id int32
	Country     string
	City        string
	Street      string
	ZipCode     string
	IsPrimary   bool
}
