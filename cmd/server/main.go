package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hrz8/gohx/assets"
	"github.com/hrz8/gohx/internal/pkg/static"
	"github.com/hrz8/gohx/internal/repo/db"
	"github.com/hrz8/gohx/view/error_page"
	"github.com/hrz8/gohx/view/page"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	ctx := context.Background()
	dbConn := connect(ctx)
	defer func() {
		dbConn.Close(ctx)
		log.Println("cleaning up...")

		os.Exit(0)
	}()

	e := echo.New()

	e.RouteNotFound("*", ErrorHandler)
	e.GET("/", HomeHandler(dbConn))
	e.GET("/login", LoginHandler)
	e.GET("/assets/*", StaticHandler(assets.PublicFiles))
	e.Logger.Fatal(e.Start(":4000"))
}

func HomeHandler(dbConn *pgx.Conn) func(c echo.Context) error {
	return func(c echo.Context) error {
		time.Sleep(500 * time.Millisecond)
		q := db.New(dbConn)
		app, err := q.GetApps(c.Request().Context(), uuid.MustParse("7c5ba257-6851-46b3-8a2b-3490bf68a1af"))
		if err == pgx.ErrNoRows {
			c.Response().Header().Set("HX-Push-Url", "/")
			return page.HomePage(nil).Render(c.Request().Context(), c.Response().Writer)
		}

		c.Response().Header().Set("HX-Push-Url", "/")
		return page.HomePage(app).Render(c.Request().Context(), c.Response().Writer)
	}
}

func LoginHandler(c echo.Context) error {
	time.Sleep(500 * time.Millisecond)
	c.Response().Header().Set("HX-Push-Url", "/login")
	return page.LoginPage().Render(c.Request().Context(), c.Response().Writer)
}

func ErrorHandler(c echo.Context) error {
	return error_page.NotFoundPage().Render(c.Request().Context(), c.Response().Writer)
}

func StaticHandler(embedded embed.FS) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Param("*")

		tmpPath, err := url.PathUnescape(path)
		if err != nil {
			return fmt.Errorf("failed to unescape path variable: %w", err)
		}

		path = tmpPath
		name := "public/" + filepath.ToSlash(filepath.Clean(strings.TrimPrefix(path, "/")))
		fileErr := static.ServeFsFile(c, name, embedded)
		if fileErr != nil && errors.Is(fileErr, echo.ErrNotFound) {
			return error_page.NotFoundPage().Render(c.Request().Context(), c.Response().Writer)
		}

		return fileErr
	}
}
