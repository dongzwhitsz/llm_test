// idl/hello.thrift
namespace go dongzwhitsz.llm_test.demo_repo

struct HelloReq {
    1: string Name (api.query="name"); // Add api annotations for easier parameter binding
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}
