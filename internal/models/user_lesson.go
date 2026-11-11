package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ApplicationLesson struct {
	ID           uuid.UUID
	Title        string
	LessonNumber int
	Course       ApplicationCourse
}

type UserLesson struct {
	ID          uuid.UUID
	User        ApplicationUser
	Lesson      ApplicationLesson
	CompletedAt *time.Time
	CreatedAt   time.Time
}

func (ul UserLesson) IsCompleted() bool {
	return ul.CompletedAt != nil
}

// type UserLessonFilter struct {
// 	UserID      *uuid.UUID
// 	LessonID    *uuid.UUID
// 	CourseID    *uuid.UUID
// 	IsCompleted *bool
// }

type UserLessonInput struct {
	LessonID uuid.UUID
}

func (in UserLessonInput) Validate() error {
	if in.LessonID == uuid.Nil {
		return fmt.Errorf("%w: lesson id is required", ErrInvalidInput)
	}

	return nil
}