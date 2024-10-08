package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SupplierRepository interface {
	GetAllSuppliers() ([]models.Supplier, error)
	GetSupplierByID(id string) (*models.Supplier, error)
	CreateSupplier(supplier models.Supplier) error
	UpdateSupplier(supplier models.Supplier) error
	DeleteSupplier(id string) error
}

type supplierRepository struct {
	db *databases.Neo4jDB
}

func NewSupplierRepository(db *databases.Neo4jDB) SupplierRepository {
	return &supplierRepository{db: db}
}

func (repo *supplierRepository) GetAllSuppliers() ([]models.Supplier, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (s:Supplier) RETURN s", nil)
	if err != nil {
		return nil, err
	}

	var suppliers []models.Supplier
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("s")
		supplierNode := node.(neo4j.Node)

		supplierMap := supplierNode.Props
		supplier, err := (&models.Supplier{}).FromMap(supplierMap)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, *supplier)
	}

	return suppliers, nil
}

func (repo *supplierRepository) GetSupplierByID(id string) (*models.Supplier, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (s:Supplier {supplier_id: $id}) RETURN s", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("s")
		if !found {
			return nil, errors.New("supplier not found")
		}
		supplierNode := node.(neo4j.Node)

		supplierMap := supplierNode.Props
		supplier, err := (&models.Supplier{}).FromMap(supplierMap)
		if err != nil {
			return nil, err
		}
		return supplier, nil
	}

	return nil, errors.New("supplier with id " + id + " not found")
}

func (repo *supplierRepository) CreateSupplier(supplier models.Supplier) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	supplierMap := supplier.ToMap()

	_, err := session.Run(ctx,
		"CREATE (s:Supplier {supplier_id: $supplier_id, supplier_name: $supplier_name, supplier_email: $supplier_email, supplier_phone: $supplier_phone, default_image: $default_image})",
		supplierMap,
	)
	return err
}

func (repo *supplierRepository) UpdateSupplier(supplier models.Supplier) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	supplierMap := supplier.ToMap()

	_, err := session.Run(ctx,
		"MATCH (s:Supplier {supplier_id: $supplier_id}) SET s.supplier_name = $supplier_name, s.supplier_email = $supplier_email, s.supplier_phone = $supplier_phone, s.default_image = $default_image",
		supplierMap,
	)
	return err
}

func (repo *supplierRepository) DeleteSupplier(id string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.Run(ctx, "MATCH (s:Supplier {supplier_id: $id}) DELETE s", map[string]interface{}{
		"id": id,
	})
	return err
}
