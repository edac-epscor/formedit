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
    <link rel="stylesheet" href="https://reporting.nmepscor.org/sites/all/modules/er/static/css/bootstrap.min.css">
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/jquery.min.js"></script>
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/bootstrap.min.js"></script>
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



const InsertViewTPL string=`<html>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://reporting.nmepscor.org/sites/all/modules/er/static/css/bootstrap.min.css">
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/jquery.min.js"></script>
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/bootstrap.min.js"></script>

<style>
pre {
   background-color: ghostwhite;
   border: 1px solid silver;
   padding: 10px 20px;
   margin: 20px; 
   }
.json-key {
   color: black;
   }
.json-value {
   color: blue;
   }
.json-string {
   color: green;
   }



.jumbotron {
    padding: 0.0em 0.0em;
    h1 {
        font-size: 2em;
    }
    p {
        font-size: 1.2em;
        .btn {
            padding: 0.0em;
        }
    }
}


</style>

</head>
<body>

<div class="container">
  <div class="jumbotron">
    <h1>JSON</h1>


<pre><code id=systemmetadata></code></pre>
<br>
</div>
</div>
<div class="container">
  <div class="jumbotron">
    <h1>XML</h1>


<pre><cod>{{.XMLSTR}}</code></pre>
<br>
</div>
</div>
</body>
<script>
if (!library)
   var library = {};

library.json = {
   replacer: function(match, pIndent, pKey, pVal, pEnd) {
      var key = '<span class=json-key>';
      var val = '<span class=json-value>';
      var str = '<span class=json-string>';
      var r = pIndent || '';
      if (pKey)
         r = r + key + pKey.replace(/[": ]/g, '') + '</span>: ';
      if (pVal)
         r = r + (pVal[0] == '"' ? str : val) + pVal + '</span>';
      return r + (pEnd || '');
      },
   prettyPrint: function(obj) {
      var jsonLine = /^( *)("[\w]+": )?("[^"]*"|[\w.+-]*)?([,[{])?$/mg;
      return JSON.stringify(obj, null, 3)
         .replace(/&/g, '&amp;').replace(/\\"/g, '&quot;')
         .replace(/</g, '&lt;').replace(/>/g, '&gt;')
         .replace(jsonLine, library.json.replacer);
      }
   };

var systemmetadata = {{.JSONSTR}};

$('#systemmetadata').html(library.json.prettyPrint(systemmetadata));
</script>
<div class="text-center">
{{.INSERTBUTTON}}
</div>
</html>`

const StarterTemplate string =`<html>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://reporting.nmepscor.org/sites/all/modules/er/static/css/bootstrap.min.css">
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/jquery.min.js"></script>
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/bootstrap.min.js"></script>

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
 <div class="alert alert-info" style="text-align: center">
  <strong>{{.NUMRECORDS}} records found.</strong>
 </div>

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
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://reporting.nmepscor.org/sites/all/modules/er/static/css/bootstrap.min.css">
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/jquery.min.js"></script>
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/bootstrap.min.js"></script>
<script>
$('form').submit(function(){
var valid = true;
  $('textarea[required]').each(function(i, el){
    if(valid && $(el).val()=='' ) valid = false;
  })

return valid;
})
</script>

<style>



.btn-insert {  margin: 10px 40px 32px 40px; background-color: hsl(0, 0%, 16%) !important; background-repeat: repeat-x; filter: progid:DXImageTransform.Microsoft.gradient(startColorstr="#5b5b5b", endColorstr="#282828"); background-image: -khtml-gradient(linear, left top, left bottom, from(#5b5b5b), to(#282828)); background-image: -moz-linear-gradient(top, #5b5b5b, #282828); background-image: -ms-linear-gradient(top, #5b5b5b, #282828); background-image: -webkit-gradient(linear, left top, left bottom, color-stop(0%, #5b5b5b), color-stop(100%, #282828)); background-image: -webkit-linear-gradient(top, #5b5b5b, #282828); background-image: -o-linear-gradient(top, #5b5b5b, #282828); background-image: linear-gradient(#5b5b5b, #282828); border-color: #282828 #282828 hsl(0, 0%, 11%); color: #fff !important; text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.33); -webkit-font-smoothing: antialiased; }



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

.datasettitle {
  background-color: #04488b;
  color: white;

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
      <div class="col-sm-6">
        <input type="hidden" class="form-control" name="id" value="{{.ID}}" readonly>
      </div>

      <div class="row">
         <div class="col-lg-6 col-lg-offset-3 ">
            <div class="form-group customborder datasettitle">Dataset ID: {{.ID}}</br> Submitted by: {{.SUBMITTERNAME}} </br> E-Mail: {{.SUBMITTEREMAIL}} </br><small class="text-muted">Dataset Created: {{.DATECREATED}}</small></div>
          </div>
        </div>

        <div class="form-group">
        <div class="col-sm-12 alert alert-info text-center ">Check the boxes if the content is valid.</div>
        </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="datasetname">Dataset Name :</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="datasetname" value="{{.DATASETNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=datasetnamebool value=true {{.DATASETNAMEBOOL}}> Dataset name is descriptive of the data.</label>
    </div>
    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstname">First Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="firstname" value="{{.FIRSTNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=firstnamebool value=true {{.FIRSTNAMEBOOL}}> First Name name appears to be valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastname">Last Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="lastname" value="{{.LASTNAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=lastnamebool value=true {{.LASTNAMEBOOL}}> Last name appears to be valid.</label>
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
      <label><input type="checkbox" name=phonebool value=true {{.PHONEBOOL}}> Phone number appears to be valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="firstnamepi">P.I. First Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="firstnamepi" value="{{.FIRSTNAMEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=firstnamepibool value=true {{.FIRSTNAMEPIBOOL}}> PI first name appears to be valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="lastnamepi">P.I. Last Name:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="lastnamepi" value="{{.LASTNAMEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=lastnamepibool value=true {{.LASTNAMEPIBOOL}}>PI last name appears to be valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="emailpi">P.I. E-Mail:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="emailpi" value="{{.EMAILPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=emailpibool value=true {{.EMAILPIBOOL}}>E-Mail appears to be valid.</label>
    </div>


    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="phonepi">P.I. Phone Number:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="phonepi" value="{{.PHONEPI}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=phonepibool value=true {{.PHONEPIBOOL}}> Phone number appears to be valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="abstract">Abstract:</label>
      <div class="col-sm-6">
        <textarea readonly rows="8" class="form-control" name="abstract" >{{.ABSTRACT}}</textarea>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=abstractbool value=true {{.ABSTRACTBOOL}}> Abstract is informitave.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="collectiontitle">Collection:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="collectiontitle" value="{{.COLLECTIONTITLE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=collectiontitlebool value=true {{.COLLECTIONTITLEBOOL}}> Collection Title appears to be valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="categorytitle">Category:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="categorytitle" value="{{.CATEGORYTITLE}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=categorytitlebool value=true {{.CATEGORYTITLEBOOL}}> Category Title appears to be valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="subcategorytitle">Subcategory:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="subcategorytitle" value="{{.SUBCATEGORYTITLE}}" readonly>
      </div>

    <div class="col-sm-4">
      <label><input type="checkbox" name="subcategorytitlebool" value=true {{.SUBCATEGORYTITLEBOOL}}> Subcategory appears to be valid.</label>
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
      <label><input type="checkbox" name=otherinfobool value=true {{.OTHERINFOBOOL}}>Other Info is descriptive (can be NULL)</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="keywords">Keywords:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="keywords" value="{{.KEYWORDS}}" readonly>
      </div>

    <div class="col-sm-4">
      <label><input type="checkbox" name=keywordsbool value=true {{.KEYWORDSBOOL}}> Keywords appears to be valid.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="placenames">Placenames:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="placenames" value="{{.PLACENAMES}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=placenamesbool value=true {{.PLACENAMESBOOL}}> Placenames are descriptive of places.</label>
    </div>

    </div>
    <div class="form-group">
      <label class="control-label col-sm-2" for="filename">Filename:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filename" value="{{.FILENAME}}" readonly>
      </div>
    <div class="col-sm-4">
      <label><input type="checkbox" name=filenamebool value=true {{.FILENAMEBOOL}}> Filename is present</label>
    </div>

    </div>

    <div class="form-group">
      <label class="control-label col-sm-2" for="filetype">File Type:</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="filetype" value="{{.FILETYPE}}" readonly>
      </div>
      <div class="col-sm-4">
      <label><input type="checkbox" name=filetypebool value=true {{.FILETYPEBOOL}}> Filetype is present</label>
     </div>
   </div>


    <div class="form-group">
      <label class="control-label col-sm-2" for="filename">Data:</label>
      <div class="control-label.text-left col-sm-6" > {{.DATA}}</div>
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
      <label><input type="checkbox" name=filedescriptionbool value=true {{.FILEDESCRIPTIONBOOL}}> Describes the data.</label>
    </div>
    </div>
{{.FIELDS}}

  </div>

    <div class="form-group customborder">
      <h1><span class="label label-default">Quality Control Note</span></h1>
	<br>
    <div class="form-group">
      <label class="control-label col-sm-2" for="Note"></label>
      <div class="col-sm-9">
        <textarea rows="8" class="form-control" name="note" required></textarea>
      </div>
    </div>


<div class="panel-footer row"><!-- panel-footer -->


    <div class="text-center">
        <div class="next">
          <button type="submit" name="button" value="reject" class="btn btn-danger">
              <span class="glyphicon glyphicon-chevron-left"></span>Reject
          </button>
          <button type="submit" name="button" value="note" class="btn btn-primary">
              Note Only  <span class="glyphicon glyphicon-file"></span>
          </button>
          <button type="submit" name="button" value="accept" class="btn btn-success">
              Accept<span class="glyphicon glyphicon-chevron-right"></span>
          </button>
        </div>
    </div>




</div>
</div>

 </div>
</form>
<div class="text-center">
{{.INSERTBUTTON}}
</div>
</div>
</body>
</html>`

