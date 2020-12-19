package main

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)
}
