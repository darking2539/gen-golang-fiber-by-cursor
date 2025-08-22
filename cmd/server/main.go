package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/user/gen-golang-fiber-by-cursor/docs"
	appPkg "github.com/user/gen-golang-fiber-by-cursor/internal/app"
	domainUser "github.com/user/gen-golang-fiber-by-cursor/internal/domain/user"
	httpHandler "github.com/user/gen-golang-fiber-by-cursor/internal/handler/http"
	"github.com/user/gen-golang-fiber-by-cursor/internal/infrastructure/db"
	"github.com/user/gen-golang-fiber-by-cursor/internal/middleware"
	"github.com/user/gen-golang-fiber-by-cursor/internal/repository"
	authUsecase "github.com/user/gen-golang-fiber-by-cursor/internal/usecase/auth"
	profileUsecase "github.com/user/gen-golang-fiber-by-cursor/internal/usecase/profile"
	"gorm.io/gorm"
)

// @title           Go Fiber Clean Architecture API
// @version         1.0
// @description     API documentation for the Go Fiber Clean Architecture sample.
// @BasePath        /
// @schemes         http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Use format: Bearer <token>

func main() {
	server := appPkg.NewServer()

	// Configs
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret"
	}

	// DB init and migration
	database, err := db.Init("app.db", &domainUser.User{})
	if err != nil {
		log.Fatalf("init db: %v", err)
	}

	// Seed default user if not exists
	seedDefaultUser(database)

	// DI
	userRepo := repository.NewUserRepository(database)
	authUC := authUsecase.New(userRepo, jwtSecret)
	profileUC := profileUsecase.New(userRepo)

	// Handlers
	authH := httpHandler.NewAuthHandler(authUC)
	profileH := httpHandler.NewProfileHandler(profileUC)

	// Routes
	r := server.App()
	// Public
	authH.Register(r)

	// Swagger docs
	r.Get("/swagger/*", swagger.HandlerDefault)
	r.Get("/swagger", func(c *fiber.Ctx) error { return c.Redirect("/swagger/index.html", 302) })

	// Protected using JWT middleware
	r.Use(middleware.JWTAuth(
		func(token string) (interface{}, error) {
			return authUC.ParseToken(token)
		},
		func(claims interface{}) (string, error) {
			if rc, ok := claims.(*jwt.RegisteredClaims); ok {
				return rc.Subject, nil
			}
			return "", fiberErr("invalid claims type")
		},
	))

	profileH.Register(r)

	if err := server.Start(":8080"); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}

func seedDefaultUser(gormDB *gorm.DB) {
	var count int64
	gormDB.Model(&domainUser.User{}).Where("username = ?", "demo").Count(&count)
	if count == 0 {
		_ = gormDB.Create(&domainUser.User{
			Username: "demo",
			Password: "password",
			Name:     "Demo User",
			Email:    "demo@example.com",
			Avatar:   "https://i.pravatar.cc/150?img=3",
			Bio:      "Hello, I'm a demo user",
		}).Error
	}
}

func fiberErr(msg string) error { return &fiberError{msg: msg} }

type fiberError struct{ msg string }

func (e *fiberError) Error() string { return e.msg }
