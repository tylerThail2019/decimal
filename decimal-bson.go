package decimal

import (
	"gopkg.in/mgo.v2/bson"
)

// GetBSON implements the bson.Getter interface
func (d Decimal) GetBSON() (interface{}, error) {
	// Pass through string to create Mongo Decimal128 type
	dec128, err := bson.ParseDecimal128(d.String())
	if err != nil {
		return nil, err
	}
	return dec128, nil
}

// SetBSON implements the bson.Setter interface
func (d *Decimal) SetBSON(raw bson.Raw) error {
	// Unmarshal as Mongo Decimal128 first then pass through string to obtain Decimal
	var dec128 bson.Decimal128
	berr := raw.Unmarshal(&dec128)
	if berr != nil {
		return berr
	}
	dec, derr := NewFromString(dec128.String())
	if derr != nil {
		return derr
	}
	*d = dec
	return nil
}