const Bouncer string=` data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAa0AAAFNCAYAAACzEjC1AAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAFbwAABW8BpVExDwAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAACAASURBVHic7d15fJTlofbx3z0z2UhCwpqEZdhlk7C54IZUURT3Wutel7ZG67GLrdVpe9p62hp93x6rdiOnp6e+x7ba1lqXqkVtrcUF1FIFXFDWQWSHRALZZuZ+/5gEAiaQZWbueWau7+eTD5CZeZ4Lwsw19zP3cz/GWouIdE24JtQfqASmAiOBEqD0oF9zgY+AWqCu3a/bgOXAMmBVsKo6luL4Ip5nVFoiHQvXhHKATwAnEy+pqcCwBG1+L/sL7DXgyWBV9YcJ2rZIxlJpibQTrgmVAPOB84Azgb4p2rUFXgceAx4PVlUvT9F+RTxFpSVZL1wT8gMXANcBc4Acp4Hi1gK/BX4erKre6DqMSLpQaUnWav186vPAjcBwx3E6EwEeBu4LVlW/4jqMiGsqLck64ZrQeOBm4EqgwHGc7ngNuAd4SJM4JFuptCRrhGtC/YDvAl8AAm7T9MobwJeDVdUvuA4ikmoqLcl4rZ9ZXQ/cDgxwHCeRHgZuCVZVr3OcQyRlVFqS0cI1oVOBe4HJrrMkSRNwN/D9YFX1XtdhRJJNpSUZKVwTygPuAr7kOkuKvANcGqyqftN1EJFk8rkOIJJo4ZrQBGAx2VNYABOBJeGa0BddBxFJJo20JKOEa0KfI344sI/rLA49CVwTrKre5jqISKKptCQjhGtC+cCvgEtcZ0kTm4ELglXVi10HEUkklZZ4XutU9seBE11nSTMNwCXBqurHXQcRSRR9piWeFq4JjQBeQoXVkQLgkXBN6AbXQUQSRaUlnhWuCU0DXiE+CUE65gd+Fq4J3eE6iEgi6PCgeFK4JnQy8ARQ7DqLh9wPXBusqtaTXjxLIy3xnHBNaCYqrJ64Gvix6xAivaHSEk9pXez2aVRYPXVjuCb0H65DiPSUDg+KZ4RrQsOJT7pI18uIeMmXg1XV97oOIdJdKi3xhHBNaCCwCJjgOkuGsMBVwarqB1wHEekOlZakvXBNKAf4O3C84yiZJgKcGqyq/ofrICJdpc+0xAvuRIWVDAHgoXBNqMx1EJGu0khL0lq4JnQB8IjrHBnub8BpuhqyeIFGWpK2wjWh0cTXE5TkOoX4BTJF0p5GWpKWWq+H9Qow3XWWLGGBM4NV1QtdBxE5FI20JF3diQorlQzwQOssTZG0pdKStBOuCR0N6GKGqTcIuNt1CJFD0eFBSSvhmlAAeB2Y6jpLFjstWFX9nOsQIh3RSEvSzc2osFxbEK4JFbgOIdIRlZakjdbZgt91nUMYA3zHdQiRjqi0JJ38nPiFC8W9r4ZrQpWuQ4gcTKUlaSFcEzoNON11DtknAFS7DiFyMJWWpIvvuQ4gHzM/XBOa5TqESHsqLXEuXBM6CzjWdQ7pkN5MSFpRaUk60EUJ09fccE1otusQIm1UWuJUuCZ0PjDDdQ45JL2pkLSh0hLXtFBr+js5XBM6xXUIEVBpiUPhmtAcQNOqvUHLaklaUGmJS1WuA0iXnR2uCQ11HUJEpSVOhGtCg4BPus4hXeYHPus6hIhKS1y5Bsh1HUK65XPhmpDfdQjJbiotSblwTcgAn3edQ7ptODDfdQjJbiotceFUYKzrENIj+hxSnFJpiQuXuw4gPXZGuCY0wHUIyV4qLUmpcE3IB5ztOof0mB84y3UIyV4qLUm1E4CBrkNIr5znOoBkL5WWpJpe8LxvXrgmlOc6hGQnlZak2rmuA0ivFRKfTCOSciotSZlwTWgiMM51DkkIjZjFCZWWpNI5rgNIwuhnKU6otCSVTnQdQBKmIlwTGu06hGQflZakkq5OnFn085SUU2lJSoRrQqOAwa5zSELNch1Aso9KS1JFL3CZRz9TSTmVlqSKDiVlnmk6X0tSTaUlqaJ35ZknF5juOoRkF5WWJF24JhRAL26Z6mjXASS7qLQkFYahCz5mqjGuA0h2UWlJKoxyHUCSRj9bSSmVlqTCSNcBJGlGug4g2UWlJamgd+OZSz9bSSmVlqTCSNcBJGmKwzWh/q5DSPZQaUkq6N14ZtPPV1JGpSWpMMJ1AEkq/XwlZVRakgolrgNIUunnKykTcB1AskJhTx5kLby/aTsrP9zG5trdbN61m821u6lvaibg8+Fv/coJ+AgOLGVc+UDGlPdndNkA+uTlJPrv4AnWwsaddazavINVm3ewevMO6vY2Eo1ZotEYURvDZwyDS4ooLy2mol8xwweUMHXkkN78m/Xo5yvSEyotSarWten8Xb1/Q3MLf1ryFv94ey2vrtrAjt17u71PY2BQ3yLGDxnIEUMGMaZ8AGPLBzCuYiD9iwq6vb10FInG9pVS269vb9zK+q27aI5Eu729gN9H5YgKjh8f5PxjJjOuYmB3Hq7SkpRRaUmyFQLsaWymdm8Du+ob2LWnkfrGJorz8+hXVEC/wgL8PsPvXlrGr55/nZ31Db3aobWwta6erXX1LHpn3b7vGwNHjRnGeUdP5qyZEzxXYDFrWfJemD+9+jZPLX2X3Q1NCdt2JBpj6ZqNLF2zkZ/+5RXmThlH1enHMnJwP3bVN7Czfi91exvJCwToV1RA/6IC+hUVUJSfByotSSFjrXWdQTLUuvtvL31txftPbdr10XH3/PlF1m7d5TrSPn5jOGnyKD55zJGcNnVcWh9OXB7ezGOvvs0ji1ewo777I89kuuj4Sq6cPW39jNFDTx/6+Tvec51HMp9KS5Lj9tsD4fLG54ETATbt2s3Z1fez/aM9joN9XEFuDqdPG8f5R09m9qRRBPzu5yet27aLx159m0dffYs1W3a6jtOhm885iS+ddULbH7caEz1++HX/Z7XLTJL5dHhQkmJ9WeN5prWwACr6FXP1J2byw8f+4TJWhxqaW3js1bd57NW36VdYwKeOn8IX5h3n5PDh0/9ayYKFi3lj3aaU77s7ykuLufHM49p/azCxwH8AlzuKJFnC/VtKyUjG8IWDv3fZidPIDXR5ToYTu/Y08ItnX+XEb/2cH/35RfY0Nqdkv4veWcc51fdzfc2f0r6wAK44eToB34EvH9bYT61ecMtgR5EkS6i0JOHW/dc3JgCnHPz9AcV9OO/oSQ4Sdd+exmbu+fOLnPitn/O/LywlWUfRV23ewSU/epAr7n2IZes3J2cnCZYb8HPZidM6vCnX5/9cqvNIdlFpScL5rL2+s9sun+2ta0HurG/g3x98hsvueZCNO+sStt2YtdQ8s4Qzv/8/vLJyfcK2mwpnz5zIgOI+Hd9oTRW3367XFUkaTcSQhAvXhJYDR3Z0mwVWrN/MRw2NLF3zIT98PP0+4+pMYX4u3/30XD59fGWvtrNu2y6+8qs/s3TNxgQlS77xQwZx+8VzycsJMG7IIIrzO7+mp8/YymHX3bk8hfEki2gihiTUuvtvz/fBhM5uN8CUEeUAzBg9lJ8/szhlnxv11p7GZm7536dYu3Unt54/p0fbWLZ+M1f9+He9Phct1T4zZwbHje/aEoMxy3RApSVJoWG8JFSgcW8lXXwzVJCb45nPuNr72V8Wc+uvnyYa695RipfeXccld//Wc4VVmJ/b3Z+Tt44Bi6eotCShYphuvWDdev4chvTvm6w4SfPQi29y26+f7vL9X1v1AVf95A/safLGqLK9Oy6bR3FBXjceYWYkLYxkPR0elMTy+aZ3Z6pdaWE+P/nseVz4w1/jtc9Xf//yMp7+10rycg79NLLWsmtPA7FujszSwRnTx3P+MZO7+7Bp8UWzPPYDFU9QaUli2e4fGpo5ZihHBstY7pEp3+3tbmhK6BqA6ea2C07uycP6bvivW0YPvw6tjiEJp8ODkmC2oiePmjFqaKKDSC8F/D4q+vXs0G2MnGCC44gAKi1JvNKePOi0qeMSnUN66aSJo8g/zKHPzviwujCkJIVKSxInflJpUU8eetz4oOcuFZLpzprZ6ZkLh2V1NWNJEpWWJMy6EfQlfipWtwV8vp584C9JUlyQx7zejX5VWpIUmoghCWP3tJT05n/Ul846kUdffctz5zEBTB5e1uG08OXrN7GnqcVBot752rmz6dsnv8ePtzo8KEmi0pKEyfHZklgvHl9amM/Xzz+Z2379l4RlSoWSPvk8+Y1rMB2MMe9+YhH3PvlS6kP1wsShg7ny5N6dauWLaaQlyaHDg5Iw1hft0SSM9i4+YSonTBiZgDSpc85REzssLIB/O/N4fJ3dmIbycgLceeWZ+H29zGxsr/8viHREpSUJE7Oms5F7A4aPurINnzHUVF3A+CGDEpgsuT4395hOb8sN+Dn2iOEpTNNzPmO479pzmDayy2ctWGBXxzeY7iyhIdJlKi1JGGPp7NodBVi6fMJPcUEe9990EWUlPZqImFKjy/ozanC/Q97n7JkTU5Smd7578VzOmD6+Ow8xQId/eUvX3qSIdJdKSxLG13lpdduQfn158CuXMnLQoQvBpYDfxz3XnHPY+1120rQur5Dugt9n+M6n53LVnJkJ26Yhcf8XRNpTaUnCBApMQl+oxpQP4InQVZw8aVQiN5swN59zElO7cCjNZwz3XHM2pYU9n42XLKWF+TzwxYu59pSjEr1plZYkhUpLEmZnUV7CX6j69snn/ps+zQ3zZqXVhIaTJ43ihnmzunz/8tJi/u9nziLgT5+n3OThZTwRujo5E1+srU38RkVUWpJAky76TjOQ8JOsfMZw2wVzeDx0VXcmCSSFMXDDvFn8z79d1O0SPX3qOP7w1cspKylOUrquKczP5d8vOpXHQ1cRHJicSX46PCjJotKSREvai9WUYDl/uvUzVF9+hpNDbYV5ufzqxou47YI5BHw9e+rMGD2UZ779WU6Y4OYzrrNnTuD526/jc6ce3eO/Q1fE/CotSQ7jtWsYSXoL14RWAElfj2lPYzMPL17O/c//kzVbdiZ1Xz5jOG3qOG6/+DQq+iVmlGQtPPDCUv7ziUXU7knuCiD5OQHOP3YyV39iJhOHDk7qvtrEYrHjRt5w1+KU7EyyikpLEipcE3oIuDhV+7MWFr2zlv/9+1JeeHsNzZFowrad4/dz/rGTuf3TcynMz03Ydg/2xOvvcMcjz/PhzsTOEh9d1p+LT6jkkhOmpXpkaiPG9ht93Z0abUnCqbQkocILvvF1jL3Lxb73NrXw8sr1/P2t1Ty/Yg0f7Oj+a2ZRfi5zJo9mbuU4Tps6lqL81J0ju3TNRha+8T7PLXufVZt3dPvx+TkBZo0P8onJY/jEkWMYMcjZohRrglXVY1ztXDKbSksSKlxz62nge+Ywd6unh5cw6Y5Nu3azZssOVm/ZydotOwlvr6U5EiVmLbGYxQKDSwoJDixlxKB+jB7cn6kjK9Jiht/6bbW8u3Er67ftIry9lg3b62iKRPEZg8/ED1kOKili1OD+jCnvz6jB/Rld1p/cgN91dMA+HKy68yLXKSQzacFcSagcE/1Xiz3Mi75hEZYzk52lol8xFf2KPbeWIcCIQaUuR0qHZK3da4zp09ntxvCvVOaR7OL+LaVklIrrfrgd2HCo+5gYNSmKIwkWjUax1m471H1iKi1JIpWWJIFd0ulNhlXDr69+DOzGFAaSBIlEWojZ2KFmjNhINPrPlAWSrKPSkoTzYX7Z+a32lwDRaOyFVOWRxLDWEo1GaWlp2UF8hfePMZa/jLn+/25NcTTJIiotSbhhVXcuBLO6g5uac3N8vwRoaWl+JcWxpJcikfgVmG0s1oTlrx3dJ+azP01pKMk6Ki1JAmstLOjghj+UX3PHNoBYLFYfiURSnEt6o/3Py8LPOrjLuhGbCp5OXSLJRiotSQp/oPF/gPYnG0UMsXva30el5R3RaIT2p8eMGLD6cbBr29/HYH7Md74TS3k4ySoqLUmKYZ+9e6fPmLNaX9i2G/jM8Kq7Xm9/n1gsSiym1zgvaGk56A3GRb+P4uMcYCPxz7fuH15V/SMH0STLqLQkaYZdd8eSYNVdY4Kb8yuGV1U/2NF9WlqaUx1LuikajRKLfXx5rODn73wr2H/1iKjfVxGsqr4mvqiWSHLp5GJJMmv5Dp0eB4y/IMbwJXHFcemdQ76xuOj30VGwJXVpJNvplUKc02grfbW9qRBJFyotca6zw0/int5QSLpRaUlaaG5ucR1BDqJRlqQjlZakhVgsSjSq0VY60ShL0pFKS9JGc3MTulROemhpadYoS9KSSkvShrWWlhYdJkwH+jlIulJpSVqJRFr0Dl9EOqXSkrTT3NzkOkLWM8a4jiDSIZWWpJ1YLKbPtkSkQyotSUsqLRHpiEpLREQ8Q6UlIiKeodISERHPUGmJiIhnqLRERMQzVFoiIuIZKi0REfEMlZaIiHiGSktERDxDpSUiIp6h0hIREc9QaYmIiGeotERExDNUWpKWtMq7iHREpSVpSqXlkt40SLpSaUla0mumOyosSWcqLUlLeuF0Sf/2kr5UWpKWVFru6N9e0plKS9JSLBZ1HSFrRaMx1xFEOqXSkrQUi8XQYSo39IZB0plKS9KW3vG7EX/DIJKeVFqStvSOP/VisZg+05K0ptKStBWNqrRSTf/mku5UWpK2YrGYXkRTLBJpcR1B5JBUWpLWIpGI6whZIxqN6NCgpD2VlqQ1vZCmjt4giBeotCTttbQ0u46Q8aLRqA7FiieotCTtRSIRzSRMsubmJtcRRLpEpSWe0Nys0VaytLQ06xCseIZKSzwhFovpMGESxP9dNWNQvEOlJZ7R0tKiKdkJFIvFaGpqdB1DpFtUWuIpzc3NRKOa5dZb1lqamhp1WFA8R6UlntPU1KTp2b0Qi8VobFRhiTcFXAcQ6Ynm5iai0Qi5uXkYY1zH8YxIpEWTWsTTNNISz4pGozQ2NmjU1QWxWJSmpkYVlnieRlriadZampubaGlpxu8PEAgE8Pn0XqxNLBqhuaVFlxuRjKHSkoxgrSUS2T+70BiDMT58Ph+5ubmO07lhraWxuRn02ZVkEL0llZQxxgw2xpxrjLkx2fuy1hKLReNFlqXnITU3NyW1sIwxlcaY0UnbgUgHNNKSpDDGBICpwHHALOB4YFTrzRuAb6cqS0ukGX8gkFUTNiLRSCrWEvw6cLkxZhvwCrC49es1a219sncu2UmlJQlhjKlgf0EdBxwF5Le7S/u3/ANTGA1roampkbz8fAyZX1yxWIyWppRMuGj7OQ4Czm39AogaY1ZwYJG9ZzXHXhJApSXdZozJBaZz4ChqeLu7WPhYO7T/c8GuPQ05/QoLkpqzvVgsRlNjI3l5+Rk94oq2/j0PfI+QNAM6+b4fqCQ+0r6+9Xu7jDGL2V9kr1pr65IfUTKNSksOyxgznP0jqFnATKD97IaDXyEP2wprtuwqmjk6daUFrSfVNjWSl5eHz2Tex7nRaDTVyzJ1Vlrw8f8D/YAzW78AYsaYd4gXWFuRvWOt1TRHOSSVlhzAGJNH/NBeW0kdBwxpd5fDjaK6JLyjrmjm6CGHv2OC2ViMpsYGcnJy8QdyMuJgobXQ0uJklZDeHOb1AZOAycBnW7+3u3U01lZkS6y1O3sXUTKNSkswxgwD5gNnAXOBPu1u7vYoqis219YXJ2I7PWFtfA1DXyRCbk4ePr93R12RSISWlqaUz2o3xuQAvf0ZHvx/qRg4rfULwBpjlgJPAk8Rn+ChkViWU2llIWOMn/jnUPNbvyoPdfdkZHhv047UD7MOEj9c2EDAHyCQk+OZk5It8ZOGW1qcXhxzRAr2YYAZxA9HfxvYbox5mniBLbTW7kpBBkkzKq0sYYwZBJxBfDQ1DyhtvcnJjK61W3eNdLHfjkSiESLRCD6fn0DAH58en4YHDuMnUEeIRFrSYbHbmSnaT/sfxEDgytavmDHmZeIF9qS1dlmK8ohjKq0MZeJT5GYQL6n5wNHsP5m8/Suei1dnu7m2ftTh75ZasViU5uYopqUFvz+A3+/D7/ODw9mG1lqi0SjR1Jx31R2pKq3O+IATgBOBO4wxG2ktMOCvOk8sc6m0Mogxpi9wOvsP+5W13nTw5AnXwwjz0d7GiobmFgpycxxH+bj9S0LF/+zz+/D7Avj9fozPJHUUZrHYaIxIa1GlwYiqMzPpeFJOKrXf9xDg861fLcaYF2j9LMxa+56LcJIcKi2PM8ZMYv9o6kT2/0xdj6YOyYJZsWErR48Z6jrKYcWiMWLRZtpWgzLG4PP59q1taIyJl5kFfAZjTaf/4tZaLBasjf8+ZolZSywWw9qoJ5YJbC3SmaTX/6v2WXKITyiaC/zIGLMG+DPxkdjfrbVNDvJJgqi0PMgYMwO4CjiPzj8QT6cXlA49t3y1J0rrYG2H7OAQh+v29ZbZ95hMsejd9aVAiesc3TAK+GLrV4Mx5jniBfaItXar02TSbd6YLiUYYwYZY75sjHkT+CfxJ2DQcazesE/+632iscx5MT9AfDAVH01lUGEBPPrauxWuM3TTAauxAOcAPwc+NMY8aow5r3WtTPEAlVYaM8YEjDHnGGMeAT4EfgRMaX8XN8kSwuzYvZdXV33gOod0QzQWY9G76ytwNOs0wfzEj1Y8SrzA/tMYc6TjTHIYKq00ZIyZbIz5IbAReBy4gP2Hcr1cVB/z56UrXUeQbnhpZZj6xuYcMuz/IfHp9DcDy40xrxpjbjDGlB7uQZJ6Kq00YYwpbX2ivAqsAL5KfPXsjGWAPy99j60f7XEdRbroNy8udx0hWdqX8FHAz4DNxpgHjTGnGZOBi1V6lH4QDhljfMaYecaYh4DNxJ8oR7W/i5tkqWGB5kiUBc++5jqKdMHStZv4xzvrXMdIhbbnXR5wCfAMsN4Y8z1jzBh3sQRUWk4YY8YZY34AhIG/ABcTf4JAhhdVR/6w+C027drtOoYcxj1PveI6gktDgW8B7xtjXjDGXGWMKXQdKhuptFLEGFNsjLnWGLMIeA/4Bgeunp61ItEYP33mVdcx5BBefm8Dr63e6DqGS6bdr7OB+4EtxphfGmNOdJYqC6m0kswYM8MYcz/xw3+/JL70zL6bnYRKQ48seZsX3w27jiEdaGyJ8INHXnAdIx31Aa4FFhlj3jfGhIwx3jvx0GNUWknS+lnVc8TPqbqK/Zf7UFF1xMBtDz7LzvoG10nkIHc9tog1W7WgegfaP5fHAHcA64wx/88YM6WTx0gvqbQSqPW8qiuMMcuIf1Z1qutMXmEt7Ni9l2889JzrKNLOX1es4aGXV7iO4QVtBeYHPgMsM8Y8ZYyZ4y5SZlJpJUDr51U3A2uBBwCdoNhDL7y9jvueXuw6hgAbdtTxzYeewzhc5d6D2v9jnQk833re10WaNp8YJtOWmEklY0wF8eWUbiC+FpvrVa89zxD/R/z6uSdyzZzpruNkrS119Vz+4z+ycedHrqNkkjXAD4H7rbU6Dt5DKq0eMMaUAd8EqoBcVFYJZYzBWsv3Lz6VC4+d5DpO1tlZ38AVP/kja/U5VqK1vU5sJ74k273WWp1Z300qrW4wxvQHvk58dFWAyipp2g5J3XruiVx18jTHabLHlrp6qn7xBCs/3O46SiZre93YRnzyxs91uZSuU2l1gTGmGPgK8aWV+qKySglj4hM05k4Zwx2XzqU4P9d1pIz24rthbvn1Qmr3NrqOki3aXkc2ArcDv7LWRtxGSn8qrUMwxhQANwK3AQNQWTkzrH9f7rn6TCYPG+w6SsaJxiw/XbgkvpxW66FZSam215XVwLeBh6y1MbeR0pdKqwPGmBzil+3+FtB2GQaVlWM5fj+h80/i0hN0CkyibN+9l689sJAlqz7YNwlGnGl7nVkB/Lu19lHHedJS1peWMaYPcAQwod3XCcAwVFZppW2CxpnTxvG9i0+hME+HC3tjyaoP+OoDC9mxe6/rKHKgtted14DfAytbv9bo8GEWlpYxphw4vfXrJGA4Hy8mlVWaGzGolG9eMJuTJoxwHcVzGpsj/OJv/2TBs69hsWTZS4CXRYBVxAtsKbAQeC3bDiVmfGkZY/zAHGAe8aKa6jSQJNTM0UP4yvzjmDlaaw8fTks0yu9eXsGCZ19jR33Dvoku4ikHv6HeBTxLvMAWWmszflXjjC0tY8xg4HPET/wd1vptjaAy1EkTRvCl+bM0UaMD0Zjl0dfe4acLl7Cptn7fYVbJCAe/pr0M/BR42Frb7CZScmVcaRljjic+4+8iIAcVVdYwwGmVY/jimbMYU9bfdRznrIW/vPk+9z29mHXbajWyyi7bgV8ANdba9a7DJFLGlJYx5ljiZ5kf5zqLuGWM4ZyZ47n8xEoqg2Wu46RcNGb5+1tr+fHCJaz8cLtmBWantjfrMeBh4DZr7Vq3kRLD86VljAkCdxK/LLZGVHKA8RUD+dSsSZx71AT6FuQd/gEe9uGu3Ty85C3+uORtttZpdSA5QDPxN/U/sNZ6+jLhni2t1qnq3wC+xv5L1YscoG2UkRvwM2/qWD517GSOGZs51+mLxmL8bcVa/rD4LV5cGcZaq8OA0pG2kddW4uef/tKrsw49WVqtF1j7HTARfWYl3TRiYCmfmjWJC46eyIDiPod/QBoK76jj4cVv8cir7+g8K+mOttfLV4BLvfh5l+dKyxhzPXAPGl1JD7WNRPw+wzFjhzFv6ljmThnDgKIC19EOaUtdPc8sW83CN1exdM2H8Vcfjaqk5+qAa621j7gO0h2eKS1jTCnw38CFxD9c1AXVJGF8xnD0mKHMmzqW0yrHMDBNRmCbdu3eV1RvrNukCRWSSG2jrp8BX7XWemKlZE+UljFmBPET6Ma5ziKZzxjDzNFDmFc5htOnjmVw38KU7n/jzo9Y+OYqFi5bzbL1m1O6b8laK4BzrLXrXAc5nLQvLWPMkcAzxBeuFUkpA0wfVcG8qWM5vXIs5aVFSdlPeEcdz7y5ioVvrmLFhq1J2YfIYWwCTrPWvuU6yKGkdWm1nij8JFDqOouIASpHlDNv6ljmTR3LkH7Fvdreum218RHVm6t4Z+O2fftI32ekZIFdwHxr7WLXQTqTtqVljJkP/JH4hAvNDpS0M37IQGZPGMnsSUEqSrtWYOu36GoegQAAEvtJREFU17HonfX84511rGm9nL2KStKIBRqAT1prF7oO05G0LC1jzFHAIlRYIiKpZoEWYLa1donrMAdLu9IyxlQArxP/DEuFJSKSehbYBsy01n7gOkx7aTVt3BiTDzwKDEGFJSLiigEGA0+0rj6UNtKqtIivSnyM6xAiIgLANOB/jTFpM4hIm9IyxtwKXOE6h4iIHOBC4LuuQ7RJi8+0jDFnA48RH5KmTaOLiMi+lTMusdb+znUY56VljJkELAEKUWGJiKQjCzQBJ1lrX3cZxGlpGWP6A68Bo52FEBGRrtoEHGWt/dBVANefaf0WFZaIiFdUAI8aYwKuAjgrLWPMpcA8V/sXEZEeORr4oqudOzk8aIwpAVYSPw9An2OJiHiHBfYC4621G1O9c1cjre8DZaiwRES8xhCfOPcjJztP9UjLGDOD+OQLTW8XEfG2M1K9sG5KS8sY4wMWEz8mKiIi3mWBNcBka21Tqnaa6sODVaiwREQygQHGAKGU7jRVIy1jzGDgPaAvOiwoIpIJ2i5jMtlauyoVO0zlSOuHQAkqLBGRTGGAXOAnKdthKkZaxpg5wPNJ35GIiLjyaWvtH5K9k6SXVuuS9m8AlUndkYiIuGKBzcA4a+2eZO4oFYcHz0eFJSKSyQzxJZ5uSPqOUjDSWgpMT+pORETENQtsA0ZaaxuStZOkjrRar5OlwhIRyXyG+NJ8n0/qTpI50jLGvIrOyxIRyRaW+OVLRifrhOOkjbSMMWegwhIRySYGGAJck7QdJGukZYx5CTg+KRsXEZF0ZYENwFhrbUuiN56UkZYx5lRUWCIi2cgAQeDKpGw8GSMtY8wLwOyEb1hERLzAAmuBI6y10URuOOEjLWPMyaiwRESymQFGA5cmfMOJHmkZY/4KnJLQjYqIiNdY4oukT7LWxhK10YSOtIwxR6HCEhGR+GhrPPDJRG400YcH/y3B2xMREW+7KZEbS9jhQWPMQGAjkIMuPyIiIvtVWmuXJ2JDiRxpfY74dVVUWCIi0l7CjsIlZKRljPETn944DJWWiIjsZ4EGYKi1tra3G0vUSOscYDgqLBEROZAB+pCgpZ0SNdJ6Dji193FERCQDWWAN8YtE9qp0ej3SMsZMRIUlIiKdM8AYYF5vN5SIw4M3JmAbIiKS+Xo9IaNXhweNMcXAh0Ah+jxLREQOLUb8EOGanm6gtyOtq4AiVFgiInJ4PuALvdlAb0da7xBfpkOlJSIih2OBOuLT3/f2ZAM9HmkZY+YCE1BhiYhI1xigFLispxvozeFBrTMoIiLdZelFf/To8KAxZgSwGvD3dMciIpLVTrLWvtjdB/V0pHUDKiwREem5Ho22uj3SMsbkE1/NvR/6PEtERHomAgSttZu686CejLQuAfqjwhIRkZ4LAFXdfVBPRlqvAzO7uyMREZF2LLAVGG6tbenqg7o10jLGzEKFJSIivWeAMuDC7jyou4cHNc1dREQSqVu90uXDg8aYwcAHQE4PQomIiHRmurX2ja7csTsjrS+jwhIRkcT7Slfv2KWRljFmDPA28dLSrEEREUmkGHCstfb1w92xqyOtHwG5qLBERCTxfMBPjDGH7ZjDlpYxZj5wTiJSJYNaVEQyVSJe3w5fA2njWODqw93pkIcHjTG5PmPejlk7JoHBusUY6CiiMYaykiKCA0spLcynMD+XovxcCvPyKMrPpU9eTvzP+XkU5eWSl+MnErNEolEisRiRaLuvWIxINEp9YzNb6urZWlvPlrp6NtfuZkttPQ3NXT6FQESkywzxk5U6UlpYQHlpEWUlRZQWFuD3+Qj4feT4ffj9PnL8fvy+/X+ORmPUNzazp6mZPa2/1je2/30Te5taaI5EU/lX7LLWbt1uYay1tq6z+wUOtZEcv/+Wlmg0qYVljKGz4uxXWEBwUCnDB5QwfGBp/Kv198P69yUnkJrlD+sbm9hSW8+6bbtYtn4zy9dvZtn6TWz7aE9K9i8i3nfwG/C8nACTh5dxZLCMseUDGFwSL6iy0iIG9y1K2uvbzvoGNmyvZcOOOj7YUcuG7XVs2F5LeHsdH+yso8VRqbX+0wz0GXM78Yl/Hep0pGWMGZrj962NxGxOby4UCZ2/m+iTl0NwYCnDB5QyfGBbMZXs+3NhXm6v9ptsW+vqWR6Ol9jSNRt5eeV6WqIx17FEJM3k5QSoHFHOkcFypgTLW4tqIH5feh27sxa21O2OF1m7Qtuwo47wtlo21e7udJCRCAbAmJi1ttJa+1aH9+ksQEW/vs9urt09NxFB+hUVMHVEBZUjKxg/ZNC+0VL/ooJEbD5tfLS3kWfefJ8nl77LorfXqsBEslheToBTjhzD2UdN5NQpYyjI9f4ZQ5FojI07P2otsniprd6ykzfXfcimXbsTtp+SPvmv1+5pOLqj2zosreEDS8/4YEfd013Z+MGjqKL8XKaMqGDqiAqmjiynckQFwwaU9Ci4l+1uaOLZZe/z5D/f5e9vrSESjR3y+LWIeFfbxxy5AT+faFdU6X60KJG2f7SHN9dv4s11m3hz/SaWrdvEzvqGHm+vpE/+FbV7Gn5z8Pc/VlrGGH95afHmzbW7B37szgcdk83PCTA5WNZaUBVUjqhg1OD+XpqtkhIbttdy31Mv88fFK4jGVF4imaKtrPoXFXDDvFlcftJ0CvOzp6gO54MddftK7M11m1i+fjN7mpoP+zhjoDg/b/dHDU3l1tq9B9x2cGmV9yv+xpba+h8cvJEcv48JQwczdWS8oKaMKOeIikFpd0w2nYW31/JjlZeI57WVVWlhPtefPour5sykT573D/8lm7WwZsuO/SOydZt4a8OWA2Y0tn9dLOmTf3ftnoavtt/GAaVljBmUnxMIN0ej+ePKB+4bPVWOKGfSsMEpm62X6dZvq+W+p17ij68sh0PMnhSR9FTSJ5/rTjuWa06ZmVWHAJMhEo2x8sNtB4zI3vtwO9FYDL/PF43GYuOttavb7n9AaU0fPfS40AVznq0cUVGodw3J97flq/nK/U9Qu6fRdRQROYy2j0fmz5jAXVecQd8++a4jZazGlghvhbfwr3UfNixYuPimrXX1v2y77YDSCteEvgl830XIbLW5djc3/ffjvLpqwyHPWRMRt3IDfr590VyuPHm66yjZ5pvBquo72v5w8DJO3b70sfROeWkxD918GTfNPx7w1JIrIlljVFl/Hr/tKhWWG9e3/8O+kVa4JnQU8JqLRBL30rvr+NzPH6GhuUUjLpE0cd4xk6i+/Ax9duXWzGBV9VI4cKT1SUdhpNUJE0by3zd8koDfRxcWOxaRJDvvmEnce825Kiz3zm/7TfvSusBBEDnICRNGct+18UX1VVwi7pw8aRR3X3W2DtmnhwNLK1wTmghMcBZHDjB/xgR+cNk8rLV6wog4MG3UEBZUxY96SFqYEq4JjYb9I62zHYaRDlx+0jS+du7sDi/LIiLJM7Z8APf/20U6WTj9nAX7S0tTYtLQTfOP59QpY13HEMkKxhhyA37+58aL6FeYWYt5Z4iZsL+0Kh0GkUO4/eK55OUEdJhQJMmstdx05vGMGFTqOop0rBLAf21FUy5wNx8/Z0vSQEnrWfcvr1zvOIlIZhs1uD/3fvZc/D69FKap/nX//OudPuITMA55BWNx6/rTj2XU4H5osCWSPN+/9HRytb5qOssDJvjQocG0lxPw871L52lFeJEkOeeoiZw4caTrGHJ4lT5gousUcngnTRzJ7EmjXMcQyTh5OQG+fdGprmNI10zyAf1dp5CuueSEqa4jiGSc+TPGM7ikyHUM6ZoBPqDEdQrpmtOnjqOkT75mEook0EXHTXEdQbquRKXlITkBPxccO1knHIskgDEwtH9fjh8/0nUU6TqVltdcfLzmzYgkgrVw4XFTdOTCW1RaXjNpeBmTh5dp+rtILxl0aNCDSnxAX9cppHsuOGaypr+L9NIx44YTHKjVLzymr4/4CVviIdNHD3EdQcTzTpgw0nUE6b48rVfiQZOGleHz6QChSG8cGSxzHUF6QKXlQX3ychhbPkAfIIv0wpHBctcRpAdUWh41JViuqe8iPTSwuA9lOqHYk1RaHjVF7xJFemzKiArXEaSHVFoeNWWESkukp/R5lneptDxq0jA96UR6SkcqvEul5VF98nLok5fjOoaIJw0boDUVvEql5WFF+TrFTqQn9NzxLpWWh/UtyNNyTiI9UJSf6zqC9JBKy8OKC/RuUaQn9NzxLpWWhxUX5GkNQpFuCvh95Ab8rmNID6m0PEzH5UW6T88bb1NpeZgOcYh0n5433qbS8rBifZgs0m1FeXreeJlKy8OOGTecUWX9XccQ8ZTR5XrOeFnAdQDpuXnTjmDetCNYs2Unf12+iueWreK1VR8QjcVcRxNxyhj2LSgd8PuYdUSQuVPGcmrlWF340ePM+gW3rQeCroNIYny0t5G/v72Whf9aycI336clEj3gCSySDUr65DO3ciynThnLyZNHafJF5gib9QtuWwZMcZ1EEm9nfQN/fGU5v33xDdZs2ek6jkhStH9TNnPMUK6YPZ2zZkwgL0cHkjLQcrN+wW2LgBNdJ5HkWvxemN8ueoOn/rVSoy/JCMYYrLUU5udy4bFHcvns6UwYOsh1LEmuFwNAnesUknyzjggy64gg3/5oDz9buJgHXlhKcySKAZ2gLJ7SVlblpcXcNP94LjhmshaPzh51Zv2C234DXOY6iaTWlrp6fvLUyzz44hu0RGMqL0l7bWU1uKSIm+Yfz6UnTCVHK1tkm99qpJWlykqK+N6lp3PDvFnc+9RL/O6lZTpeKGmp7Q1Vv8ICbjzzOK6cPV2fV2WvugCww3UKcWdI/77cdcWZXH7SdEK/eZoV4S373tGKpAVjuOYTM7nl3NkU6oT6bLcjALztOoW4VzminMdvu5r7n3+dHz6+iL1Nza4jSRZre+M0JVhO9RVn6ErD0maFWb/gtsnACtdJJH1s2rWbbz24kOeWrdKoS5wozMvllvNm85k5M/H7dNU42WeSWb/gtgCwB9C4Ww6w4JnF3PWnF7BYfdwlKTNx2GBqqj7JiEFauUIO0AgUGWst4ZrQm0Cl60SSfl5ZuZ4v/OIxdtbvdR1FMljbiP7CWUdyx+VnkK+JFvJxrwerqo9uWzB3mdMokraOGz+Cp791DTNGDwXiM7lEEi3gM/zgsnncffXZKizpzDLYv8r7UodBJM2Vlxbzh69ezmfmzMASf1cskigV/Yp5+JYruWL2dNdRJL29DvtL6wmHQcQDAn4f37vkdO699hzydEKn9FLb254TJozkqW9ey7SRFU7ziCc8CWDaZoaFa0JaOFe65N2N27huwSOs37bLdRTxINO68OUXzjiOr547W7MDpSuWBquqZ8KBF4H8k6Mw4jEThg7iyW9czSlTxriOIh5jDOTnBPivGy7k6+efrMKSrnq07TftS+sRB0HEo4oL8vjFDRdyzlETXUcRjzBAUX4eD37lUk6fOs51HPGWfaVl2p84Gq4JrQZGu0gk3hSzllsfeJrfv6wJqHJo/YoK+O2XLmHS8DLXUcRb1gSrqvcd1vEddOPPUhxGPM5nDP/nyvlcNWcmoCnx0rHBJUU8/LUrVFjSEz9t/4eDS6sG0CVupVuMgf+45DRumDerdUp8ivabmt1ILw3p35eHv3YFY8sHuI4i3rOdeC/tYw5eVy5cE/oO8N3UZZJMct+TL/GfTyzq1ZqFXbmqcnFBHuWlxZSVFtHYHGFn/V6iMUtzJEJDcwu7G5qJxmI92r/X5Ab89CssoLSwgD55OeTnBjhq9DBKiwrYWlfPltp6ttTVs7Wuns21u9nTmLrFkEcO6seDN1/KkH59U7ZPySjfClZV/6D9Nzo69fzHwNeAopREkozyxbNOoCAvh+8//LePFZdpvTDSofqoKD+P8tIiBpcUUVZaRFlJvJgGlxRR1vq9wSVFXVo1oTkSzfjFfo0x5HbzvLmG5ha2tJbZ1rr9hZbocjtiyEB+++VLGdS3sFfbkaxVB/zk4G9+bKQFEK4J3QGEUhBKMtRvFr3BN3/zl30FVZiXGy+h0mLKStqXUtEBpVSQq8ump4u9TS1s/ejAcju46DortynBch740sX0KyxwkFwyxHeCVdX/cfA3OyutfOJLZkxOQTDJUCs/3EZuIEBZSRF98lRGmergcttaV8+nj6+kuCDPdTTxrjeAY4NV1R97R9RhaQGEa0LTgCXokiUiIpI6jcDMYFV1hxcoPnj24D7Bquo3gG8lK5WIiEgHbu2ssOAQpdXqP9FiuiIikhoPE58M2KlDllawqjoGXAg8lMBQIiIiB/t/wCXBqupDTvk93EiLYFV1C3A5sCBBwURERNq7B7gmWFUdPdwdO52I0ZFwTehm4NtASc+ziYiIALAL+Hawqvpj52N1plulBRCuCQ0Avgl8AdCcVhER6a5G4D6gOlhVXdudB3a7tNqEa0JB4HRgEvHzucYD+T3amIiIZLIGYCXwVuvXM8Gq6g96sqH/D2Thp/sjVd/AAAAAAElFTkSuQmCC`
const Logo string = ` data:image/png;base64,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`
const LogoBlack string =`data:image/png;base64,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`


