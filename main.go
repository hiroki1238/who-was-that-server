package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	setupRoutes(e)

	// Check if running in Lambda
	if isLambda() {
		lambda.Start(lambdaHandler(e))
	} else {
		// Local development server
		e.Logger.Fatal(e.Start(":8080"))
	}
}

func setupRoutes(e *echo.Echo) {
	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "healthy",
			"message": "Who Was That API is running!",
			"version": "1.0.0",
		})
	})

	// API routes group
	api := e.Group("/api")

	// Sample protected route
	api.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello from Who Was That API!",
		})
	})

	// Auth routes group
	auth := e.Group("/auth")

	// Sample auth routes
	auth.POST("/login", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login endpoint - Coming soon!",
		})
	})

	auth.POST("/register", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Register endpoint - Coming soon!",
		})
	})
}

func isLambda() bool {
	return os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != ""
}

func lambdaHandler(e *echo.Echo) func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// This is a simplified version
		// In production, you would need proper request/response conversion
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"message": "Lambda function is working!"}`,
		}, nil
	}
}
