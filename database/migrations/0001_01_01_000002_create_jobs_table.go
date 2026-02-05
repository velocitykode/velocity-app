package migrations

import "github.com/velocitykode/velocity/pkg/orm/migrate"

func init() {
	migrate.Register(&migrate.Migration{
		Version:     "20010101000002",
		Description: "create jobs table",
		Up: func(m *migrate.Migrator) error {
			if err := m.CreateTable("jobs", func(t *migrate.TableBuilder) {
				t.ID()
				t.String("queue", 255)
				t.String("payload", 10000)
				t.Integer("attempts").Default("0")
				t.String("scheduled_at", 50)
				t.String("reserved_at", 50).Nullable()
				t.String("reserved_by", 255).Nullable()
				t.String("failed_at", 50).Nullable()
				t.String("failed_reason", 5000).Nullable()
				t.Timestamps()
			}); err != nil {
				return err
			}

			return m.CreateTable("failed_jobs", func(t *migrate.TableBuilder) {
				t.ID()
				t.String("queue", 255)
				t.String("payload", 10000)
				t.String("exception", 10000)
				t.Timestamps()
			})
		},
		Down: func(m *migrate.Migrator) error {
			if err := m.DropTable("failed_jobs"); err != nil {
				return err
			}
			return m.DropTable("jobs")
		},
	})
}
