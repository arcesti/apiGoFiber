package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
);

var DB *sql.DB;

func ConnectDb() {
    var err error;
    connStr := "user=postgres password=postgres123 dbname=golang sslmode=disable";
    DB, err = sql.Open("postgres", connStr);
    if err != nil {
        log.Fatal("Erro ao conectar com o banco de dados. Erro: ", err);
    }

    if err = DB.Ping(); err!=nil {
        log.Fatal("Erro ao verificar conexão: ", err);
    }

    log.Println("Banco de dados conectado!");
}

func CloseDb() {
    if DB != nil {
        if err := DB.Close(); err!=nil {
            log.Println("Erro ao fechar a conexão com o banco: ", err);
        } else {
            log.Println("Conexão com o banco fechada!");
        }
    }
}