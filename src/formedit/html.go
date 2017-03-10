package main

const StarterTemplate string =`<html>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
{{.STYLE}}
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
              <li{{.DISABLED1}}><a{{.STATUS1}}>In Progress</a></li>
              <li{{.DISABLED2}}><a{{.STATUS2}}>Submitted for Approval</a></li>
              <li{{.DISABLED3}}><a{{.STATUS3}}>Approved</a></li>
              <li{{.DISABLED4}}><a{{.STATUS4}}>Certified</a></li>
              <li{{.DISABLED5}}><a{{.STATUS5}}>Inserted into GSToRE</a></li>
            </ul>
          </div>
        </div>
      </nav>
</div>
{{.BODY}}
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
              <li{{.DISABLED1}}><a{{.STATUS1}}>In Progress</a></li>
              <li{{.DISABLED2}}><a{{.STATUS2}}>Submitted for Approval</a></li>
              <li{{.DISABLED3}}><a{{.STATUS3}}>Approved</a></li>
              <li{{.DISABLED4}}><a{{.STATUS4}}>Certified</a></li>
              <li{{.DISABLED5}}><a{{.STATUS5}}>Inserted into GSToRE</a></li>
            </ul>
          </div>
        </div>
      </nav>

<div class="container">


  <form class="form-horizontal" action="/formedit/save" method="post">
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
      <label><input type="checkbox" name=datasetnamebool value=true> Dataset name is descriptive of the data</label>
    </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstname">First Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="firstname" value="{{.FIRSTNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=firstnamebool value=true> First Name name is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastname">Last Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="lastname" value="{{.LASTNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=lastnamebool value=true> Last name is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="email">E-Mail:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="email" value="{{.EMAIL}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=emailbool value=true> E-Mail appears to be valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phone">Phone:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="phone" value="{{.PHONE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=phonebool value=true> Phone appears to be valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstnamepi">P.I. First Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="firstnamepi" value="{{.FIRSTNAMEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=firstnamepibool value=true> PI first name appears valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastnamepi">P.I. Last Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="lastnamepi" value="{{.LASTNAMEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=lastnamepibool value=true> This is a known PI</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="emailpi">P.I. E-Mail:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="emailpi" value="{{.EMAILPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=emailpibool value=true>E-Mail is a know address of a PI</label>
    </div>


    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phonepi">P.I. Phone Number:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="phonepi" value="{{.PHONEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=phonepibool value=true> Phone number is valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="abstract">Abstract:</label>
      <div class="col-sm-6">
        <textarea readonly rows="8" class="form-control" name="abstract" >{{.ABSTRACT}}</textarea>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=abstractbool value=true> Abstract is valid and informitave.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="collectiontitle">Collection:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="collectiontitle" value="{{.COLLECTIONTITLE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=collectiontitlebool value=true> Collection Title is valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="categorytitle">Category:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="categorytitle" value="{{.CATEGORYTITLE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=categorytitlebool value=true> Category Title is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="subcategorytitle">Subcategory:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="subcategorytitle" value="{{.SUBCATEGORYTITLE}}" readonly>
      </div>

    <div class="col-sm-4">
      <label><input type="checkbox" name="subcategorytitlebool" value=true> Subcategory is valid.</label>
    </div>
    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="purpose">Purpose:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="purpose" value="{{.PURPOSE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=purposebool value=true> Purpose is descriptive.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="otherinfo">Other Info:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="otherinfo" value="{{.OTHERINFO}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=otherinfobool value=true> Other Info is descriptive (can be N/A)</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="keywords">Keywords:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="keywords" value="{{.KEYWORDS}}" readonly>
      </div>

    <div class="col-sm-4">
      <label><input type="checkbox" name=keywordsbool value=true> Keywords are meaningful.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="placenames">Placenames:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="placenames" value="{{.PLACENAMES}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=placenamesbool value=true> Placenames are descriptive.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filename">Filename:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filename" value="{{.FILENAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filenamebool value=true> Filename is correct(this will be auto later...)</label>
    </div>

    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="filetype">File Type:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filetype" value="{{.FILETYPE}}" readonly>
   </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filetypebool value=true> Will also be automated later</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filedescription">File Description:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filedescription" value="{{.FILEDESCRIPTION}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filedescriptionbool value=true> Describes the data</label>
    </div>

    </div>
  </div>

    <div class="form-group customborder">
      {{.NOTES}}
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
</form>
</div>
</body>
</html>`
