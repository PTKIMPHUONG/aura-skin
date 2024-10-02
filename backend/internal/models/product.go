package models

type Product struct {
  ProductID   string `json:"product_id"` 
  ProductName  string `json:"product_name"`
  Description  string `json:"description"`
  Price        float64 `json:"price"`
  Category     string `json:"category"`  
  Stock        int    `json:"stock"`     


  Instructions  string `json:"instructions"`
  ManufacturedIn string `json:"manufactured_in"`
  Origin        string `json:"origin"`
  Usage         string `json:"usage"`
  Storage       string `json:"storage"`
  ExpirationDate string `json:"expiration_date"`
  Capacity      string `json:"capacity"`
  Features      string `json:"features"`
  Ingredients   string `json:"ingredients"`
  DefaultImage  string `json:"default_image"`
  TargetCustomers string `json:"target_customers"`
}
