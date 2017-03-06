package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func SimpleForm(w http.ResponseWriter, r *http.Request) {
	type secretdict struct {
		BODY  string
		STYLE string
	}

	token := getCookieByName(r.Cookies(), cookieid)
	if token != "" {
		var htmlstring string
		auth := isAuthorized(token)
		if auth == true {
			status := r.URL.Query().Get("status")
			if status == "" {

				params := &secretdict{BODY: "", STYLE: ""}

				t := template.New("test")
				t, err := t.Parse(StarterTemplate)
				if err != nil {
					log.Fatal(err)
				}
				err = t.Execute(w, params)
				if err != nil {
					panic(err)
				}
			} else {
				var statnumber string
				switch status {
				case "done":
					statnumber = "5"
				case "ready":
					statnumber = "4"
				case "three":
					statnumber = "3"
				case "two":
					statnumber = "2"
				case "one":
					statnumber = "1"
				default:
					fmt.Fprintln(w, "bad status")
				}
				if statnumber != "" {
					var count int
					query := `select id from datasets where status = '` + statnumber + `';`
					rows, err := formdb.Query(query)
					if err != nil {
						log.Fatal(err)
					}
					for rows.Next() {
						var id string
						err = rows.Scan(&id)
						htmlstring = htmlstring + `<a class="btn btn-primary custom" href="/formedit/edit?id=` + id + `">` + id + `</a>`
						count++
						if count == 15 {
							htmlstring = htmlstring + `</p>`
							count = 0
						}
					}
					htmlstring = `<div class="centered">` + htmlstring + `</div>`
					style := `<style>
                                                 .custom{
                                                  width:76px!important;
                                                 }
                                                  .centered
                                                 {
                                                   text-align:center;
                                                 }
                                                 </style>`

					params := &secretdict{BODY: htmlstring, STYLE: style}
					t := template.New("test")
					t, err = t.Parse(StarterTemplate)
					if err != nil {
						log.Fatal(err)
					}
					err = t.Execute(w, params)
					if err != nil {
						panic(err)
					}

				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("You are not on the list of cool kids."))
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You need to be logged in to use this API :("))
	}

}

func SaveEdit(w http.ResponseWriter, r *http.Request) {
	token := getCookieByName(r.Cookies(), cookieid)
	auth := isAuthorized(token)
	if auth == true {
		id := r.FormValue("id")
		datasetname := r.FormValue("datasetname")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		firstnamepi := r.FormValue("firstnamepi")
		lastnamepi := r.FormValue("lastnamepi")
		emailpi := r.FormValue("emailpi")
		phonepi := r.FormValue("phonepi")
		collectiontitle := r.FormValue("collectiontitle")
		categorytitle := r.FormValue("categorytitle")
		subcategorytitle := r.FormValue("subcategorytitle")
		purpose := r.FormValue("purpose")
		otherinfo := r.FormValue("otherinfo")
		keywords := r.FormValue("keywords")
		placenames := r.FormValue("placenames")
		filename := r.FormValue("filename")
		filetype := r.FormValue("filetype")
		filedescription := r.FormValue("filedescription")
		step := r.FormValue("step")
		fmt.Fprintln(w, id+" "+datasetname+" "+firstname+" "+lastname+" "+email+" "+phone+" "+firstnamepi+" "+lastnamepi+" "+emailpi+" "+phonepi+" "+collectiontitle+" "+categorytitle+" "+subcategorytitle+" "+purpose+" "+otherinfo+" "+keywords+" "+placenames+" "+filename+" "+filetype+" "+filedescription+" "+step)
	}
}

func ChangeStatus(w http.ResponseWriter, r *http.Request) {
	token := getCookieByName(r.Cookies(), cookieid)
	auth := isAuthorized(token)
	if auth == true {
		id := r.FormValue("id")
		email := r.FormValue("email")
		fmt.Fprintln(w, id+"has been rejected. Sending email to "+email)
	}

}

func SimpleEdit(w http.ResponseWriter, r *http.Request) {
	type valuedict struct {
		ID               string
		DATASETNAME      string
		FIRSTNAME        string
		LASTNAME         string
		EMAIL            string
		PHONE            string
		FIRSTNAMEPI      string
		LASTNAMEPI       string
		EMAILPI          string
		PHONEPI          string
		ABSTRACT	 string
		COLLECTIONTITLE  string
		CATEGORYTITLE    string
		SUBCATEGORYTITLE string
		PURPOSE          string
		OTHERINFO        string
		KEYWORDS         string
		PLACENAMES       string
		FILENAME         string
		FILETYPE         string
		FILEDESCRIPTION  string
		STEP             string
	}

	token := getCookieByName(r.Cookies(), cookieid)
	auth := isAuthorized(token)
	if auth == true {

		var datasetname, id, userid, status, collectiontitle, categorytitle, subcategorytitle, firstname, lastname, email, phone, firstnamepi, lastnamepi, emailpi, phonepi, abstract, purpose, otherinfo, keywords, placenames, filename, filetype, filedescription, step string
		ID := r.URL.Query().Get("id")

		query := `SELECT
                  datasets.id, datasets.datasetname, datasets.userid, datasets.status, collections.collectiontitle, categorys.categorytitle, subcategorys.subcategorytitle, datasets.firstname, datasets.lastname, datasets.email, datasets.phone, datasets.firstnamepi, datasets.lastnamepi, datasets.emailpi, datasets.phonepi, datasets.abstract, datasets.purpose, datasets.otherinfo, datasets.keywords, datasets.placenames, datasets.filename, datasets.filetype, datasets.filedescription, datasets.step
                  FROM datasets, collections, categorys, subcategorys
                  WHERE datasets.id = '` + ID + `' AND datasets.collectionid=collections.id AND datasets.categoryid=categorys.id AND datasets.subcategoryid=subcategorys.id;`

		rows, err := formdb.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			// var id string
                        err = rows.Scan(&id, &datasetname, &userid, &status, &collectiontitle, &categorytitle, &subcategorytitle, &firstname, &lastname, &email, &phone, &firstnamepi, &lastnamepi, &emailpi, &phonepi, &abstract, &purpose, &otherinfo, &keywords, &placenames, &filename, &filetype, &filedescription, &step)

			params := &valuedict{ID: id, DATASETNAME: datasetname, COLLECTIONTITLE: collectiontitle, CATEGORYTITLE: categorytitle, SUBCATEGORYTITLE: subcategorytitle, FIRSTNAME: firstname, LASTNAME: lastname, EMAIL: email, PHONE: phone, FIRSTNAMEPI: firstnamepi, LASTNAMEPI: lastnamepi, EMAILPI: emailpi, PHONEPI: phonepi, ABSTRACT: abstract, PURPOSE: purpose, OTHERINFO: otherinfo, KEYWORDS: keywords, PLACENAMES: placenames, FILENAME: filename, FILETYPE: filetype, FILEDESCRIPTION: filedescription, STEP: step}

// 			params := &secretdict{BODY: htmlstring, STYLE: style}
                                        t := template.New("edit")
                                        t, err = t.Parse(EditTemplate)
                                        if err != nil {
                                                log.Fatal(err)
                                        }
                                        err = t.Execute(w, params)
                                        if err != nil {
                                                panic(err)
                                        }

		}

	}

}

func getCookieByName(cookie []*http.Cookie, name string) string {
	cookieLen := len(cookie)
	result := ""
	for i := 0; i < cookieLen; i++ {
		if cookie[i].Name == name {
			result = cookie[i].Value
		}
	}
	return result
}

func isAuthorized(token string) bool {
	var username string
	query := "SELECT name FROM users u INNER JOIN sessions s ON u.uid = s.uid WHERE s.sid = '" + token + "';"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
	}

	if username == "gvalentin" || username == "hbarrett" || username == "sdiller" {
		return true
	} else {
		return false
	}

}
