package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
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
		LOGO      string
	}

	token := getCookieByName(r.Cookies(), cookieid)
	if token != "" {
		var body string
		var statusi int
		auth, username := isAuthorized(token)
		status := r.URL.Query().Get("status")
		if status != "" {
			var err error
			statusi, err = strconv.Atoi(status)
			LogErr(err)
		}
		if auth >= 1 && auth+1 >= statusi && statusi != 1 {
			//sort out what the viewer can see based on auth number returned by isAuthorized
			authmap := AuthMap(auth)
			//	authmap := SetCurrent(authmap, status)
			if status == "" {

				params := &secretdict{BODY: "", STYLE: "", STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"], LOGO: Logo}
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
					authmap["disabled5"] = ` class="active"`
				case "4":
					statnumber = "4"
					authmap["disabled4"] = ` class="active"`
				case "3":
					statnumber = "3"
					authmap["disabled3"] = ` class="active"`
				case "2":
					statnumber = "2"
					authmap["disabled2"] = ` class="active"`
				case "1":
					statnumber = "1"
					authmap["disabled1"] = ` class="active"`
				default:
					fmt.Fprintln(w, "bad status")
				}

				if statnumber != "" {
					query := `select id, datasetname, firstname, lastname, email, datecreated from datasets where status = '` + statnumber + `' ORDER BY id;`
					rows, err := formdb.Query(query)
					LogErr(err)
					for rows.Next() {
						var id, datasetname, firstname, lastname, email, datecreated string
						err = rows.Scan(&id, &datasetname, &firstname, &lastname, &email, &datecreated)
						body = body + `<a href="/formedit/edit?id=` + id + `" class="list-group-item list-group-item-action"><span class="badge badge-default badge-pill">Dataset ID:` + id + `</span></p>Dataset Name: <strong>` + datasetname + `</strong></p>Submitted by: <strong>` + firstname + ` ` + lastname + `</p></strong>E-Mail:<strong>` + email + `  </p></strong><small class="text-muted">Dataset Created: ` + datecreated + `</small></a>`
					}
					body = `<div class="list-group">` + body + `</div>`
					style := `<style>
                                                 .custom{
                                                  width:76px!important;
                                                 }
                                                  .centered
                                                 {
                                                   text-align:center;
                                                 }
                                                 </style>`

					params := &secretdict{BODY: body, STYLE: style, STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"], LOGO: Logo}
					t := template.New("test")
					t, err = t.Parse(StarterTemplate)
					LogErr(err)
					err = t.Execute(w, params)
					LogErr(err)
				}
			}
		} else {

			authmap := AuthMap(auth)
			w.WriteHeader(http.StatusUnauthorized)
			style := `.footer-images {
                        margin: auto;
                        display: block;
                        position: absolute;
                        bottom: 0;
                        text-align: center;
                        }
                        .footer-images .copyright {
                        text-align: center;
                        }`

			body := `<div class="container">
                             <div class="jumbotron">
                             <h1>Not Authorized</h1>
                             <p>` + username + ` is not autorized for this view.</p>
                             <div class="footer-images">
                             <img alt="Bouncer" src="` + Bouncer + `"/>
                            </div>
                            </div>
                            </div>`
			params := &secretdict{BODY: body, STYLE: style, STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"], LOGO: Logo}
			t := template.New("error")
			t, err := t.Parse(StarterTemplate)
			LogErr(err)
			err = t.Execute(w, params)
			LogErr(err)

		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		style := `.footer-images {
                        margin: auto;
                        display: block;
                        position: absolute;
                        bottom: 0;
 			text-align: center;
                        }
                        .footer-images .copyright {
                        text-align: center;
                        }`

		body := `<div class="container">
                             <div class="jumbotron">
                             <h1>Not Authorized</h1>
                             <p>You must be logged into the reporting site to use this tool</p>
                             <div class="footer-images">
                             <img alt="Bouncer" src="` + Bouncer + `"/>
                            </div>
                            </div>
                            </div>`
		params := &secretdict{BODY: body, STYLE: style, LOGO: Logo}
		t := template.New("error")
		t, err := t.Parse(Bootstrap)
		LogErr(err)
		err = t.Execute(w, params)
		LogErr(err)

	}

}

func SaveEdit(w http.ResponseWriter, r *http.Request) {
        host:=r.Host
	type save struct {
		RURL string
	}
	token := getCookieByName(r.Cookies(), cookieid)
	auth, username := isAuthorized(token)
	id := r.FormValue("id")
	var stat string
	query := `select status from datasets where id = '` + id + `';`
	rows, err := formdb.Query(query)
	LogErr(err)
	for rows.Next() {

		err = rows.Scan(&stat)
		LogErr(err)
	}

	var statusi int
	statusi, err = strconv.Atoi(stat)
	LogErr(err)
	if auth >= 1 && auth+1 >= statusi && statusi != 1 {

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
		var datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool, abstractbool bool
		datasetnamebool = r.FormValue("datasetnamebool") == "true"
		firstnamebool = r.FormValue("firstnamebool") == "true"
		lastnamebool = r.FormValue("lastnamebool") == "true"
		emailbool = r.FormValue("emailbool") == "true"
		phonebool = r.FormValue("phonebool") == "true"
		firstnamepibool = r.FormValue("firstnamepibool") == "true"
		lastnamepibool = r.FormValue("lastnamepibool") == "true"
		emailpibool = r.FormValue("emailpibool") == "true"
		phonepibool = r.FormValue("phonepibool") == "true"
		abstractbool = r.FormValue("abstractbool") == "true"
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

		log.Println(id + " " + datasetname + " " + firstname + " " + lastname + " " + email + " " + phone + " " + firstnamepi + " " + lastnamepi + " " + emailpi + " " + phonepi + " " + collectiontitle + " " + categorytitle + " " + subcategorytitle + " " + purpose + " " + otherinfo + " " + keywords + " " + placenames + " " + filename + " " + filetype + " " + filedescription + " " + step + "#" + strconv.FormatBool(datasetnamebool) + "-" + strconv.FormatBool(firstnamebool) + "-" + strconv.FormatBool(lastnamebool) + "-" + strconv.FormatBool(emailbool) + "-" + strconv.FormatBool(phonebool) + "-" + strconv.FormatBool(firstnamepibool) + "-" + strconv.FormatBool(lastnamepibool) + "-" + strconv.FormatBool(emailpibool) + "-" + strconv.FormatBool(phonepibool) + "-" + strconv.FormatBool(collectiontitlebool) + "-" + strconv.FormatBool(categorytitlebool) + "-" + strconv.FormatBool(subcategorytitlebool) + "-" + strconv.FormatBool(purposebool) + "-" + strconv.FormatBool(otherinfobool) + "-" + strconv.FormatBool(keywordsbool) + "-" + strconv.FormatBool(placenamesbool) + "-" + strconv.FormatBool(filenamebool) + "-" + strconv.FormatBool(filetypebool) + "-" + strconv.FormatBool(filedescriptionbool) + "-" + note + "-" + button + "-" + host)

		if button == "accept" {
			UpdateStatus(id, button, username)
			MakeNote(id, note, t, button, username)
			SetBools(id, datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, abstractbool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool)
//			SendMail(id, datasetname, firstname, t, note, stat, button, emailbool, emailpibool, host)
		} else if button == "reject" {
			UpdateStatus(id, button, username)
			MakeNote(id, note, t, button, username)
			SetBools(id, datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, abstractbool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool)
//			SendMail(id, datasetname, firstname, t, note, stat, button, emailbool, emailpibool, host)
		} else if button == "note" {
			MakeNote(id, note, t, button, username)
		}
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
		ID                   string
		DATASETNAME          string
		FIRSTNAME            string
		LASTNAME             string
		EMAIL                string
		PHONE                string
		FIRSTNAMEPI          string
		LASTNAMEPI           string
		EMAILPI              string
		PHONEPI              string
		ABSTRACT             string
		COLLECTIONTITLE      string
		CATEGORYTITLE        string
		SUBCATEGORYTITLE     string
		PURPOSE              string
		OTHERINFO            string
		KEYWORDS             string
		PLACENAMES           string
		FILENAME             string
		FILETYPE             string
		FILEDESCRIPTION      string
		DATA                 string
		STEP                 string
		DISABLED1            string
		DISABLED2            string
		DISABLED3            string
		DISABLED4            string
		DISABLED5            string
		STATUS1              string
		STATUS2              string
		STATUS3              string
		STATUS4              string
		STATUS5              string
		NOTES                string
		DATASETNAMEBOOL      string
		FIRSTNAMEBOOL        string
		LASTNAMEBOOL         string
		EMAILBOOL            string
		PHONEBOOL            string
		FIRSTNAMEPIBOOL      string
		LASTNAMEPIBOOL       string
		EMAILPIBOOL          string
		PHONEPIBOOL          string
		ABSTRACTBOOL         string
		COLLECTIONTITLEBOOL  string
		CATEGORYTITLEBOOL    string
		SUBCATEGORYTITLEBOOL string
		PURPOSEBOOL          string
		OTHERINFOBOOL        string
		KEYWORDSBOOL         string
		PLACENAMESBOOL       string
		FILENAMEBOOL         string
		FILETYPEBOOL         string
		DATABOOL             string
		FILEDESCRIPTIONBOOL  string
		LOGO                 string
	}

	token := getCookieByName(r.Cookies(), cookieid)
	auth, _ := isAuthorized(token)
	if auth >= 1 {
		var params *valuedict
		t := template.New("edit")
		var notes, color string
		ID := r.URL.Query().Get("id")
		authmap := AuthMap(auth)
		notesquery := `SELECT note, date, decision, userid FROM notes WHERE datasetid='` + ID + `' ORDER BY date;`
		log.Println(notesquery)
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
			LogErr(err)
		}

		var id, datasetname, userid, status, categorytitle, subcategorytitle, firstname, lastname, email, phone, firstnamepi, lastnamepi, emailpi, phonepi, abstract, purpose, otherinfo, keywords, placenames, filename, filetype, filedescription, step sql.NullString
		collectionisworking := true
		var data, collectiontest, query, collectiontitle string
		var datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool, abstractbool, databool bool

		collidquery := `SELECT
		                collections.collectiontitle
		                FROM datasets,
		                collections
		                WHERE datasets.id = '` + ID + `' AND datasets.collectionid=collections.id;`

		err = formdb.QueryRow(collidquery).Scan(&collectiontest)
		if err != nil {
			if err == sql.ErrNoRows {
				collectionisworking = false
			} else {
				LogErr(err)

			}
		} else {
			collectionisworking = true
			log.Println(collectiontest + " found")
		}
		if collectionisworking {
			query = `SELECT datasets.id, datasets.datasetname, datasets.userid, datasets.status, collections.collectiontitle, categorys.categorytitle, subcategorys.subcategorytitle, datasets.firstname, datasets.lastname, datasets.email, datasets.phone, datasets.firstnamepi, datasets.lastnamepi, datasets.emailpi, datasets.phonepi, datasets.abstract, datasets.purpose, datasets.otherinfo, datasets.keywords, datasets.placenames, datasets.filename, datasets.filetype, datasets.filedescription, datasets.step FROM datasets, collections, categorys, subcategorys WHERE datasets.id = '` + ID + `' AND datasets.collectionid=collections.id AND datasets.categoryid=categorys.id AND datasets.subcategoryid=subcategorys.id;`
		} else {
			query = `SELECT datasets.id, datasets.datasetname, datasets.userid, datasets.status, categorys.categorytitle, subcategorys.subcategorytitle, datasets.firstname, datasets.lastname, datasets.email, datasets.phone, datasets.firstnamepi, datasets.lastnamepi, datasets.emailpi, datasets.phonepi, datasets.abstract, datasets.purpose, datasets.otherinfo, datasets.keywords, datasets.placenames, datasets.filename, datasets.filetype, datasets.filedescription, datasets.step FROM datasets, categorys, subcategorys WHERE datasets.id = '` + ID + `' AND datasets.categoryid=categorys.id AND datasets.subcategoryid=subcategorys.id;`
		}

		log.Println(query)
		rows, err = formdb.Query(query)
		LogErr(err)
		for rows.Next() {
			query = `SELECT datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, abstractbool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool, databool FROM checks where datasetid='` + ID + `';`
			checkrows, err := formdb.Query(query)
			LogErr(err)
			for checkrows.Next() {
				err = checkrows.Scan(&datasetnamebool, &firstnamebool, &lastnamebool, &emailbool, &phonebool, &firstnamepibool, &lastnamepibool, &emailpibool, &phonepibool, &abstractbool, &collectiontitlebool, &categorytitlebool, &subcategorytitlebool, &purposebool, &otherinfobool, &keywordsbool, &placenamesbool, &filenamebool, &filetypebool, &filedescriptionbool, &databool)

				LogErr(err)

			}
			if collectionisworking {
				err = rows.Scan(&id, &datasetname, &userid, &status, &collectiontitle, &categorytitle, &subcategorytitle, &firstname, &lastname, &email, &phone, &firstnamepi, &lastnamepi, &emailpi, &phonepi, &abstract, &purpose, &otherinfo, &keywords, &placenames, &filename, &filetype, &filedescription, &step)
			} else {
				err = rows.Scan(&id, &datasetname, &userid, &status, &categorytitle, &subcategorytitle, &firstname, &lastname, &email, &phone, &firstnamepi, &lastnamepi, &emailpi, &phonepi, &abstract, &purpose, &otherinfo, &keywords, &placenames, &filename, &filetype, &filedescription, &step)
				collectiontitle = "INVALID COLLECTION TITLE!!!"
			}
			LogErr(err)
			if _, err := os.Stat("/uploads/" + IDtoName(Null2String(userid)) + "/" + Null2String(filename)); err == nil {
				data = `<a href="/formedit/download/` + ID + `">Data Download</a></div>`
			} else {
				data = `<font color="red"><strong>File not found!</strong></font></div>`
			}

			params = &valuedict{ID: Null2String(id), DATASETNAME: Null2String(datasetname), COLLECTIONTITLE: collectiontitle, CATEGORYTITLE: Null2String(categorytitle), SUBCATEGORYTITLE: Null2String(subcategorytitle), FIRSTNAME: Null2String(firstname), LASTNAME: Null2String(lastname), EMAIL: Null2String(email), PHONE: Null2String(phone), FIRSTNAMEPI: Null2String(firstnamepi), LASTNAMEPI: Null2String(lastnamepi), EMAILPI: Null2String(emailpi), PHONEPI: Null2String(phonepi), ABSTRACT: Null2String(abstract), PURPOSE: Null2String(purpose), OTHERINFO: Null2String(otherinfo), KEYWORDS: Null2String(keywords), PLACENAMES: Null2String(placenames), FILENAME: Null2String(filename), FILETYPE: Null2String(filetype), FILEDESCRIPTION: Null2String(filedescription), STEP: Null2String(step), STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"], NOTES: notes, DATASETNAMEBOOL: ischecked(datasetnamebool), FIRSTNAMEBOOL: ischecked(firstnamebool), LASTNAMEBOOL: ischecked(lastnamebool), EMAILBOOL: ischecked(emailbool), PHONEBOOL: ischecked(phonebool), FIRSTNAMEPIBOOL: ischecked(firstnamepibool), LASTNAMEPIBOOL: ischecked(lastnamepibool), EMAILPIBOOL: ischecked(emailpibool), PHONEPIBOOL: ischecked(phonepibool), ABSTRACTBOOL: ischecked(abstractbool), COLLECTIONTITLEBOOL: ischecked(collectiontitlebool), CATEGORYTITLEBOOL: ischecked(categorytitlebool), SUBCATEGORYTITLEBOOL: ischecked(subcategorytitlebool), PURPOSEBOOL: ischecked(purposebool), OTHERINFOBOOL: ischecked(otherinfobool), KEYWORDSBOOL: ischecked(keywordsbool), PLACENAMESBOOL: ischecked(placenamesbool), FILENAMEBOOL: ischecked(filenamebool), FILETYPEBOOL: ischecked(filetypebool), FILEDESCRIPTIONBOOL: ischecked(filedescriptionbool), LOGO: Logo, DATA: data, DATABOOL: ischecked(databool)}

			t, err = t.Parse(EditTemplate)
			LogErr(err)
		}

		err = t.Execute(w, params)
		LogErr(err)

	}

}

func Download(w http.ResponseWriter, r *http.Request) {
	var userid, filename string
	vars := mux.Vars(r)
	id := vars["id"]
	query := `SELECT userid, filename FROM datasets where id='` + id + `';`
	rows, err := formdb.Query(query)
	LogErr(err)
	for rows.Next() {
		err = rows.Scan(&userid, &filename)
		LogErr(err)
	}
	pathtofile := "/uploads/" + IDtoName(userid) + "/" + filename
	attachment := `attachment; filename=` + filename
	fmt.Println(attachment)
	w.Header().Set("Content-Disposition", attachment)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	http.ServeFile(w, r, pathtofile)
}

func InsertView(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	token := getCookieByName(r.Cookies(), cookieid)
	if token != "" {
		auth, username := isAuthorized(token)
		if auth >= 3 {
			w.Write([]byte(username + id))

			/*
			   buf := new(bytes.Buffer)
			   t := template.Must(template.New("json").Parse(JSON))
			   err = t.Execute(buf, '')
			   fmt.Println(buf.String())
			*/

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

	if stringInSlice(username, strings.Fields(configf.Admins)) {
		return 4, username
	} else if stringInSlice(username, strings.Fields(configf.Managers)) {
		return 2, username
	} else if stringInSlice(username, strings.Fields(configf.Users)) {
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
	} else if auth == 4 {
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

func send(recipient string, subject string, firstname string, topmessage string, buttonlink string, buttontext string, note string) {
	type maildict struct {
		USERFIRSTNAME string
		TOPMESSAGE    string
		BUTTONLINK    string
		BUTTONTEXT    string
		NOTES         string
	}

	from := "qualityassurance@edac.unm.edu"
	pass := "fluhelk"
	to := recipient
	params := &maildict{USERFIRSTNAME: firstname, TOPMESSAGE: topmessage, BUTTONLINK: buttonlink, BUTTONTEXT: buttontext, NOTES: note}
	buf := new(bytes.Buffer)
	t := template.New("mail")
	t, err := t.Parse(MAIL)
	LogErr(err)
	err = t.Execute(buf, params)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"Content-Type: text/html\n\n" +
		buf.String()
	log.Println(msg)
	toslc := strings.Split(to, ",")
	fmt.Println(toslc[0])
	err = smtp.SendMail("edacmail.unm.edu:587",
		smtp.PlainAuth("", from, pass, "edacmail.unm.edu"),
              from, toslc, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent")
}

func ischecked(param bool) string {

	if param {
		return "checked"
	} else {
		return ""
	}
}

func CheckRowExists(id string) bool {
	var check bool
	err := formdb.QueryRow(`SELECT EXISTS (SELECT * FROM checks WHERE datasetid ='` + id + `');`).Scan(&check)
	log.Println(id)
	log.Println(check)
	log.Println(id)

	if err != nil {
		log.Println(err)
		return false
	}
	return check
}

func IDtoName(id string) string {

	var username string
	query := `select name from users where uid ='` + id + `';`
	rows, err := db.Query(query)
	LogErr(err)
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
	}
	return username
}

func ContactFromUserID(name string) (string, string) {
	var realname, email string
	query := `select realname.realname, users.mail from realname, users where users.uid=realname.uid and users.name='` + name + `';`
	err := db.QueryRow(query).Scan(&realname, &email)
	LogErr(err)
	return realname, email

}

func ContactFromID(id string) (string, string) {
	var userid, realname, email string
	formquery := `select userid from datasets where id=` + id + `;`
	err := formdb.QueryRow(formquery).Scan(&userid)
	LogErr(err)
	query := `select realname.realname, users.mail from realname, users where users.uid=realname.uid and users.uid='` + userid + `';`
	err = db.QueryRow(query).Scan(&realname, &email)
	LogErr(err)
	return realname, email
}

func Null2String(str sql.NullString) string {
	var returnstring string
	if str.Valid {
		returnstring = str.String
	} else {
		returnstring = "NULL"
	}
	return returnstring
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MakeNote(id string, note string, t string, button string, username string) {
	stmt, err := formdb.Prepare("INSERT INTO notes (datasetid,note,date,decision,userid) VALUES (?,?,?,?,?)")
	LogErr(err)
	log.Println(username + " ran this SQL INSERT: INSERT INTO notes (datasetid,note,date,decision,userid) VALUES (" + id + "," + note + "," + t + "," + button + "," + username + ")")
	res, err := stmt.Exec(id, note, t, button, username)
	log.Println(res)
	LogErr(err)
	affect, rowerr := res.RowsAffected()
	log.Printf("ID = %d, affected = %d\n", id, affect)
	LogErr(rowerr)
}

func UpdateStatus(id string, button string, username string) {
	var addsub string
	if button == "accept" {
		addsub = "+"
	} else if button == "reject" {
		addsub = "-"
	}
	if button == "accept" || button == "reject" {
		stmt, err := formdb.Prepare("UPDATE datasets SET status = status " + addsub + " 1, rejected=1 WHERE id =?")
		LogErr(err)
		log.Println(username + " ran this SQL query: UPDATE datasets SET status = status " + addsub + " 1 WHERE id =" + id + ";")
		res, err := stmt.Exec(id)
		log.Println(res)
		LogErr(err)
		affect, rowerr := res.RowsAffected()
		log.Printf("ID = %d, affected = %d\n", id, affect)
		LogErr(rowerr)
	}
}

func SetBools(id string, datasetnamebool bool, firstnamebool bool, lastnamebool bool, emailbool bool, phonebool bool, firstnamepibool bool, lastnamepibool bool, emailpibool bool, phonepibool bool, abstractbool bool, collectiontitlebool bool, categorytitlebool bool, subcategorytitlebool bool, purposebool bool, otherinfobool bool, keywordsbool bool, placenamesbool bool, filenamebool bool, filetypebool bool, filedescriptionbool bool) {

	exists := CheckRowExists(id)
	if exists == false {
		stmt, err := formdb.Prepare("INSERT INTO checks (datasetid, datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, abstractbool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		LogErr(err)
		log.Println("setting filed status")
		res, err := stmt.Exec(id, datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, abstractbool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool)
		log.Println(res)
		LogErr(err)
	} else if exists == true {
		stmt, err := formdb.Prepare("UPDATE checks set datasetnamebool=?, firstnamebool=?, lastnamebool=?, emailbool=?, phonebool=?, firstnamepibool=?, lastnamepibool=?, emailpibool=?, phonepibool=?, abstractbool=?, collectiontitlebool=?, categorytitlebool=?, subcategorytitlebool=?, purposebool=?, otherinfobool=?, keywordsbool=?, placenamesbool=?, filenamebool=?, filetypebool=?, filedescriptionbool=? WHERE datasetid=?")
		LogErr(err)
		log.Println("setting filed status")
		res, err := stmt.Exec(datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, abstractbool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool, id)
		log.Println(res)
		LogErr(err)
	}

}

func SendMail(id string, datasetname string, username string, t string, note string, status string, button string, emailbool bool, emailpibool bool, host string) {

	var recipients, anote, buttonlink, buttontext, message, realname string
	if button == "accept" {
		if status == "2" {
			message = `Dataset ID: ` + id + ` has passed inspection by ` + username + ` and is ready for your inspection.`
			anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="green">` + button + `</font></p></div>`
			buttonlink = `https://`+host+`/formedit/edit?id=` + id
			buttontext = "Dataset:" + id

			for _, recip := range strings.Fields(configf.Managers) {
				realname, email := ContactFromUserID(recip)
				recipients = recipients + `"`+realname + `" <` + email + `>, `
			}

		} else if status == "3" {

			message = `Dataset ID: ` + id + ` has passed inspection by ` + username + ` and is ready for insertion into GSToRE.`

			anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="green">` + button + `</font></p></div>`
			buttonlink = `https://`+host+`/formedit/edit?id=` + id
			buttontext = "Dataset:" + id

			for _, recip := range strings.Fields(configf.Admins) {
				realname, email := ContactFromUserID(recip)
				recipients = recipients + realname + ` <` + email + `>, `
			}

		}
	} else if button == "reject" {
			if status == "2" {
			message = `I am sorry to inform you that dataset id: "` + id + `" with dataset name: "` + datasetname + `" has been flagged for edit and has had it status changed to "In Progress". Please log into reporting.nmepscor.org, correct any problems that were found, and resubmit for approval. Please feel free to reply to this e-mail with any questions or concerns.`
			anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username  + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="red">Correction Required</font></p></div>`
                  	buttonlink = "https://reporting.nmepscor.org/datasetentry"
                	buttontext = "DataSet Entry Form"
                        realname, email:=ContactFromID(id)
			recipients = recipients + realname + ` <` + email + `>`
			} else if status == "3" {

                        message = `Dataset id: "` + id + `" with dataset name: "` + datasetname + `" has been rejected by a manager and its status changed to "Submitted for Approval". Please see notes on further steps.`
                        anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username  + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="red">Correction Required</font></p></div>`
                        buttonlink = `https://`+host+`/formedit/edit?id=` + id
                        buttontext = "Dataset:" + id

                        for _, recip := range strings.Fields(configf.Users) {
                                realname, email := ContactFromUserID(recip)
                                recipients = recipients + realname + ` <` + email + `>, `
                        }

			} else if status == "4" {

                        message = `I am sorry to inform you that dataset id: "` + id + `" with dataset name: "` + datasetname + `" has been rejected and status set to "Approved". Please see notes for further steps.`
                        anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username  + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="red">Correction Required</font></p></div>`
                        buttonlink = `https://`+host+`/formedit/edit?id=` + id
                        buttontext = "Dataset:" + id

                        for _, recip := range strings.Fields(configf.Users) {
                                realname, email := ContactFromUserID(recip)
                                recipients = recipients + realname + ` <` + email + `>, `
                        }


			}
	}
        recipients = strings.TrimSuffix(recipients, ", ")
        fmt.Println(recipients + " EPSCoR Data Review " + realname + " " + message+ " " + buttonlink+ " " + buttontext+ " " + anote)
        send(recipients, "EPSCoR Data Review", realname, message, buttonlink, buttontext, anote)
}
