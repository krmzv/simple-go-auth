package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*User) error
	DeleteUser(int) error
	GetUsers() ([]*User, error)
	UpdateUser(*User) error
	GetUserByID(int) (*User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=reversejs sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createUserTable()
}

func (s *PostgresStore) createUserTable() error {
	query := `create table if not exists users (
			id serial primary key,
			name varchar(50),
			email varchar(50) not null unique,
			created_at timestamp
		);

		create table if not exists developer_profile (
			id serial primary key,
			developer_id integer not null,
			name varchar(50),
			skills varchar(255),
			location varchar(255),
			portfolio varchar(255),
			experience integer,
			created_at timestamp,
			foreign key (developer_id) references users (id)
		);

		create table if not exists company_profile (
			id serial primary key,
			company_id integer not null,
			name varchar(50),
			description varchar(255),
			location varchar(255),
			website varchar(255),
			created_at timestamp,
			foreign key (company_id) references users (id)
		);`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateUser(user *User) error {
	query := `insert into users
	(name, email, created_at)
	values ($1, $2, $3)`

	_, err := s.db.Query(query, user.Name, user.Email, user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateUser(*User) error {
	return nil
}

func (s *PostgresStore) DeleteUser(id int) error {
	_, err := s.db.Query("delete from users where id = $1", id)
	return err
}

func (s *PostgresStore) GetUserByID(id int) (*User, error) {
	rows, err := s.db.Query("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoUser(rows)
	}
	return nil, fmt.Errorf("User %d not found", id)
}

func (s *PostgresStore) GetUsers() ([]*User, error) {
	rows, err := s.db.Query("select * from users")
	if err != nil {
		return nil, err
	}

	users := []*User{}
	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		fmt.Println(users)
	}

	return users, nil
}

func scanIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	return user, err
}
