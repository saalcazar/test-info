package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerQuantitative interface {
	Scan(dest ...any) error
}

const (
	psqlCreateQuantitative  = `SELECT create_quantitative ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	psqlUpdateQuantitative  = `SELECT update_quantitative ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	psqlDeleteQuantitative  = `SELECT delete_quantitative ($1)`
	psqlDeleteQuantitatives = `SELECT delete_quantitatives ($1)`
	psqlGetQuantitativeAll  = `SELECT id_quantitative, achieved, day, sp_female, sp_male, f_female, f_male, na_female, na_male, p_female, p_male, id_activity, nick_user FROM quantitatives`
)

type PsqlQuantitative struct {
	db *sql.DB
}

func NewPsqlQuantitative(db *sql.DB) *PsqlQuantitative {
	return &PsqlQuantitative{db}
}

// Create
func (f *PsqlQuantitative) Create(m model.Quantitatives) error {

	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	stmt, err := f.db.Prepare(psqlCreateQuantitative)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, q := range m {
		err = stmt.QueryRow(
			q.Achieved,
			q.Day,
			q.SpFemale,
			q.SpMale,
			q.FFemale,
			q.FMale,
			q.NaFemale,
			q.NaMale,
			q.PFemale,
			q.PMale,
			q.IdActivity,
			q.NickUser,
		).Scan(&q.Day)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("Los datos cuantitativos fueron creados exitosamente")
	return nil
}

// Update
func (f *PsqlQuantitative) Update(m *model.Quantitative) error {
	stmt, err := f.db.Prepare(psqlUpdateQuantitative)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Achieved,
		&m.Day,
		&m.SpFemale,
		&m.SpMale,
		&m.FFemale,
		&m.FMale,
		&m.NaFemale,
		&m.NaMale,
		&m.PFemale,
		&m.PMale,
		&m.IdActivity,
		&m.NickUser,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe los datos cuantitativos con el id: %d", m.ID)
	}

	fmt.Println("Los datos cuantitativos se actualizaron exitosamente")
	return nil
}

// Delete
func (f *PsqlQuantitative) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteQuantitative)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El dato cuantitativo se eliminó correctamente")
	return nil
}

func (fa *PsqlQuantitative) DeleteAll(idActivity uint) error {
	stmt, err := fa.db.Prepare(psqlDeleteQuantitatives)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(idActivity)
	if err != nil {
		return err
	}

	fmt.Println("Los datos cuantitativos se eliminaron correctamente")
	return nil
}

func (p *PsqlQuantitative) GetAll() (model.Quantitatives, error) {
	stmt, err := p.db.Prepare(psqlGetQuantitativeAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Quantitatives, 0)
	for rows.Next() {
		p, err := scanRowQuantitative(rows)
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

func scanRowQuantitative(s scannerQuantitative) (*model.Quantitative, error) {
	m := &model.Quantitative{}
	err := s.Scan(
		&m.ID,
		&m.Achieved,
		&m.Day,
		&m.SpFemale,
		&m.SpMale,
		&m.FFemale,
		&m.FMale,
		&m.NaFemale,
		&m.NaMale,
		&m.PFemale,
		&m.PMale,
		&m.IdActivity,
		&m.NickUser,
	)
	if err != nil {
		return &model.Quantitative{}, err
	}
	return m, nil
}
