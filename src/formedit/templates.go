package main






const Bootstrap string = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="stylesheet" href="https://reporting.nmepscor.org/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<style>
{{.STYLE}}
body { 
    padding-top: 65px; 
}
.navbar-brand {
  padding: 0px;
}
.navbar-brand>img {
  height: 100%;
  padding: 10px;
  width: auto;
}
</style>
</head>
  <body>
{{.BODY}}
  </body>
</html>`





const StarterTemplate string =`<html>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://reporting.nmepscor.org/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<style>
body { 
    padding-top: 65px; 
}
.navbar-brand {
  padding: 0px;
}
.navbar-brand>img {
  height: 100%;
  padding: 10px;
  width: auto;
}
{{.STYLE}}
</style>
</head>
<body>



<nav class="navbar navbar-inverse navbar-fixed-top">
  <div class="container-fluid">
    <div class="navbar-header">
   <a class="navbar-brand" href="/formedit/"><img src="{{.LOGO}}" alt="NM EPSCoR Logo"></a>
    </div>
    <ul class="nav navbar-nav">
              <li{{.DISABLED1}}><a{{.STATUS1}}>In Progress</a></li>
              <li{{.DISABLED2}}><a{{.STATUS2}}>Submitted for Approval</a></li>
              <li{{.DISABLED3}}><a{{.STATUS3}}>Approved</a></li>
              <li{{.DISABLED4}}><a{{.STATUS4}}>Certified</a></li>
              <li{{.DISABLED5}}><a{{.STATUS5}}>Inserted into GSToRE</a></li>
    </ul>
  </div>
</nav>


<div class="container">

{{.BODY}}
</div>
</body>
</html>`

const RedirectTemplate string =`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html lang="en" xml:lang="en" xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta http-equiv="refresh" content="0;URL={{.RURL}}" />
<title>Redirect</title>
</head>
<body>
</body>
</html>`


const EditTemplate string = `<html>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://reporting.nmepscor.org/css/bootstrap.min.css">
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
body { 
    padding-top: 65px; 
}
.navbar-brand {
  padding: 0px;
}
.navbar-brand>img {
  height: 100%;
  padding: 10px;
  width: auto;
}
</style>


</head>
<body>

<nav class="navbar navbar-inverse navbar-fixed-top">
  <div class="container-fluid">
    <div class="navbar-header">
      <a class="navbar-brand" href="/formedit/"><img src="{{.LOGO}}" alt="NM EPSCoR Logo"></a>
    </div>
    <ul class="nav navbar-nav">
              <li{{.DISABLED1}}><a{{.STATUS1}}>In Progress</a></li>
              <li{{.DISABLED2}}><a{{.STATUS2}}>Submitted for Approval</a></li>
              <li{{.DISABLED3}}><a{{.STATUS3}}>Approved</a></li>
              <li{{.DISABLED4}}><a{{.STATUS4}}>Certified</a></li>
              <li{{.DISABLED5}}><a{{.STATUS5}}>Inserted into GSToRE</a></li>
    </ul>
  </div>
</nav>

<div class="container">


  <form class="form-horizontal" action="/formedit/save" method="post">
    <div class="form-group customborder">
<h1><span class="label label-default">Note History:</span></h1>
    {{.NOTES}}
