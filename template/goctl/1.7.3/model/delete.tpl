func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		_,err:=m.conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}}){{end}}
	return err
}

func (m *default{{.upperStartCamelObject}}Model) DeleteSoft(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}
	// 如果记录已软删除，无需再次删除
	if data.IsDeleted == 1 {
		return nil
	}
{{end}}	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `is_deleted` = 1 where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set `is_deleted` = 1 where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		_,err:=m.conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}}){{end}}
	return err
}




func (m *default{{.upperStartCamelObject}}Model) TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *default{{.upperStartCamelObject}}Model) TransOnSql(ctx context.Context, session sqlx.Session,  {{.lowerStartCamelPrimaryKey}} {{.dataType}}, sqlStr string, args ...any) (sql.Result, error) {
    query := strings.ReplaceAll(sqlStr, "{table}", m.table)
    // 如果 id != 0 并且启用了缓存逻辑
    if !isZeroValue({{.lowerStartCamelPrimaryKey}}) {
        // 查询数据（如果需要，确保数据存在）
        {{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
            if err!=nil{
                return nil,err
            }
        {{end}}

        // 缓存相关处理
        {{if .withCache}}
            {{.keys}}  // 处理缓存逻辑，例如删除或更新缓存
            // 执行带缓存处理的 SQL 操作
            return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
                return session.ExecCtx(ctx, query, args...)
            }, {{.keyValues}}) // 传递缓存相关的键值
        {{else}}
            // 不使用缓存，直接执行 SQL 操作
            return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
                return session.ExecCtx(ctx, query, args...)
            })
        {{end}}
    }

    // 如果 id == 0 或不需要缓存，直接执行 SQL
    return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        return session.ExecCtx(ctx, query, args...)
    })
}

func (m *default{{.upperStartCamelObject}}Model)  ExecSql(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}, sqlStr string, args ...any) (sql.Result, error) {
        // 如果 id != 0 并且启用了缓存逻辑
    query := strings.ReplaceAll(sqlStr, "{table}", m.table)
    if !isZeroValue({{.lowerStartCamelPrimaryKey}}) {
        // 缓存相关处理
        {{if .withCache}}
                // 查询数据（如果需要，确保数据存在）
            {{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
                if err!=nil{
                    return nil,err
                }
            {{end}}
            {{.keys}}  // 处理缓存逻辑，例如删除或更新缓存
            // 执行带缓存处理的 SQL 操作
            return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
                return conn.ExecCtx(ctx, query, args...)
            }, {{.keyValues}}) // 传递缓存相关的键值
        {{else}}
            // 不使用缓存，直接执行 SQL 操作
            return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
                return conn.ExecCtx(ctx, query, args...)
            })
        {{end}}
    }

    // 如果 id == 0 或不需要缓存，直接执行 SQL
    return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        return conn.ExecCtx(ctx, query, args...)
    })
}