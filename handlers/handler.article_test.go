package handlers

import(
	"io"
	"strings"
	"net/http"
	"net/http/httptest"
	"testing"
	
)

// Test that a GET request to the home page returns the home page with 
// the HTTP cod 2000 for  an unauthenticated user.

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := GetRouter(true)

	r.GET("/", ShowIndexPage)


	// Create a request to send to the above route
	req,_ := http.NewRequest("GET", "/", nil)

	TestHTTPResponse(t, r, req, 
		func(w *httptest.ResponseRecorder) bool{
			// Test that the http status code is 200
			statusOK := w.Code ==http.StatusOK

			// Test that the page title is "Home Page"
			p, err := io.ReadAll(w.Body)
			pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>" ) > 0

			return statusOK && pageOK
		})

}

