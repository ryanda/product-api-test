package obj

import (
	"encoding/json"
)

type ProductsResponse []ProductsResponseElement

func UnmarshalProductsResponse(data []byte) (ProductsResponse, error) {
	var r ProductsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type ProductsResponseElement struct {
	AverageRating    string      `json:"average_rating"`
	Categories       []Category  `json:"categories"`
	DateCreated      string      `json:"date_created"`
	DateModified     string      `json:"date_modified"`
	Description      string      `json:"description"`
	Dimensions       Dimensions  `json:"dimensions"`
	ID               int64       `json:"id"`
	Images           []Image     `json:"images"`
	Name             string      `json:"name"`
	OnSale           bool        `json:"on_sale"`
	Price            string      `json:"price"`
	Purchasable      bool        `json:"purchasable"`
	RatingCount      int64       `json:"rating_count"`
	RegularPrice     string      `json:"regular_price"`
	SalePrice        string      `json:"sale_price"`
	ShortDescription string      `json:"short_description"`
	Sku              string      `json:"sku"`
	Slug             string      `json:"slug"`
	StockQuantity    interface{} `json:"stock_quantity"`
	StockStatus      string      `json:"stock_status"`
	TotalSales       int64       `json:"total_sales"`
	Weight           string      `json:"weight"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Dimensions struct {
	Height string `json:"height"`
	Length string `json:"length"`
	Width  string `json:"width"`
}

type Image struct {
	Alt          string `json:"alt"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Src          string `json:"src"`
}
