package repository

import (
	"context"
	"music-libray-management/domain/entity"
	"music-libray-management/domain/repository"
	mongodb2 "music-libray-management/infra/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type documentRepository struct {
	mongodb *mongodb2.MongoDB
}

const CollectionDocument = "documents"

func NewDocumentRepository(mongodb *mongodb2.MongoDB) repository.DocumentRepository {
	return &documentRepository{mongodb}
}

func (u *documentRepository) Create(ctx context.Context, document *entity.Document) (string, error) {
	collection := u.mongodb.Collection(CollectionDocument)
	res, err := collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *documentRepository) Update(ctx context.Context, documentID string, updatedDocument *entity.Document) (int64, error) {
	collection := u.mongodb.Collection(CollectionDocument)
	objID, err := primitive.ObjectIDFromHex(documentID)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": updatedDocument,
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return res.MatchedCount, nil
}

func (u *documentRepository) GetByID(ctx context.Context, documentID string) (*entity.Document, error) {
	var document *entity.Document
	collection := u.mongodb.Collection(CollectionDocument)
	objID, err := primitive.ObjectIDFromHex(documentID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	err = collection.FindOne(ctx, filter).Decode(&document)
	return document, err
}

func (u *documentRepository) GetList(ctx context.Context) ([]*entity.Document, error) {
	var documents []*entity.Document
	collection := u.mongodb.Collection(CollectionDocument)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var document *entity.Document
		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	return documents, nil
}

func (u *documentRepository) Delete(ctx context.Context, documentID string) (int64, error) {
	collection := u.mongodb.Collection(CollectionDocument)
	objID, err := primitive.ObjectIDFromHex(documentID)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": objID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
