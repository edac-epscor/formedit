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
{{.BODY}}
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
              <li><a href="/formedit/?status=one">In Progress</a></li>
              <li><a href="/formedit/?status=two">Submitted for Approval</a></li>
              <li><a href="/formedit/?status=three">Approved</a></li>
              <li><a href="/formedit/?status=ready">Certified</a></li>
              <li><a href="/formedit/?status=done">Inserted into GSToRE</a></li>
            </ul>
          </div><!--/.nav-collapse -->
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
      <label><input type="checkbox" name=datasetnamebool value="1"> Dataset name is descriptive</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstname">First Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="firstname" value="{{.FIRSTNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=firstnamebool value="1"> First Name name is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastname">Last Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="lastname" value="{{.LASTNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=lastnamebool value="1"> Last name is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="email">E-Mail:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="email" value="{{.EMAIL}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=emailbool value="1"> E-Mail appears to be valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phone">Phone:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="phone" value="{{.PHONE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=phonebool value="1"> Phone number is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstnamepi">P.I. First Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="firstnamepi" value="{{.FIRSTNAMEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=firstnamepibool value="1"> This is a known PI.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastnamepi">P.I. Last Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="lastnamepi" value="{{.LASTNAMEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=lastnamepibool value="1"> This is a known PI</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="emailpi">P.I. E-Mail:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="emailpi" value="{{.EMAILPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=emailpibool value="1">E-Mail is a know address of a PI</label>
    </div>


    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phonepi">P.I. Phone Number:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="phonepi" value="{{.PHONEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=phonepibool value="1"> Phone number is valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="abstract">Abstract:</label>
      <div class="col-sm-6">
        <textarea readonly rows="8" class="form-control" name="abstract" >{{.ABSTRACT}}</textarea>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=abstractbool value="1"> Abstract is valid and informitave.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="collectiontitle">Collection:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="collectiontitle" value="{{.COLLECTIONTITLE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=collectiontitlebool value="1"> Collection Title number is valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="categorytitle">Category:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="categorytitle" value="{{.CATEGORYTITLE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=categorytitlebool value="1"> Category Title is valid</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="subcategorytitle">Subcategory:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="subcategorytitle" value="{{.SUBCATEGORYTITLE}}" readonly>
      </div>

    <div class="col-sm-4">
      <label><input type="checkbox" name=subcatetorytitlebool value="1"> Subcategory is valid.</label>
    </div>
    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="purpose">Purpose:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="purpose" value="{{.PURPOSE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=purposebool value="1"> Purpose is descriptive.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="otherinfo">Other Info:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="otherinfo" value="{{.OTHERINFO}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=otherinfobool value="1"> Other Info is descriptive (can be N/A)</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="keywords">Keywords:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="keywords" value="{{.KEYWORDS}}" readonly>
      </div>

    <div class="col-sm-4">
      <label><input type="checkbox" name=keywordsbool value="1"> Keywords are meaningful.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="placenames">Placenames:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="placenames" value="{{.PLACENAMES}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=placenamesbool value="1"> Placenames are descriptive.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filename">Filename:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filename" value="{{.FILENAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filenamebool value="1"> Filename is correct(this will be auto later...)</label>
    </div>

    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="filetype">File Type:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filetype" value="{{.FILETYPE}}" readonly>
   </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filetypebool value="1"> Will also be automated later</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filedescription">File Description:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filedescription" value="{{.FILEDESCRIPTION}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filedescriptionbool value="1"> What even is this?</label>
    </div>

    </div>

  </div>
    <div class="form-group customborder">
      <h1><span class="label label-default">Quality Control Note</span></h1>
	<br>
    <div class="form-group">
      <label class="control-label col-sm-2" for="Note">Note sent to researcher:</label>
      <div class="col-sm-10">
        <textarea rows="8" class="form-control" name="note" ></textarea>
      </div>
    </div>


<div class="panel-footer row"><!-- panel-footer -->
    <div class="col-xs-6 text-left">
        <div class="previous">
          <button type="submit" class="btn btn-danger">
              <span class="glyphicon glyphicon-chevron-left"></span>Reject
          </button>
        </div>
    </div>
    <div class="col-xs-6 text-right">
        <div class="next">
          <button type="submit" class="btn btn-success">
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
