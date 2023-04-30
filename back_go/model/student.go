package model

import (
	"database/sql"
	"fmt"

	"back_go/config"
)

type Student struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

var DefaultStudentDb *StudentDb

func NewStudentDb() *StudentDb {
	return &StudentDb{TbName: "student", Db: config.Db}
}

type StudentDb struct {
	TbName string
	Db     *sql.DB
}

func (s *StudentDb) Save(student Student) error {
	stmt, err := s.Db.Prepare(fmt.Sprintf("INSERT INTO %s(name,age) VALUES(?,?)", s.TbName))
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(student.Name, student.Age); err != nil {
		return err
	}
	return nil
}

func (s *StudentDb) GetAll() ([]Student, error) {
	students := make([]Student, 0)
	rows, err := s.Db.Query(fmt.Sprintf("SELECT id,name,age from %s", s.TbName))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Id, &student.Name, &student.Age); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *StudentDb) GetOne(id int32) (Student, error) {
	var student Student
	rows, err := s.Db.Query(fmt.Sprintf("SELECT id,name,age from %s WHERE id=%d", s.TbName, id))
	if err != nil {
		return student, err
	}
	if rows.Next() {
		if err := rows.Scan(&student.Id, &student.Name, &student.Age); err != nil {
			return student, err
		}
	}
	return student, nil
}

func (s *StudentDb) Remove(id int32) error {
	stmt, err := s.Db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE id=%d", s.TbName, id))
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

func (s *StudentDb) Update(student Student) error {
	stmt, err := s.Db.Prepare(fmt.Sprintf("UPDATE %s SET name=?,age=? WHERE id=?", s.TbName))
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(student.Name, student.Age, student.Id); err != nil {
		return err
	}
	return nil
}
