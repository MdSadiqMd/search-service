# Search Service 🔍

A high-performance document search engine written in Go, capable of scaling to 10 million records with sub-second query times.

<img src="https://github.com/user-attachments/assets/05c2498c-ed77-42d8-b958-0bf383ef8ca2" alt="Search Service Architecture" width="full" height="400">

## Features

- ⚡️ Lightning-fast search using inverted index
- 📈 Scales to 10M+ documents
- 🎯 Microsecond query response times
- 🔄 Support for both compressed and uncompressed XML data
- 💾 Memory-efficient index structure
- 🛠️ Easy to integrate and extend

## Performance

Based on actual benchmarks with Wikipedia dataset for the search query of `sadiq` (my name 😁):

| Operation | Documents | Time |
|-----------|-----------|------|
| Loading Documents | 693,381 | 23.86s |
| Creating Index | 693,381 | 9.85s |
| Search Query | 8 matches | 3.68µs |

The inverted index approach enables extremely fast search operations, performing full-text search across nearly 700,000 documents in just microseconds.

![Inverted Index Structure](https://github.com/user-attachments/assets/9cdbb523-934d-47d7-ad18-9de33f29a908)

## Getting Started

### Prerequisites

- Go 1.19 or higher
- At least 4GB RAM (for processing large datasets)

### Installation

```bash
git clone https://github.com/yourusername/search-service.git
cd search-service
go mod download
```

### Dataset Setup

1. Download the Wikipedia abstract dataset:
   ```bash
   wget https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz
   ```

2. Extract the dataset:
   ```bash
   gunzip enwiki-latest-abstract1.xml.gz
   ```

3. Move the XML file to the project root:
   ```bash
   mv enwiki-latest-abstract1.xml /path/to/search-service/
   ```

### Usage

Run the search service with default parameters:
```bash
go run main.go
```

Custom search parameters:
```bash
go run main.go -p path/to/data.xml -q "your search query"
```

#### Command Line Flags

- `-p`: Path to the XML data file (default: "enwiki-latest-abstract1.xml")
- `-q`: Search query (default: "Small wild cat")

## How It Works

### Inverted Index Approach

<img src="https://github.com/user-attachments/assets/6805648c-c059-46d6-95ec-f258023586f8" alt="Index Creation Process" width="full" height="400">

Instead of using simple string matching, we implement an inverted index that:
1. Processes and tokenizes document text
2. Creates a mapping of terms to document IDs
3. Enables efficient querying by looking up pre-indexed terms

This approach significantly reduces search time complexity from O(n*m) to O(1) for term lookups.
