package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pokedex-lite/assets"

	"github.com/gorilla/mux"
)

type PokemonGenerationStruct struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonStruct struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type PokemonSearchStruct struct {
	Pokemon   string
}

// Templates
var navigationBarHTML string
var indexPageTpl *template.Template
var generationViewTpl *template.Template
var searchViewTpl *template.Template
//var thirdViewTpl *template.Template

// Pokemon vars
var generationTitle string
var pokemonListOffset int
var pokemonListLimit int

// Search vars
var SearchSuccess bool
var SearchNoResults bool

// Access interface values in template, and return them as string
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return ""
	}
}

// Return ID from url
func getIdFromUrl(url string) string {
	re := regexp.MustCompile(`(/)(\d+)`)
	return re.FindString(url)
}

// Return appropriate color to template depending on type
func getTypeColor(pokemonType string) string {
	switch pokemonType {
	case "normal":
		return "A8A878"
	case "fighting":
		return "C03028"
	case "flying":
		return "A890F0"
	case "poison":
		return "A040A0"
	case "ground":
		return "E0C068"
	case "rock":
		return "B8A038"
	case "bug":
		return "A8B820"
	case "ghost":
		return "705898"
	case "steel":
		return "B8B8D0"
	case "fire":
		return "F08030"
	case "water":
		return "6890F0"
	case "grass":
		return "78C850"
	case "electric":
		return "F8D030"
	case "psychic":
		return "F85888"
	case "ice":
		return "98D8D8"
	case "dragon":
		return "7038F8"
	case "dark":
		return "705848"
	case "fairy":
		return "EE99AC"
	case "unknown":
		return "68A090"
	case "shadow":
		return "000000"
	default:
		return "3564AE"
	}
}

// Uppercase first letter of string
func toTitle(text string) string {
	return strings.Title(text)
}

func init() {
	navigationBarHTML = assets.MustAssetString("templates/navigation_bar.html")

	indexPageHtml := assets.MustAssetString("templates/index.html")
	indexPageTpl = template.Must(template.New("index_view").Funcs(template.FuncMap{"tostring": ToString, "getId": getIdFromUrl, "getColor": getTypeColor, "totitle": toTitle}).Parse(indexPageHtml))

	generationViewHtml := assets.MustAssetString("templates/generation.html")
	generationViewTpl = template.Must(template.New("generation_view").Funcs(template.FuncMap{"tostring": ToString, "getId": getIdFromUrl, "getColor": getTypeColor, "totitle": toTitle}).Parse(generationViewHtml))

	searchViewHtml := assets.MustAssetString("templates/search.html")
	searchViewTpl = template.Must(template.New("search_view").Parse(searchViewHtml))

	//thirdViewFuncMap := ThirdViewFormattingFuncMap()
	//thirdViewHTML := assets.MustAssetString("templates/third_view.html")
	//thirdViewTpl = template.Must(template.New("third_view").Funcs(thirdViewFuncMap).Parse(thirdViewHTML))
}

