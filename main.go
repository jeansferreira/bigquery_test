package main

import (
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	bq "github.com/jeansferreira/bigquery"
	"google.golang.org/api/iterator"
)

// export GOOGLE_APPLICATION_CREDENTIALS=/home/jean-ferreira/Projetos-go/projeto-bigquery-291719-accfbedbbde6.json
// export GCP_STORAGE_PROJECT_ID=projeto-bigquery-291719

func main() {
	//variavel
	projectID := os.Getenv("GCP_STORAGE_PROJECT_ID")

	_query := "SELECT full_name, age FROM projectBQID.projectDSID.pessoa"

	ctx, client, err := bq.ConnectBQ(projectID)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	it, err := bq.QueryBQ(ctx, client, _query)
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("bigquery.iterator: %v", err)
		}
		fmt.Println(values)
	}
}
