package main

type Movie struct {
	id       string    `json: "id"`
	isbn     string    `json: "isbn"`
	title    string    `json: "title"`
	director *Director `json: "director"`
}

type Director struct {
	firstName string `json: "firstname"`
	lastName  string `json: "lastname"`
}

var movies []Movie

func main() {

}
