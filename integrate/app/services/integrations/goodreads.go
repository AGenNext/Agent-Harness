// GoodReads - Developer learning resources
package goodreads

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Rating float64 `json:"rating"`
}

type Client struct {
	APIKey string
}

func NewClient(key string) *Client {
	return &Client{APIKey: key}
}

func (c *Client) Search(topic string) []Book {
	return []Book{
		{ID: "1", Title: "Clean Code", Author: "Robert Martin", Rating: 4.7},
		{ID: "2", Title: "The Pragmatic Programmer", Author: "Hunt & Thomas", Rating: 4.6},
	}
}

func (c *Client) Recommend(role string) []Book {
	m := map[string][]Book{
		"frontend": {
			{Title: "You Don't Know JS", Author: "Kyle Simpson", Rating: 4.8},
		},
		"backend": {
			{Title: "Database Internals", Author: "Alex Petrov", Rating: 4.4},
		},
		"devops": {
			{Title: "The DevOps Handbook", Author: "Gene Kim", Rating: 4.7},
		},
	}
	return m[role]
}