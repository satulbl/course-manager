package app

import (
	"context"
	"course-manager/internal/config"
	"course-manager/internal/handler"
	"course-manager/internal/pkg/logger"
	"course-manager/internal/pkg/postgres"
	"course-manager/internal/pkg/redis"
	"course-manager/internal/repository"
	courserepo "course-manager/internal/repository/course"
	lessonrepo "course-manager/internal/repository/lesson"
	userlessonrepo "course-manager/internal/repository/user_lesson"

	// mentorrepo "course-manager/internal/repository/mentor"
	sessionrepo "course-manager/internal/repository/session"
	tokenrepo "course-manager/internal/repository/token"
	userrepo "course-manager/internal/repository/user"
	usercourserepo "course-manager/internal/repository/user_course"
	"course-manager/internal/service"
	authservice "course-manager/internal/service/auth"
	courseservice "course-manager/internal/service/course"
	lessonservice "course-manager/internal/service/lesson"
	userservice "course-manager/internal/service/user"
	usercourseservice "course-manager/internal/service/user_course"
	userlessonservice "course-manager/internal/service/user_lesson"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() (err error) {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	log, closeLog, err := logger.New(logger.Config{
		Env:   cfg.App.Env,
		Level: cfg.App.LogLevel,
		File:  cfg.App.LogFile,
	})
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := closeLog(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close log file: %w", closeErr))
		}
	}()
	slog.SetDefault(log)

	defer func() {
		if err != nil {
			slog.Error("app stopped", "error", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	postgresPool, err := postgres.New(ctx, postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		Database: cfg.Postgres.DataBase,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		return err
	}
	defer postgresPool.Close()

	redisClient, err := redis.New(ctx, redis.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := redisClient.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close redis: %w", closeErr))
		}
	}()

	jwtManager, err := authservice.NewJWTManager(authservice.JWTConfig{
		Secret:     []byte(cfg.JWT.Secret),
		Issuer:     cfg.JWT.Issuer,
		AccessTTL:  cfg.JWT.AccessTTL,
		RefreshTTL: cfg.JWT.RefreshTTL,
	})
	if err != nil {
		return err
	}

	repos := &repository.Repositories{
		User:    userrepo.New(postgresPool),
		Session: sessionrepo.New(postgresPool),
		Token:   tokenrepo.New(redisClient),
		Course:  courserepo.New(postgresPool),
		// Mentor:     mentorrepo.New(postgresPool),
		Lesson:     lessonrepo.New(postgresPool),
		UserCourse: usercourserepo.New(postgresPool),
		UserLesson: userlessonrepo.New(postgresPool),
	}

	services := &service.Services{
		Auth:       authservice.New(repos.User, repos.Session, repos.Token, jwtManager),
		Course:     courseservice.New(repos.Course),
		User:       userservice.New(repos.User),
		Lesson:     lessonservice.New(repos.Lesson),
		UserCourse: usercourseservice.New(repos.UserCourse),
		UserLesson: userlessonservice.New(repos.UserLesson),
	}

	server := &http.Server{
		Addr:         ":" + cfg.HTTP.Port,
		Handler:      handler.NewRouter(services),
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	return serve(ctx, server, cfg.HTTP.ShutdownTimeout)
}

func serve(ctx context.Context, server *http.Server, shutdownTimeout time.Duration) error {
	serverErr := make(chan error, 1)
	go func() {
		slog.Info("server started", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		return fmt.Errorf("listen and serve: %w", err)
	case <-ctx.Done():
	}

	slog.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}
	return nil
}
