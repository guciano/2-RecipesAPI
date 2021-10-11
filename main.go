package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.RUN()
}

/*Each recipe should have a name, a list of ingredients, a list of instructions or steps,
and a publication date. Moreover, each recipe belongs to a set of categories or tags (for example
vegan, Italian, pastry, salads, and so on), as well an ID, which is unique identifier to
differentiate each recipe in the database. Also specify the tags on each field using backtick
annotation; for example, `json:"NAME"`. This allows us to map each field to a different name when
we send them as response, since JSON and GO have different naming conventions.
*/

type Recipe struct {
	Name 		 string 	`json:"name"`
	Tags 		 []string 	`json:"tags"`
	Ingredients  []string   `json:"ingredients"`
	Instructions []string   `json:"instructions"`
	PublishedAt  time.Time  `json: publishedAt`
}