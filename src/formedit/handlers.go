package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func SimpleForm(w http.ResponseWriter, r *http.Request) {
	type secretdict struct {
		BODY      string
		STYLE     string
		DISABLED1 string
		DISABLED2 string
		DISABLED3 string
		DISABLED4 string
		DISABLED5 string
		STATUS1   string
		STATUS2   string
		STATUS3   string
		STATUS4   string
		STATUS5   string
	}

	token := getCookieByName(r.Cookies(), cookieid)
	if token != "" {
		var htmlstring string
		auth, username := isAuthorized(token)
		if auth >= 1 {
			//sort out what the viewer can see based on auth number returned by isAuthorized
			authmap := AuthMap(auth)
			status := r.URL.Query().Get("status")
			if status == "" {

				params := &secretdict{BODY: "", STYLE: "", STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"]}
				t := template.New("test")
				t, err := t.Parse(StarterTemplate)
				LogErr(err)
				err = t.Execute(w, params)
				LogErr(err)
			} else {
				var statnumber string
				switch status {
				case "5":
					statnumber = "5"
				case "4":
					statnumber = "4"
				case "3":
					statnumber = "3"
				case "2":
					statnumber = "2"
				case "1":
					statnumber = "1"
				default:
					fmt.Fprintln(w, "bad status")
				}

				if statnumber != "" {
					var count int
					query := `select id from datasets where status = '` + statnumber + `';`
					rows, err := formdb.Query(query)
					LogErr(err)
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
					params := &secretdict{BODY: htmlstring, STYLE: style, STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"]}
					t := template.New("test")
					t, err = t.Parse(StarterTemplate)
					LogErr(err)
					err = t.Execute(w, params)
					LogErr(err)
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Sorry " + username + " you are not on the list of cool kids."))
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You need to be logged in to use this API :("))
	}

}

func SaveEdit(w http.ResponseWriter, r *http.Request) {


	type save struct {
		RURL string
	}
	token := getCookieByName(r.Cookies(), cookieid)
	auth, username := isAuthorized(token)
	if auth >= 1 {
                id := r.FormValue("id")
                var stat string
                query := `select status from datasets where id = '` + id + `';`
                rows, err := formdb.Query(query)
                LogErr(err)
                for rows.Next() {

                        err = rows.Scan(&stat)
                       	LogErr(err)
                }




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
		var datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool bool
		datasetnamebool = r.FormValue("datasetnamebool") == "true"
		firstnamebool = r.FormValue("firstnamebool") == "true"
		lastnamebool = r.FormValue("lastnamebool") == "true"
		emailbool = r.FormValue("emailbool") == "true"
		phonebool = r.FormValue("phonebool") == "true"
		firstnamepibool = r.FormValue("firstnamepibool") == "true"
		lastnamepibool = r.FormValue("lastnamepibool") == "true"
		emailpibool = r.FormValue("emailpibool") == "true"
		phonepibool = r.FormValue("phonepibool") == "true"
		collectiontitlebool = r.FormValue("collectiontitlebool") == "true"
		categorytitlebool = r.FormValue("categorytitlebool") == "true"
		subcategorytitlebool = r.FormValue("subcategorytitlebool") == "true"
		purposebool = r.FormValue("purposebool") == "true"
		otherinfobool = r.FormValue("otherinfobool") == "true"
		keywordsbool = r.FormValue("keywordsbool") == "true"
		placenamesbool = r.FormValue("placenamesbool") == "true"
		filenamebool = r.FormValue("filenamebool") == "true"
		filetypebool = r.FormValue("filetypebool") == "true"
		filedescriptionbool = r.FormValue("filedescriptionbool") == "true"
		button := r.FormValue("button")
		note := r.FormValue("note")
		t := time.Now().Format("2006-01-02 15:04:05")

		fmt.Println(id + " " + datasetname + " " + firstname + " " + lastname + " " + email + " " + phone + " " + firstnamepi + " " + lastnamepi + " " + emailpi + " " + phonepi + " " + collectiontitle + " " + categorytitle + " " + subcategorytitle + " " + purpose + " " + otherinfo + " " + keywords + " " + placenames + " " + filename + " " + filetype + " " + filedescription + " " + step + "#" + strconv.FormatBool(datasetnamebool) + "-" + strconv.FormatBool(firstnamebool) + "-" + strconv.FormatBool(lastnamebool) + "-" + strconv.FormatBool(emailbool) + "-" + strconv.FormatBool(phonebool) + "-" + strconv.FormatBool(firstnamepibool) + "-" + strconv.FormatBool(lastnamepibool) + "-" + strconv.FormatBool(emailpibool) + "-" + strconv.FormatBool(phonepibool) + "-" + strconv.FormatBool(collectiontitlebool) + "-" + strconv.FormatBool(categorytitlebool) + "-" + strconv.FormatBool(subcategorytitlebool) + "-" + strconv.FormatBool(purposebool) + "-" + strconv.FormatBool(otherinfobool) + "-" + strconv.FormatBool(keywordsbool) + "-" + strconv.FormatBool(placenamesbool) + "-" + strconv.FormatBool(filenamebool) + "-" + strconv.FormatBool(filetypebool) + "-" + strconv.FormatBool(filedescriptionbool) + "-" + note + "-" + button)
		if button == "accept" {
			stmt, err := formdb.Prepare("UPDATE datasets SET status = status + 1 WHERE id =?")
			LogErr(err)
			log.Println(username + " ran this SQL query: UPDATE datasets SET status = status + 1 WHERE id =" + id + ";")
			res, err := stmt.Exec(id)
			log.Println(res)
			LogErr(err)
			affect, rowerr := res.RowsAffected()
			log.Printf("ID = %d, affected = %d\n", id, affect)
			LogErr(rowerr)

			stmt, err = formdb.Prepare("INSERT INTO notes (datasetid,note,date,decision,userid) VALUES (?,?,?,?,?)")
			LogErr(err)

			log.Println(username + " ran this SQL INSERT: INSERT INTO notes (datasetid,note,date,decision) VALUES (" + id + "," + note + "," + t + "," + button + ")")

			res, err = stmt.Exec(id, note, t, button, username)
			log.Println(res)
			LogErr(err)
			affect, rowerr = res.RowsAffected()
			log.Printf("ID = %d, affected = %d\n", id, affect)
			LogErr(rowerr)

			//fmt.Fprintln(w, "Dataset ID "+id+" has been accepted!")
		} else if button == "reject" {
			stmt, err := formdb.Prepare("UPDATE datasets SET status = status - 1, rejected=1 WHERE id =?")
			LogErr(err)
			log.Println(username + " ran this SQL query: UPDATE datasets SET status = status - 1 WHERE id =" + id + ";")
			res, err := stmt.Exec(id)
			log.Println(res)
			LogErr(err)
			affect, rowerr := res.RowsAffected()
			log.Printf("ID = %d, affected = %d\n", id, affect)
			LogErr(rowerr)

			stmt, err = formdb.Prepare("INSERT INTO notes (datasetid,note,date,decision,userid) VALUES (?,?,?,?,?)")
			LogErr(err)
			log.Println(username + " ran this SQL INSERT: INSERT INTO notes (datasetid,note,date,decision) VALUES (" + id + "," + note + "," + t + "," + button + ")")
			res, err = stmt.Exec(id, note, t, button, username)
			log.Println(res)
			LogErr(err)
			affect, rowerr = res.RowsAffected()
			log.Printf("ID = %d, affected = %d\n", id, affect)
			LogErr(rowerr)

			//INSERT INTO notes (datasetid,note,date,decision) VALUES (255,"lol",'2017-03-08 07:48:58',"accepted");
			/*        params := &save{RURL: ""}
			          t := template.New("url")
			          t, err := t.Parse(RedirectTemplate)
			          LogErr(err)
			          err = t.Execute(w, params)
			          LogErr(err)

			*/
			//fmt.Fprintln(w, "Dataset ID "+id+" is rekt!")
			//RedirectTemplate RURL
		}

/*		var stat string
		query := `select status from datasets where id = '` + id + `';`
		rows, err := formdb.Query(query)
		LogErr(err)
		for rows.Next() {

			err = rows.Scan(&stat)
			LogErr(err)
		}*/
		params := &save{RURL: "/formedit/?status=" + stat}
		ut := template.New("url")
		ut, err = ut.Parse(RedirectTemplate)
		LogErr(err)
		err = ut.Execute(w, params)
		LogErr(err)

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
		ABSTRACT         string
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
		DISABLED1        string
		DISABLED2        string
		DISABLED3        string
		DISABLED4        string
		DISABLED5        string
		STATUS1          string
		STATUS2          string
		STATUS3          string
		STATUS4          string
		STATUS5          string
		NOTES            string
	}

	token := getCookieByName(r.Cookies(), cookieid)
	auth, _ := isAuthorized(token)
	if auth >= 1 {
		var params *valuedict
		t := template.New("edit")
		var notes, color string
		ID := r.URL.Query().Get("id")
		authmap := AuthMap(auth)
		/*TODO get notes from db and loop through to make this:
		  `<div class="well"><p class="lead"><font color="grey"><h5>`+username+` `+date+`:
		  </h5></font>`+note+`<br><br>Decision:<font color="red">`+decision+`</font></p></div>`
		*/
		notesquery := `SELECT note, date, decision, userid FROM notes WHERE datasetid='` + ID + `';`
		var note, date, decision, username string
		rows, err := formdb.Query(notesquery)
		LogErr(err)
		for rows.Next() {
			err = rows.Scan(&note, &date, &decision, &username)
			LogErr(err)
			if decision == "reject" {
				color = "red"
			} else {
				color = "green"
			}
			notes = notes + `<div class="well"><p class="lead"><font color="grey"><h5>` + username + ` ` + date + `:
</h5></font>` + note + `<br><br>Decision:<font color="` + color + `">` + decision + `</font></p></div>`
			log.Println("fart")
			LogErr(err)
		}

		//		authmap := AuthMap(auth)
		var datasetname, id, userid, status, collectiontitle, categorytitle, subcategorytitle, firstname, lastname, email, phone, firstnamepi, lastnamepi, emailpi, phonepi, abstract, purpose, otherinfo, keywords, placenames, filename, filetype, filedescription, step string
		//ID := r.URL.Query().Get("id")

		query := `SELECT
                  datasets.id, datasets.datasetname, datasets.userid, datasets.status, collections.collectiontitle, categorys.categorytitle, subcategorys.subcategorytitle, datasets.firstname, datasets.lastname, datasets.email, datasets.phone, datasets.firstnamepi, datasets.lastnamepi, datasets.emailpi, datasets.phonepi, datasets.abstract, datasets.purpose, datasets.otherinfo, datasets.keywords, datasets.placenames, datasets.filename, datasets.filetype, datasets.filedescription, datasets.step
                  FROM datasets, collections, categorys, subcategorys
                  WHERE datasets.id = '` + ID + `' AND datasets.collectionid=collections.id AND datasets.categoryid=categorys.id AND datasets.subcategoryid=subcategorys.id;`

		rows, err = formdb.Query(query)
		LogErr(err)
		for rows.Next() {
			// var id string
			err = rows.Scan(&id, &datasetname, &userid, &status, &collectiontitle, &categorytitle, &subcategorytitle, &firstname, &lastname, &email, &phone, &firstnamepi, &lastnamepi, &emailpi, &phonepi, &abstract, &purpose, &otherinfo, &keywords, &placenames, &filename, &filetype, &filedescription, &step)
			LogErr(err)
			params = &valuedict{ID: id, DATASETNAME: datasetname, COLLECTIONTITLE: collectiontitle, CATEGORYTITLE: categorytitle, SUBCATEGORYTITLE: subcategorytitle, FIRSTNAME: firstname, LASTNAME: lastname, EMAIL: email, PHONE: phone, FIRSTNAMEPI: firstnamepi, LASTNAMEPI: lastnamepi, EMAILPI: emailpi, PHONEPI: phonepi, ABSTRACT: abstract, PURPOSE: purpose, OTHERINFO: otherinfo, KEYWORDS: keywords, PLACENAMES: placenames, FILENAME: filename, FILETYPE: filetype, FILEDESCRIPTION: filedescription, STEP: step, STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"], NOTES: notes}
			//t := template.New("edit")
			t, err = t.Parse(EditTemplate)
			LogErr(err)
		}

		err = t.Execute(w, params)
		LogErr(err)

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

func isAuthorized(token string) (int, string) {
	var username string
	query := "SELECT name FROM users u INNER JOIN sessions s ON u.uid = s.uid WHERE s.sid = '" + token + "';"
	rows, err := db.Query(query)
	LogErr(err)
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
	}

	if username == "hbarrett" {
		return 3, username
	} else if username == "hbarrett" || username == "jsavickas" {
		return 2, username
	} else if username == "gvalentin" || username == "hbarrett" || username == "sdiller" || username == "jsavickas" {
		return 1, username
	} else {
		return 0, username
	}

}

func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AuthMap(auth int) map[string]string {
	var m map[string]string
	if auth == 1 {
		m = map[string]string{
			"disabled1": ` class="disabled"`,
			"disabled2": ``,
			"disabled3": ` class="disabled"`,
			"disabled4": ` class="disabled"`,
			"disabled5": ` class="disabled"`,
			"status1":   ``,
			"status2":   ` href="/formedit/?status=2"`,
			"status3":   ``,
			"status4":   ``,
			"status5":   ``,
		}
	} else if auth == 2 {
		m = map[string]string{
			"disabled1": ` class="disabled"`,
			"disabled2": ``,
			"disabled3": ``,
			"disabled4": ` class="disabled"`,
			"disabled5": ` class="disabled"`,
			"status1":   ``,
			"status2":   ` href="/formedit/?status=2"`,
			"status3":   ` href="/formedit/?status=3"`,
			"status4":   ``,
			"status5":   ``,
		}
	} else if auth == 3 {
		m = map[string]string{
			"disabled1": ``,
			"disabled2": ``,
			"disabled3": ``,
			"disabled4": ``,
			"disabled5": ``,
			"status1":   ` href="/formedit/?status=1"`,
			"status2":   ` href="/formedit/?status=2"`,
			"status3":   ` href="/formedit/?status=3"`,
			"status4":   ` href="/formedit/?status=4"`,
			"status5":   ` href="/formedit/?status=5"`,
		}
	}
	return m
}
