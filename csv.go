package main

import(
	"fmt"
	"encoding/csv"
	"os"
)

func saveToCsv(jobs []Job) error{
	file, err  := os.Create("jobs.csv")

	if err != nil {
		fmt.Println("Failed to create file:", err)
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{"Title", "Company", "Location", "URL"}
	err = writer.Write(headers)
	if err != nil {
		fmt.Println("Failed to write headers:", err)
		return err
	}
	for _, job := range jobs {
		row := []string{
			job.Title,
			job.Company,
			job.Location,
			job.URL,
		}

		err := writer.Write(row)
		if err != nil {
			fmt.Println("Failed to write job:", err)
			return err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}