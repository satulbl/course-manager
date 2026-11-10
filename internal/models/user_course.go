package models

import (
	"fmt"

	"github.com/google/uuid"
)

type ApplicationCourse struct {
	ID    uuid.UUID
	Title string
}

type ApplicationUser struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type UserCourse struct {
	ID         uuid.UUID
	User       ApplicationUser
	Course     ApplicationCourse
	IsApproved bool
}

type UserCourseFilter struct {
	UserID     *uuid.UUID
	IsApproved *bool
}

type UserCourseInput struct {
	UserID   uuid.UUID
	CourseID uuid.UUID
}

func (in UserCourseInput) Validate() error {
	if in.CourseID == uuid.Nil {
		return fmt.Errorf("%w: course id is required", ErrInvalidInput)
	}
	if in.UserID == uuid.Nil {
		return fmt.Errorf("%w: user id is required", ErrInvalidInput)
	}

	return nil
}
