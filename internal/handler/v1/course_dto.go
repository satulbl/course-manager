package v1

import (
	"course-manager/internal/models"
	"time"

	"github.com/google/uuid"
)

type courseRequest struct {
	Title       string    `json:"title"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	MentorID    uuid.UUID `json:"mentor_id"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

func (req *courseRequest) toInput() models.CourseInput {
	return models.CourseInput{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		MentorID:    req.MentorID,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}
}

type courseResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	AuthorID    uuid.UUID `json:"authorId"`
	MentorID    uuid.UUID `json:"mentorID"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type coursesListResponse struct {
	Courses    []courseResponse `json:"courses"`
	TotalCount int64            `json:"totalCount"`
}

func newCourseResponse(c *models.Course) courseResponse {
	return courseResponse{
		ID:          c.ID,
		Title:       c.Title,
		Price:       c.Price,
		Description: c.Description,
		AuthorID:    c.AuthorID,
		MentorID:    c.MentorID,
		StartDate:   c.StartDate,
		EndDate:     c.EndDate,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func newCourseListResponse(courses []models.Course, total int64) coursesListResponse {
	items := make([]courseResponse, 0, len(courses))
	for _, c := range courses {
		items = append(items, newCourseResponse(&c))
	}

	return coursesListResponse{
		Courses:    items,
		TotalCount: total,
	}
}
