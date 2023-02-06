package pagination

func GetLimitOffset(pageNo, pageSize int) (int, int) {
	offset := (pageNo - 1) * pageSize
	limit := pageSize

	return limit, offset
}
