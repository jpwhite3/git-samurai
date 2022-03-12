package dbutils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DatabaseName string = "samurai.db"

type Commit struct {
	gorm.Model
	Hash           string `gorm:"unique"`
	TreeHash       string
	ParentHashes   string
	IsMergeCommit  bool
	AuthorName     string
	AuthorEmail    string
	AuthorWhen     *time.Time
	CommitterName  string
	CommitterEmail string
	CommitterWhen  *time.Time
	Message        string
	Signature      string
}

type CommitLineage struct {
	gorm.Model
	Hash       string
	ParentHash string
}

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("samurai.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to create database")
	}

	// Migrate the schema
	db.AutoMigrate(&Commit{})
	db.AutoMigrate(&CommitLineage{})

	return db
}

func DeleteDatabase() error {
	e := os.Remove(DatabaseName)
	if e != nil {
		log.Fatal(e)
	}
	return nil
}

func recordCommitLineage(commitHash string, parentHash string) error {
	db := GetDatabase()
	commitLineage := CommitLineage{
		Hash:       commitHash,
		ParentHash: parentHash,
	}
	result := db.Create(&commitLineage)
	if result.Error != nil {
		panic("Unable to record commit lineage")
	}
	return nil
}

func InsertCommit(c *object.Commit) error {
	db := GetDatabase()
	commit := Commit{
		Hash:           c.Hash.String(),
		TreeHash:       c.TreeHash.String(),
		IsMergeCommit:  c.NumParents() > 1,
		AuthorName:     c.Author.Name,
		AuthorEmail:    c.Author.Email,
		AuthorWhen:     &c.Author.When,
		CommitterName:  c.Committer.Name,
		CommitterEmail: c.Committer.Email,
		CommitterWhen:  &c.Committer.When,
		Message:        c.Message,
		Signature:      c.PGPSignature,
	}

	// Record commit data
	result := db.Create(&commit)
	if result.Error != nil {
		panic("Unable to record commit data")
	}

	// Record Commit lineage data
	c.Parents().ForEach(func(parent *object.Commit) error {
		recordCommitLineage(c.Hash.String(), parent.Hash.String())
		return nil
	})

	return nil
}

func RecordBlame(c *object.Commit) error {
	fileIter, err := c.Files()
	if err != nil {
		return err
	}

	// File data
	fileIter.ForEach(func(f *object.File) error {
		fmt.Println(f.Name)
		return nil
	})

	return nil
}
