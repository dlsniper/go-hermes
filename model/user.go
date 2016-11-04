package model

// User type represents a user (customer) in our system.
type User struct {
	ID         int
	Username   string
	Servers    *[]Server
	MobileApps *[]MobileApp
}
