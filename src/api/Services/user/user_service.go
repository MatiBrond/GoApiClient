package user

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/taller-go/src/api/Domain"
	"github.com/mercadolibre/taller-go/src/api/Utils"
	"io/ioutil"
	"net/http"
)


var usuario Domain.User
const url = "https://api.mercadolibre.com/users/"



func GetUserByIdSer(id int64) (*Domain.User, *Utils.ApiError){

	var data []byte
	var user Domain.User

	urlFinal := fmt.Sprintf("%s%d", url, id)

	response, err := http.Get( urlFinal)
	if err != nil {

		return nil, &Utils.ApiError{
			Message: "Fatal URL",
			Status: http.StatusBadRequest}

	}
	data, error := ioutil.ReadAll(response.Body)

	if error != nil{
		return nil, &Utils.ApiError{
			Message: "Fallo aca",
			Status: http.StatusInternalServerError}
	}

	if err1 := json.Unmarshal([]byte(data), &user); err1 != nil{
		return nil, &Utils.ApiError{
			Message: "Id is empty",
			Status: http.StatusBadRequest}
	}
	return &user, nil
}

