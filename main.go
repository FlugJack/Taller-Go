package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// getAlbums responds with the list of all albums as JSON.

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("album.sqlite"), &gorm.Config{})

	if err != nil {
		return
	}

	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	db.Create(&newAlbum)
	c.IndentedJSON(200, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	db, err := gorm.Open(sqlite.Open("album.sqlite"), &gorm.Config{})

	if err != nil {
		return
	}

	var album album
	err = db.First(&album, id).Error

	if err != nil {
		return
	}

	c.IndentedJSON(200, &album)
}
func getAlbums(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("album.sqlite"), &gorm.Config{})
	if err != nil {
		return
	}

	var albums []album
	result := db.Find(&albums)

	if result.Error != nil {
		return
	}
	c.IndentedJSON(200, albums)

}

func deletAlbumByID(c *gin.Context) {
	id := c.Param("id")
	db, err := gorm.Open(sqlite.Open("album.sqlite"), &gorm.Config{})

	if err != nil {
		return
	}

	err = db.Delete(&album{}, id).Error

	c.IndentedJSON(200, gin.H{"mensaje": "Elementos borrados correctamente"})

}

func main() {

	db, err := gorm.Open(sqlite.Open("album.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&album{})

	router := gin.Default()
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deletAlbumByID)
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
