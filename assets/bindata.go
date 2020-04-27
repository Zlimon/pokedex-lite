// Code generated by go-bindata.
// sources:
// templates/generation.html
// templates/index.html
// templates/navigation_bar.html
// templates/search.html
// templates/third_view.html
// static/navigation_bar.css
// static/style.css
// static/third_view.css
// DO NOT EDIT!

package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesGenerationHtml = []byte(`<!DOCTYPE html>
<html lang="nb-NO">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Pokedex Lite | Søkeresultater</title>
    <!-- Scripts -->
    <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
    <!-- Styles -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <link rel="stylesheet" href="/static/style.css">
    <link rel="stylesheet" href="/static/navigation_bar.css">
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.4.1/css/all.css" integrity="sha384-5sAR7xN1Nv6T6+dT2mhtzEpVJvfS3NScPQTrOxhwjIuvcA67KV2R5Jz6kr4abQsz" crossorigin="anonymous">
</head>
<style type="text/css">
@media (min-width: 1200px) {
    .container {
        max-width: 1800px;
    }
}

html,
body {
    font-family: 'Ubuntu', sans-serif;
    background-color: #2a4e89;
}

h1,
p {
    color: #ffcb05;
}

.bg-left-custom {
    background-color: #4f7ec9;
}

.bg-right-custom {
    background-color: #3564AE;
}

.item-container {
    display: flex;
    flex-flow: row wrap;
    justify-content: space-around;
}
</style>

<body>
    <div id="app">
        <main class="py-4">
            <div class="container">
                <div class="row">
                    <div class="col-md-5">
                        <div class="col-md-12 rounded-left mb-4 bg-right-custom">
                            {{.NavigationBar}}
                        </div>
                        <div class="col-md-12 rounded mb-4 p-4 bg-left-custom">
                            <h1 class="border-bottom mb-4 text-center">Pokedex Lite</h1>
                            <div class="row d-flex justify-content-center">
                                <div class="col-md-6">
                                    <form method="POST">
                                        <div class="form-group">
                                            <label>Søk etter Pokemon med navn eller ID</label>
                                            <input type="text" class="form-control" id="pokemon" name="pokemon" placeholder="Prøv Pikachu eller 25">
                                        </div>
                                        <button type="submit" class="btn btn-block btn-success">Søk</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-md-7">
                        <div class="col-md-12 rounded-right bg-right-custom p-2">
                            <h1 class="text-center">Generasjon {{.GenerationTitle}}</h1>
                            <div class="item-container">
                                {{range $i, $pokedexEntry := .PokemonPokedexEntries}}
                                <div>
                                    <div class="text-center bg-left-custom rounded m-2 p-2">
                                        <img class="w-100" src="/static/images/sprites/front/{{$pokedexEntry.URL | getId}}.png">
                                    </div>
                                    <p class="text-left ml-2">
                                        <span class="text-dark">
                                            #{{$pokedexEntry.URL | getId}}
                                        </span>
                                        <br>
                                        {{$pokedexEntry.Name | tostring | totitle}}
                                    </p>
                                </div>
                                {{end}}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </div>
</body>

</html>`)

func templatesGenerationHtmlBytes() ([]byte, error) {
	return _templatesGenerationHtml, nil
}

