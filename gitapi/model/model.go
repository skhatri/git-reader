package model

type Project struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          string `json:"id"`
}

type Repository struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

type Tag struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Hash         string `json:"hash"`
	LatestCommit string `json:"latestCommit"`
	Display      string `json:"display"`
}

type Branch struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	LatestCommit string `json:"latestCommit"`
	Display      string `json:"display"`
	IsDefault    bool   `json:"isDefault"`
}
