package models

// User struct to represent a user in the database
type User struct {
	Email       string `bson:"email"`       //User email
	Name        string `bson:"name"`        //User name
	DOB         string `bson:"dob"`         //User Date of birth
	Address     string `bson:"address"`     //User address
	Description string `bson:"description"` //User description
	Follower    string `bson:"follower"`
	Following   string `bson:"following"`
	Password    string `bson:"password"`   //User password
	CreatedAt   string `bson:"created_at"` //User created date
}
