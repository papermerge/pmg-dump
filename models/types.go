package models

import (
	"github.com/google/uuid"
)

const (
	FolderModelName   = "folder"
	DocumentModelName = "document"
)

type User struct {
	ID       int
	UUID     uuid.UUID
	Username string
	EMail    string
}

type Node struct {
	ID         int
	UUID       uuid.UUID
	Title      string
	Model      string
	UserID     int       `yaml:"user_id"`
	FileName   *string   `yaml:"file_name"`
	PageCount  *int      `yaml:"page_count"`
	ParentID   *int      `yaml:"parent_id"`
	ParentUUID uuid.UUID `yaml:"parent_uuid"`
	Version    *int
}

type Folder struct {
	ID         int
	UUID       uuid.UUID
	Title      string
	UserID     int        `yaml:"user_id"`
	UserUUID   uuid.UUID  `yaml:"user_uuid"`
	ParentID   *int       `yaml:"parent_id"`
	ParentUUID *uuid.UUID `yaml:"parent_uuid"`
}

type DocumentVersion struct {
	Number    int
	UUID      uuid.UUID
	FileName  *string `yaml:"file_name"`
	PageCount *int    `yaml:"page_count"`
	Pages     []Page
}

type Document struct {
	ID       int
	UUID     uuid.UUID
	Title    string
	UserID   int  `yaml:"user_id"`
	ParentID *int `yaml:"parent_id"`
	Version  *DocumentVersion
	Versions []DocumentVersion
}

type Page struct {
	UUID   uuid.UUID
	Number int
	Text   string
}

type Data struct {
	Users     []User
	Documents []Document
	Folders   []Folder
}

type FilePath struct {
	Source string
	Dest   string
}
