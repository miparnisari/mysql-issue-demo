// Package assets contains database migration scripts and test files
package assets

import "embed"

const (
	MySQLMigrationDir = "migrations/mysql"
)

// EmbedMigrations within the openfga binary.
//
//go:embed migrations/*
var EmbedMigrations embed.FS
