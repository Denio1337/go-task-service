package task

import (
	cerror "app/router/types/error"
	"app/router/types/response"
	"app/router/validator"
	"app/service/task"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GET /tasks/dates
func GetDates(c *fiber.Ctx) error {
	// Extract pagination parameters from query
	page := uint(c.QueryInt(PageName, DefaultPage))
	pageSize := uint(c.QueryInt(PageSizeName, DefaultPageSize))

	// Extract date parameter from query
	date, err := getDateParam(c, DateName, time.Time{})
	if err != nil {
		return ErrInvalidDate
	}

	// Validate extracted query parameters
	dto := &GetDatesDTO{
		Page:     page,
		PageSize: pageSize,
		Date:     date,
	}
	if errs := validator.Validate(dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Route to service
	result, err := task.GetDates(&task.GetDatesParams{
		Page:     dto.Page,
		PageSize: dto.PageSize,
		Date:     dto.Date,
	})

	// Handle error from service
	if err != nil {
		return err
	}

	return c.JSON(response.SuccessResponse(&TaskDatesResponse{
		Dates: result.Dates,
		Meta: Meta{
			Page:       page,
			PageSize:   pageSize,
			TotalPages: result.TotalPages,
		},
	}))
}
