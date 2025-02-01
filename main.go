package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	utils "github.com/MdSadiqMd/search-service/utils"
)

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.2f µs", float64(d.Nanoseconds())/1000)
	} else if d < time.Second {
		return fmt.Sprintf("%.2f ms", float64(d.Nanoseconds())/1000000)
	}
	return fmt.Sprintf("%.2f s", d.Seconds())
}

func logWithEmoji(emoji, msg string, start time.Time) {
	log.Printf("%s %s in %s", emoji, msg, formatDuration(time.Since(start)))
}

func main() {
	log.SetFlags(log.Ltime)
	query := flag.String("q", "Small wild cat", "search query")
	dumpPath := flag.String("p", "enwiki-latest-abstract1.xml", "wiki abstract dump path")
	flag.Parse()

	start := time.Now()
	log.Printf("🚀 Starting search service for term: %q ...", *query)

	docs, err := utils.LoadDocuments(*dumpPath)
	if err != nil {
		log.Fatalf("❌ Failed to load documents: %v", err)
	}
	logWithEmoji("📚", fmt.Sprintf("Loaded %d documents", len(docs)), start)

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	logWithEmoji("📇", fmt.Sprintf("Indexed %d documents", len(docs)), start)

	start = time.Now()
	matches := idx.Search(*query)
	logWithEmoji("🔍", fmt.Sprintf("Found %d matches", len(matches)), start)

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
