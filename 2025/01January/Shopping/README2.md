### 接口文档

#### 1. 用户相关接口

##### 1.1 注册用户
- **请求方法**: `POST`

- **URL**: `/User/Register`

- **请求体**:
    ```json
    {
        "username": "关注嘉然，顿顿解馋",
        "password": "我喜欢武宇恒"
    }
    ```
    
- **响应**:
  
    - 成功:
    
        ```json
        {
            "code": 200,
         	"info": "success"
        }
        ```
    
        
    
    - 用户已存在:
        ```json
        {
            "code": 406,
            "info": "Exist"
        }
        ```
        
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "bind error"
        }
        ```

##### 1.2 用户登录
- **请求方法**: `GET`
- **URL**: `/User/Login`
- **请求体**:
    ```json
    {
        "username": "string",
        "password": "string"
    }
    ```
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "token": "token"
        }
        ```
    - 登录失败:
        ```json
        {
            "code": 406,
            "info": "login error"
        }
        ```
    - 绑定错误/其他错误:
        ```json
        {
            "code": 500,
            "info": "bind error"
        }
        ```

##### 1.3 获取用户信息
- **请求方法**: `GET`
- **URL**: `/User/GetUserInfo`
- **请求头**: 
    - `Authorization`: `Bearer token`
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": {
                "id": "string",
                "username": "string",
                "balance": 0.0,
                "avatar": "string",
                "nickname": "string",
                "bio": "string"
            }
        }
        ```
    - 未授权:
        ```json
        {
            "code": 401,
            "info": "unauthorized"
        }
        ```

##### 1.4 充值
- **请求方法**: `PUT`
- **URL**: `/User/ReCharge`
- **请求体**: **money(表单填写)**
- **响应**:
  
    - 未认证
    
        ```json
        {
        	"code": 401,
        	"info": "unauthorized"
        }
        ```
    
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 1.5 修改用户信息
- **请求方法**: `PUT`

- **URL**: `/User/AlterUserInfo`

- **请求体**(无需全部填写):
  
    ```json
    {
        "nickname": "string",
        "avatar": "string",
        "bio": "string",
        "password": "string"
    }
    ```
    
- **响应**:
  
    - 未认证
    
      ```json
      {
      	"code": 401,
      	"info": "unauthorized"
      }
      ```
    
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 1.6 删除用户
- **请求方法**: `DELETE`
- **URL**: `/User/DelUser`
- **响应**:
  
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 未授权:
        ```json
        {
            "code": 401,
            "info": "unauthorized"
        }
        ```

#### 2. 购物车相关接口

##### 2.1 添加商品到购物车
- **请求方法**: `POST`
- **URL**: `/User/AddGoodsToCart`
- **请求体**:
    ```json
    {
        "goods_id": "string",
        "number": 1
    }
    ```
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 商品缺货:
        ```json
        {
            "code": 406,
            "info": "lack"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 2.2 删除购物车中的商品
- **请求方法**: `DELETE`
- **URL**: `/User/DelGoodsFromCart`
- **请求体**:
    ```json
    {
        "goods_id": "string"
    }
    ```
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 商品未找到:
        ```json
        {
            "code": 406,
            "info": "no goods found"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 2.3 获取购物车信息
- **请求方法**: `GET`
- **URL**: `/User/GetCartInfo`
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": {
                "id": "string",
                "sum": 100.0,
                "goods": [
                    {
                        "id": "goods_id",
                        "goods_name": "string",
                        "number": 5,
                        "price": 20.0
                    }
                ]
            }
        }
        ```
    - 购物车为空或其他错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

#### 3. 商品相关接口

##### 3.1 注册商铺
- **请求方法**: `POST`
- **URL**: `/Shop/RegisterMall`
- **请求体**:
    ```json
    {
        "shop_name": "string",
        "password": "string"
    }
    ```
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 商铺已存在:
        ```json
        {
            "code": 406,
            "info": "Exist"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "bind error"
        }
        ```

##### 3.2 商铺登录
- **请求方法**: `GET`
- **URL**: `/Shop/LoginMall`
- **请求体**:
    ```json
    {
        "shop_name": "string",
        "password": "string"
    }
    ```
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "token": "jwtTokenString"
        }
        ```
    - 登录失败:
        ```json
        {
            "code": 406,
            "info": "error"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "bind error"
        }
        ```

##### 3.3 注册商品
- **请求方法**: `POST`

- **URL**: `/Shop/RegisterGoods`

- **请求体**(都必须要填写):
  
    ```json
    {
        "goods_name": "string",
        "number": 1,
        "price": 20.0,
        "avatar": "string",
        "content": "string",
        "type": "string",
        "shop_id": "string"
    }
    ```
    
- **响应**:

    - 未认证

      ```json
      {
      	"code": 401,
      	"info": "unauthorized",
      }
      ```

    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 3.4 修改商品信息
- **请求方法**: `PUT`
- **URL**: `/Shop/AlterGoodsInfo`
- **请求体**:
  
    ```json
    {
        "id": "goods_id",
        "goods_name": "string",
        "number": 2,
        "price": 25.0,
        "avatar": "string",
        "content": "string",
        "type": "string",
        "shop_id": "string"
    }
    ```
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 3.5 删除商品
- **请求方法**: `DELETE`

- **URL**: `/Shop/DelGoods`

- **请求体**:
    ```json
    {
        "id": "goods_id"
    }
    ```
    
- **响应**:

    - 未认证

      ```json
      {
      	"code": 401,
      	"info": "unauthorized",
      }
      ```

    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 商品未找到:
        ```json
        {
            "code": 404,
            "info": "goods not found"
        }
        ```

#### 4. 订单相关接口

##### 4.1 提交订单
- **请求方法**: `POST`

- **URL**: `/User/SubmitOrder`

- **响应**:

    - 未认证

      ```json
      {
      	"code": 401,
      	"info": "unauthorized",
      }
      ```

    - 成功

       ```json
    {
        "code": 200,
        "info": "ok",
        "data": {
            "order_id": "string",
            "user_id": "string",
            "sum": 100.0,
            "goods": [
                {
                    "goods_id": "string",
                    "goods_name": "string",
                    "number": 2,
                    "price": 20.5
                }
             ]
          }
    }
    
    - **订单提交失败**:
    
       ```json
    {
    "code": 500,
    "info": "error"
    }

##### 4.2 确认订单

- **请求方法**: `PUT`

- **URL**: `/User/ConfirmOrder`

- **请求体**:
    ```json
    {
        "id": "order_id"
    }
    ```
    
- **响应**:

    - **未授权**

      ```json
      {
          "code": 401,
          "info": "unauthorized"
      }
      ```

    - **绑定错误**

      ```json
      {
          "code": 500,
          "info": "error: <error_message>"
      }
      ```

    - **缺货**

      ```json
      {
          "code": 406,
          "info": "lack goods",
          "data": [
              {
                  "goods_id": "string",
                  "current_num": 3,
                  "query_num": 5
              }
          ]
      }
      ```

    - **余额不足**

      ```json
      {
          "code": 406,
          "info": "balance lack"
      }
      ```

    - **订单已删除**

      ```json
      {
          "code": 406,
          "info": "order deleted"
      }
      ```

    - **确认成功**

      ```json
      {
          "code": 200,
          "info": "ok"
      }
      ```

    - **未知错误**

      ```json
      {
          "code": 500,
          "info": "Unknown error"
      }
      ```
##### 4.3 取消订单
- **请求方法**: `DELETE`

- **URL**: `/User/CancelOrder`

- **请求体**:
    ```json
    {
        "id": "order_id"
    }
    ```
    
- **响应**:

    - 未认证

      ```json
      {
      	"code": 401,
      	"info": "unauthorized",
      }
      ```

    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 订单未找到:
        ```json
        {
            "code": 404,
            "info": "order not found"
        }
        ```

#### 5. 消息相关接口

##### 5.1 发布评论
- **请求方法**: `POST`
- **URL**: `/User/PubMsg`
- **请求体**:
    ```json
    {
        "goods_id": "string",
        "content": "string"
    }
    ```
- **响应**:

    - 未授权

      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```

    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 5.2 回复评论
