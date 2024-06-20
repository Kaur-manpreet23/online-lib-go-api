package domain

type Data1 struct {
    Search_val string `json:"search_val"`
	Search_category string `json:"search_category"`
	Search_data string `json:"search_data"`
}

type Data2 struct {
    Uid int `json:"uid"`
	Bid int `json:"bid"`
	Borrow_id int `json:"borrowId"`
}

type Data3 struct {
	IsValid bool `json:isValid`
	UserId int `json:userid`
}