package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
//	"os"
)

func SimpleForm(w http.ResponseWriter, r *http.Request) {

	token := getCookieByName(r.Cookies(), cookieid)
	if token !=""{
	var htmlstring string
	auth := isAuthorized(token)

	if auth == true {
		status := r.URL.Query().Get("status")
		if status == "" {

        type secretdict struct {
                BODY string
        }


                params := &secretdict{BODY: "lol"}

        	t := template.New("test")
        	t, err := t.Parse(StarterTemplate)
        	if err != nil {
        	        log.Fatal(err)
        	}
		err = t.Execute(w, params)
		if err != nil {
			panic(err)
		}


/*
page := `<html>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>
<div class="container">
<nav class="navbar navbar-default">
        <div class="container">
          <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
              <span class="sr-only">Toggle navigation</span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">Form Editor</a>
          </div>
          <div class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
              <li><a href="/formedit/?status=one">In Progress</a></li>
              <li><a href="/formedit/?status=two">Submitted for Approval</a></li>
              <li><a href="/formedit/?status=three">Approved</a></li>
              <li><a href="/formedit/?status=ready">Certified</a></li>
              <li><a href="/formedit/?status=done">Inserted into GSToRE</a></li>
            </ul>
          </div>
        </div>
      </nav>
</div>
</body>
</html>
`
*/
			


//			fmt.Fprintln(w, StarterTemplate)

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
//htmlstring = htmlstring + `<a class="btn btn-primary custom" href="/formedit/edit?id=`+id+`">`+id+`</a>`
htmlstring = htmlstring + `<a class="btn btn-primary custom" href="/formedit/edit?id=`+id+`">`+id+`</a>`
					count++
					if count==15{
					htmlstring = htmlstring + `</p>`
					count = 0
					}
				}

				uppage := `<html>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  <style>
   .custom {
     width: 76px !important;
     }
     .centered
     {
       text-align:center;
      }
  </style>

</head>
<body>


<div class="container">


<nav class="navbar navbar-default">
        <div class="container">
          <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
              <span class="sr-only">Toggle navigation</span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">Form Editor</a>
          </div>
          <div class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
              <li><a href="/formedit/?status=one">In Progress</a></li>
              <li><a href="/formedit/?status=two">Submitted for Approval</a></li>
              <li><a href="/formedit/?status=three">Approved</a></li>
              <li><a href="/formedit/?status=ready">Certified</a></li>
              <li><a href="/formedit/?status=done">Inserted into GSToRE</a></li>

            </ul>
          </div>
        </div>
      </nav>
</div>
<div class="centered">
`+htmlstring+`
</div>
</body>
</html>`
				fmt.Fprintln(w, uppage)
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not on the list of cool kids."))
	}


}else{
w.WriteHeader(http.StatusUnauthorized)
w.Write([]byte("You need to be logged in to use this API :("))
}

}

func SaveEdit(w http.ResponseWriter, r *http.Request) {
        token := getCookieByName(r.Cookies(), cookieid)
       	auth := isAuthorized(token)
       	if auth == true {
        id:=r.FormValue("id")
	datasetname:=r.FormValue("datasetname")
	firstname:=r.FormValue("firstname")
	lastname:=r.FormValue("lastname")
	email:=r.FormValue("email")
	phone:=r.FormValue("phone")
	firstnamepi:=r.FormValue("firstnamepi")
	lastnamepi:=r.FormValue("lastnamepi")
	emailpi:=r.FormValue("emailpi")
	phonepi:=r.FormValue("phonepi")
	collectiontitle:=r.FormValue("collectiontitle")
	categorytitle:=r.FormValue("categorytitle")
	subcategorytitle:=r.FormValue("subcategorytitle")
	purpose:=r.FormValue("purpose")
	otherinfo:=r.FormValue("otherinfo")
	keywords:=r.FormValue("keywords")
	placenames:=r.FormValue("placenames")
	filename:=r.FormValue("filename")
	filetype:=r.FormValue("filetype")
	filedescription:=r.FormValue("filedescription")
	step:=r.FormValue("step")
	fmt.Fprintln(w, id + " " +datasetname + " " +firstname + " " +lastname + " " +email + " " +phone + " " +firstnamepi + " " +lastnamepi + " " +emailpi + " " +phonepi + " " +collectiontitle + " " +categorytitle + " " +subcategorytitle + " " +purpose + " " +otherinfo + " " +keywords + " " +placenames + " " +filename + " " +filetype + " " +filedescription + " " + step)
	}

}


