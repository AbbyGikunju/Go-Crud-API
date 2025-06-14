package router

import (
	"net/http"
	"go-crud-api/db"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovie)
	r.POST("/movies", postMovie)
	r.PUT("/movies/:id", putMovie)
	r.DELETE("/movies/:id",deleteMovie)
	return r	
}


//posts a movie to the db
func postMovie(ctx *gin.Context) {
	var movie db.Movie
	err := ctx.Bind(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"movie": res,
	})
}


//gets a movie from the db
func getMovies(ctx *gin.Context) {
	res, err := db.GetMovies()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movies": res,
	})
}

func getMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movie": res,
	})
}

//update a movie in the db
func putMovie(ctx *gin.Context) {
	var updatedMovie db.Movie
	err := ctx.Bind(&updatedMovie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbMovie, err := db.GetMovie(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbMovie.Name = updatedMovie.Name
	dbMovie.Description = updatedMovie.Description

	res, err := db.UpdateMovie(dbMovie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"task" : res,
	})
}

//delete a movie from the db
func deleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteMovie(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})

}