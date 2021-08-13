package moduleapi

type DatabaseQueryRequest struct {
	Operation  string
	SQL        string
	Parameters [][]interface{}
}

type DatabaseQueryResult struct {
	Results []struct {
		ColumnName []string
		Rows       [][]interface{}
	}
}
