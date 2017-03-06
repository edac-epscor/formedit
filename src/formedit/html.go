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
        <input type="text" class="form-control" name="id" value="{{.ID}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="datasetname">Dataset Name :</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="datasetname" value="{{.DATASETNAME}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstname">First Name:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="firstname" value="{{.FIRSTNAME}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastname">Last Name:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="lastname" value="{{.LASTNAME}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="email">E-Mail:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="email" value="{{.EMAIL}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phone">Phone:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="phone" value="{{.PHONE}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstnamepi">P.I. First Name:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="firstnamepi" value="{{.FIRSTNAMEPI}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastnamepi">P.I. Last Name:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="lastnamepi" value="{{.LASTNAMEPI}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="emailpi">P.I. E-Mail:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="emailpi" value="{{.EMAILPI}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phonepi">P.I. Phone Number:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="phonepi" value="{{.PHONEPI}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="abstract">Abstract:</label>
      <div class="col-sm-10">
        <textarea readonly rows="8" class="form-control" name="abstract" >{{.ABSTRACT}}</textarea>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="collectiontitle">Collection:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="collectiontitle" value="{{.COLLECTIONTITLE}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="categorytitle">Category:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="categorytitle" value="{{.CATEGORYTITLE}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="subcategorytitle">Subcategory:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="subcategorytitle" value="{{.SUBCATEGORYTITLE}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="purpose">Purpose:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="purpose" value="{{.PURPOSE}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="otherinfo">Other Info:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="otherinfo" value="{{.OTHERINFO}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="keywords">Keywords:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="keywords" value="{{.KEYWORDS}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="placenames">Placenames:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="placenames" value="{{.PLACENAMES}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filename">Filename:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="filename" value="{{.FILENAME}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filetype">File Type:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="filetype" value="{{.FILETYPE}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filedescription">File Description:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="filedescription" value="{{.FILEDESCRIPTION}}" readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="step">Step:</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" name="step" value="{{.STEP}}" readonly>
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
