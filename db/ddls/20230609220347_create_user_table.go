package ddls

import (
	"github.com/m4hi2/capsule71/internal/db/models"
	"github.com/m4hi2/capsule71/pkg/migrator"
	"github.com/m4hi2/capsule71/pkg/migrator/ddl"
	"gorm.io/gorm"
)

func init() { ddl.Register(NewCreateUserTable20230609220347) }

type CreateUserTable20230609220347 struct{}

func NewCreateUserTable20230609220347() migrator.IMigration { return &CreateUserTable20230609220347{} }
func (m *CreateUserTable20230609220347) Name() string       { return "20230609220347_create_user_table" }

// Up Called when "migrate ddl up" command
func (m *CreateUserTable20230609220347) Up(tx *gorm.DB) error {
	// Your gorm migration code goes here
	// Example: tx.Migrator().CreateTable(&models.User{})
	if tx.Migrator().HasTable(models.User{}) == false {
		err := tx.Migrator().CreateTable(models.User{})
		if err != nil {
			return err
		}
	}
	return nil
}

// Down Called when "migrate ddl down" command
func (m *CreateUserTable20230609220347) Down(tx *gorm.DB) error {
	// Your gorm migration code goes here
	// Example: tx.Migrator().DropTable(&models.User{})
	if tx.Migrator().HasTable(models.User{}) == true {
		err := tx.Migrator().DropTable(models.User{})
		if err != nil {
			return err
		}
	}
	return nil
}
