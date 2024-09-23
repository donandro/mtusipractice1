package controllers

import (
	"pillsreminder/apps/mobileapi/api/rest/models"
	commongin "pillsreminder/pkg/rest/gin"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type PillsController struct {
	logger zerolog.Logger
}

func NewPillsController(logger zerolog.Logger) *AuthController {
	return &AuthController{
		logger: logger,
	}
}

// GetPills godoc
//
//	@Summary		GetPills
//	@Description	Get pills with filter by name.
//	@Accept			json
//	@Produce		json
//	@Param			searchstring query string		false	"Search string name"
//	@Success		200 {object}	models.GetPillsResponse
//	@Failure		500	{object}	models.MessageResponse
//	@Router			/pills [get]
func (ac *AuthController) GetPills(c *gin.Context) {
	commongin.Ok(c, models.GetPillsResponse{
		Pills: []models.Pill{
			{
				ID:          1,
				Name:        "Нурофен-ЭКСТРА",
				Dosage:      "200 мг",
				Instruction: "Инструкция Нурофен-ЭКСТРА",
			},
			{
				ID:          2,
				Name:        "Тирозол",
				Dosage:      "5 мг",
				Instruction: "Инструкция Тирозол",
			},
		},
	})
}

// GetPillByID godoc
//
//	@Summary		GetPillByID
//	@Description	Get pills with filter by name.
//	@Accept			json
//	@Produce		json
//	@Param			pill_id	path	int		true	"Pill Id"
//	@Success		200 {object}	models.Pill
//	@Failure		500	{object}	models.MessageResponse
//	@Router			/pills/:pill_id [get]
func (ac *AuthController) GetPillByID(c *gin.Context) {
	commongin.Ok(c, models.Pill{
		ID:          2,
		Name:        "Тирозол",
		Dosage:      "5 мг",
		Instruction: "Инструкция Тирозол",
	})
}
