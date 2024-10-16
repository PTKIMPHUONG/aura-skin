package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SaleRepository interface {
	GetAllSales(page int, pageSize int) ([]map[string]interface{}, error)
	GetSaleByID(id string) (map[string]interface{}, error)
	GetSalesByDateStart(dateStart string, page int, pageSize int) ([]map[string]interface{}, error)
	GetSalesByDateEnd(dateStart string, page int, pageSize int) ([]map[string]interface{}, error)
	CreateSale(sale models.Sale, variantID string) error
	UpdateSale(id string, sale models.Sale) error
	DeleteSale(id string) error
	GetExpiredSales(page int, pageSize int) ([]map[string]interface{}, error) 
	SearchSalesByDescription(description string, page int, pageSize int) ([]map[string]interface{}, error) 
	GetSalesByStatus(isActive bool, page int, pageSize int) ([]map[string]interface{}, error)
}

type saleRepository struct {
	db *databases.Neo4jDB
}

func NewSaleRepository(db *databases.Neo4jDB) SaleRepository {
	return &saleRepository{db: db}
}

func convertSalesToResponse(sales []models.Sale) []map[string]interface{} {
	var response []map[string]interface{}
	for _, sale := range sales {
		response = append(response, sale.ToResponseMap())
	}
	return response
}

func (repo *saleRepository) GetAllSales(page int, pageSize int) ([]map[string]interface{}, error) {
    ctx := context.Background()
    session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close(ctx)

    skip := (page - 1) * pageSize

    query := `
        MATCH (s:Sale)
        RETURN s ORDER BY s.dateStart SKIP $skip LIMIT $limit
    `
    params := map[string]interface{}{
        "skip":  skip,
        "limit": pageSize,
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

	return convertSalesToResponse(sales), nil
}

func (repo *saleRepository) GetSaleByID(id string) (map[string]interface{}, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (s:Sale {sale_id: $id}) RETURN s", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("s")
		saleNode := node.(neo4j.Node)

		saleMap := saleNode.Props
		sale, err := (&models.Sale{}).FromMap(saleMap)
		if err != nil {
			return nil, err
		}
		return sale.ToResponseMap(), nil 
	}

	return nil, errors.New("sale with id " + id + " not found")
}

func (repo *saleRepository) GetSalesByDateStart(dateStart string, page int, pageSize int) ([]map[string]interface{}, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	skip := (page - 1) * pageSize

	query := `
    MATCH (s:Sale)
    WHERE s.date_start >= $dateStart
    RETURN s ORDER BY s.date_start SKIP $skip LIMIT $limit
	`

	params := map[string]interface{}{
		"dateStart": dateStart,
		"skip":      skip,
		"limit":     pageSize,
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

	return convertSalesToResponse(sales), nil
}

func (repo *saleRepository) GetSalesByDateEnd(dateEnd string, page int, pageSize int) ([]map[string]interface{}, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	skip := (page - 1) * pageSize

	query := `
    MATCH (s:Sale)
    WHERE s.date_end <= $dateEnd
    RETURN s ORDER BY s.date_end SKIP $skip LIMIT $limit
	`

	params := map[string]interface{}{
		"dateEnd": dateEnd,
		"skip":      skip,
		"limit":     pageSize,
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

	return convertSalesToResponse(sales), nil
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
			`CREATE (s:Sale {sale_id: $sale_id, date_start: $date_start, date_end: $date_end, percent_sale: $percent_sale,
              description: $description, is_active: $is_active})`,
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
         SET s.date_start = $date_start, s.date_end = $date_end, 
             s.percent_sale = $percent_sale, s.description = $description, 
             s.is_active = $is_active  
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
         SET s.is_active = false 
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

func (repo *saleRepository) GetExpiredSales(page int, pageSize int) ([]map[string]interface{}, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	skip := (page - 1) * pageSize

	query := `
        MATCH (s:Sale)
        WHERE date(s.date_end) < date() AND s.is_active = true
        RETURN s ORDER BY s.date_end SKIP $skip LIMIT $limit
    `
	params := map[string]interface{}{
		"skip":  skip,
		"limit": pageSize,
	}
	fmt.Println("Running query with params:", params)
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
	return convertSalesToResponse(sales), nil
}

func (repo *saleRepository) SearchSalesByDescription(description string, page int, pageSize int) ([]map[string]interface{}, error) {
    ctx := context.Background()
    session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close(ctx)

    skip := (page - 1) * pageSize

    queryStr := `
        MATCH (s:Sale)
        WHERE s.description CONTAINS $description
        RETURN s ORDER BY s.date_start SKIP $skip LIMIT $limit
    `
    params := map[string]interface{}{
        "description": description,
        "skip":  skip,
        "limit": pageSize,
    }

    result, err := session.Run(ctx, queryStr, params)
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

    return convertSalesToResponse(sales), nil
}

func (repo *saleRepository) GetSalesByStatus(isActive bool, page int, pageSize int) ([]map[string]interface{}, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	skip := (page - 1) * pageSize
	query := "MATCH (s:Sale {is_active: $isActive}) RETURN s SKIP $skip LIMIT $limit"

	params := map[string]interface{}{
		"isActive": isActive,
		"skip":     skip,
		"limit":    pageSize,
	}
	result, err := session.Run(ctx, query, params)
	if err != nil {
		fmt.Println("Error mapping users:", err)
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
		return nil, errors.New("no sales found")
	}

	return convertSalesToResponse(sales), nil
}