package Domain

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/taller-go/src/api/Utils"
	"io/ioutil"
	"net/http"
)

var usuario User
const url = "https://api.mercadolibre.com/users/"

func (user *User) Get() (*Utils.ApiError ){

	var data []byte
	if user.ID == 0 {
		return  &Utils.ApiError{
			"userId empty",
			http.StatusBadRequest,
		}
	}

	urlFinal := fmt.Sprintf("%s%d", url, user.ID)

	response, err := http.Get( urlFinal)
	if err != nil {

		return &Utils.ApiError{
			Message: "bar request",
			Status: http.StatusBadRequest}

	}
	data, error := ioutil.ReadAll(response.Body)

	if error != nil{
		return  &Utils.ApiError{
			Message: "Internal Error",
			Status: http.StatusInternalServerError}
	}

	if err1 := json.Unmarshal([]byte(data), &user); err1 != nil{
		return  &Utils.ApiError{
			Message: "exit",
			Status: http.StatusInternalServerError}
	}
	return nil

}