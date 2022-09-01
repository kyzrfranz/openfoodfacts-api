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
	Nutriments  Nutriment    `json:"nutriments"`
}

type Ingredient struct {
	Id              string  `json:"id"`
	PercentEstimate float64 `json:"percent_estimate"`
	PercentMax      float64 `json:"percent_max"`
	PercentMin      float64 `json:"percent_min"`
	Text            string  `json:"text"`
	Vegan           bool    `json:"vegan"`
	Vegetarian      bool    `json:"vegetarian"`
}

type Nutriment struct {
	Carbohydrates                                      int     `json:"carbohydrates"`
	Carbohydrates100G                                  int     `json:"carbohydrates_100g"`
	CarbohydratesUnit                                  string  `json:"carbohydrates_unit"`
	CarbohydratesValue                                 int     `json:"carbohydrates_value"`
	Energy                                             int     `json:"energy"`
	EnergyKcal                                         int     `json:"energy-kcal"`
	EnergyKcal100G                                     int     `json:"energy-kcal_100g"`
	EnergyKcalUnit                                     string  `json:"energy-kcal_unit"`
	EnergyKcalValue                                    int     `json:"energy-kcal_value"`
	EnergyKj                                           int     `json:"energy-kj"`
	EnergyKj100G                                       int     `json:"energy-kj_100g"`
	EnergyKjUnit                                       string  `json:"energy-kj_unit"`
	EnergyKjValue                                      int     `json:"energy-kj_value"`
	Energy100G                                         int     `json:"energy_100g"`
	EnergyUnit                                         string  `json:"energy_unit"`
	EnergyValue                                        int     `json:"energy_value"`
	Fat                                                int     `json:"fat"`
	Fat100G                                            int     `json:"fat_100g"`
	FatUnit                                            string  `json:"fat_unit"`
	FatValue                                           int     `json:"fat_value"`
	FruitsVegetablesNutsEstimateFromIngredients100G    int     `json:"fruits-vegetables-nuts-estimate-from-ingredients_100g"`
	FruitsVegetablesNutsEstimateFromIngredientsServing int     `json:"fruits-vegetables-nuts-estimate-from-ingredients_serving"`
	NutritionScoreFr                                   int     `json:"nutrition-score-fr"`
	NutritionScoreFr100G                               int     `json:"nutrition-score-fr_100g"`
	Proteins                                           int     `json:"proteins"`
	Proteins100G                                       int     `json:"proteins_100g"`
	ProteinsUnit                                       string  `json:"proteins_unit"`
	ProteinsValue                                      int     `json:"proteins_value"`
	Salt                                               float64 `json:"salt"`
	Salt100G                                           float64 `json:"salt_100g"`
	SaltUnit                                           string  `json:"salt_unit"`
	SaltValue                                          float64 `json:"salt_value"`
	SaturatedFat                                       int     `json:"saturated-fat"`
	SaturatedFat100G                                   int     `json:"saturated-fat_100g"`
	SaturatedFatUnit                                   string  `json:"saturated-fat_unit"`
	SaturatedFatValue                                  int     `json:"saturated-fat_value"`
	Sodium                                             float64 `json:"sodium"`
	Sodium100G                                         float64 `json:"sodium_100g"`
	SodiumUnit                                         string  `json:"sodium_unit"`
	SodiumValue                                        float64 `json:"sodium_value"`
	Sugars                                             int     `json:"sugars"`
	Sugars100G                                         int     `json:"sugars_100g"`
	SugarsUnit                                         string  `json:"sugars_unit"`
	SugarsValue                                        int     `json:"sugars_value"`
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
