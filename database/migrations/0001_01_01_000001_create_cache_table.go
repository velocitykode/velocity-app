package migrations

import "github.com/velocitykode/velocity/pkg/orm/migrate"

func init() {
	migrate.Register(&migrate.Migration{
		Version:     "20010101000001",
		Description: "create cache table",
		Up: func(m *migrate.Migrator) error {
			return m.CreateTable("cache", func(t *migrate.TableBuilder) {
				t.String("key", 255).Unique()
				t.String("value", 10000)
				t.Integer("expiration")
			})
		},
		Down: func(m *migrate.Migrator) error {
			return m.DropTable("cache")
		},
	})
}
