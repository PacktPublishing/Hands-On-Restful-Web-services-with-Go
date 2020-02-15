package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter7/jsonstore/helper"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// DBClient stores the database session imformation. Needs to be initialized once
type DBClient struct {
	db *gorm.DB
}

// PackageResponse is the response to be send back for Package
type PackageResponse struct {
	Package helper.Package `json:"Package"`
}

// GetPackage fetches a package
func (driver *DBClient) GetPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	vars := mux.Vars(r)
	// Handle response details
	driver.db.First(&Package, vars["id"])
	var PackageData interface{}
	// Unmarshal JSON string to interface
	json.Unmarshal([]byte(Package.Data), &PackageData)
	var response = PackageResponse{Package: Package}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}

// GetPackagesbyWeight fetches all packages with given weight
func (driver *DBClient) GetPackagesbyWeight(w http.ResponseWriter, r *http.Request) {
	var packages []helper.Package
	weight := r.FormValue("weight")
	// Handle response details
	var query = "select * from \"Package\" where data->>'weight'=?"
	driver.db.Raw(query, weight).Scan(&packages)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(packages)
	w.Write(respJSON)
}

// PostPackage saves a package
func (driver *DBClient) PostPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	postBody, _ := ioutil.ReadAll(r.Body)
	Package.Data = string(postBody)
	driver.db.Save(&Package)
	responseMap := map[string]interface{}{"id": Package.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

func main() {
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create a new router
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/v1/package/{id:[a-zA-Z0-9]*}", dbclient.GetPackage).Methods("GET")
	r.HandleFunc("/v1/package", dbclient.PostPackage).Methods("POST")
	r.HandleFunc("/v1/package", dbclient.GetPackagesbyWeight).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