const GSTORETABLEXML string = `<?xml version='1.0' encoding='UTF-8'?>
<metadata>
	<idinfo>
		<citation>
			<citeinfo>
				<origin>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</origin>
				<pubdate>{{.RELEASEDATE}}</pubdate>
				<title>{{.TITLE}}</title>
				<geoform>{{.GEOFORM}}</geoform>
				<onlink>http://nmepscor.org/dataportal</onlink>
			</citeinfo>
		</citation>
		<descript>
			<abstract>{{.ABSTRACT}}</abstract>
			<purpose>{{.PURPOSE}}</purpose>
			<supplinf>These data are made available through New Mexico's Experimental Program to
                                         Stimulate Competitive Research (NM EPSCoR) data portal. The portal provides access
                                         to data generated by and of use to the EPSCoR community including researchers,
                                         educators, and policy leaders, as well as the general public.</supplinf>
		</descript>
		<timeperd>
			<timeinfo>
				<sngdate>
					<caldate>Unknown</caldate>
				</sngdate>
			</timeinfo>
			<current>Publication date</current>
		</timeperd>
		<status>
			<progress>Complete</progress>
			<update>Unknown</update>
		</status>
		<keywords>
			<theme>
				<themekt>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</themekt>
				<themekey>{{.CATEGORYTITLE}}</themekey>
			</theme>
			<theme>
				<themekt>ISO 19115 Topic Category</themekt>
				<themekey>utilitiesCommunications</themekey>
			</theme>
			<theme>
				<themekt>Project</themekt>
				<themekey>EPSCoR T1R4</themekey>
				<themekey>{{.SUBCATEGORYTITLE}}</themekey>
				<themekey>{{.CATEGORYTITLE}}</themekey>
			</theme>
			<place>
				<placekt>GNIS</placekt>
			</place>
		</keywords>
		<accconst>None</accconst>
		<useconst>None</useconst>
		<ptcontac>
			<cntinfo>
				<cntorgp>
					<cntorg>{{.GROUPNAME}}</cntorg>
					<cntper>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</cntper>
				</cntorgp>
				<cntaddr>
					<addrtype>mailing address</addrtype>
					{{.ADDRESS}}
					<city>{{.CITY}}</city>
					<state>{{.STATE}}</state>
					<postal>{{.ZIP}}</postal>
					<country>USA</country>
				</cntaddr>
				<cntvoice>{{.PHONEPI}}</cntvoice>
				<cntemail>{{.EMAILPI}}</cntemail>
			</cntinfo>
		</ptcontac>
		<datacred>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</datacred>
		<native>Unknown</native>
	</idinfo>
	<dataqual>
		<logic>Not provided</logic>
		<complete>Not provided</complete>
		<lineage>
			<procstep>
				<procdesc>Unknown</procdesc>
				<procdate>Unknown</procdate>
			</procstep>
		</lineage>
	</dataqual>
	<spdoinfo>
		<indspref>United States</indspref>
	</spdoinfo>
	<eainfo>
		<detailed>
			<enttyp>
				<enttypl>{{.COLLECTIONTITLE}}</enttypl>
				<enttypd>{{.TITLE}}</enttypd>
				<enttypds>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</enttypds>
			</enttyp>{{.ATTRIBUTES}}</detailed>
	</eainfo>
	<distinfo>
		<distrib>
			<cntinfo>
				<cntorgp>
					<cntorg>Earth Data Analysis Center</cntorg>
				</cntorgp>
				<cntpos>Clearinghouse Manager</cntpos>
				<cntaddr>
					<addrtype>mailing and physical address</addrtype>
					<address>MSC01 1110</address>
					<address>1 University of New Mexico</address>
					<city>Albuquerque</city>
					<state>NM</state>
					<postal>87131-0001</postal>
					<country>USA</country>
				</cntaddr>
				<cntvoice>505-277-3622 ext. 230</cntvoice>
				<cntfax>505-277-3614</cntfax>
				<cntemail>clearinghouse@edac.unm.edu</cntemail>
				<hours>0800 - 1700 MT, M-F -7 hours GMT</hours>
			</cntinfo>
		</distrib>
		<resdesc>Downloadable Data</resdesc>
		<distliab>The material on this site is made available as a public service. Maps and data are to be used for reference purposes only and the Earth Data Analysis Center (EDAC), New Mexico EPSCoR (NM EPSCoR) and The University of New Mexico are not responsible for any inaccuracies herein contained. No responsibility is assumed for damages or other liabilities due to the accuracy, availability, use or misuse of the information herein provided. Unless otherwise indicated in the documentation (metadata) for individual data sets, information on this site is public domain and may be copied without permission; citation of the source is appreciated.</distliab>
		<stdorder>
			<digform>
				<digtinfo>
					<formname>Spreadsheet</formname>
					<transize>1</transize>
				</digtinfo>
				<digtopt>
					<onlinopt>
						<computer>
							<networka>
								<networkr>http://nmepscor.org</networkr>
							</networka>
						</computer>
						<accinstr>Download from New Mexico EPSCoR (NM EPSCoR) at http://nmepscor.org.</accinstr>
					</onlinopt>
				</digtopt>
			</digform>
			<fees>None. The files are available to download from New Mexico EPSCoR (NM EPSCoR) (http://nmepscor.org).</fees>
			<ordering>Contact Earth Data Analysis Center at clearinghouse@edac.unm.edu</ordering>
		</stdorder>
		<custom>Contact Earth Data Analysis Center at clearinghouse@edac.unm.edu</custom>
		<techpreq>Adequate computer capability is the only technical prerequisite for viewing data in digital form.</techpreq>
	</distinfo>
	<metainfo>
		<metd>{{.RELEASEDATE}}</metd>
		<metc>
			<cntinfo>
				<cntorgp>
					<cntorg>Earth Data Analysis Center</cntorg>
				</cntorgp>
				<cntpos>Clearinghouse Manager</cntpos>
				<cntaddr>
					<addrtype>mailing and physical address</addrtype>
					<address>MSC01 1110</address>
					<address>1 University of New Mexico</address>
					<city>Albuquerque</city>
					<state>NM</state>
					<postal>87131-0001</postal>
					<country>USA</country>
				</cntaddr>
				<cntvoice>505-277-3622 ext. 230</cntvoice>
				<cntfax>505-277-3614</cntfax>
				<cntemail>clearinghouse@edac.unm.edu</cntemail>
				<hours>0800 - 1700 MT, M-F -7 hours GMT</hours>
			</cntinfo>
		</metc>
		<metstdn>FGDC Content Standards for Digital Geospatial Metadata</metstdn>
		<metstdv>FGDC-STD-001-1998</metstdv>
		<mettc>local time</mettc>
	</metainfo>
</metadata>`


