package main

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users = []user{
	{ID: "1", FirstName: "Emalee", LastName: "Creigan", Email: "ecreigan0@nature.com"},
	{ID: "2", FirstName: "Mariellen", LastName: "Peerless", Email: "mpeerless1@dot.gov"},
	{ID: "3", FirstName: "Gabie", LastName: "Brims", Email: "gbrims2@cnbc.com"},
	{ID: "4", FirstName: "Francesco", LastName: "Morman", Email: "fmorman3@rediff.com"},
	{ID: "5", FirstName: "Ario", LastName: "Denerley", Email: "adenerley4@state.gov"},
}
