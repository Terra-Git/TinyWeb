# 这是一个简单的mysql客户端

目标暂时是完成mysql连接，插入，删除，更新，查询的接口封装

暂时不考虑连接池以及多线程情况

golang 关于并发-同步这一块还需要学习，存疑


## 2023-03-03
增加 mysql insert接口的简单封装

目前是使用二维数组来控制的多条数据的插入

增加 mysql delete 接口，只支持删除单条数据，暂时未考虑复杂情况的删除