const GSTOREXML string = `<?xml version='1.0' encoding='UTF-8'?>
<metadata>
	<idinfo>
		<citation>
			<citeinfo>
				<origin>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</origin>
				<pubdate>{{.RELEASEDATE}}</pubdate>
				<title>{{.TITLE}}</title>
				<geoform>{{.GEOFORM}}</geoform>
				<onlink>http://nmepscor.org/dataportal</onlink>
			</citeinfo>
		</citation>
		<descript>
			<abstract>{{.ABSTRACT}}</abstract>
			<purpose>{{.PURPOSE}}</purpose>
			<supplinf>These data are made available through New Mexico's Experimental Program to
                                         Stimulate Competitive Research (NM EPSCoR) data portal. The portal provides access
                                         to data generated by and of use to the EPSCoR community including researchers,
                                         educators, and policy leaders, as well as the general public.</supplinf>
		</descript>
		<timeperd>
			<timeinfo>
				<sngdate>
					<caldate>Unknown</caldate>
				</sngdate>
			</timeinfo>
			<current>Publication date</current>
		</timeperd>
		<status>
			<progress>Complete</progress>
			<update>Unknown</update>
		</status>
		<spdom>
			<bounding>
				<westbc>{{.WESTBC}}</westbc>
				<eastbc>{{.EASTBC}}</eastbc>
				<northbc>{{.NORTHBC}}</northbc>
				<southbc>{{.SOUTHBC}}</southbc>
			</bounding>
		</spdom>
		<keywords>
			<theme>
				<themekt>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</themekt>
				<themekey>{{.CATEGORYTITLE}}</themekey>
			</theme>
			<theme>
				<themekt>ISO 19115 Topic Category</themekt>
				<themekey>utilitiesCommunications</themekey>
			</theme>
			<theme>
				<themekt>Project</themekt>
				<themekey>EPSCoR T1R4</themekey>
				<themekey>{{.SUBCATEGORYTITLE}}</themekey>
				<themekey>{{.CATEGORYTITLE}}</themekey>
			</theme>
			<place>
				<placekt>GNIS</placekt>
			</place>
		</keywords>
		<accconst>None</accconst>
		<useconst>None</useconst>
		<ptcontac>
			<cntinfo>
				<cntorgp>
					<cntorg>{{.GROUPNAME}}</cntorg>
					<cntper>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</cntper>
				</cntorgp>
				<cntaddr>
					<addrtype>mailing address</addrtype>
					{{.ADDRESS}}
					<city>{{.CITY}}</city>
					<state>{{.STATE}}</state>
					<postal>{{.ZIP}}</postal>
					<country>USA</country>
				</cntaddr>
				<cntvoice>{{.PHONEPI}}</cntvoice>
				<cntemail>{{.EMAILPI}}</cntemail>
			</cntinfo>
		</ptcontac>
		<datacred>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</datacred>
		<native>Unknown</native>
	</idinfo>
	<dataqual>
		<logic>Not provided</logic>
		<complete>Not provided</complete>
		<lineage>
			<procstep>
				<procdesc>Unknown</procdesc>
				<procdate>Unknown</procdate>
			</procstep>
		</lineage>
	</dataqual>
	<spdoinfo>
		<indspref>United States</indspref>
	</spdoinfo>
	<eainfo>
		<detailed>
			<enttyp>
				<enttypl>{{.COLLECTIONTITLE}}</enttypl>
				<enttypd>{{.TITLE}}</enttypd>
				<enttypds>{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}</enttypds>
			</enttyp>{{.ATTRIBUTES}}</detailed>
	</eainfo>
	<distinfo>
		<distrib>
			<cntinfo>
				<cntorgp>
					<cntorg>Earth Data Analysis Center</cntorg>
				</cntorgp>
				<cntpos>Clearinghouse Manager</cntpos>
				<cntaddr>
					<addrtype>mailing and physical address</addrtype>
					<address>MSC01 1110</address>
					<address>1 University of New Mexico</address>
					<city>Albuquerque</city>
					<state>NM</state>
					<postal>87131-0001</postal>
					<country>USA</country>
				</cntaddr>
				<cntvoice>505-277-3622 ext. 230</cntvoice>
				<cntfax>505-277-3614</cntfax>
				<cntemail>clearinghouse@edac.unm.edu</cntemail>
				<hours>0800 - 1700 MT, M-F -7 hours GMT</hours>
			</cntinfo>
		</distrib>
		<resdesc>Downloadable Data</resdesc>
		<distliab>The material on this site is made available as a public service. Maps and data are to be used for reference purposes only and the Earth Data Analysis Center (EDAC), New Mexico EPSCoR (NM EPSCoR) and The University of New Mexico are not responsible for any inaccuracies herein contained. No responsibility is assumed for damages or other liabilities due to the accuracy, availability, use or misuse of the information herein provided. Unless otherwise indicated in the documentation (metadata) for individual data sets, information on this site is public domain and may be copied without permission; citation of the source is appreciated.</distliab>
		<stdorder>
			<digform>
				<digtinfo>
					<formname>Spreadsheet</formname>
					<transize>1</transize>
				</digtinfo>
				<digtopt>
					<onlinopt>
						<computer>
							<networka>
								<networkr>http://nmepscor.org</networkr>
							</networka>
						</computer>
						<accinstr>Download from New Mexico EPSCoR (NM EPSCoR) at http://nmepscor.org.</accinstr>
					</onlinopt>
				</digtopt>
			</digform>
			<fees>None. The files are available to download from New Mexico EPSCoR (NM EPSCoR) (http://nmepscor.org).</fees>
			<ordering>Contact Earth Data Analysis Center at clearinghouse@edac.unm.edu</ordering>
		</stdorder>
		<custom>Contact Earth Data Analysis Center at clearinghouse@edac.unm.edu</custom>
		<techpreq>Adequate computer capability is the only technical prerequisite for viewing data in digital form.</techpreq>
	</distinfo>
	<metainfo>
		<metd>{{.RELEASEDATE}}</metd>
		<metc>
			<cntinfo>
				<cntorgp>
					<cntorg>Earth Data Analysis Center</cntorg>
				</cntorgp>
				<cntpos>Clearinghouse Manager</cntpos>
				<cntaddr>
					<addrtype>mailing and physical address</addrtype>
					<address>MSC01 1110</address>
					<address>1 University of New Mexico</address>
					<city>Albuquerque</city>
					<state>NM</state>
					<postal>87131-0001</postal>
					<country>USA</country>
				</cntaddr>
				<cntvoice>505-277-3622 ext. 230</cntvoice>
				<cntfax>505-277-3614</cntfax>
				<cntemail>clearinghouse@edac.unm.edu</cntemail>
				<hours>0800 - 1700 MT, M-F -7 hours GMT</hours>
			</cntinfo>
		</metc>
		<metstdn>FGDC Content Standards for Digital Geospatial Metadata</metstdn>
		<metstdv>FGDC-STD-001-1998</metstdv>
		<mettc>local time</mettc>
	</metainfo>
</metadata>`


