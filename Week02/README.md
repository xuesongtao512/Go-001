# 极客大学「Go训练营-第0期」作业提交仓库

#### Week2 作业
问题: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
为什么，应该怎么做请写出代码？

答: 我认为应该 Wrap 这个 error，抛给上层, 理由: 在调用工具包或第三方时遇到 error 应该 warp.

老师的答案:
```cassandraql
dao 层

 return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))


biz 层:

if errors.Is(err, code.NotFound} {

}
```