func main() {
	serverCfg := Config{
		Host:         "localhost:5000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	htmlServer := Start(serverCfg)
	defer htmlServer.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmt.Println("main : shutting down")
}

// Config provides basic configuration
type Config struct {
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// HTMLServer represents the web service that serves up HTML
type HTMLServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

// Start launches the HTML Server
func Start(cfg Config) *HTMLServer {
	// Setup Context
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup Handlers
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/generation/{generation}", GenerationHandler)
	router.HandleFunc("/generation/{generation}/pokemon/{number}", IndexHandler)
	router.HandleFunc("/search", SearchHandler)
	//router.HandleFunc("/third/{number}", ThirdHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Create the HTML Server
	htmlServer := HTMLServer{
		server: &http.Server{
			Addr:           cfg.Host,
			Handler:        router,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}

	// Add to the WaitGroup for the listener goroutine
	htmlServer.wg.Add(1)

	// Start the listener
	go func() {
		fmt.Printf("\nHTMLServer : Service started : Host=%v\n", cfg.Host)
		htmlServer.server.ListenAndServe()
		htmlServer.wg.Done()
	}()

	return &htmlServer
}

// Stop turns off the HTML Server
func (htmlServer *HTMLServer) Stop() error {
	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Printf("\nHTMLServer : Service stopping\n")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := htmlServer.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := htmlServer.server.Close(); err != nil {
			fmt.Printf("\nHTMLServer : Service stopping : Error=%v\n", err)
			return err
		}
	}

	// Wait for the listener to report that it is closed.
	htmlServer.wg.Wait()
	fmt.Printf("\nHTMLServer : Stopped\n")
	return nil
}

// Render a template, or server error.
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

// Push the given resource to the client.
func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

// Route Handlers

func generationUrl(id int) string {
	switch id {
	case 1:
		generationTitle = "I"
		pokemonListOffset = 0
		pokemonListLimit = 151
		return "https://pokeapi.co/api/v2/pokemon/?offset=0&limit=151"
	case 2:
		generationTitle = "II"
		pokemonListOffset = 151
		pokemonListLimit = 100
		return "https://pokeapi.co/api/v2/pokemon/?offset=151&limit=100"
	case 3:
		generationTitle = "III"
		pokemonListOffset = 251
		pokemonListLimit = 135
		return "https://pokeapi.co/api/v2/pokemon/?offset=251&limit=135"
	case 4:
		generationTitle = "IV"
		pokemonListOffset = 386
		pokemonListLimit = 107
		return "https://pokeapi.co/api/v2/pokemon/?offset=386&limit=107"
	case 5:
		generationTitle = "V"
		pokemonListOffset = 493
		pokemonListLimit = 156
		return "https://pokeapi.co/api/v2/pokemon/?offset=493&limit=156"
	case 6:
		generationTitle = "VI"
		pokemonListOffset = 649
		pokemonListLimit = 72
		return "https://pokeapi.co/api/v2/pokemon/?offset=649&limit=72"
	case 7:
		generationTitle = "VII"
		pokemonListOffset = 721
		pokemonListLimit = 86
		return "https://pokeapi.co/api/v2/pokemon/?offset=721&limit=86"
	default:
		generationTitle = "I"
		pokemonListOffset = 0
		pokemonListLimit = 151
		return "https://pokeapi.co/api/v2/pokemon/?offset=0&limit=151"
	}
}

func getGenerationData(id int) PokemonGenerationStruct {
	// Get Pokemon generation data
	pokemonGenerationData, _ := http.Get(generationUrl(id))
	readPokemonGenerationData, _ := ioutil.ReadAll(pokemonGenerationData.Body)

	// Create JSON object
	var pokemonGenerationObject PokemonGenerationStruct
	json.Unmarshal(readPokemonGenerationData, &pokemonGenerationObject)

	return pokemonGenerationObject
}

func getPokemonData(nameOrId string) PokemonStruct {
	// Get Pokemon data
	pokemonData, _ := http.Get("https://pokeapi.co/api/v2/pokemon/"+nameOrId+"/")
	readPokemonData, _ := ioutil.ReadAll(pokemonData.Body)

	// Create JSON object
	var pokemonObject PokemonStruct
	json.Unmarshal(readPokemonData, &pokemonObject)

	return pokemonObject
}

// IndexHandler renders the homepage view template
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	pathVariables := mux.Vars(r)
	generationId, err := strconv.Atoi(pathVariables["generation"]) // Get generation ID parameter from URL
	if err != nil {
		generationId = 1 // If invalid parameter in URL, set to default generation
	}

	// Get generation data
	generationDataObject := getGenerationData(generationId)

	pokemonId, err := strconv.Atoi(pathVariables["number"]) // Get Pokemon ID parameter from URL
	if err != nil {
		pokemonId = 1 // If invalid parameter in URL, set to default
	} else if pokemonId > pokemonListLimit {
		pokemonId = pokemonListLimit // If Pokemon ID exceeds limit, set to limit
	} else if pokemonId <= 0 {
		pokemonId = 1 // If less than or equal to 0, set to default
	}

	pokemonId = pokemonId -1

	// Get Pokemon data
	pokemonDataObject := getPokemonData(generationDataObject.Results[pokemonId].Name)

	fullData := map[string]interface{}{
		"NavigationBar":			template.HTML(navigationBarHTML),
		"GenerationTitle":			generationTitle,
		"GenerationId":				generationId,
		"PokemonArrayIndex":		pokemonId,
		"PokemonPokedexEntries":	generationDataObject.Results,
		"PokemonName":				strings.Title(pokemonDataObject.Name),
		"PokemonId":				appendId(pokemonDataObject.ID),
		"PokemonPreviousName":		strings.Title(generationDataObject.Results[validateArrayIndex(pokemonId - 1)].Name),
		"PokemonPreviousId":		appendId(validateId(pokemonDataObject.ID - 1)),
		"PokemonNextName":			strings.Title(generationDataObject.Results[validateArrayIndex(pokemonId + 1)].Name),
		"PokemonNextId":			appendId(validateId(pokemonDataObject.ID + 1)),
		"PokemonHeight":			pokemonDataObject.Height,
		"PokemonWeight":			pokemonDataObject.Weight,
		"PokemonTypes":				pokemonDataObject.Types,
		"PokemonAbilities":			pokemonDataObject.Abilities,
	}

	if r.Method != http.MethodPost {
		render(w, r, indexPageTpl, "index_view", fullData)
	} else {
		searchData := PokemonSearchStruct {
			Pokemon:   r.FormValue("pokemon"),
		}

		// Get Pokemon data
		pokemonDataObject := getPokemonData(searchData.Pokemon)

		if len(searchData.Pokemon) > 0 {
			if pokemonDataObject.ID != 0 {
				SearchNoResults = false
				SearchSuccess = true
			} else {
				SearchNoResults = true
				SearchSuccess = false
			}
		} else {
			SearchNoResults = false
			SearchSuccess = false
		}

		fullData := map[string]interface{}{
			"NavigationBar":	template.HTML(navigationBarHTML),
			"SearchSuccess":	SearchSuccess,
			"SearchNoResults":	SearchNoResults,
			"PokemonName":		strings.Title(pokemonDataObject.Name),
			"PokemonId":		appendId(pokemonDataObject.ID),
		}

		render(w, r, searchViewTpl, "search_view", fullData)
	}
}

func validateArrayIndex(id int) int {
	if id == -1 {
		return pokemonListLimit - 1
	}

	if id == pokemonListLimit {
		return 0
	}

	return id
}

func validateId(id int) int {
	if id == pokemonListOffset {
		return pokemonListOffset + pokemonListLimit
	}

	if id > pokemonListOffset + pokemonListLimit {
		return pokemonListOffset + 1
	}

	return id;
}

func appendId(id int) string {
	if id <= 0 {
		return "001" // Bulbasaur
	} else if id < 10 {
		return "00" + strconv.Itoa(id)
	} else if id < 100 {
		return "0" + strconv.Itoa(id)
	} else {
		return strconv.Itoa(id)
	}
}

// GenerationHandler renders the generation view template
func GenerationHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	pathVariables := mux.Vars(r)
	generationId, err := strconv.Atoi(pathVariables["generation"]) // Get generation ID parameter from URL
	if err != nil {
		generationId = 1 // If invalid parameter in URL, set to default generation
	}

	generationDataObject := getGenerationData(generationId)

	fullData := map[string]interface{}{
		"NavigationBar":			template.HTML(navigationBarHTML),
		"GenerationTitle":			generationTitle,
		"PokemonPokedexEntries":	generationDataObject.Results,
	}

	render(w, r, generationViewTpl, "generation_view", fullData)
}