- **请求方法**: `POST`
- **URL**: `/User/Response`
- **请求体**:
    ```json
    {
        "parent_id": "string",
        "content": "string"
    }
    ```
- **响应**:
  
    - 未授权
    
      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```
    
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 5.3 获取商品评论
- **请求方法**: `GET`

- **URL**: `/User/GetGoodsMsg`

- **请求体**:
    ```json
    {
        "goods_id": "string"
    }
    ```
    
- **响应**:
  
    - 未授权
    
      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```
    
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": [
                {
                    "id": "string",
                    "parent_id": "string",
                    "goods_id": "string",
                    "user_id": "string",
                    "content": "string",
                    "praised_num": 10,
                    "create_at": "2023-01-01T00:00:00Z",
                    "updated_at": "2023-01-01T00:00:00Z",
                    "response": [] // 潜在的回复
                }
            ]
        }
        ```
    - 商品未找到:
        ```json
        {
            "code": 404,
            "info": "error"
        }
        ```

##### 5.4 点赞评论

- **请求方法**: `PUT`
- **URL**: `/User/Praise`
- **请求体**:
  
    ```json
    {
        "user_id": "string",
        "message_id": "string"
    }
    ```
- **响应**:
  
    - 未授权
    
      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```
    
    - 成功 (点赞成功):
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 取消点赞 (已点赞):
        ```json
        {
            "code": 200,
            "info": "ok - unliked"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```



