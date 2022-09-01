package v1

type Product struct {
	Id          string `json:"_id"`
	ProductName string `json:"product_name"`
	ImageThumb  string `json:"image_front_thumb_url"`
	Links       []Link `json:"links"`
}

type ProductDetail struct {
	Id          string       `json:"_id"`
	ProductName string       `json:"product_name"`
	ImageThumb  string       `json:"image_front_thumb_url"`
	Image       string       `json:"image_front_url"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Id              string `json:"id"`
	PercentEstimate int    `json:"percent_estimate"`
	PercentMax      int    `json:"percent_max"`
	PercentMin      int    `json:"percent_min"`
	Text            string `json:"text"`
	Vegan           bool   `json:"vegan"`
	Vegetarian      bool   `json:"vegetarian"`
}

type ProductsResponse struct {
	Count     int       `json:"count"`
	Page      int       `json:"page"`
	PageCount int       `json:"page_count"`
	PageSize  int       `json:"page_size"`
	Products  []Product `json:"products"`
	Skip      int       `json:"skip"`
}

type ProductDetailResponse struct {
	Product ProductDetail `json:"product"`
}