const JSONlol string = `{"description": "{{.DATASETNAME}}","orig_epsg": {{.ORIGEPSG}},"taxonomy": "file","basename": "{{.BASENAME}}","is_embargoed": "{{.ISEMBARGOED}}","spatial":{"epsg": "{{.EPSG}}","geomtype": "{{.GEOMTYPE}}","features": "{{.FEATURES}}","records": {{.RECORDS}},"bbox": "{{.WESTBC}},{{.SOUTHBC}},{{.EASTBC}},{{.NORTHBC}}"},"sources":[{"mimetype": "application/x-zip-compressed","files":["/geodata/energize/uploads/{{.UPLOADUSER}}/{{.FILENAME}}"],"set": "original","external": "False","extension": "{{.EXTENSION}}"}],"active": "True","dataone_archive":"{{.DATAONEARCHIVE}}","releasedate":"{{.RELEASEDATE}}","author":"{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}","categories": [{"theme": "{{.CATEGORYTITLE}}","subtheme": "{{.SUBCATEGORYTITLE}}","groupname": "{{.GROUPNAME}}"}],"apps":["energize"],"standards":["FGDC-STD-001-1998","ISO-19115:2003"],"project": "NM EPSCoR R4","services": [],"formats": ["{{.EXTENSION}}"],"metadata": {"xml": "{{.RAWXML}}","upgrade": "true","standard": "FGDC-STD-001-1998"}}`


