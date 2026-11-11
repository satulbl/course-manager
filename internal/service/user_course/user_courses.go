package usercourse

import (
	"context"
	"course-manager/internal/models"
	"fmt"
)

func (s *Service) GetUserCourses(ctx context.Context, user models.CurrentUser, f models.UserCourseFilter) ([]models.UserCourse, error) {
	if f.IsApproved == nil && f.UserID == nil && !user.IsAdmin() {
		return nil, models.ErrForbidden
	}

	userCourses, err := s.userCourses.GetUserCourses(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("get user courses: %w", err)
	}

	return userCourses, nil
}
