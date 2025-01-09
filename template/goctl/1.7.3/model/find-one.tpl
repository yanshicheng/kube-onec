func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query :=  fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} AND `is_deleted` = 0 limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		return conn.QueryRowCtx(ctx, v, query, {{.lowerStartCamelPrimaryKey}})
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}query := fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} AND `is_deleted` = 0 limit 1", {{.lowerStartCamelObject}}Rows, m.table)
	var resp {{.upperStartCamelObject}}
	err := m.conn.QueryRowCtx(ctx, &resp, query, {{.lowerStartCamelPrimaryKey}})
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}

func (m *default{{.upperStartCamelObject}}Model) Search(ctx context.Context,  orderStr string, isAsc bool, page, pageSize uint64, queryStr string, args ...any) ([]*{{.upperStartCamelObject}}, uint64, error) {
	// 确保分页参数有效
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	// 构造查询条件
	// 添加 `is_deleted` = 0 条件，保证只查询未软删除数据
    // 初始化 WHERE 子句
    where := "WHERE `is_deleted` = 0"
    if queryStr != "" {
        where = fmt.Sprintf("WHERE %s AND `is_deleted` = 0", queryStr)
    }

	// 根据 isAsc 参数确定排序方式
	sortDirection := "ASC"
	if !isAsc {
		sortDirection = "DESC"
	}

	// 如果用户未指定排序字段，则默认使用 id
	if orderStr == "" {
		orderStr = fmt.Sprintf("ORDER BY id %s", sortDirection)
	} else {
		orderStr = strings.TrimSpace(orderStr)
		if !strings.HasPrefix(strings.ToUpper(orderStr), "ORDER BY") {
			orderStr = "ORDER BY " + orderStr
		}
		orderStr = fmt.Sprintf("%s %s", orderStr, sortDirection)
	}

    countQuery := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", m.table, where)

	var total uint64
    var resp []*{{.upperStartCamelObject}}
	err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		// 无匹配记录
		return resp, 0, ErrNotFound
	}
	offset := (page - 1) * pageSize
	dataQuery := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT %d,%d", {{.lowerStartCamelObject}}Rows, m.table, where, orderStr, offset, pageSize)
    
	err = m.QueryRowsNoCacheCtx(ctx, &resp, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (m *default{{.upperStartCamelObject}}Model) SearchNoPage(ctx context.Context, orderStr string, isAsc bool,  queryStr string, args ...any) ([]*{{.upperStartCamelObject}}, error) {
    // 初始化 WHERE 子句
    where := "WHERE `is_deleted` = 0"
    if queryStr != "" {
        where = fmt.Sprintf("WHERE %s AND `is_deleted` = 0", queryStr)
    }

	// 根据 isAsc 参数确定排序方式
	sortDirection := "ASC"
	if !isAsc {
		sortDirection = "DESC"
	}
	// 如果用户未指定排序字段，则默认使用 id
	if orderStr == "" {
		orderStr = fmt.Sprintf("ORDER BY id %s", sortDirection)
	} else {
		orderStr = strings.TrimSpace(orderStr)
		if !strings.HasPrefix(strings.ToUpper(orderStr), "ORDER BY") {
			orderStr = "ORDER BY " + orderStr
		}
		orderStr = fmt.Sprintf("%s %s", orderStr, sortDirection)
	}  
	dataQuery := fmt.Sprintf("SELECT %s FROM %s %s %s", {{.lowerStartCamelObject}}Rows, m.table, where, orderStr)
    var resp []*{{.upperStartCamelObject}}
    err := m.QueryRowsNoCacheCtx(ctx, &resp, dataQuery, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}