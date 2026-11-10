package v1

import (
	"course-manager/internal/models"

	"github.com/google/uuid"
)

type createUserCourseRequest struct {
	UserID   uuid.UUID `json:"userId"`
	CourseID uuid.UUID `json:"courseId"`
}

func (req *createUserCourseRequest) toInput() models.UserCourseInput {
	return models.UserCourseInput{
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}
}

type applicationCourse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type applicationUser struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type userCourseResponse struct {
	ID         uuid.UUID         `json:"id"`
	User       applicationUser   `json:"user"`
	Course     applicationCourse `json:"course"`
	IsApproved bool              `json:"isApproved"`
}

func newUserCourseResponse(uc models.UserCourse) userCourseResponse {
	return userCourseResponse{
		ID: uc.ID,
		User: applicationUser{
			ID:    uc.User.ID,
			Name:  uc.User.Name,
			Email: uc.User.Email,
		},
		Course: applicationCourse{
			ID:    uc.Course.ID,
			Title: uc.Course.Title,
		},
		IsApproved: uc.IsApproved,
	}
}

func newUserCoursesResponse(items []models.UserCourse) []userCourseResponse {
	res := make([]userCourseResponse, 0, len(items))
	for _, uc := range items {
		res = append(res, newUserCourseResponse(uc))
	}
	return res
}
