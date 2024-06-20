package domain

type Books struct {
    ID     int64                `json:"id"`
    Title  string               `json:"title"`
    Author string               `json:"author"`
    Publication_date string     `json:"publication_date"`
    Genre string                `json:"genre"`
    Status int                  `json:"status"`
    Quantity int                `json:"quantity"`
}