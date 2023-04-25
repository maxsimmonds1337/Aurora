package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Name string
	Age  int
	Rows []map[string]interface{}
}

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

	app.insert_data(body)

	// if no errors, then we set the headers in the http response writer
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("data"); err != nil { // here, we make a json encoder that will write to the response writer, and encode the data back into json
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (app *App) data(w http.ResponseWriter, req *http.Request) {

	data := Data{Name: "Aurora", Age: 0}

	tpl, err := template.New("test").Parse(`

	<table class="table table-striped table-sm">

		<thead>
		<tr>
		<th scope="col">Time</th>
		<th scope="col">Activities</th>
		<th scope="col">Colour</th>
		<th scope="col">Breast Milk Time</th>
		<th scope="col">Bresat Milk mls</th>
		<th scope="col">Formula Milk mls</th>
		</tr>
		</thead>
		<tbody>
		{{range $index, $row := .Rows}}
			<tr>
				<td class ="centered">{{$row.Time}}</td>
				<td class ="centered">{{$row.Activities}}</td>
				<td class ="centered">{{$row.Color}}</td>
				<td class ="centered">{{$row.BreastMilkTime}}</td>
				<td class ="centered">{{$row.BreastMilkMls}}</td>
				<td class ="centered">{{$row.FormulaMilkMls}}</td>
			</tr>
		{{end}}
		</tbody>
	</table>
			`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := app.DB.Query("SELECT * FROM baby_logs ORDER BY log_id DESC;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // make sure we remove the memory once the function closes

	var (
		LogId          int //TODO: make this match the var below
		BabyID         string
		Time           string
		Activities     string
		Color          string
		BreastMilkTime []uint8
		BreastMilkMls  []uint8
		FormulaMilkMls []uint8
	)

	var dataRows []map[string]interface{} // this will hold all the rows, once we get a row we append here
	for rows.Next() {                     //this will loop over each element in the rows object
		var row = make(map[string]interface{}) // declare a row which is a map of key strings and assorted values, but make it empty
		// next, scan
		if err := rows.Scan(&LogId, &BabyID, &Time, &Activities, &Color, &BreastMilkTime, &BreastMilkMls, &FormulaMilkMls); err != nil {
			log.Fatal(err)
		}

		row["LogId"] = LogId
		row["BabyID"] = BabyID
		row["Time"] = Time
		row["Activities"] = Activities
		row["Color"] = Color
		row["BreastMilkTime"] = string(BreastMilkTime)
		row["BreastMilkMls"] = string(BreastMilkMls)
		row["FormulaMilkMls"] = string(FormulaMilkMls)

		dataRows = append(dataRows, row)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	data.Rows = dataRows // add the row to Rows

	// execute the template, w is the response writer (which is needed to return the http repsonse) data is the data to fill the template with
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// this function returns data to the chart for the dash board.
func (app *App) chartdata(w http.ResponseWriter, req *http.Request) {

	// read the body of the request
	body, err := io.ReadAll(req.Body)
	if err != nil {
		// if we hit an error, return this TODO: this should be a json string to handle it correctly on the frontend
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	//upon exiting the function, close the body request
	defer req.Body.Close()

	// this will check to see what code we get, 7d = 7 days worth of data, etc.
	if string(body) == "7d" {
		// 7 days, so we get 7 days worth of averages

		// need to see how to handle this, it;s now a 2d array returned, so need to get the data out before charting

		rows, err := app.DB.Query("SELECT DAYNAME(time), sum(breast_milk_time) as total_breast_milk_time, sum(breast_milk_mls) as total_breast_milk_mls, sum(formula_milk_mls) as total_formula_milk_mls FROM baby_logs WHERE time > DATE_SUB(NOW(), INTERVAL 6 DAY) GROUP BY DAYNAME(time), DATE(time) ORDER BY DATE(time) ASC;")
		if err != nil {
			http.Error(w, "error executing sql2", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var (
			Day            string
			BreastMilkTime float32
			BreastMilkMls  float32
			FormulaMilkMls float32

			Days            []string
			BreastMilkTimes []float32
			BreastMilkMlss  []float32
			FormulaMilkMlss []float32
		)

		for rows.Next() {

			// scan the row into the ars
			rows.Scan(&Day, &BreastMilkTime, &BreastMilkMls, &FormulaMilkMls)

			// add them to the slices
			Days = append(Days, Day)
			BreastMilkTimes = append(BreastMilkTimes, BreastMilkTime)
			FormulaMilkMlss = append(FormulaMilkMlss, FormulaMilkMls)
			BreastMilkMlss = append(BreastMilkMlss, BreastMilkMls)

		}

		type Dataset struct {
			Data                 []float32 `json:"data"` // maybe remove these tags, as they might not be needed
			LineTension          int       `json:"lineTension"`
			BackgroundColor      string    `json:"backgroundColor"`
			BorderColor          string    `json:"borderColor"`
			BorderWidth          int       `json:"borderWidth"`
			PointBackgroundColor string    `json:"pointBackgroundColor"`
		}

		type ChartData struct {
			Labels   []string  `json:"labels"`
			Datasets []Dataset `json:"datasets"`
		}

		// hard code something in the struct and return it for a test

		chartData := ChartData{
			Labels: Days,
			Datasets: []Dataset{
				Dataset{
					Data:                 BreastMilkTimes,
					LineTension:          int(0),
					BackgroundColor:      string("transparent"),
					BorderColor:          string("#ff0000"),
					BorderWidth:          int(4),
					PointBackgroundColor: string("#ff0000"),
				},
				Dataset{
					Data:                 BreastMilkMlss,
					LineTension:          int(0),
					BackgroundColor:      string("transparent"),
					BorderColor:          string("#00ff00"),
					BorderWidth:          int(4),
					PointBackgroundColor: string("#00ff00"),
				},
				Dataset{
					Data:                 FormulaMilkMlss,
					LineTension:          int(0),
					BackgroundColor:      string("transparent"),
					BorderColor:          string("#0000ff"),
					BorderWidth:          int(4),
					PointBackgroundColor: string("#0000ff"),
				},
			},
		}

		if err := json.NewEncoder(w).Encode(chartData); err != nil { // here, we make a json encoder that will write to the response writer, and encode the data back into json
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

		// for i, row := range dataRows {
		// 	fmt.Printf("Row %d:\n", i+1)
		// 	for key, value := range row {
		// 		fmt.Printf("\t%s: %v\n", key, value)
		// 	}
		// }

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
	GCP_USER := os.Getenv("GCP_USER")
	GCP_PASS := os.Getenv("GCP_PASS")
	HOST := os.Getenv("HOST")
	DB_PORT := os.Getenv("DB_PORT")
	PORT := os.Getenv("PORT")

	fmt.Printf("user %v pass %v host %v DB_PORT %v port %v\n", GCP_USER, GCP_PASS, HOST, DB_PORT, PORT)

	db, err := sql.Open("mysql", GCP_USER+":"+GCP_PASS+"@tcp("+HOST+":"+DB_PORT+")/aurora")
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
	http.HandleFunc("/chartdata", app.chartdata)
	http.HandleFunc("/data", app.data)

	// Set up a file server to serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":"+PORT, nil)

}
