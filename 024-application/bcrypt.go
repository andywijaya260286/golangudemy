package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt" //import "golang.org/x/crypto/bcrypt"
)

func main() {
	s := `pmmp123`

	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	v, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	fmt.Println(v, err)

	//func CompareHashAndPassword(hashedPassword, password []byte) error
	loginPassword := `pmmp123`
	err = bcrypt.CompareHashAndPassword(v, []byte(loginPassword))

	if err != nil {
		fmt.Println("error = ", err)
	} else {
		fmt.Println("success login")
	}

}
