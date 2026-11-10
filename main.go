package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"encoding/json"
)

type Job struct {
    Title    string `json:"title"`
    Company  string `json:"company_name"`
    Location string `json:"candidate_required_location"`
    URL      string `json:"url"`
}


type JobResponse struct{
	Jobs [] Job `json:"jobs"`
}
func main(){
	url := "https://remotive.com/api/remote-jobs"

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil{
		fmt.Printf("Request failed: %v\n", err)		
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Unexpected error occured", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("An error occured", err)
		return
	}

	var response JobResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("JSON decoding failed:", err)
		return
	}

	fmt.Println(len(response.Jobs))

	if err := saveToCsv(response.Jobs); err != nil {
		fmt.Println("Failed to save CSV:", err)
		return
	}

	fmt.Println("Jobs saved to jobs.csv")

	for _, job := range response.Jobs{
		 fmt.Println("Title:", job.Title)
		fmt.Println("Company:", job.Company)
		fmt.Println("Location:", job.Location)
		fmt.Println("URL:", job.URL)
		fmt.Println("--------------------")
	}

}