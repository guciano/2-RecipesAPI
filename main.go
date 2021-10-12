package main

import (
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
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

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not Found"})
		return
	}
	recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)
}

func init() {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.Run()
}