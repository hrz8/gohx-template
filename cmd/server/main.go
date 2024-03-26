package main

import (
	"embed"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/hrz8/gohx/assets"
	"github.com/hrz8/gohx/internal/pkg/static"
	"github.com/hrz8/gohx/view/error_page"
	"github.com/hrz8/gohx/view/page"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.RouteNotFound("*", ErrorHandler)
	e.GET("/", HomeHandler)
	e.GET("/login", LoginHandler)
	e.GET("/assets/*", StaticHandler(assets.PublicFiles))
	e.Logger.Fatal(e.Start(":4000"))
}

func HomeHandler(c echo.Context) error {
	time.Sleep(500 * time.Millisecond)
	c.Response().Header().Set("HX-Push-Url", "/")
	return page.HomePage("World").Render(c.Request().Context(), c.Response().Writer)
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
