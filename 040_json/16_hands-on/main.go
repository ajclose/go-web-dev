package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type status struct {
	Code    int
	Descrip string
}

type statuses []status

func main() {
	var data statuses

	rcvd := `[{"Descrip":"StatusOK","Code":200},{"Descrip":"StatusMovedPermanently","Code":301},{"Descrip":"StatusFound","Code":302},{"Descrip":"StatusSeeOther","Code":303},{"Descrip":"StatusTemporaryRedirec","Code":307},{"Descrip":"StatusBadRequest","Code":400},{"Descrip":"StatusUnauthorized","Code":401},{"Descrip":"StatusPaymentRequired","Code":402},{"Descrip":"StatusForbidden","Code":403},{"Descrip":"StatusNotFound","Code":404},{"Descrip":"StatusMethodNotAllowed","Code":405},{"Descrip":"StatusTeapot","Code":418},{"Descrip":"StatusInternalServerError","Code":500}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range data {
		fmt.Println(v.Code, v.Descrip)
	}
}
