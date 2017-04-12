// Testing go-swagger generation
//
// The purpose of this application is to test go-swagger in a simple GET request.
//
//     Schemes: http
//     Host: localhost:8080
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Daniel<danielfs.ti@gmail.com>
//
//     Consumes:
//     - text/plain
//
//     Produces:
//     - text/plain
//
// swagger:meta
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/hello/{name}", index).Methods("GET")
    router.HandleFunc("/swagger.json", swagger).Methods("GET")

    handler := cors.Default().Handler(router)

    log.Fatal(http.ListenAndServe(":8080", handler))
}

func swagger(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    http.ServeFile(w, r, "swagger.json")
}

func index(w http.ResponseWriter, r *http.Request) {
    // swagger:operation GET /hello/{name} hello Hello
    //
    // Returns a simple Hello message
    // ---
    // consumes:
    // - text/plain
    // produces:
    // - text/plain
    // parameters:
    // - name: name
    //   in: path
    //   description: Name to be returned.
    //   required: true
    //   type: string
    // responses:
    //   '200':
    //     description: The hello message
    //     type: string

    log.Println("Responsing to /hello request")
    log.Println(r.UserAgent())

    vars := mux.Vars(r)
    name := vars["name"]

    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Hello:", name)
}