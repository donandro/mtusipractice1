package controllers

import (
	"pillsreminder/apps/mobileapi/api/rest/models"
	commongin "pillsreminder/pkg/rest/gin"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type PlansController struct {
	logger zerolog.Logger
}

func NewPLansController(logger zerolog.Logger) *AuthController {
	return &AuthController{
		logger: logger,
	}
}

// GetPlans godoc
//
//	@Summary		GetPills
//	@Description	Get pills with filter by name.
//	@Accept			json
//	@Produce		json
//	@Success		200 {object}	models.GetPlansResponse
//	@Failure		500	{object}	models.MessageResponse
//	@Router			/plans [get]
func (ac *AuthController) GetPlans(c *gin.Context) {
	commongin.Ok(c, models.GetPlansResponse{
		Plans: []models.Plan{
			{
				ID:           1,
				Name:         "Приём Тирозола",
				ReminderTime: "20:00",
				Frequency:    "daily",
				DaysOfWeek:   "Mon,Wed,Fri",
				StartDate:    "2024-09-01",
				EndDate:      "2024-09-30",
				Pill: models.Pill{
					ID:          2,
					Name:        "Тирозол",
					Dosage:      "5 мг",
					Instruction: "Инструкция Тирозол",
				},
			},
			{
				ID:           2,
				Name:         "Приём нурофена",
				ReminderTime: "20:00",
				Frequency:    "daily",
				DaysOfWeek:   "Mon,Wed,Fri",
				StartDate:    "2024-09-01",
				EndDate:      "2024-09-30",
				Pill: models.Pill{
					ID:          1,
					Name:        "Нурофен-ЭКСТРА",
					Dosage:      "200 мг",
					Instruction: "Инструкция Нурофен-ЭКСТРА",
				},
			},
		},
	})
}

// GetPlanByID godoc
//
//	@Summary		GetPlanByID
//	@Description	Get plan by id.
//	@Accept			json
//	@Produce		json
//	@Param			plan_id	path	int		true	"Plan Id"
//	@Success		200 {object}	models.Plan
//	@Failure		500	{object}	models.MessageResponse
//	@Router			/plans/:plan_id [get]
func (ac *AuthController) GetPlanByID(c *gin.Context) {
	commongin.Ok(c, models.Plan{
		ID:           1,
		Name:         "Приём Тирозола",
		ReminderTime: "20:00",
		Frequency:    "daily",
		DaysOfWeek:   "Mon,Wed,Fri",
		StartDate:    "2024-09-01",
		EndDate:      "2024-09-30",
		Pill: models.Pill{
			ID:          2,
			Name:        "Тирозол",
			Dosage:      "5 мг",
			Instruction: "Инструкция Тирозол",
		},
	})
}

// AddPlan godoc
//
//	@Summary		AddPlan
//	@Description	Get plan by id.
//	@Accept			json
//	@Produce		json
//	@Param			plan_id	path	int		true	"Plan Id"
//	@Success		200 {object}	models.Plan
//	@Failure		500	{object}	models.MessageResponse
//	@Router			/plans [post]
func (ac *AuthController) AddPlan(c *gin.Context) {
	commongin.Ok(c, models.Plan{
		ID:           2,
		Name:         "Приём нурофена",
		ReminderTime: "20:00",
		Frequency:    "daily",
		DaysOfWeek:   "Mon,Wed,Fri",
		StartDate:    "2024-09-01",
		EndDate:      "2024-09-30",
		Pill: models.Pill{
			ID:          1,
			Name:        "Нурофен-ЭКСТРА",
			Dosage:      "200 мг",
			Instruction: "Инструкция Нурофен-ЭКСТРА",
		},
	})
}

// DeletePlan godoc
//
//	@Summary		AddPlan
//	@Description	Get plan by id.
//	@Accept			json
//	@Produce		json
//	@Param			plan_id	path	int		true	"Plan Id"
//	@Success		200
//	@Failure		500	{object}	models.MessageResponse
//	@Router			/plans [delete]
func (ac *AuthController) DeletePlan(c *gin.Context) {
	commongin.Ok(c, models.Plan{
		ID:           2,
		Name:         "Приём нурофена",
		ReminderTime: "20:00",
		Frequency:    "daily",
		DaysOfWeek:   "Mon,Wed,Fri",
		StartDate:    "2024-09-01",
		EndDate:      "2024-09-30",
		Pill: models.Pill{
			ID:          1,
			Name:        "Нурофен-ЭКСТРА",
			Dosage:      "200 мг",
			Instruction: "Инструкция Нурофен-ЭКСТРА",
		},
	})
}