func templatesGenerationHtml() (*asset, error) {
	bytes, err := templatesGenerationHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/generation.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIndexHtml = []byte(`<!DOCTYPE html>
<html lang="nb-NO">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Pokedex Lite | Hjem</title>
    <!-- Scripts -->
    <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
    <!-- Styles -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <link rel="stylesheet" href="/static/style.css">
    <link rel="stylesheet" href="/static/navigation_bar.css">
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.4.1/css/all.css" integrity="sha384-5sAR7xN1Nv6T6+dT2mhtzEpVJvfS3NScPQTrOxhwjIuvcA67KV2R5Jz6kr4abQsz" crossorigin="anonymous">
</head>
<style type="text/css">
@media (min-width: 1200px) {
    .container {
        max-width: 1800px;
    }
}

html,
body {
    font-family: 'Ubuntu', sans-serif;
    background-color: #2a4e89;
}

h1,
p {
    color: #ffcb05;
}

.bg-left-custom {
    background-color: #4f7ec9;
}

.bg-right-custom {
    background-color: #3564AE;
}

.item-container {
    display: flex;
    flex-flow: row wrap;
    justify-content: space-around;
}
</style>

<body>
    <div id="app">
        <main class="py-4">
            <div class="container">
                <div class="row">
                    <div class="col-md-5">
                        <div class="col-md-12 rounded-left mb-4 bg-right-custom">
                            {{.NavigationBar}}
                        </div>
                        <div class="col-md-12 rounded mb-4 p-4 bg-left-custom">
                            <h1 class="border-bottom mb-4 text-center">Pokedex Lite</h1>
                            <div class="row d-flex justify-content-center">
                                <div class="col-md-6">
                                    <form method="POST">
                                        <div class="form-group">
                                            <label>Søk etter Pokemon med navn eller ID</label>
                                            <input type="text" class="form-control" id="pokemon" name="pokemon" placeholder="Prøv Pikachu eller 25">
                                        </div>
                                        <button type="submit" class="btn btn-block btn-success">Søk</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                        <div class="col-md-12 rounded p-4 bg-left-custom">
                            <h1 class="border-bottom mb-4 text-center">Generation {{.GenerationTitle}}</h1>
                            <div class="item-container">
                                {{range $i, $pokedexEntry := .PokemonPokedexEntries}}
                                <div>
                                    <div class="text-center bg-right-custom rounded m-2 p-2">
                                        <img class="w-100" src="/static/images/sprites/front/{{$pokedexEntry.URL | getId}}.png">
                                    </div>
                                    <p class="text-left ml-2">
                                        <span class="text-dark">
                                            #{{$pokedexEntry.URL | getId}}
                                        </span>
                                        <br>
                                        {{$pokedexEntry.Name | tostring | totitle}}
                                    </p>
                                </div>
                                {{end}}
                            </div>
                            <!-- <?php
								$rePattern = '/\d+/';

								echo '
									<ul class="nav nav-tabs" id="myTab" role="tablist">
								';
										foreach ($generationList as $key => $generation) {
											echo '
												<li class="nav-item">
													<a class="nav-link '.(!$generationButton++ ? 'active' : '').'" id="'.$generation->name.'-tab" data-toggle="tab" href="#'.$generation->name.'" role="tab" aria-controls="'.$generation->name.'" aria-selected="false">'.strtoupper(str_replace('generation-', '', $generation->name)).'</a>
												</li>
											';
										}
								echo '
									</ul>
									<div class="tab-content" id="myTabContent">
								';
										foreach ($generationList as $key => $generation) {
											$generationData = getJson($generation->url);

											$pokemonList = $generationData->pokemon_species;

											echo '
												<div class="tab-pane fade '.(!$generationCount++ ? 'show active' : '').'" id="'.$generation->name.'" role="tabpanel" aria-labelledby="'.$generation->name.'-tab">
													<div class="item-container">
											';
														foreach ($pokemonList as $key => $value) {
															echo '
																<div>
																	<div class="text-center bg-right rounded m-2 p-2">
																		<img src="/images/sprites/front/'.str_replace('/', '', preg($value->url)[0]).'.png" style="width: 125px;">
																	</div>
																	<p class="text-left ml-2">
																		<span class="text-dark">#'.appendId(str_replace('/', '', preg($value->url)[0])).'</span>
																		<br>
																		'.ucfirst($value->name).'
																	</span>
																</div>
															';
														}
											echo '
													</div>
												</div>
											';
										}
								echo '
									</div>
								';
							?> -->
                        </div>
                    </div>
                    <div class="col-md-7">
                        <div class="col-md-12 rounded-right bg-right-custom p-2">
                            <div class="row pb-4">
                                <div class="col-md-3">
                                    <a href="/generation/{{.GenerationId}}/pokemon/{{.PokemonPreviousId}}">
                                        <p class="text-center">{{.PokemonPreviousName}} #{{.PokemonPreviousId}}</p>
                                        <img class="w-100 bg-left-custom rounded" src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/{{.PokemonPreviousId}}.png" alt="First slide">
                                    </a>
                                </div>
                                <div class="col-md-6">
                                    <h1 class="text-center">{{.PokemonName}} #{{.PokemonId}}</h1>
                                    <img class="bg-left-custom rounded" src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/{{.PokemonId}}.png" alt="First slide">
                                </div>
                                <div class="col-md-3">
                                    <a href="/generation/{{.GenerationId}}/pokemon/{{.PokemonNextId}}">
                                        <p class="text-center">{{.PokemonNextName}} #{{.PokemonNextId}}</p>
                                        <img class="w-100 bg-left-custom rounded" src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/{{.PokemonNextId}}.png" alt="First slide">
                                    </a>
                                </div>
                            </div>
                            <!-- <div id="carouselExampleControls" class="carousel slide" data-ride="carousel">
								<div class="carousel-inner">
									<div class="carousel-item active">
										<img class="d-block image-carousel" src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/001.png" alt="First slide">
									</div>
									<div class="carousel-item">
										<img class="d-block image-carousel" src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/002.png" alt="Second slide">
									</div>
									<div class="carousel-item">
										<img class="d-block image-carousel" src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/003.png" alt="Third slide">
									</div>
								</div>
								<a class="carousel-control-prev" href="#carouselExampleControls" role="button" data-slide="prev">
									<span class="carousel-control-prev-icon" aria-hidden="true"></span>
									<span class="sr-only">Previous</span>
								</a>
								<a class="carousel-control-next" href="#carouselExampleControls" role="button" data-slide="next">
									<span class="carousel-control-next-icon" aria-hidden="true"></span>
									<span class="sr-only">Next</span>
								</a>
							</div> -->
                            <div class="col-md-6 bg-left-custom mx-auto rounded">
                                <div class="row">
                                    <div class="col-md-6">
                                        <span>Høyde:</span>
                                        <h5>{{.PokemonHeight}}</h5>
                                        <span>Vekt:</span>
                                        <h5>{{.PokemonWeight}}</h5>
                                    </div>
                                    <div class="col-md-6">
                                        <span>Kategori:</span>
                                        <h5>Seed</h5>
                                        <span>Evner:</span>
                                        {{range $i, $ability := .PokemonAbilities}}
                                        <h5>{{$ability.Ability.Name | tostring | totitle}}</h5>
                                        {{end}}
                                    </div>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-md-6">
                                    <h3>Type</h3>
                                    {{range $i, $type := .PokemonTypes}}
                                    <div class="p-2 m-2 rounded" style="background-color: #{{$type.Type.Name | getColor}}">
                                        <span>{{$type.Type.Name | tostring | totitle}}</span>
                                    </div>
                                    {{end}}
                                </div>
                                <div class="col-md-6">
                                    <h3>Svakheter</h3>
                                    <div class="bg-danger p-2 m-2 rounded">
                                        <span>Poison</span>
                                    </div>
                                    <div class="bg-danger p-2 m-2 rounded">
                                        <span>Grass</span>
                                    </div>
                                </div>
                            </div>
                            <!-- <table class="center">
								<tr>
									<td>
										<img style="vertical-align: middle;" src="http://www.pkparaiso.com/imagenes/xy/sprites/animados/magnemite.gif">
									</td>
									<td>
										<img style="vertical-align: middle;" src="http://www.pkparaiso.com/imagenes/xy/sprites/animados-shiny/magnemite.gif">
									</td>
								</tr>
							</table>

							<div class="infoBox">
								<table>
									<tr>
										<td>Høyde: <span class="highlight">{{.Height}}</span> m</td>
										<td>Vekt: <span class="highlight">{{.Weight}}</span> kg</td>
									</tr>
									<tr>
										<td>Type:</td>
										<td>Evner:</td>
									</tr>
									<tr>
										<td><div class="primaryTypeColor">{{.PrimaryType}}</div> <div class="secondaryTypeColor">{{.SecondaryType}}</div></td>
										<td>{{.Abilities}}</td>
									</tr>
								</table>
							</div> -->
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </div>
</body>

</html>`)

func templatesIndexHtmlBytes() ([]byte, error) {
	return _templatesIndexHtml, nil
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNavigation_barHtml = []byte(`<nav class="navbar navbar-expand-lg navbar-dark">
    <a class="navbar-brand" href="/">Pokedex Lite</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
            <li class="nav-item active">
                <a class="nav-link" href="/">Hjem <span class="sr-only">(current)</span></a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/">Quiz?</a>
            </li>
            <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Generasjon</a>
                <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                    <a class="dropdown-item" href="/generation/1">Generation I</a>
                    <a class="dropdown-item" href="/generation/2">Generation II</a>
                    <a class="dropdown-item" href="/generation/3">Generation III</a>
                    <a class="dropdown-item" href="/generation/4">Generation IV</a>
                    <a class="dropdown-item" href="/generation/5">Generation V</a>
                    <a class="dropdown-item" href="/generation/6">Generation VI</a>
                    <a class="dropdown-item" href="/generation/7">Generation VII</a>
                </div>
            </li>
        </ul>
    </div>
</nav>`)

func templatesNavigation_barHtmlBytes() ([]byte, error) {
	return _templatesNavigation_barHtml, nil
}

func templatesNavigation_barHtml() (*asset, error) {
	bytes, err := templatesNavigation_barHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/navigation_bar.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSearchHtml = []byte(`<!DOCTYPE html>
<html lang="nb-NO">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Pokedex Lite | Søkeresultater</title>
    <!-- Scripts -->
    <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
    <!-- Styles -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <link rel="stylesheet" href="/static/style.css">
    <link rel="stylesheet" href="/static/navigation_bar.css">
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.4.1/css/all.css" integrity="sha384-5sAR7xN1Nv6T6+dT2mhtzEpVJvfS3NScPQTrOxhwjIuvcA67KV2R5Jz6kr4abQsz" crossorigin="anonymous">
</head>
<style type="text/css">
@media (min-width: 1200px) {
    .container {
        max-width: 1800px;
    }
}

html,
body {
    font-family: 'Ubuntu', sans-serif;
    background-color: #2a4e89;
}

h1,
p {
    color: #ffcb05;
}

.bg-left-custom {
    background-color: #4f7ec9;
}

.bg-right-custom {
    background-color: #3564AE;
}

.item-container {
    display: flex;
    flex-flow: row wrap;
    justify-content: space-around;
}
</style>

<body>
    <div id="app">
        <main class="py-4">
            <div class="container">
                <div class="row">
                    <div class="col-md-5">
                        <div class="col-md-12 rounded-left mb-4 bg-right-custom">
                            {{.NavigationBar}}
                        </div>
                        <div class="col-md-12 rounded mb-4 p-4 bg-left-custom">
                            <h1 class="border-bottom mb-4 text-center">Pokedex Lite</h1>
                            <div class="row d-flex justify-content-center">
                                <div class="col-md-6">
                                    <form method="POST">
                                        <div class="form-group">
                                            <label>Søk etter Pokemon med navn eller ID</label>
                                            <input type="text" class="form-control" id="pokemon" name="pokemon" placeholder="Prøv Pikachu eller 25">
                                        </div>
                                        <button type="submit" class="btn btn-block btn-success">Søk</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-md-7">
                        <div class="col-md-12 rounded-right bg-right-custom p-2">
                            <h1 class="text-center">Søkeresultater</h1>
                            {{if .SearchSuccess}}
                            <h1 class="text-center">{{.PokemonName}} #{{.PokemonId}}</h1>
                            {{end}}
                            {{if .SearchNoResults}}
                            <h1 class="text-center">Fant ingenting</h1>
                            {{end}}
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </div>
</body>

</html>`)

func templatesSearchHtmlBytes() ([]byte, error) {
	return _templatesSearchHtml, nil
}

func templatesSearchHtml() (*asset, error) {
	bytes, err := templatesSearchHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/search.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesThird_viewHtml = []byte(`<!DOCTYPE html>
<html lang="en">
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <title>Golang HTML Server</title>

        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">

        <link rel="stylesheet" href="/static/style.css">
        <link rel="stylesheet" href="/static/navigation_bar.css">
        <link rel="stylesheet" href="/static/third_view.css">
    </head>
    <body>
        <div class="container">
            {{.NavigationBar}}

            <h1>Rendering Data</h1>
            <h4>This page takes the integer passed in and determines if it is odd or even</h4>
            <div class="result-box">
                {{if .StringQuery}}
                    <h2 class="result-underlined">You didn't enter an integer. This is what was entered:</h2>
                    <h3>{{.StringQuery}}</h3>
                {{else}}
                    <h2 class="result-underlined">The number entered is:</h2>
                    <h3>{{.Number}}</h3>
                    <h2 class="result-underlined">This number is:</h2>
                    <h3>{{.Number | formatOddOrEven}}</h3>
                {{end}}
            </div>
        </div>
    </body>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
</html>`)

func templatesThird_viewHtmlBytes() ([]byte, error) {
	return _templatesThird_viewHtml, nil
}

func templatesThird_viewHtml() (*asset, error) {
	bytes, err := templatesThird_viewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/third_view.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticNavigation_barCss = []byte(``)

func staticNavigation_barCssBytes() ([]byte, error) {
	return _staticNavigation_barCss, nil
}

func staticNavigation_barCss() (*asset, error) {
	bytes, err := staticNavigation_barCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/navigation_bar.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticStyleCss = []byte(`@media (min-width: 1200px) {
    .container {
        max-width: 1800px;
    }
}

html,
body {
    font-family: 'Ubuntu', sans-serif;
    background-color: #2a4e89;
}

h1,
p {
    color: #ffcb05;
}

.bg-left-custom {
    background-color: #4f7ec9;
}

.bg-right-custom {
    background-color: #3564AE;
}

.item-container {
    display: flex;
    flex-flow: row wrap;
    justify-content: space-around;
}`)

func staticStyleCssBytes() ([]byte, error) {
	return _staticStyleCss, nil
}

func staticStyleCss() (*asset, error) {
	bytes, err := staticStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/style.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticThird_viewCss = []byte(`.result-box {
    margin-top: 50px;
}

.result-underlined {
    text-decoration: underline;
}`)

func staticThird_viewCssBytes() ([]byte, error) {
	return _staticThird_viewCss, nil
}

func staticThird_viewCss() (*asset, error) {
	bytes, err := staticThird_viewCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/third_view.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/generation.html": templatesGenerationHtml,
	"templates/index.html": templatesIndexHtml,
	"templates/navigation_bar.html": templatesNavigation_barHtml,
	"templates/search.html": templatesSearchHtml,
	"templates/third_view.html": templatesThird_viewHtml,
	"static/navigation_bar.css": staticNavigation_barCss,
	"static/style.css": staticStyleCss,
	"static/third_view.css": staticThird_viewCss,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"static": &bintree{nil, map[string]*bintree{
		"navigation_bar.css": &bintree{staticNavigation_barCss, map[string]*bintree{}},
		"style.css": &bintree{staticStyleCss, map[string]*bintree{}},
		"third_view.css": &bintree{staticThird_viewCss, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"generation.html": &bintree{templatesGenerationHtml, map[string]*bintree{}},
		"index.html": &bintree{templatesIndexHtml, map[string]*bintree{}},
		"navigation_bar.html": &bintree{templatesNavigation_barHtml, map[string]*bintree{}},
		"search.html": &bintree{templatesSearchHtml, map[string]*bintree{}},
		"third_view.html": &bintree{templatesThird_viewHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