// SearchHandler renders the second view template
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	searchData := PokemonSearchStruct {
		Pokemon:   r.FormValue("pokemon"),
	}

	// Get Pokemon data
	pokemonData, _ := http.Get("https://pokeapi.co/api/v2/pokemon/"+searchData.Pokemon+"/")
	readPokemonData, _ := ioutil.ReadAll(pokemonData.Body)

	var pokemonObject PokemonStruct
	json.Unmarshal(readPokemonData, &pokemonObject)

	if len(searchData.Pokemon) > 0 {
		if pokemonObject.ID != 0 {
			SearchNoResults = false
			SearchSuccess = true
		} else {
			SearchNoResults = true
			SearchSuccess = false
		}
	} else {
		SearchNoResults = false
		SearchSuccess = false
	}

	fullData := map[string]interface{}{
		"NavigationBar":	template.HTML(navigationBarHTML),
		"SearchSuccess":	SearchSuccess,
		"SearchNoResults":	SearchNoResults,
		"PokemonName":		strings.Title(pokemonObject.Name),
		"PokemonId":		appendId(pokemonObject.ID),
	}

	render(w, r, searchViewTpl, "search_view", fullData)
}

// ThirdHandler renders the third view template
//func ThirdHandler(w http.ResponseWriter, r *http.Request) {
//	push(w, "/static/style.css")
//	push(w, "/static/navigation_bar.css")
//	push(w, "/static/third_view.css")
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")

//	var queryString string
//	pathVariables := mux.Vars(r)
//	queryNumber, err := strconv.Atoi(pathVariables["number"])
//	if err != nil {
//		queryString = pathVariables["number"]
//	}
//	fullData := map[string]interface{}{
//		"NavigationBar": template.HTML(navigationBarHTML),
//		"Number":        queryNumber,
//		"StringQuery":   queryString,
//	}
//	render(w, r, thirdViewTpl, "third_view", fullData)
//}