##### 5.5 修改评论
- **请求方法**: `PUT`
- **URL**: `/User/AlterMsg`
- **请求体**:
  
    ```json
    {
        "id": "string",
        "content": "string",
        "user_id": "string"
    }
    ```
- **响应**:
  
    - 未授权
    
      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```
    
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 评论未找到:
        ```json
        {
            "code": 404,
            "info": "msg not found"
        }
        ```
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

##### 5.6 删除评论
- **请求方法**: `DELETE`

- **URL**: `/User/DelMsg`

- **请求体**:
    ```json
    {
        "id": "string",
        "user_id": "string"
    }
    ```
    
- **响应**:
  
    - 未授权
    
      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```
    
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    
    - 绑定错误:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```



#### 6. 浏览记录接口

##### 6.1 获取浏览记录
- **请求方法**: `GET`
- **URL**: `/User/BrowseRecords`
- **响应**:

    - 未授权

      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```

    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": [
                {
                    "id": "string",
                    "user_id": "string",
                    "goods_id": "string",
                    "goods_name": "string",
                    "avatar": "string",
                    "browse_time": "2023-01-01T00:00:00Z"
                }
            ]
        }
        ```
    - 无浏览记录:
        ```json
        {
            "code": 204,
            "info": "no browse records"
        }
        ```

##### 6.2 浏览商品
- **请求方法**: `GET`

- **URL**: `/User/BrowseGoods`

- **请求体**:
    ```json
    {
        "goods_id": "string"
    }
    ```
    
- **响应**:

    - 用户不存在

        ```
        {
            "code": 406,
            "info": "user error"
        }
        ```

    - 未授权

        ```json
        {
        	"code": 401,
            "info": "unauthorized"
        }
        ```

    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": {
                "id": "string",
                "goods_name": "string",
                "shop_id": "string",
                "type": "string",
                "number": 10,
                "price": 25.0,
                "star": 5,
                "content": "string",
                "avatar": "string"
            }
        }
        ```
    - 商品未找到:
        ```json
        {
            "code": 500,
            "info": "error"
        }
        ```

#### 7. 收藏和搜索相关接口

##### 7.1 收藏商品
- **请求方法**: `PUT`
- **URL**: `/User/Star`
- **请求体**:
    ```json
    {
        "user_id": "string",
        "goods_id": "string"
    }
    ```
- **响应**:
  
    - 
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok"
        }
        ```
    - 商品已被收藏:
        ```json
        {
            "code": 409,
            "info": "already starred"
        }
        ```
    

##### 7.2 获取所有收藏
- **请求方法**: `GET`

- **URL**: `/User/GetAllStar`

- **响应**:
    
    - 未授权
    
      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```
    
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": [
                {
                    "goods_name": "string",
                    "type": "string",
                    "price": 20.0,
                    "star": 5,
                    "avatar": "string"
                }
            ]
        }
        ```
    - 无收藏商品:
        ```json
        {
            "code": 204,
            "info": "no starred goods"
        }
        ```

##### 7.3 搜索商品
- **请求方法**: `POST`

- **URL**: `/User/Search`

- **请求体**:
    ```json
    {
        "content": "search_query"
    }
    ```
    
- **响应**:

    - 未授权

      ```json
      {
      	"code": 401,
          "info": "unauthorized"
      }
      ```

    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": [
                {
                    "goods_id": "string",
                    "goods_name": "string",
                    "type": "string",
                    "price": 20.0,
                    "star": 5,
                    "avatar": "string"
                }
            ]
        }
        ```
    - 无结果:
        ```json
        {
            "code": 404,
            "info": "no goods found"
        }
        ```

##### 7.4 根据商品类型搜索商品
- **请求方法**: `GET`
- **URL**: `/User/GetTypeGoods`
- **请求体**:
    ```json
    {
        "type": "string"
    }
    ```
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": [
                {
                    "goods_name": "string",
                    "type": "string",
                    "price": 20.0,
                    "star": 5,
                    "avatar": "string"
                }
            ]
        }
        ```
    - 无结果:
        ```json
        {
            "code": 404,
            "info": "no goods found"
        }
        ```

##### 7.5 查看店铺和商品信息
- **请求方法**: `GET`
- **URL**: `/GetShopAndGoodsInfo`
- **请求体**:
    ```json
    {
        "shop_name": "string"
    }
    ```
- **响应**:
    - 成功:
        ```json
        {
            "code": 200,
            "info": "ok",
            "data": {
                "id": "shop_id",
                "shop_name": "string",
                "profit": 1000.0,
                "goods": [
                    {
                        "goods_name": "string",
                        "type": "string",
                        "price": 20.0,
                        "star": 5,
                        "avatar": "string"
                    }
                ]
            }
        }
        ```
    - 店铺未找到:
        ```json
        {
            "code": 404,
            "info": "shop not found"
        }
        ```

-----

以上大致为本demo实现的接口, 如有疑问,请联系我.
