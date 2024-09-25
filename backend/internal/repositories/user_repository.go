package repositories

// import (
// 	"context"
// 	"firstproject/models"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type UserRepository interface {
// 	CreateUser(user *models.User) error
// 	FindUserByEmail(email string) (*models.User, error)
// 	DeleteUser(userID string) error
// }

// type userRepository struct {
// 	collection *mongo.Collection
// }

// func NewUserRepository(db *mongo.Client) UserRepository {
// 	return &userRepository{collection: db.Database("elec_equipment").Collection("users")}
// }

// func (r *userRepository) CreateUser(user *models.User) error {
// 	user.CreatedAt = time.Now()
// 	_, err := r.collection.InsertOne(context.TODO(), user)
// 	return err
// }

// func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
// 	var user models.User
// 	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
// 	return &user, err
// }

// func (r *userRepository) DeleteUser(email string) error {
// 	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"email": email})
// 	return err
// }
