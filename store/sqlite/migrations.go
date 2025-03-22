package sqlite

import (
	"cmp"
	"embed"
	"log"
	"path"
	"slices"
	"strconv"
)

type migration struct {
    version int
    statement string
}

//go:embed migrations/*.sql
var migrationsFS embed.FS

func (s Store) runMigrations() {
    migrations, err := loadMigrations()

    if err != nil {
        log.Fatalf("Unable to load database migrations: %v", err)
    }

    for _, migration := range migrations {

    }

    // Loop through migrations and apply the versions greater than current user_version
}

func loadMigrations() ([]migration, error) {
    migrations := []migration{}

    migrationsDir := "migrations"

    migrationEntries, err := migrationsFS.ReadDir(migrationsDir)

    if err != nil {
        return []migration{}, err
    }

    for _, entry := range migrationEntries {
        if entry.IsDir() {
            continue
        }

        version := parseVersionFromFilename(entry.Name())

        statment, err := migrationsFS.ReadFile(path.Join(migrationsDir, entry.Name()))

        migrations = append(migrations, migration{version, string(statment)})

        if err != nil {
            return []migration{}, err
        }
    }

    slices.SortFunc(migrations,
        func(a, b migration) int {
            return cmp.Compare(a.version, b.version)
    })

    return migrations, nil
}

func parseVersionFromFilename(filename string) int {
    version, err := strconv.ParseInt(filename[:3], 10, 32)

    if err != nil {
        log.Fatalf("Unable to parse migration version from filename: %v", err)
    }

    return int(version)
}
