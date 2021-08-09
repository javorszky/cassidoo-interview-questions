package august082021

// FieldType is a new int type to use as enums in the field. See https://yourbasic.org/golang/iota/ for why this was
// chosen.
type FieldType int

// String renders the string representations of the FieldType enums.
func (f FieldType) String() string {
	return [...]string{"x", "_", "1", "2", "3", "4", "5", "6", "7", "8"}[f]
}
