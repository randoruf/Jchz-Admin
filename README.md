# Jchz-Admin 韭菜盒子后台管理系统
注：该项目很简陋且前端有些 bug（因为我不会前端，如果有人能帮我修一下就好了，比如弹出对话框不好好点"x"关闭会无法再打开），代码可能也有点不太标准，但这是为了纪念被学校的小组作业逼得5天从零基础爆肝的 go+vue 全栈简易后台项目：韭菜盒子后台管理系统。<br>
![](https://s2.loli.net/2022/07/07/6d9IvDypcQXR8xo.jpg)
### 简介：
前端页面使用 Vue3.2 + element plus 开发，支持中英双语（英语不全），参考了[B站视频教程](https://www.bilibili.com/video/BV1pq4y1c7oy/?spm_id_from=333.999.0.0)；后端服务器使用 Go 语言开发，使用了 Gin + Gorm 框架，采用 JWT 进行用户鉴权，支持使用 yaml 配置文件进行服务器配置，结构主要参照了 [GIN-VUE-ADMIN](https://github.com/flipped-aurora/gin-vue-admin)。 🍭🍭
### 已实现功能：
1. 用户管理：三类用户：普通用户，商家用户，管理员的增删改查功能。
2. 文章管理：文章的删改查功能（没有添加功能是因为觉得后台不需要做这种事）。
3. 标签管理：Tag 的增删改查功能。
4. 店铺管理：店铺的增删改查功能。

好了5天就做了这么点的 CURD，再多也没有了。😥😥

