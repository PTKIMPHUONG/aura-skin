package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type UserRepository interface {
	GetUsersByName(name string) ([]models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUserNode(user models.User) error
	UpdateUserNode(id string, user models.User) error
	DeleteUserNode(id string) error
	GetUserByID(id string) (models.User, error)
}

type userRepository struct {
	db *databases.Neo4jDB
}

func NewUserRepository(db *databases.Neo4jDB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) GetUserByID(id string) (models.User, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (u:User {id: $id}) RETURN u", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return models.User{}, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("u")
		if !found {
			return models.User{}, errors.New("user not found")
		}

		userNode := node.(neo4j.Node)
		user := models.User{
			ID:       userNode.Props["id"].(string),
			Username: userNode.Props["username"].(string),
			Email:    userNode.Props["email"].(string),
			Password: userNode.Props["password"].(string),
		}
		return user, nil
	}

	return models.User{}, errors.New("user with id " + id + " not found")
}

func (repo *userRepository) CreateUserNode(user models.User) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	_, err = tx.Run(ctx,
		"CREATE (u:User {id: $id, username: $username, email: $email, password: $password})",
		map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"password": user.Password,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *userRepository) UpdateUserNode(id string, user models.User) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	result, err := tx.Run(ctx,
		"MATCH (u:User {id: $id}) SET u.username = $username, u.email = $email, u.password = $password RETURN u",
		map[string]interface{}{
			"id":       id,
			"username": user.Username,
			"email":    user.Email,
			"password": user.Password,
		},
	)
	if err != nil {
		return err
	}

	if !result.Next(ctx) {
		return errors.New("user with id " + id + " not found")
	}

	return tx.Commit(ctx)
}

func (repo *userRepository) DeleteUserNode(id string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	result, err := tx.Run(ctx,
		"MATCH (u:User {id: $id}) DELETE u",
		map[string]interface{}{
			"id": id,
		},
	)
	if err != nil {
		return err
	}

	summary, err := result.Consume(ctx)
	if err != nil {
		return err
	}

	if summary.Counters().NodesDeleted() == 0 {
		return errors.New("user with id " + id + " not found")
	}

	return tx.Commit(ctx)
}

func (repo *userRepository) GetUsersByName(name string) ([]models.User, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx,
		"MATCH (u:User) WHERE u.username CONTAINS $name RETURN u",
		map[string]interface{}{
			"name": name,
		},
	)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("u")
		if !found {
			continue
		}

		userNode := node.(neo4j.Node)
		user := models.User{
			ID:       userNode.Props["id"].(string),
			Username: userNode.Props["username"].(string),
			Email:    userNode.Props["email"].(string),
			Password: userNode.Props["password"].(string),
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("no users found with the name " + name)
	}

	return users, nil
}

func (repo *userRepository) GetUserByEmail(email string) (models.User, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx,
		"MATCH (u:User {email: $email}) RETURN u",
		map[string]interface{}{
			"email": email,
		},
	)
	if err != nil {
		return models.User{}, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("u")
		if !found {
			return models.User{}, errors.New("user not found")
		}

		userNode := node.(neo4j.Node)
		user := models.User{
			ID:       userNode.Props["id"].(string),
			Username: userNode.Props["username"].(string),
			Email:    userNode.Props["email"].(string),
			Password: userNode.Props["password"].(string),
		}
		return user, nil
	}

	return models.User{}, errors.New("user with email " + email + " not found")
}
