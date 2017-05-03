package main

import (
	"bufio"
	"bytes"
	"database/sql"
	//	"io"
	//	"encoding/json"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
	//    "github.com/tealeg/xlsx"
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
                NUMRECORDS string
	}
        var numrecords int
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
                                                numrecords++
						var id, datasetname, firstname, lastname, email, datecreated string
						err = rows.Scan(&id, &datasetname, &firstname, &lastname, &email, &datecreated)
						body = body + `<a href="/formedit/edit?id=` + id + `" class="list-group-item list-group-item-action"><span class="badge badge-default badge-pill">Dataset ID:` + id + `</span></p>Dataset Name: <strong>` + datasetname + `</strong></p>Submitted by: <strong>` + firstname + ` ` + lastname + `</p></strong>E-Mail:<strong>` + email + `  </p></strong><small class="text-muted">Dataset Created: ` + datecreated + `</small></a>`
					}
					body = `<div class="list-group list-group-item-action">` + body + `</div>`
					style := `<style>
                                                 .custom{
                                                  width:76px!important;
                                                 }
                                                  .centered
                                                 {
                                                   text-align:center;
                                                 }
                                                 </style>`

					params := &secretdict{BODY: body, STYLE: style, STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"], LOGO: Logo, NUMRECORDS: strconv.Itoa(numrecords)}
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
	host := r.Host
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
                realname, _ := ContactFromUserID(username)
		if button == "accept" || button == "reject" {
			UpdateStatus(id, button, username)
			MakeNote(id, note, t, button, username)
			SetBools(id, datasetnamebool, firstnamebool, lastnamebool, emailbool, phonebool, firstnamepibool, lastnamepibool, emailpibool, phonepibool, abstractbool, collectiontitlebool, categorytitlebool, subcategorytitlebool, purposebool, otherinfobool, keywordsbool, placenamesbool, filenamebool, filetypebool, filedescriptionbool)
			SendMail(id, datasetname, realname, t, note, stat, button, emailbool, emailpibool, host)
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
                SUBMITTERNAME        string
                SUBMITTEREMAIL       string
                DATECREATED          string
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

		var id, datasetname, userid, status, categorytitle, subcategorytitle, firstname, lastname, email, phone, firstnamepi, lastnamepi, emailpi, phonepi, abstract, purpose, otherinfo, keywords, placenames, filename, filetype, filedescription, step, datecreated sql.NullString
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
			query = `SELECT datasets.id, datasets.datasetname, datasets.userid, datasets.status, collections.collectiontitle, categorys.categorytitle, subcategorys.subcategorytitle, datasets.firstname, datasets.lastname, datasets.email, datasets.phone, datasets.firstnamepi, datasets.lastnamepi, datasets.emailpi, datasets.phonepi, datasets.abstract, datasets.purpose, datasets.otherinfo, datasets.keywords, datasets.placenames, datasets.filename, datasets.filetype, datasets.filedescription, datasets.step, datasets.datecreated FROM datasets, collections, categorys, subcategorys WHERE datasets.id = '` + ID + `' AND datasets.collectionid=collections.id AND datasets.categoryid=categorys.id AND datasets.subcategoryid=subcategorys.id;`
		} else {
			query = `SELECT datasets.id, datasets.datasetname, datasets.userid, datasets.status, categorys.categorytitle, subcategorys.subcategorytitle, datasets.firstname, datasets.lastname, datasets.email, datasets.phone, datasets.firstnamepi, datasets.lastnamepi, datasets.emailpi, datasets.phonepi, datasets.abstract, datasets.purpose, datasets.otherinfo, datasets.keywords, datasets.placenames, datasets.filename, datasets.filetype, datasets.filedescription, datasets.step, datasets.datecreated FROM datasets, categorys, subcategorys WHERE datasets.id = '` + ID + `' AND datasets.categoryid=categorys.id AND datasets.subcategoryid=subcategorys.id;`
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
				err = rows.Scan(&id, &datasetname, &userid, &status, &collectiontitle, &categorytitle, &subcategorytitle, &firstname, &lastname, &email, &phone, &firstnamepi, &lastnamepi, &emailpi, &phonepi, &abstract, &purpose, &otherinfo, &keywords, &placenames, &filename, &filetype, &filedescription, &step, &datecreated)
			} else {
				err = rows.Scan(&id, &datasetname, &userid, &status, &categorytitle, &subcategorytitle, &firstname, &lastname, &email, &phone, &firstnamepi, &lastnamepi, &emailpi, &phonepi, &abstract, &purpose, &otherinfo, &keywords, &placenames, &filename, &filetype, &filedescription, &step, &datecreated)
				collectiontitle = "INVALID COLLECTION TITLE!!!"
			}
			LogErr(err)
                        submittername, submitteremail := ContactFromUserID(IDtoName(Null2String(userid)))
			if _, err := os.Stat("/uploads/" + IDtoName(Null2String(userid)) + "/" + Null2String(filename)); err == nil {
				data = `<a href="/formedit/download/` + ID + `">Data Download</a>`
			} else {
				//data = `<font color="grey">/uploads/` + IDtoName(Null2String(userid)) + "/" + Null2String(filename)+`
				data = `<h4><font color="red">File not found! </font><small class="text-muted">/uploads/` + IDtoName(Null2String(userid)) + "/" + Null2String(filename) + `</small></h4>`

				//<font color="blue"> in ` + "/uploads/" + IDtoName(Null2String(userid)) + "/" + Null2String(filename)</font>
			}

			params = &valuedict{ID: Null2String(id), DATASETNAME: Null2String(datasetname), COLLECTIONTITLE: collectiontitle, CATEGORYTITLE: Null2String(categorytitle), SUBCATEGORYTITLE: Null2String(subcategorytitle), FIRSTNAME: Null2String(firstname), LASTNAME: Null2String(lastname), EMAIL: Null2String(email), PHONE: Null2String(phone), FIRSTNAMEPI: Null2String(firstnamepi), LASTNAMEPI: Null2String(lastnamepi), EMAILPI: Null2String(emailpi), PHONEPI: Null2String(phonepi), ABSTRACT: Null2String(abstract), PURPOSE: Null2String(purpose), OTHERINFO: Null2String(otherinfo), KEYWORDS: Null2String(keywords), PLACENAMES: Null2String(placenames), FILENAME: Null2String(filename), FILETYPE: Null2String(filetype), FILEDESCRIPTION: Null2String(filedescription), STEP: Null2String(step), STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"], NOTES: notes, DATASETNAMEBOOL: ischecked(datasetnamebool), FIRSTNAMEBOOL: ischecked(firstnamebool), LASTNAMEBOOL: ischecked(lastnamebool), EMAILBOOL: ischecked(emailbool), PHONEBOOL: ischecked(phonebool), FIRSTNAMEPIBOOL: ischecked(firstnamepibool), LASTNAMEPIBOOL: ischecked(lastnamepibool), EMAILPIBOOL: ischecked(emailpibool), PHONEPIBOOL: ischecked(phonepibool), ABSTRACTBOOL: ischecked(abstractbool), COLLECTIONTITLEBOOL: ischecked(collectiontitlebool), CATEGORYTITLEBOOL: ischecked(categorytitlebool), SUBCATEGORYTITLEBOOL: ischecked(subcategorytitlebool), PURPOSEBOOL: ischecked(purposebool), OTHERINFOBOOL: ischecked(otherinfobool), KEYWORDSBOOL: ischecked(keywordsbool), PLACENAMESBOOL: ischecked(placenamesbool), FILENAMEBOOL: ischecked(filenamebool), FILETYPEBOOL: ischecked(filetypebool), FILEDESCRIPTIONBOOL: ischecked(filedescriptionbool), LOGO: Logo, DATA: data, DATABOOL: ischecked(databool), SUBMITTERNAME:submittername, SUBMITTEREMAIL:submitteremail, DATECREATED:Null2String(datecreated)}

			t, err = t.Parse(EditTemplate)
			LogErr(err)
		}

		err = t.Execute(w, params)
		LogErr(err)

	}

}

