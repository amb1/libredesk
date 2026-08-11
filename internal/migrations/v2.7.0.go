package migrations

import (
	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/stuffbin"
)

func V2_7_0(db *sqlx.DB, fs stuffbin.FileSystem, ko *koanf.Koanf) error {
	if _, err := db.Exec(`INSERT INTO settings ("key", value) VALUES ('app.sidebar_counts_enabled', 'true'::jsonb) ON CONFLICT ("key") DO NOTHING;`); err != nil {
		return err
	}
	return nil
}
