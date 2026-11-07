package models

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID
	Title       string
	Price       int
	Description string
	AuthorID    uuid.UUID
	MentorID    uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CourseInput struct {
	Title       string
	Price       int
	Description string
	MentorID    uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
}

const (
	maxCourseTitleLen       = 150
	maxCourseDescriptionLen = 5000
)

func (in *CourseInput) Validate() error {
	in.Title = strings.TrimSpace(in.Title)
	in.Description = strings.TrimSpace(in.Description)

	if in.Title == "" {
		return fmt.Errorf("%w: title is required", ErrInvalidInput)
	}
	if utf8.RuneCountInString(in.Title) > maxCourseTitleLen {
		return fmt.Errorf("%w: title must be at most %d characters", ErrInvalidInput, maxCourseTitleLen)
	}
	if in.Description == "" {
		return fmt.Errorf("%w: description is required", ErrInvalidInput)
	}
	if utf8.RuneCountInString(in.Description) > maxCourseDescriptionLen {
		return fmt.Errorf("%w: description must be at most %d characters", ErrInvalidInput, maxCourseDescriptionLen)
	}
	if in.Price < 0 {
		return fmt.Errorf("%w: price must not be negative", ErrInvalidInput)
	}
	if in.MentorID == uuid.Nil {
		return fmt.Errorf("%w: mentor id is required", ErrInvalidInput)
	}
	if in.StartDate.IsZero() || in.EndDate.IsZero() {
		return fmt.Errorf("%w: start date and end date are required", ErrInvalidInput)
	}
	if in.EndDate.Before(in.StartDate) {
		return fmt.Errorf("%w: end date must not be before start date", ErrInvalidInput)
	}
	return nil
}

func (c *Course) Edit(in CourseInput) error {
	if err := in.Validate(); err != nil {
		return fmt.Errorf("edit course: %w", err)
	}

	c.Title = in.Title
	c.Price = in.Price
	c.Description = in.Description
	c.MentorID = in.MentorID
	c.StartDate = in.StartDate
	c.EndDate = in.EndDate
	c.UpdatedAt = time.Now().UTC()

	return nil
}