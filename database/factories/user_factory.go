package factories

import (
	"github.com/velocitykode/velocity/orm"
	ormtesting "github.com/velocitykode/velocity/orm/testing"
)

// UserFactory creates test users
func UserFactory(manager *orm.Manager) *ormtesting.Factory {
	faker := ormtesting.Faker()

	factory := ormtesting.NewFactory(manager, "users", func() map[string]interface{} {
		return map[string]interface{}{
			"name":  faker.Name(),
			"email": faker.Email(),
			"role":  "user",
		}
	})

	// Define admin state
	factory.DefineState("admin", map[string]interface{}{
		"role": "admin",
	})

	return factory
}
