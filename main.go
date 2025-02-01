package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wili abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()
	log.Println("🔍 Searching...")
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %s", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Created Index in %s for docs %d", time.Since(start), len(docs))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Found %d matches in %s", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		log.Println(id, " -- ", docs[id])
	}
}
