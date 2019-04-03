package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"webapp-course-be/routes"

	"github.com/gorilla/handlers"
)

func main() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r := routes.CreateRouter()

	fmt.Println("Listening on port 6767")
	log.Fatal(http.ListenAndServe(":6767", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

/* Notes:
 * Endpoints for:
	- Insert, getAllProducts, getOneProduct, deleteProduct, updateProduct, uploadImgToProduct(/upload-file, save in upload folder and
		update the img name in db).

	respObj:
		- status
		- code: 200 0 404
		- data

*/