func Download(w http.ResponseWriter, r *http.Request) {
	token := getCookieByName(r.Cookies(), cookieid)
	auth, _ := isAuthorized(token)
	if auth >= 1 {

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
		attachment := `attachment; filename="` + filename + `"`
		fmt.Println(attachment)
		w.Header().Set("Content-Disposition", attachment)
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
		http.ServeFile(w, r, pathtofile)
	}
}

func InsertView(w http.ResponseWriter, r *http.Request) {
	type jsondict struct {
		DATASETNAME      string
		UPLOADUSER       string
		FILENAME         string
		FIRSTNAMEPI      string
		LASTNAMEPI       string
		CATEGORYTITLE    string
		SUBCATEGORYTITLE string
		RAWXML           string
		BASENAME         string
		ISEMBARGOED      string
		RELEASEDATE      string
		ORIGEPSG         string
		FEATURES         string
		GEOMTYPE         string
		RECORDS          string
		EPSG             string
		WESTBC           string
		EASTBC           string
		NORTHBC          string
		SOUTHBC          string
		EXTENSION        string
		DATAONEARCHIVE   string
		GROUPNAME        string
		COLLECTIONTITLE  string
		ABSTRACT         string
		ADDRESS          string
		CITY             string
		COLLECTIONNAME   string
		TITLE            string
		EMAILPI          string
		PHONEPI          string
		PURPOSE          string
		STATE            string
		ZIP              string
		GEOFORM          string
		ATTRIBUTES       string
	}

	var datasetname, uploaduser, firstnamepi, filetype, lastnamepi, categorytitle, subcategorytitle, rawxml, userid, filename, basename, isembargoed, releasedate, lat, lon, westbc, eastbc, northbc, southbc, dataonearchive, uploadtodataone, groupname, collectiontitle, abstract, purpose, city, state, zip, phonepi, emailpi, attributes string
	var releasedatenull, address1, address2, address3 sql.NullString

	//Bill has these hardcoded, so I am to for now, but this should be dynamic no?
	//Why is records randomly set to 201?
	epsg := "4326"
	origepsg := "26913"
	features := "1"
	geomtype := "POLYGON"
	records := "201" //<-this is the numbe of records?
	geoform := "spreadsheet"

	id := mux.Vars(r)["id"]
	post := r.URL.Query().Get("post")
	token := getCookieByName(r.Cookies(), cookieid)
	if token != "" {
		auth, _ := isAuthorized(token)
		if auth >= 3 {
			query := `SELECT userid, filename, embargoreleasedate FROM datasets where id='` + id + `';`
			err := formdb.QueryRow(query).Scan(&userid, &filename, &releasedatenull)
			if Null2String(releasedatenull) == "NULL" {
				isembargoed = "False"
				releasedate = string(time.Now().Format("2006-01-02"))
			} else {
				isembargoed = "True"
				releasedate = Null2String(releasedatenull)
			}
			LogErr(err)

			id := mux.Vars(r)["id"]
			query = `SELECT datasets.userid, datasets.filename, datasets.filetype, datasets.datasetname, institutions.latitude, institutions.longitude, datasets.uploadtodataone, datasets.firstnamepi, datasets.lastnamepi, categorys.categorytitle, subcategorys.subcategorytitle, institutions.instName_short, collections.collectiontitle, datasets.abstract, datasets.purpose, institutions.address_1, institutions.address_2, institutions.address_3, institutions.city, institutions.state, institutions.zipcode, datasets.phonepi, datasets.emailpi FROM datasets, institutions, categorys, subcategorys, collections WHERE datasets.institutionid=institutions.id AND datasets.categoryid=categorys.id AND datasets.subcategoryid=subcategorys.id AND datasets.collectionid=collections.id AND datasets.id='` + id + `';`
			err = formdb.QueryRow(query).Scan(&userid, &filename, &filetype, &datasetname, &lat, &lon, &uploadtodataone, &firstnamepi, &lastnamepi, &categorytitle, &subcategorytitle, &groupname, &collectiontitle, &abstract, &purpose, &address1, &address2, &address3, &city, &state, &zip, &phonepi, &emailpi)
			fmt.Println(filetype)
			if filetype == "*.csv" {
				if _, err := os.Stat("/uploads/" + IDtoName(userid) + "/" + filename); err == nil {
					f, _ := os.Open("/uploads/" + IDtoName(userid) + "/" + filename)

					csvr := csv.NewReader(bufio.NewReader(f))
					records, err := csvr.ReadAll()
					LogErr(err)
					firstlabel := records[0][0]

					alllabels := records[0]
					//datecolumn := ""
					var datefield, timefield int
					for i := 0; i < len(alllabels); i += 1 {
						lab := alllabels[i]
						//fmt.Println(lab)
						if strings.ToLower(lab) == "date" {
							datefield = i
						}
						if strings.ToLower(lab) == "time" {
							timefield = i
						}
					}
					fmt.Println(getcurrentdataformat(records, firstlabel, datefield))
					for i := range records {
						// Element count.

						//fmt.Printf("Elements: %v", len(records[i]))
						//fmt.Println()
						// Elements.
						if records[i][0] != firstlabel {

							var stime string
							stime = records[i][datefield] + " " + records[i][timefield]
							timeindex, err := time.Parse("1/2/2006 15:04:05", stime)
							LogErr(err)
							observed := timeindex.Format("20060102T15:04:05")
							_ = observed
							//							fmt.Println(records[i][1])
						}
					}
					//              fmt.Println(csvr[0])
					//	for {
					//		record, err := csvr.Read()
					//              fmt.Println(record[0])
					// Stop at EOF.
					//		if err == io.EOF {
					//			break
					//		}
					// Display record.
					// ... Display record length.
					// ... Display all individual elements of the slice.
					//	fmt.Println(record)
					//	fmt.Println(len(record))
					//	for value := range record {
					//		fmt.Printf("  %v\n", record[value])
					//	}
					//					}
					//
					//				} else {
					//					fmt.Println("file not found!")
				}

			} else {
				//fmt.Println("Not a csv")

				//			LogErr(err)

				basename = strings.TrimSuffix(filename, filepath.Ext(filename))
				if uploadtodataone == "Yes" {
					dataonearchive = "True"
				} else if uploadtodataone == "No" {
					dataonearchive = "False"
				}
				westbc = lon
				eastbc = lon
				northbc = lat
				southbc = lat
				uploaduser = IDtoName(userid)
				extension := strings.TrimPrefix(filetype, "*.")

				var address string
				address = MakeAddr(address1, address)
				address = MakeAddr(address2, address)
				address = MakeAddr(address3, address)
				query = `SELECT field, description, units, frequency, aggregation, nodata, dommin, dommax FROM fieldinfo WHERE datasetid='` + id + `';`
				rows, err := formdb.Query(query)
				LogErr(err)
				piname := firstnamepi + ` ` + lastnamepi
				for rows.Next() {
					var field, description, units, frequency, aggregation, nodata, dommin, dommax string
					err = rows.Scan(&field, &description, &units, &frequency, &aggregation, &nodata, &dommin, &dommax)
					attributemap := map[string]string{
						"field":       field,
						"description": description,
						"units":       units,
						"frequency":   frequency,
						"aggregation": aggregation,
						"nodata":      nodata,
						"dommin":      dommin,
						"dommax":      dommax,
						"attrdefs":    piname,
					}

					attributemap = NormalizeAttributes(attributemap)

					attributes = attributes + BuildAttributes(attributemap)
				}

				title := groupname + ` ` + categorytitle + `, ` + collectiontitle + ` - ` + datasetname

				params := &jsondict{DATASETNAME: datasetname, UPLOADUSER: uploaduser, FILENAME: filename, FIRSTNAMEPI: firstnamepi, LASTNAMEPI: lastnamepi, CATEGORYTITLE: categorytitle, SUBCATEGORYTITLE: subcategorytitle, RAWXML: rawxml, BASENAME: basename, ISEMBARGOED: isembargoed, RELEASEDATE: releasedate, ORIGEPSG: origepsg, FEATURES: features, GEOMTYPE: geomtype, RECORDS: records, EPSG: epsg, WESTBC: westbc, EASTBC: eastbc, NORTHBC: northbc, SOUTHBC: southbc, EXTENSION: extension, DATAONEARCHIVE: dataonearchive, GROUPNAME: groupname, TITLE: title, GEOFORM: geoform, ABSTRACT: abstract, PURPOSE: purpose, ADDRESS: address, CITY: city, STATE: state, ZIP: zip, PHONEPI: phonepi, EMAILPI: emailpi, COLLECTIONTITLE: collectiontitle, ATTRIBUTES: attributes}

				jsonbuf := new(bytes.Buffer)
				jt := template.Must(template.New("json").Parse(JSON))
				err = jt.Execute(jsonbuf, params)
				LogErr(err)
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonbuf.Bytes())
				xmlbuf := new(bytes.Buffer)
				xmlt := template.Must(template.New("xml").Parse(GSTOREXML))
				err = xmlt.Execute(xmlbuf, params)
				fmt.Println(xml.Marshal(xmlbuf.Bytes()))
				if strings.EqualFold(post, "True") {
					fmt.Println("Inserting")
					req, err := http.NewRequest("POST", "http://129.24.63.66/gstore_v3/apps/energize/datasets", jsonbuf)
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("Accept-Encoding", "gzip, deflate")
					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()

					fmt.Println("response Status:", resp.Status)
					fmt.Println("response Headers:", resp.Header)
					body, _ := ioutil.ReadAll(resp.Body)
					fmt.Println("response Body:", string(body))

				}
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

func isAuthorized(token string) (int, string) {
	var username string
	query := "SELECT name FROM users u INNER JOIN sessions s ON u.uid = s.uid WHERE s.sid = '" + token + "';"
        fmt.Println(query)
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
		log.Println(err)
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
        fmt.Println("***************************************************************")
        fmt.Println(params)
        fmt.Println(firstname)
       	fmt.Println("***************************************************************")
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
        fmt.Println("name to lookup is " + name)
	var realname, email string
	query := `select realname.realname, users.mail from realname, users where users.uid=realname.uid and users.name='` + name + `';`
        fmt.Println(query)
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

	var recipients, anote, buttonlink, buttontext, message, realname, email string
	if button == "accept" {
		if status == "2" {
			message = `Dataset ID "` + id + `, titled "` + datasetname + `" has passed inspection by ` + username + ` and is ready for your inspection.`
			//anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="green">` + button + `</font></p></div>`
                        anote = `<p class="lead">`+ note + `<font color="grey"><h5>Approved: `+ t + `</h5></font></p><br><br>`
			buttonlink = `https://` + host + `/formedit/edit?id=` + id
			buttontext = "Click here to view and approve"

			for _, recip := range strings.Fields(configf.Managers) {
				realname, email = ContactFromUserID(recip)
				recipients = recipients + `"` + realname + `" <` + email + `>, `
			}

		} else if status == "3" {

			message = `Dataset ID: ` + id + `, titled "` + datasetname + `" has passed inspection by ` + username + ` and is ready for insertion into GSToRE.`
			//anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="green">` + button + `</font></p></div>`
                        anote = `<p class="lead">`+ note + `<font color="grey"><h5>Approved: `+ t + `</h5></font></p><br><br>`
			buttonlink = `https://` + host + `/formedit/edit?id=` + id
			//buttontext = "Dataset:" + id
                        buttontext = "Click here to view and approve"
			for _, recip := range strings.Fields(configf.Admins) {
				realname, email = ContactFromUserID(recip)
				recipients = recipients + realname + ` <` + email + `>, `
			}

		}
	} else if button == "reject" {
		if status == "2" {
			message = `I am sorry to inform you that dataset id "` + id + `" titled "` + datasetname + `" has been rejected. </br> Please log into reporting.nmepscor.org and fix the items listed below in the "Notes" section then re-submit your dataset.`
//			anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="red">Correction Required</font></p></div>`
                        anote = `<p class="lead">`+ note + `<font color="grey"><h5>Date Rejected `+ t + `</h5></font></p><br><br>`
			buttonlink = "https://reporting.nmepscor.org"
			buttontext = "Click here to log in"
			realname, email = ContactFromID(id)
			recipients = recipients + realname + ` <` + email + `>`
		} else if status == "3" {

			message = `Dataset id: "` + id + `" with dataset name: "` + datasetname + `" has been rejected by ` + username + ` and its status changed to "Submitted for Approval". Please see notes on further steps.`
                        //anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="red">Correction Required</font></p></div>`
                        anote = `<p class="lead">`+ note + `<font color="grey"><h5>Date Rejected `+ t + `</h5></font></p><br><br>`
			buttonlink = `https://` + host + `/formedit/edit?id=` + id
			buttontext = "Click here to view and approve."

			for _, recip := range strings.Fields(configf.Users) {
				realname, email = ContactFromUserID(recip)
				recipients = recipients + realname + ` <` + email + `>, `
			}

		} else if status == "4" {

			message = `I am sorry to inform you that dataset id: "` + id + `" with dataset name: "` + datasetname + `" has been rejected and status set to "Approved". Please see notes for further steps.`
       //			anote = `<div class="well"><p class="lead"><font color="grey"><h5>` + username + ` ` + t + `: </h5></font>` + note + `<br><br>Decision:<font color="red">Correction Required</font></p></div>`
                        anote = `<p class="lead">`+ note + `<font color="grey"><h5>Date Rejected `+ t + `</h5></font></p><br><br>`
			buttonlink = `https://` + host + `/formedit/edit?id=` + id
			buttontext = "Click here to view and approve."

			for _, recip := range strings.Fields(configf.Users) {
				realname, email = ContactFromUserID(recip)
				recipients = recipients + realname + ` <` + email + `>, `
			}

		}
	}
	recipients = strings.TrimSuffix(recipients, ", ")
	subject := "EPSCoR Data Review of " + datasetname
	//fmt.Println(recipients + " EPSCoR Data Review " + realname + " " + message + " " + buttonlink + " " + buttontext + " " + anote)
	send(recipients, subject, realname, message, buttonlink, buttontext, anote)
}

func MakeAddr(addr sql.NullString, newaddr string) string {

	if Null2String(addr) != "NULL" {
		newaddr = newaddr + `<address>` + Null2String(addr) + `</address>`
	}
	return newaddr
}

func NormalizeAttributes(attr map[string]string) map[string]string {
	fmt.Println(attr["attrdefs"])

	if len(attr["rdommin"]) == 0 {
		attr["rdommin"] = "NA"
	}
	if len(attr["rdommax"]) == 0 {
		attr["rdommax"] = "NA"
	}
	if len(attr["nodata"]) == 0 {
		attr["nodata"] = "NA"
	}
	if len(attr["aggregation"]) == 0 {
		attr["aggregation"] = "NA"
	}
	if len(attr["frequency"]) == 0 {
		attr["frequency"] = "NA"
	}
	if len(attr["units"]) == 0 {
		attr["units"] = "NA"
	}
	if len(attr["description"]) == 0 {
		attr["description"] = "NA"
	}
	if len(attr["field"]) == 0 {
		attr["field"] = "NA"
	}
	//	for k, v := range attr {
	//		if len(v) == 0 {
	//			attr[k] = "NA"
	//		}
	//	fmt.Println(attr["rdommin"])
	//	}
	//      fmt.Println(attr["rdommin"])
	return attr
}

func BuildAttributes(attr map[string]string) string {

	//for k, v := range attr{
	xml := `<attr>
        <attrlabl>` + attr["field"] + `</attrlabl>
        <attrdef>` + attr["description"] + `</attrdef>
        <attrdefs>` + attr["attrdefs"] + `</attrdefs>
        <attrdomv>
          <rdom>
            <rdommin>` + attr["rdommin"] + `</rdommin>
            <rdommax>` + attr["rdommax"] + `</rdommax>
            <attrunit>` + attr["units"] + `</attrunit>
          </rdom>
        </attrdomv>
      </attr>`
	//	}
	return xml
}

func getcurrentdataformat(records [][]string, firstlabel string, datefield int) string {
	var format, sdate string
	for i := range records {
		if records[i][0] != firstlabel {
			sdate = records[i][datefield] // + " " + records[i][timefield]
			matched, err := regexp.MatchString("[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]", sdate)
			LogErr(err)
			if matched == true {
				//split the string by -
				// grab the second set ie position 1
				//are there any larger numbers than 12?
				//yes than format is "2006-2-1"
				//else if
				// grab 3rd set pos 2
				//are there any larger numbers than 12?
				//yes than format is "2006-1-2"

				format = "2006-1-2"

			}
			matched, err = regexp.MatchString("[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9]", sdate)
			LogErr(err)
			if matched == true {
				format = "2006/1/2"

			}
			matched, err = regexp.MatchString("[0-9][0-9]/[0-9][0-9]/[0-9][0-9][0-9][0-9]", sdate)
			LogErr(err)
			if matched == true {
				format = "1/2/2006"

			}
			matched, err = regexp.MatchString("[0-9][0-9]-[0-9][0-9]-[0-9][0-9][0-9][0-9]", sdate)
			LogErr(err)
			if matched == true {
				format = "1-2-2006"

			}

			LogErr(err)

			//     timeindex, err := time.Parse("1/2/2006 15:04:05", stime)
			//   LogErr(err)
			// observed := timeindex.Format("20060102T15:04:05")
			// fmt.Println(observed)
			//                                                      fmt.Println(records[i$
		}
	}

	return format
}
