package dbmodel

import (
	"database/sql"
	"time"
)

type Blog struct {
	Id int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) Insert(title string,content string) (int,error){
	var res int
	stmt := `insert into blogs(title,content,created,expires) values ($1,$2,NOW() , NOW() + INTERVAL '365 days') returning id;`
	err := m.DB.QueryRow(stmt, title, content).Scan(&res)

	if err != nil {
		return 0, err
	}

	return res,nil
}

func (m *BlogModel) Get(id int) (*Blog,error) {
	stmt := `select id,title,content,created,expires from blogs where expires > NOW() and id=$1`

	row := m.DB.QueryRow(stmt,id)

	s := &Blog{}

	err := row.Scan(&s.Id,&s.Title,&s.Content,&s.Created,&s.Expires)

	if err != nil {
		return nil, err
	}

	return s,nil
}

func (m *BlogModel) Latest() ([]*Blog, error) {
	stmt := `select id,title,content,created from blogs where expires > NOW() order by id desc limit 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	blogs := []*Blog{}

	for rows.Next() {
		s := &Blog{}
		err = rows.Scan(&s.Id,&s.Title, &s.Content, &s.Created)
		if err != nil {
			return nil,err
		}

		blogs = append(blogs, s)
	}

	if err = rows.Err(); err != nil {
		return nil,err
	}

	return blogs,nil
}
