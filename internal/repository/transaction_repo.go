package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TransactionRepo struct {
	col *mongo.Collection
}

func NewTransactionRepo(db *mongo.Database) *TransactionRepo {
	return &TransactionRepo{
		col: db.Collection("transactions"),
	}
}

func (r *TransactionRepo) GetGGR(ctx context.Context, from, to time.Time) ([]bson.M, error) {
	pipeline := mongo.Pipeline{
		{{"$match", bson.M{"createdAt": bson.M{"$gte": from, "$lte": to}}}},
		{{"$group", bson.M{
			"_id": "$currency",
			"wager": bson.M{"$sum": bson.M{
				"$cond": []interface{}{bson.M{"$eq": []string{"$type", "Wager"}}, "$usdAmount", 0},
			}},
			"payout": bson.M{"$sum": bson.M{
				"$cond": []interface{}{bson.M{"$eq": []string{"$type", "Payout"}}, "$usdAmount", 0},
			}},
		}}},
		{{"$project", bson.M{
			"currency": "$_id",
			"ggr": bson.M{"$subtract": []string{"$wager", "$payout"}},
		}}},
	}

	cur, err := r.col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var res []bson.M
	err = cur.All(ctx, &res)
	return res, err
}