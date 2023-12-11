package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// 字首大寫
type matched struct {
	Id             int
	Gender         string
	Height         int
	Age            int
	Ask_gender     string
	Ask_height_up  int
	Ask_height_low int
	Ask_age_up     int
	Ask_age_low    int
}
type people struct {
	Person []matched
}

var pe = people{}
var tmpl = template.Must(template.ParseFiles("index.html"))

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		pe.Person = nil
		r.ParseForm()

		gender := r.Form.Get("gender")

		heightstr := r.Form.Get("height")
		height, err := strconv.Atoi(heightstr)
		if err != nil {
			//log.Fatal("Error converting height to integer:", err)關掉server
			http.Error(w, "Invalid input, please enter a valid number", http.StatusBadRequest)
			return
		}

		agestr := r.Form.Get("age")
		age, err := strconv.Atoi(agestr)
		if err != nil {
			http.Error(w, "Invalid input, please enter a valid number", http.StatusBadRequest)
			return
		}

		ask_gender := r.Form.Get("ask_gender")

		ask_height_upstr := r.Form.Get("ask_height_up")
		ask_height_up, err := strconv.Atoi(ask_height_upstr)
		if err != nil {
			http.Error(w, "Invalid input, please enter a valid number", http.StatusBadRequest)
			return
		}

		ask_height_lowstr := r.Form.Get("ask_height_low")
		ask_height_low, err := strconv.Atoi(ask_height_lowstr)
		if err != nil {
			http.Error(w, "Invalid input, please enter a valid number", http.StatusBadRequest)
			return
		}

		ask_age_upstr := r.Form.Get("ask_age_up")
		ask_age_up, err := strconv.Atoi(ask_age_upstr)
		if err != nil {
			http.Error(w, "Invalid input, please enter a valid number", http.StatusBadRequest)
			return
		}

		ask_age_lowstr := r.Form.Get("ask_age_low")
		ask_age_low, err := strconv.Atoi(ask_age_lowstr)
		if err != nil {
			http.Error(w, "Invalid input, please enter a valid number", http.StatusBadRequest)
			return
		}

		db, err := sql.Open("mysql", "user:1234@tcp(localhost:3306)/kkbox?charset=utf8")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		rows, err := db.Query("select * from matched_table where "+
			"ask_gender = ? and gender = ? "+
			"and ask_height_up >= ? and ask_height_low <= ? "+
			"and height <= ? and height >= ? "+
			"and ask_age_up >= ? and ask_age_low <= ? "+
			"and age <= ? and age >= ? ",
			gender, ask_gender, height, height, ask_height_up, ask_height_low,
			age, age, ask_age_up, ask_age_low)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			ma := matched{}
			err := rows.Scan(&ma.Id, &ma.Gender, &ma.Height, &ma.Age,
				&ma.Ask_gender, &ma.Ask_height_up, &ma.Ask_height_low,
				&ma.Ask_age_up, &ma.Ask_age_low)
			if err != nil {
				panic(err)
			}
			pe.Person = append(pe.Person, ma)
		}

	}
	err := tmpl.Execute(w, pe)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)

}
