package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json.Id`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getCasebooks(w http.ResponseWriter, r *http.Request) {
	query := `{
		launchesPast(limit: 10) {
		  mission_name
		  launch_date_local
		  launch_site {
			site_name_long
		  }
		  links {
			article_link
			video_link
		  }
		  rocket {
			rocket_name
			first_stage {
			  cores {
				flight
				core {
				  reuse_count
				  status
				}
			  }
			}
			second_stage {
			  payloads {
				payload_type
				payload_mass_kg
				payload_mass_lbs
			  }
			}
		  }
		  ships {
			name
			home_port
			image
		  }
		}
	  }`

	json_provider.post(query)
	fmt.Fprintf(w, "Under Construction")
}

func getCasebooksAnnotateDC(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under Construction")
}

func getCasebooksAnnotateDT(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under Construction")
}

func getCasebooksAnnotateDCDT(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under Construction")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches on of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)

	for index, article := range Articles {
		if article.Id == id {
			var updatedArticle Article
			json.Unmarshal(reqBody, &updatedArticle)
			Articles[index] = updatedArticle
		}
	}
}

func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	// add our new DELETE endpoint here

	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	myRouter.HandleFunc("/casebooks/{asset-id}", getCasebooks)
	myRouter.HandleFunc("/casebooks/{asset-id}/annotate-dc", getCasebooksAnnotateDC)
	myRouter.HandleFunc("/casebooks/{asset-id}/annotate-dt", getCasebooksAnnotateDT)
	myRouter.HandleFunc("/casebooks/{asset-id}/annotate-dc-dt", getCasebooksAnnotateDCDT)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
