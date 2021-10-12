package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

/*Each recipe should have a name, a list of ingredients, a list of instructions or steps,
and a publication date. Moreover, each recipe belongs to a set of categories or tags (for example
vegan, Italian, pastry, salads, and so on), as well an ID, which is unique identifier to
differentiate each recipe in the database. Also specify the tags on each field using backtick
annotation; for example, `json:"NAME"`. This allows us to map each field to a different name when
we send them as response, since JSON and GO have different naming conventions.
*/

var recipes []Recipe
func init() {
	recipes = make([]Recipe, 0)
}

type Recipe struct {
	ID 			 string  	`json:"id"`
	Name 		 string 	`json:"name"`
	Tags 		 []string 	`json:"tags"`
	Ingredients  []string   `json:"ingredients"`
	Instructions []string   `json:"instructions"`
	PublishedAt  time.Time  `json:"publishedAt"`
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()

	recipes = append(recipes, recipe)

	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.Run()
}