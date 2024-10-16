package repositories

import (
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type UserRepository interface {
	Repository[models.User]
	GetUsersByName(name string) ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetOrdersByUserID(id string) ([]models.Order, error)
	UploadProfilePicture(userID string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	GetAllUsers() ([]models.User, error)
	GetUserByRole(isAdmin bool) ([]models.User, error)
	GetProductVariantsByUserID(userID string) ([]models.ProductVariant, error)
	AddToWishlist(userID, variantID string) error
	RemoveFromWishlist(userID, variantID string) error
	GetUserWishlist(userID string) ([]models.ProductVariant, error)
}

type userRepository struct {
	db          *databases.Neo4jDB
	storageRepo StorageRepository
}

func NewUserRepository(db *databases.Neo4jDB, storageRepo StorageRepository) UserRepository {
	return &userRepository{
		db:          db,
		storageRepo: storageRepo,
	}
}

func (repo *userRepository) GetByID(id string) (*models.User, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (u:User {id: $id, is_active: true}) RETURN u", map[string]interface{}{
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
			ID:          userNode.Props["id"].(string),
			Username:    userNode.Props["username"].(string),
			Email:       userNode.Props["email"].(string),
			Password:    userNode.Props["password"].(string),
			PhoneNumber: userNode.Props["phone_number"].(string),
			IsActive:    userNode.Props["is_active"].(bool),
			IsAdmin:     userNode.Props["is_admin"].(bool),
			Gender:      userNode.Props["gender"].(string),
			BirthDate:   userNode.Props["birth_date"].(string),
			User_image:  userNode.Props["user_image"].(string),
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
	user.IsActive = true
	user.User_image = ""

	user.ID = uuid.NewString()
	_, err = tx.Run(ctx,
		"CREATE (u:User {id: $id, username: $username, email: $email, password: $password, phone_number: $phone_number, is_active: $is_active, is_admin: $is_admin, gender: $gender, birth_date: $birth_date, user_image: $user_image})",
		map[string]interface{}{
			"id":           user.ID,
			"username":     user.Username,
			"email":        user.Email,
			"password":     user.Password,
			"phone_number": user.PhoneNumber,
			"is_active":    user.IsActive,
			"is_admin":     user.IsAdmin,
			"gender":       user.Gender,
			"birth_date":   user.BirthDate,
			"user_image":   user.User_image,
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

	if user.Gender != "" {
		params["gender"] = user.Gender
	}

	if user.BirthDate != "" {
		params["birth_date"] = user.BirthDate
	}

	result, err := tx.Run(ctx,
		"MATCH (u:User {id: $id}) SET u.username = $username, u.phone_number = $phone_number, u.password = $password, u.gender = $gender, u.birth_date = $birth_date RETURN u",
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
		"MATCH (u:User {id: $id}) SET u.is_active = false RETURN u",
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
			ID:          userNode.Props["id"].(string),
			Username:    userNode.Props["username"].(string),
			Email:       userNode.Props["email"].(string),
			Password:    userNode.Props["password"].(string),
			PhoneNumber: userNode.Props["phone_number"].(string),
			IsActive:    userNode.Props["is_active"].(bool),
			IsAdmin:     userNode.Props["is_admin"].(bool),
			Gender:      userNode.Props["gender"].(string),
			BirthDate:   userNode.Props["birth_date"].(string),
			User_image:  userNode.Props["user_image"].(string),
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
		user := models.User{
			ID:          userNode.Props["id"].(string),
			Username:    userNode.Props["username"].(string),
			Email:       userNode.Props["email"].(string),
			Password:    userNode.Props["password"].(string),
			PhoneNumber: userNode.Props["phone_number"].(string),
			IsActive:    userNode.Props["is_active"].(bool),
			IsAdmin:     userNode.Props["is_admin"].(bool),
			Gender:      userNode.Props["gender"].(string),
			BirthDate:   userNode.Props["birth_date"].(string),
			User_image:  userNode.Props["user_image"].(string),
		}

		return &user, nil
	}

	return nil, errors.New("user with email " + email + " not found")
}


func (repo *userRepository) GetOrdersByUserID(id string) ([]models.Order, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (u:User {id: $id})-[:PLACED_ORDER]->(o:Order)
		RETURN o
	`
	result, err := session.Run(ctx, query, map[string]interface{}{
		"id": id,
	})
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

	if len(orders) == 0 {
		return nil, errors.New("no orders found for the specified user")
	}

	return orders, nil
}
func (repo *userRepository) UploadProfilePicture(userID string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	profilePictureUrl, err := repo.storageRepo.UploadFile(file, fileHeader, "user_images")
	if err != nil {
		return "", err
	}

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Close(ctx)

	_, err = tx.Run(ctx,
		"MATCH (u:User {id: $id}) SET u.user_image = $profile_picture_url RETURN u",
		map[string]interface{}{
			"id":                  userID,
			"profile_picture_url": profilePictureUrl,
		},
	)
	if err != nil {
		fmt.Println("Error in updating user image in Update function:", err)
		return "", err
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return "", err
	}
	fmt.Println("Successfully updated profile picture in Neo4j with URL:", profilePictureUrl)
	return profilePictureUrl, nil
}

func (repo *userRepository) GetAllUsers() ([]models.User, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := "MATCH (u:User {is_active: true}) RETURN u"
	result, err := session.Run(ctx, query, nil)
	if err != nil {
		return nil, err
	}

	return mapUsers(result, ctx)
}

func mapUsers(result neo4j.ResultWithContext, ctx context.Context) ([]models.User, error) {
	var users []models.User

	for result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("u")
		if !found {
			continue
		}

		userNode := node.(neo4j.Node)
		user := models.User{
			ID:          getStringProp(userNode.Props, "id"),
			Username:    getStringProp(userNode.Props, "username"),
			Email:       getStringProp(userNode.Props, "email"),
			Password:    getStringProp(userNode.Props, "password"),
			PhoneNumber: getStringProp(userNode.Props, "phone_number"),
			User_image:  getStringProp(userNode.Props, "user_image"),
			IsAdmin:     getBoolProp(userNode.Props, "is_admin"),
			IsActive:    getBoolProp(userNode.Props, "is_active"),
			Gender:      getStringProp(userNode.Props, "gender"),
			BirthDate:   getStringProp(userNode.Props, "birth_date"),
		}
		users = append(users, user)
	}

	fmt.Println("Mapped users:", users)

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

// Helper functions for property extraction
func getStringProp(props map[string]interface{}, key string) string {
	if val, ok := props[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return ""
}

func getBoolProp(props map[string]interface{}, key string) bool {
	if val, ok := props[key]; ok {
		if boolVal, ok := val.(bool); ok {
			return boolVal
		}
	}
	return false
}

func (repo *userRepository) GetUserByRole(isAdmin bool) ([]models.User, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (u:User {is_admin: $isAdmin, is_active: true}) RETURN u", map[string]interface{}{
		"isAdmin": isAdmin,
	})
	if err != nil {
		fmt.Println("Error mapping users:", err)
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
			ID:          userNode.Props["id"].(string),
			Username:    userNode.Props["username"].(string),
			Email:       userNode.Props["email"].(string),
			Password:    userNode.Props["password"].(string),
			PhoneNumber: userNode.Props["phone_number"].(string),
			User_image:  userNode.Props["user_image"].(string),
			IsAdmin:     userNode.Props["is_admin"].(bool),
			IsActive:    userNode.Props["is_active"].(bool),
			Gender:      userNode.Props["gender"].(string),
			BirthDate:   userNode.Props["birth_date"].(string),
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

//Get product variant by userID
func (repo *userRepository) GetProductVariantsByUserID(userID string) ([]models.ProductVariant, error) {
    ctx := context.Background()
    session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close(ctx)

    query := `
        MATCH (u:User {id: $userID})-[:PLACED_ORDER]->(o:Order)-[:CONTAINS]->(pv:ProductVariant)
        RETURN pv
    `
    result, err := session.Run(ctx, query, map[string]interface{}{
        "userID": userID,
    })
    if err != nil {
        return nil, err
    }

    var productVariants []models.ProductVariant
    for result.Next(ctx) {
        record := result.Record()
        node, _ := record.Get("pv")
        productVariantNode := node.(neo4j.Node)
        productVariantMap := productVariantNode.Props

        productVariant, err := (&models.ProductVariant{}).FromMap(productVariantMap)
        if err != nil {
            return nil, err
        }
        productVariants = append(productVariants, *productVariant)
    }

    if len(productVariants) == 0 {
        return nil, errors.New("no product variants found for the specified user")
    }

    return productVariants, nil
}

func (repo *userRepository) AddToWishlist(userID, variantID string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(ctx, `
			MATCH (u:User {id: $userID, is_active: true}), (v:ProductVariant {variant_id: $variantID, is_active: true})
			MERGE (u)-[r:WISHES_FOR {since: timestamp()}]->(v)
			RETURN u, v
		`, map[string]interface{}{
			"userID":    userID,
			"variantID": variantID,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) RemoveFromWishlist(userID, variantID string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(ctx, `
			MATCH (u:User {id: $userID})-[r:WISHES_FOR]->(v:ProductVariant {variant_id: $variantID})
			DELETE r
			RETURN u, v
		`, map[string]interface{}{
			"userID":    userID,
			"variantID": variantID,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetUserWishlist(userID string) ([]models.ProductVariant, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close(ctx)

	result, err := session.Run(ctx, `
		MATCH (u:User {id: $userID})-[r:WISHES_FOR]->(v:ProductVariant)
		RETURN v
	`, map[string]interface{}{
		"userID": userID,
	})
	if err != nil {
		return nil, err
	}

	var variants []models.ProductVariant
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("v")
		variantNode := node.(neo4j.Node)

		variantMap := variantNode.Props
		variant, err := (&models.ProductVariant{}).FromMap(variantMap)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}

	return variants, nil
}