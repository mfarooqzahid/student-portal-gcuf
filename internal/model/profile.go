package model

type Profile struct {
	Name       string `json:"name"`
	FatherName string `json:"fathername"`
	RollNo     string `json:"rollno"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	DOB        string `json:"dob"`
	Phone      string `json:"phone"`
	City       string `json:"city"`
	Address    string `json:"address"`
	ImageURL   string `json:"imageurl"`
	Program    string `json:"program"`
	Session    string `json:"session"`
	RegNo      string `json:"regno"`
}
