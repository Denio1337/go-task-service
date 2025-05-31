package task

import (
	"app/storage"
	"math"
)

func GetDates(params *GetDatesParams) (*GetDatesResult, error) {
	// Calculate offset and limit for pagination
	offset := (params.Page - 1) * params.PageSize
	limit := params.PageSize

	// Get dates from storage
	result, err := storage.GetDates(&storage.GetDatesParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	// Calculate total pages based on total tasks and page size
	var totalPages uint
	switch result.Total {
	case 0:
		totalPages = 1 // If no tasks, we still return one page
	default:
		totalPages = uint(math.Ceil(float64(result.Total) / float64(params.PageSize)))
	}

	// Check if requested page exceeds total pages
	if params.Page > totalPages {
		return nil, ErrExceedsPageCount
	}

	return &GetDatesResult{
		Dates:      result.Dates,
		TotalPages: totalPages,
	}, nil
}
