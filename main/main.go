package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/omarmendozaaa/backend-go/server"
)

func main() {
	s := server.New()
	fmt.Println("Welcome . . .")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}

// package main

// import (
// 	"fmt"

// 	"github.com/omarmendozaaa/backend-go/server"
// )

// var observations []server.Node = []server.Node{
// 	{20.0, 20.0, 20.0, 20.0},
// 	{21.0, 21.0, 21.0, 21.0},
// 	{100.5, 100.5, 100.5, 100.5},
// 	{50.1, 50.1, 50.1, 50.1},
// 	{64.2, 64.2, 64.2, 64.2},
// }

// func main() {
// 	// Get a list of centroids and output the values
// 	_, centroids := server.Train(observations, 3, 50)
// 	// Show the centroids
// 	fmt.Println("The centroids are")
// 	for _, centroid := range centroids {
// 		fmt.Println(centroid)
// 	}
// 	// Output the clusters
// 	fmt.Println("...")
// 	for _, observation := range observations {
// 		index := server.Nearest(observation, centroids)
// 		fmt.Println(observation, "tbelongs in cluser", index+1, ".")
// 	}
// }
