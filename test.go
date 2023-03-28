package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// function to get the contents of the posted data
// this should be a json string of activities, colour, milk inputs, etc
func (app *App) get_post_content(w http.ResponseWriter, req *http.Request) {

	// read the body of the request
	body, err := io.ReadAll(req.Body)
	if err != nil {
		// if we hit an error, return this TODO: this should be a json string to handle it correctly on the frontend
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	//upon exiting the function, close the body request
	defer req.Body.Close()

	// declare a variable called data which is a list of string keys with an undefined values (since the json can be anything)
	// var data map[string]interface{}
	// if err := json.Unmarshal(body, &data); err != nil { // here we unflatten the body into the data variable, and capture any errors
	// 	http.Error(w, "Error decoding request body", http.StatusBadRequest)
	// 	return
	// }

	// body_json, err := json.Marshal(body)
	// if err != nil {
	// 	http.Error(w, "Error encoding request body", http.StatusBadRequest)
	// }

	app.insert_data(body)

	// if no errors, then we set the headers in the http response writer
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("data"); err != nil { // here, we make a json encoder that will write to the response writer, and encode the data back into json
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// method of the app struct, used to insert data into the db
func (app *App) insert_data(body_json []byte) {

	// here, we make a struct that will hold the values parsed in json format
	var baby_log struct {
		Baby_id          int      `json:"baby_id"`
		Activities       []string `json:"activities"`
		Colour           string   `json:"colour"`
		Breast_milk_time int      `json:"breast_milk_time"`
		Breast_milk_mls  int      `json:"breast_milk_mls"`
		Formula_milk_mls int      `json:"formula_milk_mls"`
		Time             string   `json:"time"`
	}

	//here we take the json, and put it into our struct
	err := json.Unmarshal(body_json, &baby_log)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare INSERT statement
	stmt, err := app.DB.Prepare("INSERT INTO baby_logs (baby_id, time, activities, color, breast_milk_time, breast_milk_mls, formula_milk_mls) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Set values for the INSERT statement
	baby_id := baby_log.Baby_id

	activities := strings.Join(baby_log.Activities, ",")

	fmt.Printf("activities is %s\n", activities)

	// activities = "has_blood"

	// fmt.Printf("activities is %s and is of type %T\n", activities, activities)

	colour := baby_log.Colour
	breast_milk_time := baby_log.Breast_milk_time
	breast_milk_mls := baby_log.Breast_milk_mls
	formula_milk_mls := baby_log.Formula_milk_mls
	time := baby_log.Time

	// Execute INSERT statement with values
	_, err = stmt.Exec(baby_id, time, activities, colour, breast_milk_time, breast_milk_mls, formula_milk_mls)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted data into baby_logs table.")

}

type App struct {
	DB *sql.DB
}

func main() {

	// Open a database connection
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/aurora")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	app := &App{DB: db}

	// Test the database connection
	err = app.DB.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to the database!")

	http.HandleFunc("/get_post_content", app.get_post_content)

	// Set up a file server to serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8090", nil)

}
