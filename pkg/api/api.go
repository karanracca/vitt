package api

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"vitt/pkg/store"

// 	"github.com/urfave/cli/v2"
// )

// func getTransactions(ctx *cli.Context, w http.ResponseWriter, r *http.Request) {

// 	db := ctx.Context.Value("db").(*store.Store)

// 	transactions, err := db.GetTransactions()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Convert the struct to JSON
// 	jsonResponse, err := json.Marshal(transactions)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	//Allow CORS here By * or specific origin
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	// Set the content type as application/json
// 	w.Header().Set("Content-Type", "application/json")
// 	// Write the JSON response
// 	w.Write(jsonResponse)

// }

// func Init(ctx *cli.Context) error {

// 	http.HandleFunc("/api/transactions", func(w http.ResponseWriter, r *http.Request) {
// 		getTransactions(ctx, w, r)
// 	})

// 	// Start the server on port 8080
// 	log.Println("Server listening on port 8080...")
// 	return http.ListenAndServe(":8080", nil)
// }
