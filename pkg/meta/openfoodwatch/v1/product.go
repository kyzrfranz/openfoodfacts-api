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
	Vegan           string  `json:"vegan"`
	Vegetarian      string  `json:"vegetarian"`
}

type Nutriment struct {
	Carbohydrates                                      float64 `json:"carbohydrates"`
	Carbohydrates100G                                  float64 `json:"carbohydrates_100g"`
	CarbohydratesUnit                                  string  `json:"carbohydrates_unit"`
	CarbohydratesValue                                 float64 `json:"carbohydrates_value"`
	Energy                                             float64 `json:"energy"`
	EnergyKcal                                         float64 `json:"energy-kcal"`
	EnergyKcal100G                                     float64 `json:"energy-kcal_100g"`
	EnergyKcalUnit                                     string  `json:"energy-kcal_unit"`
	EnergyKcalValue                                    float64 `json:"energy-kcal_value"`
	EnergyKj                                           float64 `json:"energy-kj"`
	EnergyKj100G                                       float64 `json:"energy-kj_100g"`
	EnergyKjUnit                                       string  `json:"energy-kj_unit"`
	EnergyKjValue                                      float64 `json:"energy-kj_value"`
	Energy100G                                         float64 `json:"energy_100g"`
	EnergyUnit                                         string  `json:"energy_unit"`
	EnergyValue                                        float64 `json:"energy_value"`
	Fat                                                float64 `json:"fat"`
	Fat100G                                            float64 `json:"fat_100g"`
	FatUnit                                            string  `json:"fat_unit"`
	FatValue                                           float64 `json:"fat_value"`
	FruitsVegetablesNutsEstimateFromIngredients100G    float64 `json:"fruits-vegetables-nuts-estimate-from-ingredients_100g"`
	FruitsVegetablesNutsEstimateFromIngredientsServing float64 `json:"fruits-vegetables-nuts-estimate-from-ingredients_serving"`
	NutritionScoreFr                                   float64 `json:"nutrition-score-fr"`
	NutritionScoreFr100G                               float64 `json:"nutrition-score-fr_100g"`
	Proteins                                           float64 `json:"proteins"`
	Proteins100G                                       float64 `json:"proteins_100g"`
	ProteinsUnit                                       string  `json:"proteins_unit"`
	ProteinsValue                                      float64 `json:"proteins_value"`
	Salt                                               float64 `json:"salt"`
	Salt100G                                           float64 `json:"salt_100g"`
	SaltUnit                                           string  `json:"salt_unit"`
	SaltValue                                          float64 `json:"salt_value"`
	SaturatedFat                                       float64 `json:"saturated-fat"`
	SaturatedFat100G                                   float64 `json:"saturated-fat_100g"`
	SaturatedFatUnit                                   string  `json:"saturated-fat_unit"`
	SaturatedFatValue                                  float64 `json:"saturated-fat_value"`
	Sodium                                             float64 `json:"sodium"`
	Sodium100G                                         float64 `json:"sodium_100g"`
	SodiumUnit                                         string  `json:"sodium_unit"`
	SodiumValue                                        float64 `json:"sodium_value"`
	Sugars                                             float64 `json:"sugars"`
	Sugars100G                                         float64 `json:"sugars_100g"`
	SugarsUnit                                         string  `json:"sugars_unit"`
	SugarsValue                                        float64 `json:"sugars_value"`
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
