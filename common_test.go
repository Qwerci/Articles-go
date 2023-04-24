package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/Qwerci/article/models"
	"github.com/gin-gonic/gin"
)

var tmpArticleList []models.Article

// This function is used for setup before executing the test functions

func TestMain(m *testing.M){
	// sSet Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("template/*")
	}
	return r
}

func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, 
	f func(w *httptest.ResponseRecorder) bool) {

		// Create a response recorder
		w := httptest.NewRecorder()

		// Create the service and process the above request
		r.ServeHTTP(w, req)

		if !f(w){
			t.Fail()
		}
	}

// This function is used to store the main lists into the temporary one
func SaveLists(){
	tmpArticleList = models.ArticleList
}

// This function is used to restore the main lists from the temporary one

func RestoreLists(){
	models.ArticleList = tmpArticleList
}