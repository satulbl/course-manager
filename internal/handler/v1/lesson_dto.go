package v1

import (
	"course-manager/internal/models"
	"time"

	"github.com/google/uuid"
)

type lessonRequest struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Content      string    `json:"content"`
	LessonNumber int       `json:"lessonNumber"`
	CourseID     uuid.UUID `json:"courseId"`
}

func (l *lessonRequest) toInput() models.LessonInput {
	return models.LessonInput{
		Title:        l.Title,
		Description:  l.Description,
		Content:      l.Content,
		LessonNumber: l.LessonNumber,
		CourseID:     l.CourseID,
	}
}

type createLessonResponse struct {
	ID uuid.UUID `json:"id"`
}

func newCreateLessonResponse(id uuid.UUID) createLessonResponse {
	return createLessonResponse{ID: id}
}

type lessonCourseResponse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type lessonMentorResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type lessonResponse struct {
	ID           uuid.UUID            `json:"id"`
	Title        string               `json:"title"`
	Description  string               `json:"description"`
	Content      string               `json:"content"`
	LessonNumber int                  `json:"lessonNumber"`
	Course       lessonCourseResponse `json:"course"`
	Mentor       lessonMentorResponse `json:"mentor"`
	CreatedAt    time.Time            `json:"createdAt"`
	UpdatedAt    time.Time            `json:"updatedAt"`
}

func newLessonResponse(l *models.Lesson) lessonResponse {
	return lessonResponse{
		ID:           l.ID,
		Title:        l.Title,
		Description:  l.Description,
		Content:      l.Content,
		LessonNumber: l.LessonNumber,
		Course: lessonCourseResponse{
			ID:    l.Course.ID,
			Title: l.Course.Title,
		},
		Mentor: lessonMentorResponse{
			ID:   l.Mentor.ID,
			Name: l.Mentor.Name,
		},
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}

type lessonListItemResponse struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	LessonNumber int       `json:"lessonNumber"`
}

func newLessonListItemResponse(l models.Lesson) lessonListItemResponse {
	return lessonListItemResponse{
		ID:           l.ID,
		Title:        l.Title,
		Description:  l.Description,
		LessonNumber: l.LessonNumber,
	}
}

type lessonListResponse struct {
	Items      []lessonListItemResponse `json:"items"`
	TotalCount int64                    `json:"total"`
}

func newLessonListResponse(lessons []models.Lesson, total int64) lessonListResponse {
	items := make([]lessonListItemResponse, 0, len(lessons))
	for _, l := range lessons {
		items = append(items, newLessonListItemResponse(l))
	}
	return lessonListResponse{Items: items, TotalCount: total}
}
