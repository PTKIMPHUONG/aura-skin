package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SaleRepository interface {
	GetAllSales(page int, pageSize int, search string) ([]models.Sale, error)
	GetSaleByID(id string) (models.Sale, error)
	GetSalesByDateStart(dateStart string) ([]models.Sale, error)
	GetSalesByDateEnd(dateEnd string) ([]models.Sale, error)
	CreateSale(sale models.Sale, variantID string) error
	UpdateSale(id string, sale models.Sale) error
	DeleteSale(id string) error
}

type saleRepository struct {
	db *databases.Neo4jDB
}

func NewSaleRepository(db *databases.Neo4jDB) SaleRepository {
	return &saleRepository{db: db}
}

func (repo *saleRepository) GetAllSales(page int, pageSize int, search string) ([]models.Sale, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	skip := (page - 1) * pageSize

	query := `
		MATCH (s:Sale)
		WHERE s.sale_id CONTAINS $search OR s.description CONTAINS $search
		RETURN s ORDER BY s.dateStart SKIP $skip LIMIT $limit
	`
	params := map[string]interface{}{
		"search": search,
		"skip":   skip,
		"limit":  pageSize,
	}

	result, err := session.Run(ctx, query, params)
	if err != nil {
		return nil, err
	}

	var sales []models.Sale
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("s")
		saleNode := node.(neo4j.Node)

		saleMap := saleNode.Props
		sale, err := (&models.Sale{}).FromMap(saleMap)
		if err != nil {
			return nil, err
		}
		sales = append(sales, *sale)
	}

	if len(sales) == 0 {
		return repo.GetAllSales(page, pageSize, "")
	}

	return sales, nil
}

func (repo *saleRepository) GetSaleByID(id string) (models.Sale, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (s:Sale {sale_id: $id}) RETURN s", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return models.Sale{}, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("s")
		saleNode := node.(neo4j.Node)

		saleMap := saleNode.Props
		sale, err := (&models.Sale{}).FromMap(saleMap)
		if err != nil {
			return models.Sale{}, err
		}
		return *sale, nil
	}

	return models.Sale{}, errors.New("sale with id " + id + " not found")
}

func (repo *saleRepository) GetSalesByDateStart(dateStart string) ([]models.Sale, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (s:Sale {dateStart: $dateStart}) RETURN s", map[string]interface{}{
		"dateStart": dateStart,
	})
	if err != nil {
		return nil, err
	}

	var sales []models.Sale
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("s")
		saleNode := node.(neo4j.Node)

		saleMap := saleNode.Props
		sale, err := (&models.Sale{}).FromMap(saleMap)
		if err != nil {
			return nil, err
		}
		sales = append(sales, *sale)
	}

	return sales, nil
}

func (repo *saleRepository) GetSalesByDateEnd(dateEnd string) ([]models.Sale, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (s:Sale {dateEnd: $dateEnd}) RETURN s", map[string]interface{}{
		"dateEnd": dateEnd,
	})
	if err != nil {
		return nil, err
	}

	var sales []models.Sale
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("s")
		saleNode := node.(neo4j.Node)

		saleMap := saleNode.Props
		sale, err := (&models.Sale{}).FromMap(saleMap)
		if err != nil {
			return nil, err
		}
		sales = append(sales, *sale)
	}

	return sales, nil
}

func (repo *saleRepository) CreateSale(sale models.Sale, variantID string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	// Check if ProductVariant exists and is active
	variantExistsResult, err := tx.Run(ctx,
		"MATCH (v:ProductVariant {variant_id: $variantID, is_active: true}) RETURN v",
		map[string]interface{}{
			"variantID": variantID,
		},
	)
	if err != nil {
		return err
	}

	if !variantExistsResult.Next(ctx) {
		return errors.New("variant does not exist or is not active")
	}

	saleMap := sale.ToMap()
	sale.IsActive = true // Default active sale

	// Create Sale if not exists
	saleExistsResult, err := tx.Run(ctx,
		"MATCH (s:Sale {sale_id: $sale_id}) RETURN s",
		map[string]interface{}{
			"sale_id": saleMap["sale_id"],
		},
	)
	if err != nil {
		return err
	}

	if !saleExistsResult.Next(ctx) {
		_, err = tx.Run(ctx,
			`CREATE (s:Sale {sale_id: $sale_id, dateStart: $dateStart, dateEnd: $dateEnd, percentSale: $percentSale,
              description: $description, isActive: $isActive})`,
			saleMap,
		)
		if err != nil {
			return err
		}
	}

	// Create relationship APPLIES_TO_VARIANT with ProductVariant
	_, err = tx.Run(ctx,
		`
        MATCH (v:ProductVariant {variant_id: $variantID}), (s:Sale {sale_id: $sale_id})
        MERGE (s)-[r:APPLIES_TO_VARIANT]->(v)
        RETURN r
        `,
		map[string]interface{}{
			"variantID": variantID,
			"sale_id":   saleMap["sale_id"],
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *saleRepository) UpdateSale(id string, sale models.Sale) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	saleMap := sale.ToMap()

	// Check if Sale exists
	saleExistsResult, err := tx.Run(ctx,
		"MATCH (s:Sale {sale_id: $sale_id}) RETURN s",
		map[string]interface{}{
			"sale_id": id,
		},
	)
	if err != nil {
		return err
	}

	if !saleExistsResult.Next(ctx) {
		return errors.New("sale does not exist")
	}

	_, err = tx.Run(ctx,
		`MATCH (s:Sale {sale_id: $sale_id}) 
         SET s.dateStart = $dateStart, s.dateEnd = $dateEnd, 
             s.percentSale = $percentSale, s.description = $description, 
             s.isActive = $isActive  
         RETURN s`,
		saleMap,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *saleRepository) DeleteSale(id string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	// Check if Sale exists
	saleExistsResult, err := tx.Run(ctx,
		"MATCH (s:Sale {sale_id: $sale_id}) RETURN s",
		map[string]interface{}{
			"sale_id": id,
		},
	)
	if err != nil {
		return err
	}

	if !saleExistsResult.Next(ctx) {
		return errors.New("sale does not exist")
	}

	_, err = tx.Run(ctx,
		`MATCH (s:Sale {sale_id: $sale_id}) 
         SET s.isActive = false 
         RETURN s`,
		map[string]interface{}{
			"sale_id": id,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}