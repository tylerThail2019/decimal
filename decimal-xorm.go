package decimal

func (d *Decimal) FromDB(data []byte) error {
	dd, err := NewFromString(string(data))
	*d = dd
	return err
}

func (d Decimal) ToDB() ([]byte, error) {
	return []byte(d.String()), nil
}
