package lib

type User struct {
	Username string `json:"username" bson:"username"`
	ApiKey   string `json:"apikey" bson:"apikey"`
}
