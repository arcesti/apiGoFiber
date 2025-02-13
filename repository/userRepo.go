package repository

import (
	"testeGO/config"
	"testeGO/models"
	"log"
)

func GetUsers() ([]models.User, error) {
	rows, err := config.DB.Query("SELECT * FROM usuarios");
	if err!= nil {
		return nil, err;
	}
	defer rows.Close();

	var users []models.User;
	for rows.Next() {
		var user models.User;
		if err := rows.Scan(&user.Cpf, &user.Nome); err!=nil {
			return nil, err;
		}
		users = append(users, user);
	}
	return users, nil;
}

func AddUser(user models.User) error {
	_, err := config.DB.Exec("INSERT INTO usuarios (nome, cpf) VALUES ($1, $2)", user.Nome, user.Cpf);
	log.Println(user.Nome, user.Cpf);
	if err != nil {
		return err;
	}
	log.Println("Usuario adicionado com sucesso!");
	return nil;
}

func AlterUser(user models.User) error {
	_, err := config.DB.Exec("UPDATE usuarios SET nome = $1 WHERE cpf = $2", user.Nome, user.Cpf);
	if err != nil {
		return err;
	}
	log.Println("Usuario alterado com sucesso!");
	return nil;
}

func DeleteUser(cpf string) error {
	_, err := config.DB.Exec("DELETE FROM usuarios WHERE cpf = $1", cpf);
	if err != nil {
		log.Fatal("Erro: ", err);
		return err;
	}
	log.Println("Usuario excluido com suesso!");
	return nil;
}