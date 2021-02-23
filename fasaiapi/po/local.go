package fasaiapi

import "database/sql"


// PO type structure data model from local database
type PO struct {
	name string
}


// ReadPOLocal get po data from local database
func ReadPOLocal(db *sql.DB, datetime string, isGenesis bool) (map[string]PO, error) {

	return map[string]PO{}, nil
}
