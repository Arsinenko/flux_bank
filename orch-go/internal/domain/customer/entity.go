package customer

import "time"

type Customer struct {
	Id        int32
	FirstName string
	LastName  string
	Email     string
	Phone     *string
	BirthDate *time.Time
	CreatedAt *time.Time
}
