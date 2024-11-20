package location

import (
	httpDtos "SleekSpace/dtos/HttpServices"
	generalUtilities "SleekSpace/utilities/funcs/general"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LocationReverseGeoCoding(c *gin.Context) {
	var coords httpDtos.ReverseGeoCodingDTO
	validateModelFields := validator.New()
	c.BindJSON(&coords)
	modelFieldsValidationError := validateModelFields.Struct(coords)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	url := "https://api.locationiq.com/v1/reverse.php?key=" + generalUtilities.GetEnvVariables().LocationIQToken + "&lat=" + coords.Lat + "&lon=" + coords.Lon + "&format=json"

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching data"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not convert response body"})
		return
	}

	var responseObject httpDtos.IReverseLocation
	json.Unmarshal(body, &responseObject)
	c.JSON(http.StatusOK, gin.H{"response": responseObject})
}