</div>
    <div class="form-group customborder">
      <h1><span class="label label-default">Dataset ID: {{.ID}}</span></h1>
    <div class="form-group">
      <!--<label class="control-label col-sm-2" for="id">ID:{{.ID}}</label>-->
      <div class="col-sm-6">
        <input type="hidden" class="form-control" name="id" value="{{.ID}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="datasetname">Dataset Name :</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="datasetname" value="{{.DATASETNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=datasetnamebool value=true {{.DATASETNAMEBOOL}}> Dataset name is descriptive of the data</label>
    </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstname">First Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="firstname" value="{{.FIRSTNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=firstnamebool value=true {{.FIRSTNAMEBOOL}}> First Name name is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastname">Last Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="lastname" value="{{.LASTNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=lastnamebool value=true {{.LASTNAMEBOOL}}> Last name is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="email">E-Mail:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="email" value="{{.EMAIL}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=emailbool value=true {{.EMAILBOOL}}> E-Mail appears to be valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phone">Phone:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="phone" value="{{.PHONE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=phonebool value=true {{.PHONEBOOL}}> Phone appears to be valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstnamepi">P.I. First Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="firstnamepi" value="{{.FIRSTNAMEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=firstnamepibool value=true {{.FIRSTNAMEPIBOOL}}> PI first name appears valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastnamepi">P.I. Last Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="lastnamepi" value="{{.LASTNAMEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=lastnamepibool value=true {{.LASTNAMEPIBOOL}}> This is a known PI</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="emailpi">P.I. E-Mail:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="emailpi" value="{{.EMAILPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=emailpibool value=true {{.EMAILPIBOOL}}>E-Mail is a know address of a PI</label>
    </div>


    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phonepi">P.I. Phone Number:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="phonepi" value="{{.PHONEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=phonepibool value=true {{.PHONEPIBOOL}}> Phone number is valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="abstract">Abstract:</label>
      <div class="col-sm-6">
        <textarea readonly rows="8" class="form-control" name="abstract" >{{.ABSTRACT}}</textarea>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=abstractbool value=true {{.ABSTRACTBOOL}}> Abstract is valid and informitave.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="collectiontitle">Collection:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="collectiontitle" value="{{.COLLECTIONTITLE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=collectiontitlebool value=true {{.COLLECTIONTITLEBOOL}}> Collection Title is valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="categorytitle">Category:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="categorytitle" value="{{.CATEGORYTITLE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=categorytitlebool value=true {{.CATEGORYTITLEBOOL}}> Category Title is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="subcategorytitle">Subcategory:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="subcategorytitle" value="{{.SUBCATEGORYTITLE}}" readonly>
      </div>

    <div class="col-sm-4">
      <label><input type="checkbox" name="subcategorytitlebool" value=true {{.SUBCATEGORYTITLEBOOL}}> Subcategory is valid.</label>
    </div>
    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="purpose">Purpose:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="purpose" value="{{.PURPOSE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=purposebool value=true {{.PURPOSEBOOL}}> Purpose is descriptive.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="otherinfo">Other Info:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="otherinfo" value="{{.OTHERINFO}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=otherinfobool value=true {{.OTHERINFOBOOL}}> Other Info is descriptive (can be blank)</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="keywords">Keywords:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="keywords" value="{{.KEYWORDS}}" readonly>
      </div>

    <div class="col-sm-4">
      <label><input type="checkbox" name=keywordsbool value=true {{.KEYWORDSBOOL}}> Keywords are meaningful.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="placenames">Placenames:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="placenames" value="{{.PLACENAMES}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=placenamesbool value=true {{.PLACENAMESBOOL}}> Placenames are descriptive.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filename">Filename:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filename" value="{{.FILENAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filenamebool value=true {{.FILENAMEBOOL}}> Filename is correct(this will be auto later...)</label>
    </div>

    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="filetype">File Type:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filetype" value="{{.FILETYPE}}" readonly>
      </div>
      <div class="col-sm-4">
      <label><input type="checkbox" name=filetypebool value=true {{.FILETYPEBOOL}}> Will also be automated later</label>
     </div>
   </div>


    <div class="form-group">
      <label class="control-label col-sm-2" for="filename">Data:</label>
      <div class="control-label col-sm-6"> {{.DATA}}
      <div class="col-sm-4">
        <label><input type="checkbox" name=filenamebool value=true {{.DATABOOL}}> Data is well formated.</label>
      </div>
    </div>


    
    <div class="form-group">
      <label class="control-label col-sm-2" for="filedescription">File Description:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filedescription" value="{{.FILEDESCRIPTION}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filedescriptionbool value=true {{.FILEDESCRIPTIONBOOL}}> Describes the data</label>
    </div>

    </div>
  </div>

    <div class="form-group customborder">
      <h1><span class="label label-default">Quality Control Note</span></h1>
	<br>
    <div class="form-group">
      <label class="control-label col-sm-2" for="Note">Note sent to researcher:</label>
      <div class="col-sm-9">
        <textarea rows="8" class="form-control" name="note" ></textarea>
      </div>
    </div>


<div class="panel-footer row"><!-- panel-footer -->
    <div class="col-xs-6 text-left">
        <div class="previous">
          <button type="submit" name="button" value="reject" class="btn btn-danger">
              <span class="glyphicon glyphicon-chevron-left"></span>Reject
          </button>
        </div>
    </div>
    <div class="col-xs-6 text-right">
        <div class="next">
          <button type="submit" name="button" value="accept" class="btn btn-success">
              Accept<span class="glyphicon glyphicon-chevron-right"></span>
          </button>
        </div>
    </div>
</div>
</div>

 </div>
</form>
</div>
</body>
</html>`

const Bouncer string=` data:image/png;base64,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`
const Logo string = ` data:image/png;base64,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`

const JSON string = `{
        "description": "{{.DATASETNAME}}",
        "orig_epsg": 26913,
        "taxonomy": "file",
        "basename": "08182015_LRF_NetIntensities-ICPMS_AnimasRiver_Directonly",
        "is_embargoed": "False",
        "spatial": {
            "epsg": 4326,
            "geomtype": "POLYGON",
            "features": 1,
            "records": 201,
            "bbox": "-106.61860000,35.08390000,-106.61860000,35.08390000"
        },
        "sources": [
            {
                "mimetype": "application/x-zip-compressed",
                "files": [
                    "/geodata/energize/uploads/{{.UPLOADUSER}}/{{.FILENAME}}"
                ],
                "set": "original",
                "external": "False",
                "extension": "xlsx"
            }
        ],
        "active": "True",
        "dataone_archive":"True",
        "releasedate":"2017-02-20",
        "author":"{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}",
        "categories": [
            {
                "theme": "{{.CATEGORYTITLE}}",
                "subtheme": "{{.SUBCATEGORYTITLE}}",
                "groupname": "UNM"
            }
        ],
        "apps": ["energize"],
        "standards": [
            "FGDC-STD-001-1998",
            "ISO-19115:2003"
        ],
        "project": "NM EPSCoR R4",
        "services": [],
        "formats": ["xlsx"],
        "metadata": {
            "xml": "{{.RAWXML}}",
            "upgrade": "true",
            "standard": "FGDC-STD-001-1998"
        }
}`

/*params = &valuedict{
COLLECTIONTITLE: collectiontitle, 

FIRSTNAME: Null2String(firstname), LASTNAME: Null2String(lastname), EMAIL: Null2String(email), PHONE: Null2String(phone), 

PURPOSE: Null2String(purpose), OTHERINFO: Null2String(otherinfo), 

KEYWORDS: Null2String(keywords), PLACENAMES: Null2String(placenames), FILENAME: Null2String(filename), FILETYPE: Null2String(filetype), FILEDESCRIPTION: Null2String(filedescription), STEP: Null2String(step), STATUS1: authmap["status1"], DISABLED1: authmap["disabled1"], STATUS2: authmap["status2"], DISABLED2: authmap["disabled2"], STATUS3: authmap["status3"], DISABLED3: authmap["disabled3"], STATUS4: authmap["status4"], DISABLED4: authmap["disabled4"], STATUS5: authmap["status5"], DISABLED5: authmap["disabled5"], NOTES: notes, DATASETNAMEBOOL: ischecked(datasetnamebool), FIRSTNAMEBOOL: ischecked(firstnamebool), LASTNAMEBOOL: ischecked(lastnamebool), EMAILBOOL: ischecked(emailbool), PHONEBOOL: ischecked(phonebool), FIRSTNAMEPIBOOL: ischecked(firstnamepibool), LASTNAMEPIBOOL: ischecked(lastnamepibool), EMAILPIBOOL: ischecked(emailpibool), PHONEPIBOOL: ischecked(phonepibool), ABSTRACTBOOL: ischecked(abstractbool), COLLECTIONTITLEBOOL: ischecked(collectiontitlebool), CATEGORYTITLEBOOL: ischecked(categorytitlebool), SUBCATEGORYTITLEBOOL: ischecked(subcategorytitlebool), PURPOSEBOOL: ischecked(purposebool), OTHERINFOBOOL: ischecked(otherinfobool), KEYWORDSBOOL: ischecked(keywordsbool), PLACENAMESBOOL: ischecked(placenamesbool), FILENAMEBOOL: ischecked(filenamebool), FILETYPEBOOL: ischecked(filetypebool), FILEDESCRIPTIONBOOL: ischecked(filedescriptionbool), LOGO: Logo, DATA: data, DATABOOL: ischecked(databool)}
*/
