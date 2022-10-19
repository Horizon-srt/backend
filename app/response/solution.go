package response

import (
	"github.com/EduOJ/backend/app/response/resource"
)

type GetSolutionResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    struct {
		*resource.Solution `json:"solution"`
	} `json:"data"`
}

type CreateSolutionResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    struct {
		*resource.Solution `json:"solution"`
	} `json:"data"`
}

type UpdateSolutionResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    struct {
		*resource.Solution `json:"solution"`
	} `json:"data"`
}
