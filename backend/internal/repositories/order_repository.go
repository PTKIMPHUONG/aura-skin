package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
	CreateOrder(order models.Order, userID string, productVariantID string) error
	CancelOrder(orderID string) error
	GetOrderByID(orderID string) (models.Order, error)
	UpdateOrder(orderID string, order models.Order) error
}

type orderRepository struct {
	db *databases.Neo4jDB
}

func NewOrderRepository(db *databases.Neo4jDB) OrderRepository {
	return &orderRepository{db: db}
}

func (repo *orderRepository) GetAllOrders() ([]models.Order, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := "MATCH (o:Order) RETURN o"
	result, err := session.Run(ctx, query, nil)
	if err != nil {
		return nil, err
	}

	var orders []models.Order
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("o")
		orderNode := node.(neo4j.Node)

		orderMap := orderNode.Props
		order, err := (&models.Order{}).FromMap(orderMap)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}

	return orders, nil
}

func (repo *orderRepository) CreateOrder(order models.Order, userID string, productVariantID string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	// Kiểm tra sự tồn tại của User
	userExistsResult, err := tx.Run(ctx,
		"MATCH (u:User {id: $userID}) RETURN u",
		map[string]interface{}{
			"userID": userID,
		},
	)
	if err != nil {
		return err
	}

	if !userExistsResult.Next(ctx) {
		return errors.New("user does not exist")
	}

	// Kiểm tra sự tồn tại của ProductVariant
	productVariantExistsResult, err := tx.Run(ctx,
		"MATCH (pv:ProductVariant {variant_id: $variantID}) RETURN pv",
		map[string]interface{}{
			"variantID": productVariantID,
		},
	)
	if err != nil {
		return err
	}

	if !productVariantExistsResult.Next(ctx) {
		return errors.New("product variant does not exist")
	}

	// Tạo Order
	orderMap := order.ToMap()
	_, err = tx.Run(ctx,
		`CREATE (o:Order {order_id: $order_id, country: $country, 
          delivery_fee: $delivery_fee, address_line: $address_line, province: $province, 
          total_amount: $total_amount, district: $district, ward: $ward, 
          recipient_name: $recipient_name, contact_number: $contact_number, status: $status, 
          created_at: $created_at})`,
		orderMap,
	)
	if err != nil {
		return err
	}

	// Tạo relationship giữa Order, User và ProductVariant
	_, err = tx.Run(ctx,
		`
        MATCH (u:User {id: $userID}), (o:Order {order_id: $order_id}), 
              (pv:ProductVariant {variant_id: $variantID})
        MERGE (u)-[:PLACED_ORDER]->(o)
        MERGE (o)-[:CONTAINS]->(pv)
        RETURN o
        `,
		map[string]interface{}{
			"userID":      userID,
			"order_id":    order.OrderID,
			"variantID":   productVariantID,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *orderRepository) UpdateOrder(orderID string, order models.Order) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	orderMap := order.ToMap()

	_, err = tx.Run(ctx,
		`MATCH (o:Order {order_id: $order_id}) 
		 SET o.country = $country, 
			 o.delivery_fee = $delivery_fee, 
			 o.address_line = $address_line, 
			 o.province = $province, 
			 o.total_amount = $total_amount, 
			 o.district = $district, 
			 o.ward = $ward, 
			 o.recipient_name = $recipient_name, 
			 o.contact_number = $contact_number, 
			 o.status = $status, 
			 o.created_at = $created_at 
		 RETURN o`,
		orderMap,
	)
	
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *orderRepository) CancelOrder(orderID string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	// Kiểm tra sự tồn tại của Order
	orderExistsResult, err := tx.Run(ctx,
		"MATCH (o:Order {order_id: $order_id}) RETURN o",
		map[string]interface{}{
			"order_id": orderID,
		},
	)
	if err != nil {
		return err
	}

	if !orderExistsResult.Next(ctx) {
		return errors.New("order does not exist")
	}

	// Cập nhật trạng thái của đơn hàng
	_, err = tx.Run(ctx,
		"MATCH (o:Order {order_id: $order_id}) SET o.status = 'canceled'",
		map[string]interface{}{
			"order_id": orderID,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}


func (repo *orderRepository) GetOrderByID(id string) (models.Order, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (o:Order {order_id: $id}) RETURN o", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return models.Order{}, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("o")
		if !found {
			return models.Order{}, errors.New("order not found")
		}
		orderNode := node.(neo4j.Node)

		orderMap := orderNode.Props
		order, err := (&models.Order{}).FromMap(orderMap)
		if err != nil {
			return models.Order{}, err
		}
		return *order, nil
	}

	return models.Order{}, errors.New("order with id " + id + " not found")
}
