package main

import (
	"fmt"

	"github.com/koosie0507/pluralsight-go-getting-started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