const TABLEJSON string = `{
        "description": "{{.DATASETNAME}}",
        "taxonomy": "table",
        "basename": "{{.BASENAME}}",
        "is_embargoed": "{{.ISEMBARGOED}}",
	"sources": [
            {
             	"mimetype": "{{.MIMETYPE}}",
                "files": [
                    "/geodata/energize/uploads/{{.UPLOADUSER}}/{{.FILENAME}}"
                ],
                "set": "original",
                "external": "False",
                "extension": "{{.EXTENSION}}"
            }
	],
	"active": "True",
        "dataone_archive":"{{.DATAONEARCHIVE}}",
        "releasedate":"{{.RELEASEDATE}}",
        "author":"{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}",
        "categories": [
            {
             	"theme": "{{.CATEGORYTITLE}}",
                "subtheme": "{{.SUBCATEGORYTITLE}}",
                "groupname": "{{.GROUPNAME}}"
            }
	],
	"apps": ["energize"],
        "standards": [
            "FGDC-STD-001-1998",
            "ISO-19115:2003"
        ],
	"project": "NM EPSCoR R4",
        "services": [],
        "formats": [{{.XLSX}}"json","xls","csv"],
        "metadata": {
            "xml": "{{.RAWXML}}",
            "upgrade": "true",
            "standard": "FGDC-STD-001-1998"
        }
}`


