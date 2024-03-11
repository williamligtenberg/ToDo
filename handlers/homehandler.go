package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func HomeHandler(c echo.Context) error {
	// Retrieve all cookies from the request
	cookies := c.Cookies()

	// Pass the cookies to the template
	data := map[string]interface{}{
		"Cookies": cookies,
	}

	// Render the HTML template with the data
	if err := c.Render(http.StatusOK, "home.jet.html", data); err != nil {
		// Handle the error, e.g., log it
		log.Println("Error rendering template:", err)
		return err
	}

	return nil
}
