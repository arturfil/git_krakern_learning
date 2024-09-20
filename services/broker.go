package services

import "database/sql"

type Service struct {
    db *sql.DB
}

func (s *Service) GetAll() {
    s.db.Exec("SELECT * FROM requests;");
}
