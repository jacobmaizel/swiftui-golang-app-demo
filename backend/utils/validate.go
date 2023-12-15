package utils

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Namespace string `json:"namespace"`
	Field     string `json:"field"`
	Message   string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This is required"
	case "lte":
		return "should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func BindAndValidate(c *gin.Context, toBind any) []ErrorMsg {

	log.Println("toBind", toBind)

	// log.Println(toBind)

	if err := c.ShouldBindJSON(toBind); err != nil {

		log.Println("Bind Errors:::::::::::::", err)

		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {

				log.Println(fe.Namespace())
				log.Println(fe.Field())
				log.Println("bingo", fe.StructNamespace())
				log.Println("struct field", fe.StructField())
				log.Println(fe.Tag())
				log.Println("actual tag", fe.ActualTag())
				log.Println(fe.Kind())
				log.Println(fe.Type())
				log.Println(fe.Value())
				log.Println(fe.Param())

				out[i] = ErrorMsg{fe.StructNamespace(), fe.Field(), GetErrorMsg(fe)}
			}

			return out

		}

		return []ErrorMsg{
			{
				Namespace: "Error",
				Field:     "Error",
				Message:   err.Error(),
			},
		}

	}

	log.Println("no errors, life is good in bind and validate")
	return nil

}
