package migrations

import (
	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/stuffbin"
)

// V2_7_0 is retained as a no-op so environments that already applied this
// version stay compatible. Sidebar counts ship without a settings toggle.
func V2_7_0(db *sqlx.DB, fs stuffbin.FileSystem, ko *koanf.Koanf) error {
	_, err := db.Exec(`DELETE FROM settings WHERE "key" = 'app.sidebar_counts_enabled'`)
	return err
}
