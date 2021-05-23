package repository

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
	"user/internal/dao"
	"user/internal/model"
)

type IUserRepository interface {

	FindUserByName(string) (*model.User, error)

	FindUserByID(int) (*model.User, error)

	CreateUser(*model.User) (int, error)

	//DeleteUserByID(int) error

	//UpdateUser(*model2.User) error

	//FindAll() ([]*model.User, error)
}

type UserRepository struct {
	Dao *dao.Dao
}

func NewUserRepository() IUserRepository {
	return &UserRepository{Dao: dao.NewDao()}
}


func (u *UserRepository) FindUserByID(id int) (user *model.User, err error) {
	querySQL, err := u.Dao.DB.Prepare("SELECT * FROM user WHERE id = ?")
	if err != nil {
		logrus.Errorf("Prepare Select SQL Error: %s", err.Error())
		return nil, err
	}

	defer querySQL.Close()
	v := new(model.User)
	err = querySQL.QueryRow(id).Scan(&v.ID, &v.UserName, &v.Email, &v.Telephone, &v.Age, &v.CreateTime)
	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query Video Error: %s", err.Error())
	}
	return v, nil
}

func (u *UserRepository) CreateUser (user *model.User) (int, error) {
	userName := user.UserName
	email := user.Email
	telephone := user.Telephone
	age := user.Age
	createTime := time.Now().Second()

	insertSql, err := u.Dao.DB.Prepare("INSERT INTO user (user_name, email, telephone, age, create_time) VALUES (?, ?, ?, ?)")
	if err != nil {
		logrus.Errorf("Prepare Insert SQL Error: %s", err.Error())
	}

	res, err := insertSql.Exec(userName, email, telephone, age, createTime)
	if err != nil {
		logrus.Errorf("Insert Video Info SQL Error: %s", err.Error())
		return 0, err
	}
	defer insertSql.Close()
	id, _ := res.LastInsertId()
	return int(id), nil
}

func (u *UserRepository) FindUserByName(userName string) (*model.User, error) {
	querySQL, err := u.Dao.DB.Prepare("SELECT * FROM user WHERE user_name = ?")
	if err != nil {
		logrus.Errorf("Prepare Select SQL Error: %s", err.Error())
		return nil, err
	}

	defer querySQL.Close()
	v := new(model.User)
	err = querySQL.QueryRow(userName).Scan(&v.ID, &v.UserName, &v.Email, &v.Telephone, &v.Age, &v.CreateTime, &v.HashPassword)
	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query Video Error: %s", err.Error())
	}
	return v, nil
}