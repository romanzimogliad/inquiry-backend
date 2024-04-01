package database

import "context"

func (d *Database) Ping(ctx context.Context) (string, error) {
	row := d.pool.QueryRow(ctx, "SELECT name FROM subject WHERE id=1")
	var res string
	err := row.Scan(&res)
	if err != nil {
		return "", err
	}
	return res, nil
}
