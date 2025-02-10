namespace go user

include "base.thrift"

struct RegisterRequest {
    1: string username
    2: string password
}

struct LoginRequest {
    1: string username
    2: string password
}


service User {
    base.Response Register(1: RegisterRequest Request)
    base.Response Login(1: LoginRequest Request)
}