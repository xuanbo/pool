## pool

> golang 协程池

### 安装

```
go get github.com/xuanbo/pool
```

### 使用

```
package main

import (
    "fmt"
    "time"

    "github.com/xuanbo/pool"
)

func main() {
    // 创建协程池，5个goroutines、任务队列长度为10
    wp := pool.NewWorkerPool(5, 10)
    // 运行
    wp.Start()
    for i := 0; i < 100; i++ {
        count := i
        // 提交任务，如果任务队列满了会阻塞
        wp.Add(func() {
            fmt.Printf("%d\n", count)
        })
    }
    // 停止
    wp.Stop()
    time.Sleep(2 * time.Second)
}
```

### 参考

[使用Go语言每分钟处理1百万请求（译）](https://mp.weixin.qq.com/s?__biz=MjM5OTcxMzE0MQ==&mid=2653369770&idx=1&sn=044be64c577a11a9a13447b373e80082&chksm=bce4d5b08b935ca6ad59abb5cc733a341a5126fefc0e6600bd61c959969c5f77c95fbfb909e3&mpshare=1&scene=1&srcid=1010dpu0DlPHi6y1YmrixifX#rd)