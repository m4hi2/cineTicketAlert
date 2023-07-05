package dml

import (
	"fmt"
	"github.com/m4hi2/capsule71/pkg/migrator"
)

func createMigrationTemplate(name, version string) string {
	key := fmt.Sprintf("%s_%s", version, name)
	structName := migrator.SnakeCaseToCamelCase(fmt.Sprintf("%s_%s", name, version))
	return fmt.Sprintf(
		`package dmls

import (
	"gorm.io/gorm"
	"github.com/m4hi2/capsule71/pkg/migrator"
	"github.com/m4hi2/capsule71/pkg/migrator/dml"
)

func init() {dml.Register(New%s)}
type %s struct {}
func New%s() migrator.IMigration {return &%s{}}
func (m *%s) Name() string {return "%s"}
		
func (m *%s) Up(tx *gorm.DB) error{
	// Your gorm dml code goes here
	// Example: err := tx.Save(&datum).Error
	return nil
}
		
func (m *%s) Down(tx *gorm.DB) error{
	// Your gorm dml code goes here
	// Example: tx.Unscoped().Where(&datum).Delete(&datum).Error
	return nil
}
	`, structName, structName, structName, structName, structName, key, structName, structName,
	)
}

func createPkgTemplate(name string) string {
	return fmt.Sprintf(
		`package %s
`, name)
}
