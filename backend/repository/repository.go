package repository

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Repository struct {
	Config Config
	DB     *gorm.DB
}

type Config struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Dbname   string `mapstructure:"DB_NAME"`
	Schema   string `mapstructure:"DB_SCHEMA"`
}

var once sync.Once
var repository *Repository

func (r Repository) GetDB() *gorm.DB {
	return r.DB
}

func GetInstance() *Repository {
	once.Do(func() {
		repository = &Repository{}
		db, err := connectToDB("PG")
		if err != nil {
			fmt.Println(err)
		}
		repository.DB = db
	})
	return repository
}

// connectGenerator handle for multiple database type in test just use postgres
func connectGenerator(typeDataBase string) string {
	switch typeDataBase {
	case "PG":
		return fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable`,
			repository.Config.Host, repository.Config.User, repository.Config.Password,
			repository.Config.Dbname, repository.Config.Port)
	default:
		return ""
	}
}

// ConnectToDB connect to db with name. Example: "PG"
func connectToDB(typeDatabase string) (*gorm.DB, error) {
	err := setConfig()
	if err != nil {
		return nil, err
	}
	// if err = repository.migrateDB(); err != nil {
	// 	return nil, err
	// }
	// log.Printf("Applied migrations!")
	db, err := gorm.Open(postgres.Open(connectGenerator(typeDatabase)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: repository.Config.Schema + ".",
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ConnectToDB connect to db with name. Example: "PG"
func (repo *Repository) ConnectToDBWithConfig(typeDatabase string, config *Config) (*gorm.DB, error) {
	repo.Config = *config
	db, err := gorm.Open(postgres.Open(connectGenerator(typeDatabase)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: repo.Config.Schema + ".",
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// SetConfig CreateConfig load config in env with pg
func setConfig() error {
	config, err := loadConfig() // for simple make just one file .env for postgres
	if err != nil {
		return err
	}
	repository.Config = *config
	return nil
}

// load config from .env to app
func loadConfig() (*Config, error) {
	config := &Config{}
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// init data base in docker
func (repo *Repository) migrateDB() error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	db, err := sql.Open("postgres", connectGenerator("PG"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
