package model

type ArticleType struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}
type ArticleType1 struct {
	Model     Model
	TagID     int32 `json:"tag_id"`
	ArticleID int32 `json:"article_id"`
}
