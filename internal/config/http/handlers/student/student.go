package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/gaurav25kapoor/students-api/internal/config/storage"
	"github.com/gaurav25kapoor/students-api/internal/config/types"
	"github.com/gaurav25kapoor/students-api/internal/config/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a new student")
		var student types.Student
		err:=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF){
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return
		}

		if err!=nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return
		}
		
		//request validation
		if err:=validator.New().Struct(student);err!=nil{
			validateErrs:=err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
			return
		}

		

		lastId,err:=storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully",slog.String("userId",fmt.Sprint(lastId)))

		if err!=nil{
			response.WriteJson(w,http.StatusInternalServerError,err)
		}
  

		response.WriteJson(w,http.StatusCreated,map[string]int64{"id":lastId})
	}
}