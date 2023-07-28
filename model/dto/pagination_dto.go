package dto

//ini buat paging di taruh di paramaeter
type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

//ini buat paging di taruh di return
type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

//ini buat di taruh di response
type Paging struct {
	Page        int
	RowsPerPage int
	TotalRows   int
	TotalPages  int
}

//example pagination -> product 100
//paging {page: 1, RowsPerPage: 10, TotalRows: 100, TotalPages: 10}
