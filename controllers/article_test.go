package controllers

import(
	"testing"
	"github.com/Qwerci/article/models"
)


// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := GetAllArticles()

	// check that the lenth of the list of articles returned is the 
	// same as the lenght of the dlobal variable holdinging the list

	for i, v := range alist {
		if v.Content != models.ArticleList[i].Content || v.Title != models.ArticleList[i].Title {
			t.Fail()
			break
		}
	}
}