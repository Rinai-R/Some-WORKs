在demo的基本功能完善之后，我又在某网站上面看见了一些sql优化的文章，令我大受震撼，反思我在这个demo中对应的mysql的表结构有很多重复的内容，并且还有很多的问题都凸显出来，我可能接下来一段时间学习的重点就会变成mySQL的优化吧

-----

- **用户**
  - id(唯一自增id)
  - username(唯一，不为空)
  - password(not null)
  - balance(初始值为0)
  - avatar(默认头像)

CREATE TABLE user(

id INT AUTO_INCREMENT PRIMARY KEY,

username VARCHAR(255) NOT NULL UNIQUE,

password VARCHAR(255) NOT NULL,

balance  DECIMAL(10,2) DEFAULT 0.00,

avatar VARCHAR(255) DEFAULT 'default'

);

- **浏览记录**
  - id
  - user_id
  - browse_time
  - goods_id

CREATE TABLE browse_records (

id INT AUTO_INCREMENT NOT NULL,

user_id INT NOT NULL,

browse_time DATETIME DEFAULT CURRENT_TIMESTAMP,

goods_id INT NOT NULL,

goods_name VARCHAR(255) NOT NULL,

avatar VARCHAR(255) NOT NULL,

FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,

FOREIGN KEY (goods_id) REFERENCES goods(id) ON DELETE CASCADE

)

- **购物车**
  - user_id(归属于谁？)
  - sum(价格总和)
  - 外键，绑定user，用户被删，购物车也会被删。

CREATE TABLE shopping_cart (

user_id INT NOT NULL,

sum DECIMAL(10,2) DEFAULT 0.00,

FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE

);

- **被加入购物车的商品实体**
  - user_id(表面是归属用户，实际绑定外键是归属于购物车，更有逻辑性)
  - goods_id(商品的id)
  - number
  - price(数量乘上单价的总价格)
  - 外键，绑定到购物车实体。

CREATE TABLE cart_goods (

user_id INT NOT NULL,

goods_id INT NOT NULL,

goods_name VARCHAR(255) NOT NULL,

number INT NOT NULL,

price DECIMAL(10,2) NOT NULL,

FOREIGN KEY (goods_id) REFERENCES goods(id) ON DELETE CASCADE,

FOREIGN KEY (user_id) REFERENCES shopping_cart(user_id) ON DELETE CASCADE

);

- **商家**
  - id
  - shop_name
  - password
  - profit

CREATE TABLE shop (

id INT AUTO_INCREMENT PRIMARY KEY,

shop_name VARCHAR(255) NOT NULL UNIQUE,

password VARCHAR(255) NOT NULL,

profit DECIMAL(10,2) DEFAULT 0.00

);

- **商家的商品**
  - shop_id
  - id
  - type
  - number
  - price
  - star
  - avatar
  - content
  - star
  - 外键，绑定到商家

CREATE TABLE goods (

id INT AUTO_INCREMENT PRIMARY KEY,

goods_name VARCHAR(255) NOT NULL,

shop_id INT NOT NULL,

type VARCHAR(255) NOT NULL,

number INT NOT NULL,

price DECIMAL(10,2) NOT NULL,

content VARCHAR(255) NOT NULL,

star INT DEFAULT 0,

avatar VARCHAR(255) DEFAULT 'default',

FOREIGN KEY (shop_id) REFERENCES shop(id) ON DELETE CASCADE

);

- 评论
  - id
  - goods_id
  - user_id
  - content
  - praised_num
  - parent_id
  - create_at
  - updated_at

CREATE TABLE msg (

id INT AUTO_INCREMENT PRIMARY KEY,

goods_id INT NOT NULL,

user_id INT NOT NULL,

content VARCHAR(255) NOT NULL,

praised_num INT DEFAULT 0,

parent_id INT NOT NULL,

create_at DATETIME DEFAULT CURRENT_TIMESTAMP,

updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,

FOREIGN KEY (parent_id) REFERENCES msg(id) ON DELETE CASCADE,

FOREIGN KEY (goods_id) REFERENCES goods(id) ON DELETE CASCADE

)

- 点赞表（相当于一条链子，将用户和点赞的内容串起来）
  - 点赞者
  - 点赞评论
  - 外键-点赞人id
  - 外键-评论id

CREATE TABLE praise (

user_id INT NOT NULL,

message_id INT NOT NULL,

FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,

FOREIGN KEY (message_id) REFERENCES msg(id) ON DELETE CASCADE

)

- 收藏表
  - 收藏人
  - 收藏商品
  - 外键 收藏人id
  - 外键 商品id

CREATE TABLE star (

user_id INT NOT NULL,

goods_id INT NOT NULL,

FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,

FOREIGN KEY (goods_id) REFERENCES goods(id) ON DELETE CASCADE

)

- 搜索操作
  - id
  - 搜索内容

CREATE TABLE search (

id INT AUTO_INCREMENT PRIMARY KEY,

content VARCHAR(255) NOT NULL

)

- 商品相关性表
  - 搜索操作id
  - 商品id
  - 商品名称
  - 头像
  - 相关系数

CREATE TABLE association (

search_id INT NOT NULL,

goods_id INT NOT NULL,

goods_name VARCHAR(255) NOT NULL,

value INT NOT NULL,

avatar VARCHAR(255) DEFAULT 'default',

type VARCHAR(255) NOT NULL,

star INT NOT NULL,

FOREIGN KEY (search_id) REFERENCES search(id) ON DELETE CASCADE,

FOREIGN KEY (goods_id) REFERENCES goods(id) ON DELETE CASCADE

)

- 订单
  - 订单id
  - 用户id
  - 店铺id
  - 总价格

CREATE TABLE orders (

id INT AUTO_INCREMENT PRIMARY KEY,

user_id INT NOT NULL,

shop_id INT NOT NULL,

sum INT NOT NULL

);

- 订单商品(此处如果利用购物车商品或者原商品绑定的话，还是感觉不太好看)
  - 商品id
  - 单价price
  - 数量
  - 订单id
  - 外键订单id

CREATE TABLE order_goods (

id INT NOT NULL,

price DECIMAL(10,2) NOT NULL,

number INT NOT NULL,

order_id INT NOT NULL,

goods_name VARCHAR(255) NOT NULL,

FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE

)
