from locust import HttpUser, TaskSet, task


# 定义用户行为，继承TaskSet类，用于描述用户行为
# (这个类下面放各种请求，请求是基于requests的，每个方法请求和requests差不多，请求参数、方法、响应对象和requests一样的使用，url这里写的是路径)
# client.get===>requests.get
# client.post===>requests.post

class test_1(TaskSet):
    # task装饰该方法为一个事务方法的参数用于指定该行为的执行权重。参数越大，每次被虚拟用户执行概率越高，不设置默认是1，
    @task
    def test_api(self):
        # 定义requests的请求头
        header = {"User-Agent": "Mozilla/5.0 "
                                "(Windows NT 6.1; Win64; x64) AppleWebKit/537.36 "
                                "(KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36"}
        # r是包含所有响应内容的一个对象
        r = self.client.get("/giftCode/login", timeout=30, headers=header)
        # 这里可以使用assert断言请求是否正确，也可以使用if判断
        if r.status_code == 200:
            print("success")

class test_2(TaskSet):
    # task装饰该方法为一个事务方法的参数用于指定该行为的执行权重。参数越大，每次被虚拟用户执行概率越高，不设置默认是1，
    @task
    def test_api(self):
        # 定义requests的请求头
        header = {"User-Agent": "Mozilla/5.0 "
                                "(Windows NT 6.1; Win64; x64) AppleWebKit/537.36 "
                                "(KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36"}
        # r是包含所有响应内容的一个对象
        r = self.client.get("/giftCode/VerGiftCode", timeout=30, headers=header)
        # 这里可以使用assert断言请求是否正确，也可以使用if判断
        if r.status_code == 200:
            print("success")


# 这个类类似设置性能测试，继承HttpLocust
class websitUser(HttpUser):
    # 指向一个上面定义的用户行为类
    tasks = [test_1,test_2]
    # 执行事物之间用户等待时间的下界，单位毫秒，相当于lr中的think time
    min_wait = 3000
    max_wait = 6000


if __name__ == '__main__':
    import os

    os.system('locust -f giftCode_04.py --host=http://127.0.0.1:8080/')
