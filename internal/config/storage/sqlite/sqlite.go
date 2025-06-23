package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/gaurav25kapoor/students-api/internal/config"
	"github.com/gaurav25kapoor/students-api/internal/config/types"
	_ "github.com/mattn/go-sqlite3"
)

// Sqlite wraps the database connection
type Sqlite struct {
	Db *sql.DB
}

// New initializes the SQLite database connection and creates the students table
func New(cfg *config.Config) (*Sqlite, error) {
	// Use the correct driver name: "sqlite3"
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	// Create the table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)
	if err != nil {
		return nil, err
	}

	// Return the wrapped db connection
	return &Sqlite{
		Db: db,
	}, nil

}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stat, err := s.Db.Prepare("INSERT INTO students (name,email,age) VALUES(?,?,?)")

	if err != nil {
		return 0, err
	}

	defer stat.Close()

	result, err := stat.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil

}

func (s *Sqlite) GetStudentById(id int64) (types.Student,error){
	stmt,err:=s.Db.Prepare("SELECT * FROM students WHERE id=? LIMIT 1")
	if err!=nil{
		return types.Student{},err
	}
	defer stmt.Close()

	var student types.Student
	err=stmt.QueryRow(id).Scan(&student.Id,&student.Name,&student.Email,&student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}


func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}
