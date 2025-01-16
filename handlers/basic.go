package handlers

import (
	"example/api-template/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var basics = []models.Basic{}

func GetBasics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, basics)
}

func SearchDetail(c *gin.Context) {
	detail := c.Query("detail")
	var results []models.Basic
	for _, basic := range basics {
		if basic.Detail == detail || strings.Contains(basic.Detail, detail) {
			results = append(results, basic)
		}
	}
	if len(results) > 0 {
		c.IndentedJSON(http.StatusOK, results)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Detail not found"})
	}
}

func RedirectDetails(c *gin.Context) {
	c.Redirect(301, "/basics")
}

func CreateBasic(c *gin.Context) {
	var newBasic models.Basic
	if err := c.BindJSON(&newBasic); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newBasic.ID = GenerateID()
	fmt.Println(newBasic)
	basics = append(basics, newBasic)
	c.IndentedJSON(http.StatusCreated, newBasic)
}

func BasicInfo(c *gin.Context) {
	id := c.Param("id")
	for i, basic := range basics {
		if basic.ID == id {
			basics[i].Visits += 1
			c.IndentedJSON(http.StatusOK, basics[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Basic not found"})
}

//	func uploadFile(c *gin.Context) {
//	    file, err := c.FormFile("file")
//	    if err != nil {
//	        c.JSON(400, gin.H{"error": err.Error()})
//	        return
//	    }
//	    err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
//	    if err != nil {
//	        c.JSON(500, gin.H{"error": "Failed to upload file"})
//	        return
//	    }
//	    c.JSON(200, gin.H{"status": "File uploaded successfully", "filename": file.Filename})
//	}
