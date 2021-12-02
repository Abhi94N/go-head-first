package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"rsc.io/quote"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlubums(c *gin.Context) {
	var newAlbum album

	//Call BindJSON to bind the recieved JSON to
	// newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func updateAlbum(c *gin.Context) {
	var updateAlbum album
	id := c.Param("id")
	//Call BindJSON to bind the recieved JSON to
	// newAlbum

	if err := c.BindJSON(&updateAlbum); err != nil {
		return
	}
	fmt.Println(updateAlbum)
	for i, a := range albums {
		if a.ID == id {
			albums[i] = updateAlbum
			c.IndentedJSON(http.StatusOK, gin.H{"data": updateAlbum})
			return
		}
	}
}

//Patch
func patchAlbum(c *gin.Context) {
	var updateAlbum album
	id := c.Param("id")
	//Call BindJSON to bind the recieved JSON to
	// newAlbum

	if err := c.BindJSON(&updateAlbum); err != nil {
		return
	}
	fmt.Println(updateAlbum)
	for i, a := range albums {
		if a.ID == id {
			// albums[i] = updateAlbum
			if updateAlbum.Title != "" {
				albums[i].Title = updateAlbum.Title
			}
			if updateAlbum.Artist != "" {
				albums[i].Artist = updateAlbum.Artist
			}

			if updateAlbum.Price != 0 {
				albums[i].Price = updateAlbum.Price
			}
			c.IndentedJSON(http.StatusCreated, gin.H{"data": albums[i], "index": i})
			return
		}
	}
}

func remove(slice []album, s int) []album {
	return append(slice[:s], slice[s+1:]...)
}

//Delete
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	//Call BindJSON to bind the recieved JSON to
	// newAlbum
	for i, a := range albums {
		if a.ID == id {
			tmpAlbums := albums
			albums = remove(tmpAlbums, i)
			fmt.Println(albums)
			c.IndentedJSON(http.StatusOK, gin.H{"data": albums[i], "deleted": true})
			return
		}
	}
}

func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbum)
	router.PATCH("/albums/:id", patchAlbum)
	router.DELETE("/albums/delete/:id", DeleteAlbum)

	router.Run("localhost:8080")
}
