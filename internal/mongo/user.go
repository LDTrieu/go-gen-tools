package mongo

import (
	"context"
	"go-gen-tools/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	coll *mongo.Collection
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	if _, err := r.coll.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var result models.User
	err := r.coll.FindOne(ctx, &models.User{Email: email}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) FindByUserID(ctx context.Context, userID primitive.ObjectID) (*models.User, error) {
	var result models.User
	if err := r.coll.FindOne(ctx, &models.User{ID: userID}).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// func (r *userRepository) FindAll(ctx context.Context,
// 	offset, limit int) ([]*models.User, error) {
// 	var result []*models.User
// 	opt := &options.FindOptions{}
// 	opt.SetSkip(int64(offset))
// 	opt.SetLimit(int64(limit))
// 	cur, err := r.coll.Find(ctx, bson.M{
// 		"role": bson.M{
// 			"$ne": pb.UserRole_SUPER_ADMIN.String(),
// 		},
// 	}, opt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }
