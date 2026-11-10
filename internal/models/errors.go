package models

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrForbidden          = errors.New("forbidden")
	ErrInvalidInput       = errors.New("invalid input")
	ErrEmailTaken         = errors.New("email already registered")
	ErrLessonNumberTaken  = errors.New("lesson number already taken in this course")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidToken       = errors.New("invalid or expired token")

	ErrFileTooLarge        = errors.New("file too large")
	ErrUnsupportedFileType = errors.New("unsupported file type")
	ErrFileNotAvailable    = errors.New("file not available or not found")

	ErrUserAlreadyEnrolled = errors.New("user already enrolled in course")
	ErrCourseNotFound      = errors.New("course not found")
	ErrAuthorNotFound      = errors.New("author not found")
	ErrMentorNotFound      = errors.New("mentor not found")
	ErrUserNotFound        = errors.New("user not found")
)
