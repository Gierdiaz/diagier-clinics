package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/Gierdiaz/diagier-clinics/config"
)

func InitDatabase(config *config.Config) (*sqlx.DB, error) {
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.Database.DB_HOST, config.Database.DB_PORT, config.Database.DB_USER, config.Database.DB_NAME, config.Database.DB_PASSWORD)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao pingar o banco de dados: %w", err)
	}
	log.Println("Conexão com PostgreSQL estabelecida com sucesso")
	return db, nil
}