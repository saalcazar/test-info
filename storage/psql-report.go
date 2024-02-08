package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerReport interface {
	Scan(dest ...any) error
}

const (
	psqlCreateReport  = `SELECT create_report ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	psqlUpdateReport  = `SELECT update_report ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	psqlDeleteReport  = `SELECT delete_report ($1)`
	psqlGetReportByID = `SELECT id_report, issues, results, obstacle, conclusions, anexos, approved, name_user, name_proyect, signature, cod_founder, nick_user, id_activity FROM reports WHERE id_report = $1`
	psqlGetAllReport  = `SELECT id_report, issues, results, obstacle, conclusions, anexos, approved, name_user, name_proyect, signature, cod_founder, nick_user, id_activity FROM reports`
)

type PsqlReport struct {
	db *sql.DB
}

func NewPsqlReport(db *sql.DB) *PsqlReport {
	return &PsqlReport{db}
}

// Create
func (p *PsqlReport) Create(m *model.Report) error {
	stmt, err := p.db.Prepare(psqlCreateReport)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Issues,
		m.Results,
		m.Obstacle,
		m.Conclusions,
		m.Anexos,
		m.Approved,
		m.NameUser,
		m.NameProyect,
		m.Signature,
		m.CodFounder,
		m.NickUser,
		m.IdActivity,
	).Scan(&m.Results)
	if err != nil {
		return err
	}
	fmt.Println("El informe se creo correctamente")
	return nil
}

// Update
func (p *PsqlReport) Update(m *model.Report) error {
	stmt, err := p.db.Prepare(psqlUpdateReport)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Issues,
		&m.Results,
		&m.Obstacle,
		&m.Conclusions,
		&m.Anexos,
		&m.Approved,
		&m.NameUser,
		&m.NameProyect,
		&m.Signature,
		&m.CodFounder,
		&m.NickUser,
		&m.IdActivity,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe el inform con id: %d", m.ID)
	}

	fmt.Println("Se actualizó el informe correctamente")
	return nil
}

// Delete
func (p *PsqlReport) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteReport)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Se elimino el informe correctamente")
	return nil
}

// GetByID
func (p *PsqlReport) GetByID(id uint) (*model.Report, error) {
	stmt, err := p.db.Prepare(psqlGetReportByID)
	if err != nil {
		return &model.Report{}, err
	}
	defer stmt.Close()
	return scanRowReports(stmt.QueryRow(id))
}

func (p *PsqlReport) GetAll() (model.Reports, error) {
	stmt, err := p.db.Prepare(psqlGetAllReport)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Reports, 0)
	for rows.Next() {
		p, err := scanRowReports(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, *p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func scanRowReports(s scannerReport) (*model.Report, error) {
	m := &model.Report{}
	err := s.Scan(
		&m.ID,
		&m.Issues,
		&m.Results,
		&m.Obstacle,
		&m.Conclusions,
		&m.Anexos,
		&m.Approved,
		&m.NameUser,
		&m.NameProyect,
		&m.Signature,
		&m.CodFounder,
		&m.NickUser,
		&m.IdActivity,
	)
	if err != nil {
		return &model.Report{}, err
	}
	return m, nil
}
