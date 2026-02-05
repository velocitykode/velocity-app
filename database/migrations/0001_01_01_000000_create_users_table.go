package migrations

import "github.com/velocitykode/velocity/pkg/orm/migrate"

func init() {
	migrate.Register(&migrate.Migration{
		Version:     "20010101000000",
		Description: "create users table",
		Up: func(m *migrate.Migrator) error {
			return m.CreateTable("users", func(t *migrate.TableBuilder) {
				t.ID()
				t.String("name")
				t.String("email").Unique()
				t.String("password")
				t.String("role").Default("user")
				t.Timestamps()
			})
		},
		Down: func(m *migrate.Migrator) error {
			return m.DropTable("users")
		},
	})
}