const FILEJSON string = `{
        "description": "{{.DATASETNAME}}",
        "orig_epsg": {{.ORIGEPSG}},
        "taxonomy": "file",
        "basename": "{{.BASENAME}}",
        "is_embargoed": "{{.ISEMBARGOED}}",
        "spatial": {
            "epsg": {{.EPSG}},
            "geomtype": "{{.GEOMTYPE}}",
            "features": "{{.FEATURES}}",
            "records": {{.RECORDS}},
            "bbox": "{{.WESTBC}},{{.SOUTHBC}},{{.EASTBC}},{{.NORTHBC}}"
        },
        "sources": [
            {
                "mimetype": "application/x-zip-compressed",
                "files": [
                    "/geodata/energize/uploads/{{.UPLOADUSER}}/{{.FILENAME}}"
                ],
                "set": "original",
                "external": "False",
                "extension": "{{.EXTENSION}}"
            }
        ],
        "active": "True",
        "dataone_archive":"{{.DATAONEARCHIVE}}",
        "releasedate":"{{.RELEASEDATE}}",
        "author":"{{.FIRSTNAMEPI}} {{.LASTNAMEPI}}",
        "categories": [
            {
                "theme": "{{.CATEGORYTITLE}}",
                "subtheme": "{{.SUBCATEGORYTITLE}}",
                "groupname": "{{.GROUPNAME}}"
            }
        ],
        "apps": ["energize"],
        "standards": [
            "FGDC-STD-001-1998",
            "ISO-19115:2003"
        ],
        "project": "NM EPSCoR R4",
        "services": [],
        "formats": ["{{.EXTENSION}}"],
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




const MAIL string =`<!doctype html>
<html>
  <head>
    <meta name="viewport" content="width=device-width" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>Simple Transactional Email</title>
    <style>
      /* -------------------------------------
          GLOBAL RESETS
      ------------------------------------- */
      img {
        border: none;
        -ms-interpolation-mode: bicubic;
        max-width: 100%; }

      body {
        background-color: #f6f6f6;
        font-family: sans-serif;
        -webkit-font-smoothing: antialiased;
        font-size: 14px;
        line-height: 1.4;
        margin: 0;
        padding: 0;
        -ms-text-size-adjust: 100%;
        -webkit-text-size-adjust: 100%; }

      table {
        border-collapse: separate;
        mso-table-lspace: 0pt;
        mso-table-rspace: 0pt;
        width: 100%; }
        table td {
          font-family: sans-serif;
          font-size: 14px;
          vertical-align: top; }

      /* -------------------------------------
          BODY & CONTAINER
      ------------------------------------- */

      .body {
        background-color: #f6f6f6;
        width: 100%; }

      /* Set a max-width, and make it display as block so it will automatically stretch to that width, but will also shrink down on a phone or something */
      .container {
        display: block;
        Margin: 0 auto !important;
        /* makes it centered */
        max-width: 580px;
        padding: 10px;
        width: 580px; }

