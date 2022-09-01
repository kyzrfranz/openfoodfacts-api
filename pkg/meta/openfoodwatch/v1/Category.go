package v1

type Category struct {
	Id       string `json:"id"`
	Known    int    `json:"known"`
	Name     string `json:"name"`
	Products int    `json:"products"`
	Url      string `json:"url"`
}

type CategoryResponse struct {
	Count int        `json:"count"`
	Tags  []Category `json:"tags"`
}
