package dbutils

import (
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Commit struct {
	gorm.Model
	Hash           string `gorm:"unique"`
	TreeHash       string
	ParentHashes   string
	AuthorName     string
	AuthorEmail    string
	AuthorWhen     *time.Time
	CommitterName  string
	CommitterEmail string
	CommitterWhen  *time.Time
	Message        string
	Signature      string
}

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("samurai.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to create database")
	}

	// Migrate the schema
	db.AutoMigrate(&Commit{})

	return db
}

func InsertCommit(c *object.Commit) error {
	db := GetDatabase()
	commit := Commit{
		Hash:           c.Hash.String(),
		TreeHash:       c.TreeHash.String(),
		AuthorName:     c.Author.Name,
		AuthorEmail:    c.Author.Email,
		AuthorWhen:     &c.Author.When,
		CommitterName:  c.Committer.Name,
		CommitterEmail: c.Committer.Email,
		CommitterWhen:  &c.Committer.When,
		Message:        c.Message,
		Signature:      c.PGPSignature,
	}
	result := db.Create(&commit)
	if result.Error != nil {
		panic("Unable to create database")
	}
	return nil
}
