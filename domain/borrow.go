package domain

type Manage_borrows struct {
    Bid int64 `json:"bid"`
	Uid int64 `json:"uid"`
	IssueDate string `json:"issue_date"`
	ReturnDate string `json:"return_date"`
}