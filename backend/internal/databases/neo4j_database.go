package databases

import (
	"context"
	config "auraskin/internal/configs/dev"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jDB struct {
	Driver neo4j.DriverWithContext
}

var dbInstance *Neo4jDB

func Instance() *Neo4jDB {
	return dbInstance
}

// Khởi tạo kết nối Neo4j
func NewNeo4jDB(cfg *config.Neo4jConfig) (*Neo4jDB, error) {
	if dbInstance != nil {
		return dbInstance, nil 
	}

	// Tạo driver kết nối đến Neo4j
	driver, err := neo4j.NewDriverWithContext(cfg.URI, neo4j.BasicAuth(cfg.Username, cfg.Password, ""))
	if err != nil {
		return nil, err
	}

	dbInstance = &Neo4jDB{
		Driver: driver,
	}

	return dbInstance, nil
}

//Disconnect Neo4j
func (db *Neo4jDB) Disconnect() {
	if db.Driver != nil {
		err := db.Driver.Close(context.Background())
		if err != nil {
			log.Printf("Failed to disconnect Neo4j connection: %v", err)
		} else {
			log.Println("Neo4j connection disconnected!")
		}
	}
}
