package utils

type Index map[string][]int

func (idx Index) Add(docs []document) {
	for _, doc := range docs {
		var token string
		for _, token = range analyze(doc.Text) {
			ids := idx[token]
		}
	}
}
