package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"
	"fmt"
	"net/url"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
	GetVariantsByProductID(productID string) ([]models.ProductVariant, error)
	GetVariantsByProductName(productName string) ([]models.ProductVariant, error)
	CreateProduct(product models.Product, categoryID string, supplierID string) error
	UpdateProduct(id string, product models.Product) error
	DeleteProduct(id string) error
}

type productRepository struct {
	db *databases.Neo4jDB
}

func NewProductRepository(db *databases.Neo4jDB) ProductRepository {
	return &productRepository{db: db}
}

func (repo *productRepository) GetAllProducts() ([]models.Product, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (p:Product) RETURN p", nil)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("p")
		productNode := node.(neo4j.Node)

		productMap := productNode.Props
		product, err := (&models.Product{}).FromMap(productMap)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}

func (repo *productRepository) GetProductByID(id string) (models.Product, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (p:Product {product_id: $id}) RETURN p", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return models.Product{}, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("p")
		if !found {
			return models.Product{}, errors.New("product not found")
		}
		productNode := node.(neo4j.Node)

		productMap := productNode.Props
		product, err := (&models.Product{}).FromMap(productMap)
		if err != nil {
			return models.Product{}, err
		}
		return *product, nil
	}

	return models.Product{}, errors.New("product with id " + id + " not found")
}

func (repo *productRepository) GetVariantsByProductID(productID string) ([]models.ProductVariant, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, `
        MATCH (p:Product {product_id: $productID})<-[:BELONGS_TO]-(v:ProductVariant) 
        RETURN v
    `, map[string]interface{}{
		"productID": productID,
	})
	if err != nil {
		return nil, err
	}

	var variants []models.ProductVariant
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("v")
		variantNode := node.(neo4j.Node)

		fmt.Println("Variant Node Props:", variantNode.Props)

		variantMap := variantNode.Props
		variant, err := (&models.ProductVariant{}).FromMap(variantMap)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}

	return variants, nil
}

func (repo *productRepository) GetVariantsByProductName(productName string) ([]models.ProductVariant, error) {
	decodedProductName, err := url.QueryUnescape(productName)
	if err != nil {
		return nil, fmt.Errorf("error decoding product name: %v", err)
	}
	
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, `
        MATCH (p:Product {product_name: $productName})<-[:BELONGS_TO]-(v:ProductVariant) 
        RETURN v
    `, map[string]interface{}{
		"productName": decodedProductName,
	})
	if err != nil {
		return nil, err
	}

	var variants []models.ProductVariant
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("v")
		variantNode := node.(neo4j.Node)

		fmt.Println("Variant Node Props:", variantNode.Props)

		variantMap := variantNode.Props
		variant, err := (&models.ProductVariant{}).FromMap(variantMap)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}

	return variants, nil
}

func (repo *productRepository) CreateProduct(product models.Product, categoryID string, supplierID string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	productMap := product.ToMap()
	product.IsActive = true // Mặc định isActive là true

	// Kiểm tra sự tồn tại của Category và thuộc tính is_active
	categoryExistsResult, err := tx.Run(ctx,
		"MATCH (c:Category {category_id: $categoryID, is_active: true}) RETURN c",
		map[string]interface{}{
			"categoryID": categoryID,
		},
	)
	if err != nil {
		return err
	}

	if !categoryExistsResult.Next(ctx) {
		return errors.New("category does not exist or is not active")
	}

	// Kiểm tra sự tồn tại của Supplier và thuộc tính is_active
	supplierExistsResult, err := tx.Run(ctx,
		"MATCH (s:Supplier {supplier_id: $supplierID, is_active: true}) RETURN s",
		map[string]interface{}{
			"supplierID": supplierID,
		},
	)
	if err != nil {
		return err
	}

	if !supplierExistsResult.Next(ctx) {
		return errors.New("supplier does not exist or is not active")
	}

	// Tạo Product nếu chưa tồn tại
	productExistsResult, err := tx.Run(ctx,
		"MATCH (p:Product {product_id: $product_id}) RETURN p",
		map[string]interface{}{
			"product_id": productMap["product_id"],
		},
	)
	if err != nil {
		return err
	}

	if !productExistsResult.Next(ctx) {
		_, err = tx.Run(ctx,
			`CREATE (p:Product {product_id: $product_id, product_name: $product_name, description: $description, 
              default_price: $default_price, capacity: $capacity, ingredients: $ingredients, features: $features,
			  origin: $origin, manufactured_in: $manufactured_in, usage: $usage,
              default_image: $default_image, expiration_date: $expiration_date, storage: $storage, 
              created_at: $created_at, target_customers: $target_customers, is_active: $is_active})`,
			productMap,
		)
		if err != nil {
			return err
		}
	}

	// Tạo relationship BELONGS_TO với Category
	_, err = tx.Run(ctx,
		`
        MATCH (c:Category {category_id: $categoryID}), (p:Product {product_id: $product_id})
        MERGE (p)-[r:BELONGS_TO]->(c)
        RETURN r
        `,
		map[string]interface{}{
			"categoryID": categoryID,
			"product_id": productMap["product_id"],
		},
	)
	if err != nil {
		return err
	}

	// Tạo relationship SUPPLIER_OF với Supplier
	_, err = tx.Run(ctx,
		`
        MATCH (s:Supplier {supplier_id: $supplierID}), (p:Product {product_id: $product_id})
        MERGE (p)-[r:SUPPLIER_OF]->(s)
        RETURN r
        `,
		map[string]interface{}{
			"supplierID": supplierID,
			"product_id": productMap["product_id"],
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *productRepository) UpdateProduct(id string, product models.Product) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	productMap := product.ToMap()

	// Kiểm tra sản phẩm có tồn tại hay không
	productExistsResult, err := tx.Run(ctx,
		"MATCH (p:Product {product_id: $product_id}) RETURN p",
		map[string]interface{}{
			"product_id": id,
		},
	)
	if err != nil {
		return err
	}

	if !productExistsResult.Next(ctx) {
		return errors.New("product does not exist")
	}

	_, err = tx.Run(ctx,
		`MATCH (p:Product {product_id: $product_id}) 
         SET p.product_name = $product_name, p.description = $description, 
             p.default_price = $default_price, p.capacity = $capacity, 
             p.ingredients = $ingredients, p.features = $features,
			 p.origin = $origin, p.manufactured_in = $manufactured_in, p.usage = $usage, p.default_image = $default_image, 
             p.expiration_date = $expiration_date, p.storage = $storage, 
             p.created_at = $created_at, p.target_customers = $target_customers, 
             p.is_active = $is_active  
         RETURN p`,
		productMap,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *productRepository) DeleteProduct(id string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	// Kiểm tra sản phẩm có tồn tại và is_active = true
	productExistsResult, err := tx.Run(ctx,
		"MATCH (p:Product {product_id: $product_id, is_active: true}) RETURN p",
		map[string]interface{}{
			"product_id": id,
		},
	)
	if err != nil {
		return err
	}

	if !productExistsResult.Next(ctx) {
		return errors.New("product does not exist or is not active")
	}

	// Cập nhật is_active từ true thành false
	_, err = tx.Run(ctx,
		`MATCH (p:Product {product_id: $product_id, is_active: true}) 
         SET p.is_active = false 
         RETURN p`,
		map[string]interface{}{
			"product_id": id,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
