package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type CategoryRepository interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id string) (models.Category, error)
	GetProductsByCategoryID(categoryID string) ([]models.Product, error)
	CreateCategory(category models.Category) error
	UpdateCategory(id string, category models.Category) error
	DeleteCategory(id string) error
}

type categoryRepository struct {
	db *databases.Neo4jDB
}

func NewCategoryRepository(db *databases.Neo4jDB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (repo *categoryRepository) GetAllCategories() ([]models.Category, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (c:Category) RETURN c", nil)
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("c")
		categoryNode := node.(neo4j.Node)

		categoryMap := categoryNode.Props
		category, err := (&models.Category{}).FromMap(categoryMap)
		if err != nil {
			return nil, err
		}
		categories = append(categories, *category)
	}

	return categories, nil
}

func (repo *categoryRepository) GetCategoryByID(id string) (models.Category, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (c:Category {category_id: $id}) RETURN c", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return models.Category{}, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("c")
		if !found {
			return models.Category{}, errors.New("category not found")
		}
		categoryNode := node.(neo4j.Node)

		categoryMap := categoryNode.Props
		category, err := (&models.Category{}).FromMap(categoryMap)
		if err != nil {
			return models.Category{}, err
		}
		return *category, nil
	}

	return models.Category{}, errors.New("category with id " + id + " not found")
}

func (repo *categoryRepository) GetProductsByCategoryID(categoryID string) ([]models.Product, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (c:Category {category_id: $category_id})<-[:BELONGS_TO]-(p:Product)
		RETURN p
	`
	result, err := session.Run(ctx, query, map[string]interface{}{
		"category_id": categoryID,
	})
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

func (repo *categoryRepository) CreateCategory(category models.Category) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	categoryMap := category.ToMap()

	_, err := session.Run(ctx,
		"CREATE (c:Category {category_id: $category_id, category_name: $category_name})",
		categoryMap,
	)
	return err
}

func (repo *categoryRepository) UpdateCategory(id string, category models.Category) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	categoryMap := category.ToMap()

	_, err := session.Run(ctx,
		"MATCH (c:Category {category_id: $category_id}) SET c.category_name = $category_name",
		categoryMap,
	)
	return err
}

func (repo *categoryRepository) DeleteCategory(id string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.Run(ctx, "MATCH (c:Category {category_id: $id}) DELETE c", map[string]interface{}{
		"id": id,
	})
	return err
}
