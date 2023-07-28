package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "os"
    "fmt"
)

type Version struct {
    CurrentVersion string `json:"current_version"`
}

func getVersionData(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var v Version
    json.Unmarshal(body, &v)

    return v.CurrentVersion, nil
}

func readLocalVersionFile(file_path string) (string, error) {
    file, err := os.Open(file_path)
    if err != nil {
        return "", err
    }
    defer file.Close()

    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        return "", err
    }

    var v Version
    json.Unmarshal(bytes, &v)

    return v.CurrentVersion, nil
}

func main() {
    url := "http://127.0.0.1/data/version.json"
    file_path := "current_version.json"

    version_from_url, err := getVersionData(url)
    if err != nil {
        fmt.Println("Failed to fetch data from URL: ", err)
        return
    }

    version_from_file, err := readLocalVersionFile(file_path)
    if err != nil {
        fmt.Println("Failed to open file: ", err)
        return
    }

    if version_from_url == version_from_file {
        fmt.Println("Version match. The version is: ", version_from_url)
    } else {
        fmt.Println("Version mismatch. URL version: ", version_from_url, ". File version: ", version_from_file)
    }
}