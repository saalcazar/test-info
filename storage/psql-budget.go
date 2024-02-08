package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerBudget interface {
	Scan(dest ...any) error
}

const (
	psqlCreateBudget  = `SELECT create_budget ($1, $2, $3, $4, $5, $6, $7, $8)`
	psqlUpdateBudget  = `SELECT update_budget ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	psqlDeleteBudget  = `SELECT delete_budget ($1)`
	psqlDeleteBudgets = `SELECT delete_budgets ($1)`
	psqlGetAllBudgets = `SELECT id_budget, quantity, code, description, import_usd, import_bob, id_activity, cod_founder, nick_user FROM budgets`
)

type PsqlBudget struct {
	db *sql.DB
}

func NewPsqlBudget(db *sql.DB) *PsqlBudget {
	return &PsqlBudget{db}
}

// Create
func (f *PsqlBudget) Create(m model.Budgets) error {

	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	stmt, err := f.db.Prepare(psqlCreateBudget)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, bd := range m {
		err = stmt.QueryRow(
			bd.Quantity,
			bd.Code,
			bd.Description,
			bd.ImportUSD,
			bd.ImportBOB,
			bd.IdActivity,
			bd.CodFounder,
			bd.NickUser,
		).Scan(&bd.Code)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("El item de la solicitud fue creado exitosamente")
	return nil
}

// Update
func (f *PsqlBudget) Update(m *model.Budget) error {
	stmt, err := f.db.Prepare(psqlUpdateBudget)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Quantity,
		&m.Code,
		&m.Description,
		&m.ImportUSD,
		&m.ImportBOB,
		&m.IdActivity,
		&m.CodFounder,
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
		return fmt.Errorf("no existe el item de la solicitud con el id: %d", m.ID)
	}

	fmt.Println("El item de la solicitud se actualizó exitosamente")
	return nil
}

// Delete
func (f *PsqlBudget) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteBudget)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El item de la solicitud se eliminó correctamente")
	return nil
}

func (fa *PsqlBudget) DeleteAll(idActivity uint) error {
	stmt, err := fa.db.Prepare(psqlDeleteBudgets)
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

func (p *PsqlBudget) GetAll() (model.Budgets, error) {
	stmt, err := p.db.Prepare(psqlGetAllBudgets)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Budgets, 0)
	for rows.Next() {
		p, err := scanRowBudgets(rows)
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

func scanRowBudgets(s scannerBudget) (*model.Budget, error) {
	m := &model.Budget{}
	err := s.Scan(
		&m.ID,
		&m.Quantity,
		&m.Code,
		&m.Description,
		&m.ImportUSD,
		&m.ImportBOB,
		&m.IdActivity,
		&m.CodFounder,
		&m.NickUser,
	)
	if err != nil {
		return &model.Budget{}, err
	}
	return m, nil
}
