package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	movieID := "tt0372784"
	resp, err := http.Get("http://www.omdbapi.com/?i=" + movieID + "&plot=short&r=json")
	if err != nil {
		fmt.Println("error in file retrieval", err)
	}
	defer resp.Body.Close()

	fmt.Println("status code is", resp.Status)

	var m MoviesOMDB
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		fmt.Println("error parsing body", err)
		return
	}

	if m.ImdbRating == "N/A" {
		fmt.Printf("N/A \n")
	} else {
		ratingFloat, err := strconv.ParseFloat(m.ImdbRating, 64)
		ratingInt := int(ratingFloat * 10)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		fmt.Printf("The movie : %s was released in %s - the IMBD rating is %d%% with %s votes\n", m.Title, m.Year, ratingInt, m.ImdbVotes)
	}
}
