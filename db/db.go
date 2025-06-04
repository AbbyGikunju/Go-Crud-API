package db

import (
	//"errors"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Movie struct {
	ID			string		`json:"id" gorm:"primarykey"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
}

func INITPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var(
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		dbUser = os.Getenv("DB_USER")
		dbName = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	host,
	port,
	dbUser,
	dbName,
	password,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database",err)
	}
	DB.AutoMigrate(Movie{})

}

//Create movie operation
func CreateMovie(movie *Movie) (*Movie, error) {
	movie.ID = uuid.New().String()
	res := DB.Create(&movie)
	if res.Error != nil {
		return nil, res.Error
	}
	return movie,nil
}

//Read movie operation by id
func GetMovie(id string) (*Movie, error) {
	var movie Movie
	res := DB.First(&movie, "id= ?",id)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("movie of id %s not found", id)
	}
	return &movie, nil
}


//Read all movies operation
func GetMovies() ([]*Movie, error) {
	var movies []*Movie
	res := DB.Find(&movies)
	if res.Error != nil {
		return nil,errors.New("no movies found")
	}
	return movies, nil
}

//Update movie operation
func UpdateMovie(movie *Movie) (*Movie, error) {
	var movieToUpdate Movie
	result := DB.Model(&movieToUpdate).Where("id = ?", movie.ID).Updates(movie)
	if result.RowsAffected == 0{
		return &movieToUpdate, errors.New("movie not updated")
	}
	return movie, nil
}

//Delete Movie Operation
func DeleteMovie(id string) error {
var deletedMovie Movie
result := DB.Where("id = ?",id).Delete(&deletedMovie)
if result.RowsAffected == 0 {
	return errors.New("movie not deleted")
	}
	return nil
}