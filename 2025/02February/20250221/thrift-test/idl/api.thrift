namespace go api

include "base.thrift"

// User
struct RegisterRequest {
    1:required string username (api.vd="len($)>5 && len($)<32"),
    2:required string password (api.vd="len($)>6 && len($)<32"),
}

struct RegisterResponse {
    1:base.BaseResponse base_resp,
    2:string token,
}

struct LoginRequest {
    1:required string username (api.vd="len($)>5 && len($)<32"),
    2:required string password (api.vd="len($)>6 && len($)<32"),
}

struct LoginResponse {
    1:base.BaseResponse base_resp,
    2:string token,
}



service ApiService {
    RegisterResponse Register(1:RegisterRequest req)(api.post="/user/register"),
    LoginResponse Login(1:LoginRequest req)(api.post="/user/login"),
}
