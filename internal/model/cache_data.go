package model

type CacheData[T any] struct {
	Data  T             `json:"data"`
	Pages *PageMetadata `json:"pages"`
}
