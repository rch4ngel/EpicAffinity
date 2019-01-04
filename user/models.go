package user

import (
	"github.com/pkg/errors"
	"github.com/ryu/epic_affinity/config"
	"net/http"
)

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Username  string
	Password  string
	Role      string
}

func AllUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	xu := []User{}
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Username, &u.Password, &u.Role)
		if err != nil {
			return nil, err
		}
		xu = append(xu, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return xu, nil
}

func OneUser(r *http.Request) (User, error) {
	u := User{}
	un := r.FormValue("username")
	if un == "" {
		return u, errors.New("400. Bad Request")
	}
	row := config.DB.QueryRow("SELECT * FROM users where username = $1", un)
	err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Username, &u.Password, &u.Role)
	if err != nil {
		return u, err
	}
	return u, nil
}

func PutUser(r *http.Request) (User, error) {
	u := User{}
	u.Firstname = r.FormValue("firstname")
	u.Lastname = r.FormValue("lastname")
	u.Username = r.FormValue("username")
	u.Password = r.FormValue("password")
	u.Role = r.FormValue("role")

	if u.Firstname == "" || u.Lastname == "" || u.Username == "" || u.Password == "" || u.Role == "" {
		return u, errors.New("400. Bad Request")
	}
	_, err := config.DB.Exec("UPDATE users SET firstname = $1, lastname = $2, username = $3, password = $4, role= $5 WHERE username = $3",
		u.Firstname,
		u.Lastname,
		u.Username,
		u.Password,
		u.Role,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

func CreateUser(r *http.Request) (User, error) {
	u := User{}
	u.Firstname = r.FormValue("firstname")
	u.Lastname = r.FormValue("lastname")
	u.Username = r.FormValue("username")
	u.Password = r.FormValue("password")
	u.Role = r.FormValue("role")

	if u.Firstname == "" || u.Lastname == "" || u.Username == "" || u.Password == "" || u.Role == "" {
		return u, errors.New("400. Bad Request")
	}

	_, err := config.DB.Exec("INSERT INTO users (firstname, lastname, username, password, role) VALUES ($1, $2, $3, $4, $5)",
		u.Firstname,
		u.Lastname,
		u.Username,
		u.Password,
		u.Role,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}
