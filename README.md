# Pokedex Lite
A lightweight Pokedex web application

![Pokedex Lite](https://i.imgur.com/EB5kgU4.png)

## How to run
To initialize the templates, get the **jteeuwen/go-bindata** library by running:

`go get -u github.com/jteeuwen/go-bindata/...`

*This will generate a static file with all the HTML templates and CSS information*

Run this to generate the file (do this in the root directory):

`go-bindata -o=assets/bindata.go --nocompress --nometadata --pkg=assets templates/... static/...`

Run `go build` to start the local web server

## Features
* Pokemon information such as type, miscellaneous data and weakness
* Browsing generations
* Search functionality

![Pokemon information](https://i.imgur.com/CGeGHqo.png)
