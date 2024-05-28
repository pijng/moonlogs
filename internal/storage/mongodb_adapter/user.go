package mongodb_adapter

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserStorage(db *mongo.Database) *UserStorage {
	return &UserStorage{
		db:         db,
		collection: db.Collection("users"),
	}
}

func (s *UserStorage) CreateUser(ctx context.Context, user entities.User) (*entities.User, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, "users")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	user.ID = nextValue
	update := bson.M{
		"id": user.ID, "email": user.Email, "password": "", "password_digest": user.PasswordDigest,
		"name": user.Name, "role": user.Role, "tag_ids": user.Tags, "token": "", "is_revoked": user.IsRevoked,
	}

	result, err := s.collection.InsertOne(ctx, update)
	if err != nil {
		return nil, fmt.Errorf("failed inserting user: %w", err)
	}

	var u entities.User
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&u)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying inserted user: %w", err)
	}

	return &u, nil
}

func (s *UserStorage) GetUserByID(ctx context.Context, id int) (*entities.User, error) {
	var u entities.User

	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	return &u, nil
}

func (s *UserStorage) GetUsersByTagID(ctx context.Context, tagID int) ([]*entities.User, error) {
	filter := bson.M{"tag_ids": bson.M{"$elemMatch": bson.M{"$eq": tagID}}}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	users := make([]*entities.User, 0)

	for cursor.Next(ctx) {
		var u entities.User
		if err := cursor.Decode(&u); err != nil {
			return nil, fmt.Errorf("failed decoding user: %w", err)
		}

		users = append(users, &u)
	}

	return users, nil
}

func (s *UserStorage) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var u entities.User

	err := s.collection.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying user by email: %w", err)
	}

	return &u, nil
}

func (s *UserStorage) GetUserByToken(ctx context.Context, token string) (*entities.User, error) {
	var u entities.User

	err := s.collection.FindOne(ctx, bson.M{"token": token}).Decode(&u)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying user by token: %w", err)
	}

	return &u, nil
}

func (s *UserStorage) DeleteUserByID(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting user: %w", err)
	}

	return err
}

func (s *UserStorage) UpdateUserByID(ctx context.Context, id int, user entities.User) (*entities.User, error) {

	update := bson.M{
		"email": user.Email, "name": user.Name, "role": user.Role, "tag_ids": user.Tags, "is_revoked": user.IsRevoked,
	}

	if len(user.PasswordDigest) > 0 {
		update["password_digest"] = user.PasswordDigest
		update["token"] = user.Token
	}

	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": update})
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	return s.GetUserByID(ctx, id)
}

func (s *UserStorage) UpdateUserTokenByID(ctx context.Context, id int, token string) error {
	update := bson.M{"$set": bson.M{"token": token}}

	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return fmt.Errorf("failed updating user token: %w", err)
	}

	return err
}

func (s *UserStorage) GetAllUsers(ctx context.Context) ([]*entities.User, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying users: %w", err)
	}
	defer cursor.Close(ctx)

	users := make([]*entities.User, 0)

	for cursor.Next(ctx) {
		var u entities.User
		if err := cursor.Decode(&u); err != nil {
			return nil, fmt.Errorf("failed decoding user: %w", err)
		}

		users = append(users, &u)
	}

	return users, nil
}

func (s *UserStorage) CreateInitialAdmin(ctx context.Context, admin entities.User) (*entities.User, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, "users")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	document := bson.M{
		"id": nextValue, "name": admin.Name, "email": admin.Email, "password": "",
		"password_digest": admin.PasswordDigest, "role": "Admin", "token": "", "is_revoked": admin.IsRevoked,
	}

	result, err := s.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("failed inserting admin: %w", err)
	}

	var u entities.User
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&u)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying inserted admin: %w", err)
	}

	return &u, nil
}
