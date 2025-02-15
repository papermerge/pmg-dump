package exporter

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/papermerge/pmg-dump/models"
)

var NodeID2UUID map[int]uuid.UUID
var UserID2UUID map[int]uuid.UUID

func GetUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, username, email FROM core_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	UserID2UUID = make(map[int]uuid.UUID)

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.EMail)
		if err != nil {
			return nil, err
		}
		user.UUID = uuid.New()
		UserID2UUID[user.ID] = user.UUID
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetNodes(db *sql.DB) ([]models.Node, error) {
	query := `
    SELECT
      tree.id,
      tree.title,
      tree.parent_id,
      tree.user_id,
      ct.model,
      doc.version,
      doc.file_name,
      doc.page_count
    FROM core_basetreenode tree
    JOIN django_content_type ct ON ct.id = tree.polymorphic_ctype_id
    LEFT JOIN core_document doc ON doc.basetreenode_ptr_id = tree.id
  `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []models.Node

	NodeID2UUID = make(map[int]uuid.UUID)

	for rows.Next() {
		var node models.Node
		err = rows.Scan(
			&node.ID,
			&node.Title,
			&node.ParentID,
			&node.UserID,
			&node.Model,
			&node.Version,
			&node.FileName,
			&node.PageCount)
		if err != nil {
			return nil, err
		}
		node.UUID = uuid.New()
		NodeID2UUID[node.ID] = node.UUID
		nodes = append(nodes, node)
	}
	return nodes, nil
}
