package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type UserRepository interface {
	Repository[models.User]
	GetUsersByName(name string) ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *databases.Neo4jDB
}

func NewUserRepository(db *databases.Neo4jDB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) GetByID(id string) (*models.User, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (u:User {id: $id, isActive: true}) RETURN u", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("u")
		if !found {
			return nil, errors.New("user not found")
		}

		userNode := node.(neo4j.Node)
		user := models.User{
			ID:       userNode.Props["id"].(string),
			Username: userNode.Props["username"].(string),
			Email:    userNode.Props["email"].(string),
			Password: userNode.Props["password"].(string),
		}
		return &user, nil
	}

	return nil, errors.New("user with id " + id + " not found")
}

func (repo *userRepository) Create(user models.User) error {
	exists, _ := repo.GetUserByEmail(user.Email)
	if exists != nil {
		return errors.New("email already exists")
	}
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	user.ID = uuid.NewString()
	_, err = tx.Run(ctx,
		"CREATE (u:User {id: $id, username: $username, email: $email, password: $password, phone_number: $phone_number})",
		map[string]interface{}{
			"id":           user.ID,
			"username":     user.Username,
			"email":        user.Email,
			"password":     user.Password,
			"phone_number": user.PhoneNumber,
		},
	)
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (repo *userRepository) Update(user models.User) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	params := map[string]interface{}{}

	params["id"] = user.ID
	if user.Username != "" {
		params["username"] = user.Username
	}

	if user.PhoneNumber != "" {
		params["phone_number"] = user.PhoneNumber
	}

	if user.Password != "" {
		params["password"] = user.Password
	}

	result, err := tx.Run(ctx,
		"MATCH (u:User {id: $id}) SET u.username = $username, u.phone_number = $phone_number, u.password = $password RETURN u",
		params,
	)
	if err != nil {
		return err
	}

	if !result.Next(ctx) {
		return errors.New("user with id " + user.ID + " not found")
	}

	return tx.Commit(ctx)
}

func (repo *userRepository) Delete(id string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	result, err := tx.Run(ctx,
		"MATCH (u:User {id: $id}) SET u.isActive = false RETURN u",
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

	if summary.Counters().PropertiesSet() == 0 {
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

func (repo *userRepository) GetUserByEmail(email string) (*models.User, error) {
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
		return nil, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("u")
		if !found {
			return nil, errors.New("user not found")
		}

		userNode := node.(neo4j.Node)
		user := models.User{}
		if id, ok := userNode.Props["id"].(string); ok {
			user.ID = id
		} else {
			return nil, errors.New("invalid id")
		}

		if username, ok := userNode.Props["username"].(string); ok {
			user.Username = username
		} else {
			return nil, errors.New("invalid username")
		}

		if email, ok := userNode.Props["email"].(string); ok {
			user.Email = email
		} else {
			return nil, errors.New("invalid email")
		}

		if phoneNumber, ok := userNode.Props["phone_number"].(string); ok {
			user.PhoneNumber = phoneNumber
		} else {
			return nil, errors.New("invalid phone number")
		}

		if password, ok := userNode.Props["password"].(string); ok {
			user.Password = password
		} else {
			return nil, errors.New("invalid password")
		}

		return &user, nil
	}

	return nil, errors.New("user with email " + email + " not found")
}
