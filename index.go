package main

var testTpl = `{{.Title}}`

var htmlTPL = `<!doctype html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css"
        integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <title>{{ .Conf.Name }}</title>
    <style type="text/css">
        h1 {
            text-align: center;
            padding-bottom: 5%;
        }

        body {
            background: linear-gradient(rgb(245, 245, 213), rgb(245, 245, 220));
        }

        .head {
            background: linear-gradient(rgb(245, 245, 213), rgb(245, 245, 220));
        }

        pre {
            background: #000000;
            color: #ffffff;
        }

        table {
            border-collapse: collapse;
            width: 100%;
            margin-bottom: 1rem;
            color: #212529;
            border: 1px solid #dee2e6;
        }
    </style>
</head>

<body>
	<div class="container">
        <div class="row">
            <div class="jumbotron col-12 head" style="text-align: center;">
                <h1 class="display-4">{{.Conf.Name}}</h1>
                <p class="lead">{{.Conf.Summary1}}</p>
                <hr class="my-4">
                <p>{{.Conf.Summary2}}</p>
                <a class="btn btn-primary btn-lg" href="#" role="button">Learn more</a>
            </div>
        </div>
    </div>
    <div class="container">
        <div class="row">
            <div class="col-12">
                {{.Content}}
            </div>
        </div>
    </div>


    <script src="https://cdn.jsdelivr.net/npm/jquery@3.4.1/dist/jquery.slim.min.js"
        integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n"
        crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js"
        integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo"
        crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/js/bootstrap.min.js"
        integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6"
        crossorigin="anonymous"></script>


</body>

</html>`
var indexTPL = `<!doctype html>
<html lang="en">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css"
        integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">

    <title>{{ .Conf.Name }}</title>
	<style type="text/css">
        h1 {
            text-align: center;
        }

		body {
            background: linear-gradient(rgb(245, 245, 213), rgb(245, 245, 220));
        }

        .head {
            background: linear-gradient(rgb(245, 245, 213), rgb(245, 245, 220));
        }
    </style>
</head>

<body>
    <div class="container">
        <div class="row">
            <div class="jumbotron col-12 head" style="text-align: center;">
                <h1 class="display-4">{{.Conf.Name}}</h1>
                <p class="lead">{{.Conf.Summary1}}</p>
                <hr class="my-4">
                <p>{{.Conf.Summary2}}</p>
                <a class="btn btn-primary btn-lg" href="#" role="button">Learn more</a>
            </div>
        </div>
    </div>
    <div class="container">
        <div class="row">
<!--
            <div class="col-2">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <a class="nav-link active" href="#">Active</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Link</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Link</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">Disabled</a>
                    </li>
                </ul>
            </div>
-->			
            <div class="col-12">
				{{if .Md.Title}}
				<div class="row" style="padding-bottom: 5%;">
                    <div class="media">
                        <div class="row">
                            <div class="col-3">
                                <img src="{{index .Md.Thumbnail 0}}"
                                    class="mr-3" width="100%">
                            </div>
                            <div class="clo-8 offset-1">
                                <div class="row">
                                    <h5 class="mt-0"><a href="{{.Md.Href}}">{{.Md.Title}}</a></h5>
                                </div>
                                <div class="row">
                                    {{.Md.Summary}}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
<!--
				<div class="row">
                    <div class="media" style="padding: 3%;">
                        <img src="{{index .Md.Thumbnail 0}}"
                            class="mr-3" width="20%">
                        <div class="media-body">
                            <h5 class="mt-0">{{.Md.Title}}</h5>
                            {{.Md.Summary}}
                        </div>
                    </div>
                </div>
-->
				{{end}}
				{{ range .History}}
                <div class="row" style="padding-bottom: 5%;"> 
					<div class="media">
                        <div class="row">
                            <div class="col-3">
                                <img src="{{index .Thumbnail 0}}"
                                    class="mr-3" width="100%">
                            </div>
                            <div class="clo-8 offset-1">
                                <div class="row">
                                    <h5 class="mt-0"><a href="{{.Href}}">{{.Title}}</a></h5>
                                </div>
                                <div class="row">
                                    {{.Summary}}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
				{{end}}
<!--
                <nav aria-label="Page navigation example">
                    <ul class="pagination justify-content-end">
                        <li class="page-item disabled">
                            <a class="page-link" href="#" tabindex="-1" aria-disabled="true">Previous</a>
                        </li>
                        <li class="page-item"><a class="page-link" href="#">1</a></li>
                        <li class="page-item"><a class="page-link" href="#">2</a></li>
                        <li class="page-item"><a class="page-link" href="#">3</a></li>
                        <li class="page-item">
                            <a class="page-link" href="#">Next</a>
                        </li>
                    </ul>
                </nav>
            </div>
-->
	</div>
	
	<!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://cdn.jsdelivr.net/npm/jquery@3.4.1/dist/jquery.slim.min.js"
        integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n"
        crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js"
        integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo"
        crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/js/bootstrap.min.js"
        integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6"
        crossorigin="anonymous"></script>


	</body>

</html>`