.customborder{
  width: 85%;
  color: #555;
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 20px;
}


      /* This should also be a block element, so that it will fill 100% of the .container */
      .content {
        box-sizing: border-box;
        display: block;
        Margin: 0 auto;
        max-width: 580px;
        padding: 10px; }

      /* -------------------------------------
          HEADER, FOOTER, MAIN
      ------------------------------------- */
      .main {
        background: #fff;
        border-radius: 3px;
        width: 100%; }

      .wrapper {
        box-sizing: border-box;
        padding: 20px; }

      .footer {
        clear: both;
        padding-top: 10px;
        text-align: center;
        width: 100%; }
        .footer td,
        .footer p,
        .footer span,
        .footer a {
          color: #999999;
          font-size: 12px;
          text-align: center; }

      /* -------------------------------------
          TYPOGRAPHY
      ------------------------------------- */
      h1,
      h2,
      h3,
      h4 {
        color: #000000;
        font-family: sans-serif;
        font-weight: 400;
        line-height: 1.4;
        margin: 0;
        Margin-bottom: 30px; }

      h1 {
        font-size: 35px;
        font-weight: 300;
        text-align: center;
        text-transform: capitalize; }

      p,
      ul,
      ol {
        font-family: sans-serif;
        font-size: 14px;
        font-weight: normal;
        margin: 0;
        Margin-bottom: 15px; }
        p li,
        ul li,
        ol li {
          list-style-position: inside;
          margin-left: 5px; }

      a {
        color: #3498db;
        text-decoration: underline; }

      /* -------------------------------------
          BUTTONS
      ------------------------------------- */
      .btn {
        box-sizing: border-box;
        width: 100%; }
        .btn > tbody > tr > td {
          padding-bottom: 15px; }
        .btn table {
          width: auto; }
        .btn table td {
          background-color: #ffffff;
          border-radius: 5px;
          text-align: center; }
        .btn a {
          background-color: #ffffff;
          border: solid 1px #3498db;
          border-radius: 5px;
          box-sizing: border-box;
          color: #3498db;
          cursor: pointer;
          display: inline-block;
          font-size: 14px;
          font-weight: bold;
          margin: 0;
          padding: 12px 25px;
          text-decoration: none;
          text-transform: capitalize; }

      .btn-primary table td {
        background-color: #38646f; }

      .btn-primary a {
        background-color: #8e0c0d;
        border-color: #3498db;
        color: #ffffff; }

      /* -------------------------------------
          OTHER STYLES THAT MIGHT BE USEFUL
      ------------------------------------- */
      .last {
        margin-bottom: 0; }

      .first {
        margin-top: 0; }

      .align-center {
        text-align: center; }

      .align-right {
        text-align: right; }

      .align-left {
        text-align: left; }

      .clear {
        clear: both; }

      .mt0 {
        margin-top: 0; }

      .mb0 {
        margin-bottom: 0; }

      .preheader {
        color: transparent;
        display: none;
        height: 0;
        max-height: 0;
        max-width: 0;
        opacity: 0;
        overflow: hidden;
        mso-hide: all;
        visibility: hidden;
        width: 0; }

      .powered-by a {
        text-decoration: none; }

      hr {
        border: 0;
        border-bottom: 1px solid #f6f6f6;
        Margin: 20px 0; }

      /* -------------------------------------
          RESPONSIVE AND MOBILE FRIENDLY STYLES
      ------------------------------------- */
      @media only screen and (max-width: 620px) {
        table[class=body] h1 {
          font-size: 28px !important;
          margin-bottom: 10px !important; }
        table[class=body] p,
        table[class=body] ul,
        table[class=body] ol,
        table[class=body] td,
        table[class=body] span,
        table[class=body] a {
          font-size: 16px !important; }
        table[class=body] .wrapper,
        table[class=body] .article {
          padding: 10px !important; }
        table[class=body] .content {
          padding: 0 !important; }
        table[class=body] .container {
          padding: 0 !important;
          width: 100% !important; }
        table[class=body] .main {
          border-left-width: 0 !important;
          border-radius: 0 !important;
          border-right-width: 0 !important; }
        table[class=body] .btn table {
          width: 100% !important; }
        table[class=body] .btn a {
          width: 100% !important; }
        table[class=body] .img-responsive {
          height: auto !important;
          max-width: 100% !important;
          width: auto !important; }}

      /* -------------------------------------
          PRESERVE THESE STYLES IN THE HEAD
      ------------------------------------- */
      @media all {
        .ExternalClass {
          width: 100%; }
        .ExternalClass,
        .ExternalClass p,
        .ExternalClass span,
        .ExternalClass font,
        .ExternalClass td,
        .ExternalClass div {
          line-height: 100%; }
        .apple-link a {
          color: inherit !important;
          font-family: inherit !important;
          font-size: inherit !important;
          font-weight: inherit !important;
          line-height: inherit !important;
          text-decoration: none !important; }
        .btn-primary table td:hover {
          background-color: #34495e !important; }
        .btn-primary a:hover {
          background-color: #34495e !important;
          border-color: #34495e !important; } }

	.well {
	  min-height: 0px;
	  padding: 0px;
	  margin-bottom: 0px;
	  background-color: #e6e6e6;
	  border: 1px solid #e3e3e3;
	  border-radius: 0px;
	  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, .05);
	          box-shadow: inset 0 1px 1px rgba(0, 0, 0, .05);
          border-bottom: none;
	}

    </style>
  </head>
  <body class="">
    <table border="0" cellpadding="0" cellspacing="0" class="body">
      <tr>
        <td>&nbsp;</td>
        <td class="container">
          <div class="content">

            <!-- START CENTERED WHITE CONTAINER -->
            <span class="preheader">This is preheader text. Some clients will show this text as a preview.</span>
            <table class="main">

              <!-- START MAIN CONTENT AREA -->
              <tr>
                <td class="wrapper">
                  <table border="0" cellpadding="0" cellspacing="0">
                    <tr>
		    <td>
                     <div class="well well-sm"style="padding: 0px;">
                      <img alt="Header Image" src="https://reporting.nmepscor.org/sites/all/modules/er/static/img/datareview.png" />
                     </div>
                     <br>
                     <br>
                        <p>Hello {{.USERFIRSTNAME}},</p>
                        <p>{{.TOPMESSAGE}}</p>

                        <a href="{{.BUTTONLINK}}" target="_blank">{{.BUTTONTEXT}}</a><br><br>
			<p>Notes:</p>
                        {{.NOTES}}
                        <p>Please feel free to reply to this e-mail with any questions or concerns.</p>
                        <p>Thank you in advance for your cooperation in this matter.</p>
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>

              <!-- END MAIN CONTENT AREA -->
              </table>

            <!-- START FOOTER -->
            <div class="footer">
              <table border="0" cellpadding="0" cellspacing="0">
                <tr>
                  <td class="content-block">

                    <a href="https://www.nmepscor.org">New Mexico EPSCoR</a> is funded by the National Science Foundation (NSF)
                  </td>
                </tr>
              </table>
            </div>

            <!-- END FOOTER -->
<!-- END CENTERED WHITE CONTAINER --></div>
        </td>
        <td>&nbsp;</td>
      </tr>
    </table>
  </body>
</html>
`

const Message string = `<html>
<head>
  <title>Form Edit</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://reporting.nmepscor.org/sites/all/modules/er/static/css/bootstrap.min.css">
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/jquery.min.js"></script>
  <script src="https://reporting.nmepscor.org/sites/all/modules/er/static/js/bootstrap.min.js"></script>

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

</style>
</head>
<body>



<nav class="navbar navbar-inverse navbar-fixed-top">
  <div class="container-fluid">
    <div class="navbar-header">
   <a class="navbar-brand" href="/formedit/"><img src=" data:image/png;base64,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" alt="NM EPSCoR Logo"></a>
    </div>
    <ul class="nav navbar-nav">
              <li><a href="/formedit/?status=1">In Progress</a></li>
              <li><a href="/formedit/?status=2">Submitted for Approval</a></li>
              <li><a href="/formedit/?status=3">Approved</a></li>
              <li><a href="/formedit/?status=4">Certified</a></li>
              <li><a href="/formedit/?status=5">Inserted into GSToRE</a></li>
    </ul>
  </div>
</nav>


<div class="container">
 <div class="alert alert-info" style="text-align: center">
  <strong> {{.MESSAGE}}</strong>
 </div>


</div>
</body>
</html>`
