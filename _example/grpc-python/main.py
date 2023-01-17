#! /usr/bin/env python
# coding=utf8

import grpc
import hello_pb2
import myservice_v1_pb2_grpc


def run():
    """
    模拟请求服务方法信息
    :return:
    """
    conn = grpc.insecure_channel('localhost:9609')
    client = myservice_v1_pb2_grpc.app1Stub(channel=conn)
    request = hello_pb2.Hello(message="jaronnie")
    response = client.SayHello(request)
    print("received:", response.message)


if __name__ == '__main__':
    run()
