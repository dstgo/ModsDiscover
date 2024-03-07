package entity

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	tableOptions = "gorm:table_options"
)

// Table
// represents an entity for table
type Table interface {
	TableName() string
	TableComment() string
}

var tables = []Table{
	// user
	&User{},
	// role
	&Role{},
	&Permission{},
	// dict
	&Dict{},
	&DictData{},

	// relation table
	&RolePermission{},
	&UserRole{},
}

// Migrate auto migration all the table defined
func Migrate(db *gorm.DB) error {
	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			err := db.Set(tableOptions, fmt.Sprintf(`comment '%s'`, table.TableComment())).Migrator().CreateTable(table)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
