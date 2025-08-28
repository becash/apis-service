package domain

type AutoIncrement struct {
	ID  string `bson:"_id"`
	SEQ int32  `bson:"seq"`
}
