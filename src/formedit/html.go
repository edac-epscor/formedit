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
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<style>
{{.STYLE}}
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

const Bouncer string=` data:image/png;base64,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`
