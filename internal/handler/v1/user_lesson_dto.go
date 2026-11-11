package v1

import (
	"course-manager/internal/models"

	"github.com/google/uuid"
)

type userLessonRequest struct {
	LessonID uuid.UUID `json:"lessonId"`
}

func (ul userLessonRequest) toInput() models.UserLessonInput {
	return models.UserLessonInput{
		LessonID: ul.LessonID,
	}
}
