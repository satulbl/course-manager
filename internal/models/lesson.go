package models

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

const (
	titleMaxLength       = 80
	descriptionMaxLength = 150
)

type LessonCourse struct {
    ID    uuid.UUID
    Title string
}

type LessonMentor struct {
    ID   uuid.UUID
    Name string
}

type Lesson struct {
    ID           uuid.UUID
    Title        string
    Description  string
    Content      string
    LessonNumber int
    Course       LessonCourse
    Mentor       LessonMentor
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

type LessonInput struct {
	Title        string
	Description  string
	Content      string
	LessonNumber int
	CourseID     uuid.UUID
}

func (in *LessonInput) Validate() error {
	if in.Title == "" {
		return fmt.Errorf("%w: title is required", ErrInvalidInput)
	}
	if utf8.RuneCountInString(in.Title) > titleMaxLength {
		return fmt.Errorf("%w: title must be at most %d characters", ErrInvalidInput, titleMaxLength)
	}
	if in.Description == "" {
		return fmt.Errorf("%w: description is required", ErrInvalidInput)
	}
	if utf8.RuneCountInString(in.Description) > descriptionMaxLength {
		return fmt.Errorf("%w: description must be at most %d characters", ErrInvalidInput, descriptionMaxLength)
	}
	if in.Content == "" {
		return fmt.Errorf("%w: content is required", ErrInvalidInput)
	}
	if in.LessonNumber <= 0 {
		return fmt.Errorf("%w: lesson number must be positive", ErrInvalidInput)
	}
	if in.CourseID == uuid.Nil {
		return fmt.Errorf("%w: course id is required", ErrInvalidInput)
	}
	return nil
}