func SimpleEdit(w http.ResponseWriter, r *http.Request) {
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


			page := `<html>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<style>
.customborder{
  width: 100%;
  color: #555;
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 20px;
}
</style>


</head>
<body>


<div class="container">


<nav class="navbar navbar-default">
        <div class="container">
          <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
              <span class="sr-only">Toggle navigation</span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">Form Editor</a>
          </div>
          <div class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
              <li><a href="/formedit/?status=one">In Progress</a></li>
              <li><a href="/formedit/?status=two">Submitted for Approval</a></li>
              <li><a href="/formedit/?status=three">Approved</a></li>
              <li><a href="/formedit/?status=ready">Certified</a></li>
              <li><a href="/formedit/?status=done">Inserted into GSToRE</a></li>
            </ul>
          </div><!--/.nav-collapse -->
        </div>
      </nav>
  <form class="form-horizontal" action="/formedit/save" method="post">
    <div class="form-group customborder">
      <h1><span class="label label-default">Dataset</span></h1>
<div class="form-group">
      <label class="control-label col-sm-2" for="id">ID:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="id" value="` + id + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="datasetname">Dataset Name :</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="datasetname" value="` + datasetname + `">
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstname">First Name:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="firstname" value="` + firstname + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastname">Last Name:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="lastname" value="` + lastname + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="email">E-Mail:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="email" value="` + email + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phone">Phone:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="phone" value="` + phone + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstnamepi">P.I. First Name:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="firstnamepi" value="` + firstnamepi + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastnamepi">P.I. Last Name:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="lastnamepi" value="` + lastnamepi + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="emailpi">P.I. E-Mail:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="emailpi" value="` + emailpi + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phonepi">P.I. Phone Number:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="phonepi" value="` + phonepi + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="abstract">Abstract:</label>
      <div class="col-sm-10">
        <textarea rows="8" class="form-control" name="abstract" >` + abstract + `</textarea>
      </div>
    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="collectiontitle">Collection:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="collectiontitle" value="` + collectiontitle + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="categorytitle">Category:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="categorytitle" value="` + categorytitle + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="subcategorytitle">Subcategory:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="subcategorytitle" value="` + subcategorytitle + `" readonly>
      </div>
    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="purpose">Purpose:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="purpose" value="` + purpose + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="otherinfo">Other Info:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="otherinfo" value="` + otherinfo + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="keywords">Keywords:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="keywords" value="` + keywords + `" readonly>
      </div>
    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="placenames">Placenames:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="placenames" value="` + placenames + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filename">Filename:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="filename" value="` + filename + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filetype">File Type:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="filetype" value="` + filetype + `" readonly>
      </div>
    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="filedescription">File Description:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="filedescription" value="` + filedescription + `" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="step">Step:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="step" value="` + step + `" readonly>
      </div>
    </div>



</div>
    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button type="submit" class="btn btn-danger">Save</button>
      </div>
    </div>
  </form>
</div>

</body>
</html>`

			// &emailpi, &phonepi, &abstract)
			//firstnamepi, lastnamepi, emailpi, phonepi, abstract
			//  First Name PI: <input type="text" name="firstnamepi" value="`+firstnamepi+`"><br>

			fmt.Fprintln(w, page)

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
