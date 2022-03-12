package dbutils

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/jpwhite3/git-samurai/fsutils"
	"github.com/stretchr/testify/assert"
)

func TestGetDatabase(t *testing.T) {
	db := GetDatabase()
	assert.Equal(t, "sqlite", db.Name())
	assert.True(t, fsutils.IsFile(DatabaseName))
	DeleteDatabase()
	assert.False(t, fsutils.IsFile(DatabaseName))
}

func Test_recordCommitLineage(t *testing.T) {
	type args struct {
		commitHash string
		parentHash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := recordCommitLineage(tt.args.commitHash, tt.args.parentHash); (err != nil) != tt.wantErr {
				t.Errorf("recordCommitLineage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertCommit(t *testing.T) {
	type args struct {
		c *object.Commit
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertCommit(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("InsertCommit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRecordBlame(t *testing.T) {
	type args struct {
		c *object.Commit
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RecordBlame(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RecordBlame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
