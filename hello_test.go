package main

import "testing"

//Fail: 标记失败，但继续执行当前测试函数
//FailNow: 失败，立即终止当前测试函数执行
//Log: 输出错误信息
//Error: Fail + Log
//Fatal: FailNow + Log
//Skip: 跳过当前函数，通常用于未完成的测试用例

func Add(a, b int) int {
    return a + b
}


func Test_Add1(t *testing.T) {
    if Add(2, 3) != 5 {
         t.Error("result is wrong!")
    } else {
         t.Log("result is right")
    }
}


func Test_Add2(t *testing.T) {
    if Add(2, 3) != 6 {
         t.Fatal("result is wrong!")
    } else {
         t.Log("result is right")
    }
}

//性能测试函数以Benchmark 开头，参数类型是 *testing.B，可与Test函数放在同个文件中。
//默认情况下，go test不执行Benchmark测试，必须用“-bench <pattern>”指定性能测试函数。
//go test -bench=.
func Benchmark(b *testing.B) {
    for i := 0; i < b.N; i++ { // b.N，测试循环次数
        Add(4, 5)
    }
}