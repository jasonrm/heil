package main

import (
    "fmt"
    "net/http"
    "crypto/subtle"
    // "strconv"
    "time"
    "flag"
    "os"
)

var expectedPassword = ""

func main() {
    portPtr := flag.Int("port", 8080, "port number")
    passwordPtr := flag.String("basic-auth-password", "", "Must be the same as set on oauth2_proxy")
    flag.Parse()

    expectedPassword = *passwordPtr
    if (expectedPassword == "") {
        fmt.Fprintf(os.Stderr, "Missing required -basic-auth-password argument/flag\n")
        os.Exit(2)
    }

    listenStr := fmt.Sprintf(":%d", *portPtr)

    fmt.Println(time.Now().Format(time.RFC3339), "Listening On", listenStr);
    http.HandleFunc("/auth/user", handleUser)
    http.HandleFunc("/auth/vhost", handleVirtualHost)
    http.HandleFunc("/auth/resource", handleResource)
    http.HandleFunc("/", handleError)
    http.ListenAndServe(listenStr, nil)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.FormValue("username")
    givenPassword := r.FormValue("password")

    body := "deny"
    if (subtle.ConstantTimeCompare([]byte(expectedPassword), []byte(givenPassword)) == 1) {
        body = "allow administrator"
    }

    fmt.Fprint(w, body)
    fmt.Println(time.Now().Format(time.RFC3339), r.Method, "User", body, username, r.UserAgent())
}

func handleVirtualHost(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.FormValue("username")
    vhost := r.FormValue("vhost")
    ip := r.FormValue("ip")

    body := "deny"
    if (len(username) > 0 && len(vhost) > 0) {
        body = "allow"
    }

    fmt.Fprint(w, body)
    fmt.Println(time.Now().Format(time.RFC3339), r.Method, "VirtualHost", body, username, vhost, ip, r.UserAgent())
}

func handleResource(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.FormValue("username")
    vhost := r.FormValue("vhost")
    resource := r.FormValue("resource")
    name := r.FormValue("name")
    permission := r.FormValue("permission")

    body := "deny"
    if (len(username) > 0 && len(vhost) > 0 && len(resource) > 0 && len(name) > 0 && len(permission) > 0) {
        body = "allow"
    }

    fmt.Fprint(w, body)
    fmt.Println(time.Now().Format(time.RFC3339), r.Method, "Resource", body, username, vhost, resource, name, permission, r.UserAgent())
}

func handleError(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
}
