package model

type AnchoreRequestInfo struct {
	ImageName   string `json:"tag"`
	ImageDigest string `json:"digest"`
	CreateTime  string `json:"created_at"`
}
