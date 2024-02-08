package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerSurrenders interface {
	Scan(dest ...any) error
}

const (
	psqlCreateSurrender  = `SELECT create_surrender ($1, $2, $3, $4, $5, $6, $7, $8)`
	psqlUpdateSurrender  = `SELECT update_surrender ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	psqlDeleteSurrender  = `SELECT delete_surrender ($1)`
	psqlDeleteSurrenders = `SELECT delete_surrenders ($1)`
	psqlGetAllSurrenders = `SELECT id_surrender, date_invoice, invoice_number, code, description, inport_usd, inport_bob, id_activity, nick_user FROM surrenders`
)

type PsqlSurrender struct {
	db *sql.DB
}

func NewPsqlSurrender(db *sql.DB) *PsqlSurrender {
	return &PsqlSurrender{db}
}

// Create
func (f *PsqlSurrender) Create(m model.Surrenders) error {
	tx, err := f.db.Begin()
	if err != nil {
		tx.Rollback()
	}

	stmt, err := f.db.Prepare(psqlCreateSurrender)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, sr := range m {
		err = stmt.QueryRow(
			sr.DateInvoice,
			sr.InvoiceNumber,
			sr.Code,
			sr.Description,
			sr.ImportUSD,
			sr.ImportBOB,
			sr.IdActivity,
			sr.NickUser,
		).Scan(&sr.DateInvoice)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("Los datos de la rendición fueron creados exitosamente")
	return nil
}

// Update
func (f *PsqlSurrender) Update(m *model.Surrender) error {
	stmt, err := f.db.Prepare(psqlUpdateSurrender)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.DateInvoice,
		&m.InvoiceNumber,
		&m.Code,
		&m.Description,
		&m.ImportUSD,
		&m.ImportBOB,
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
		return fmt.Errorf("no existe los datos de la rendición con el id: %d", m.ID)
	}

	fmt.Println("Los datos de la rendición se actualizaron exitosamente")
	return nil
}

// Delete
func (f *PsqlSurrender) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteSurrender)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Los datos de la rendición se eliminaron correctamente")
	return nil
}

func (fa *PsqlSurrender) DeleteAll(idActivity uint) error {
	stmt, err := fa.db.Prepare(psqlDeleteSurrenders)
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

func (p *PsqlSurrender) GetAll() (model.Surrenders, error) {
	stmt, err := p.db.Prepare(psqlGetAllSurrenders)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Surrenders, 0)
	for rows.Next() {
		p, err := scanRowSurrenders(rows)
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

func scanRowSurrenders(s scannerSurrenders) (*model.Surrender, error) {
	m := &model.Surrender{}
	err := s.Scan(
		&m.ID,
		&m.DateInvoice,
		&m.InvoiceNumber,
		&m.Code,
		&m.Description,
		&m.ImportUSD,
		&m.ImportBOB,
		&m.IdActivity,
		&m.NickUser,
	)
	if err != nil {
		return &model.Surrender{}, err
	}
	return m, nil
}
