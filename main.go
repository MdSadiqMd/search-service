package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	utils "github.com/MdSadiqMd/search-service/utils"
)

func main() {
	log.SetFlags(log.Ltime)
	query := flag.String("q", "sadiq", "search query")
	dumpPath := flag.String("p", "enwiki-latest-abstract1.xml", "wiki abstract dump path")
	flag.Parse()

	start := time.Now()
	log.Printf("🚀 Starting search service for term: %q ...", *query)

	docs, err := utils.LoadDocuments(*dumpPath)
	if err != nil {
		log.Fatalf("❌ Failed to load documents: %v", err)
	}
	utils.LogWithEmoji("📚", fmt.Sprintf("Loaded %d documents", len(docs)), start)

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	utils.LogWithEmoji("📇", fmt.Sprintf("Indexed %d documents", len(docs)), start)

	start = time.Now()
	matches := idx.Search(*query)
	utils.LogWithEmoji("🔍", fmt.Sprintf("Found %d matches", len(matches)), start)

	if len(matches) == 0 {
		log.Printf("😕 No results found for %q", *query)
		os.Exit(0)
	}

	log.Printf("📋 Top results for %q:", *query)
	for i, id := range matches {
		if i >= 10 {
			log.Printf("   ... and %d more matches", len(matches)-10)
			break
		}
		log.Printf("   %d. %s\n      URL: %s", i+1, docs[id].Title, docs[id].URL)
	}
}
