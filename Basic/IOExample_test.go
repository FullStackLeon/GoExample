package Basic

import (
	"fmt"
	"io"
	"log"
	"testing"
	"time"
)

func TestIOPipe(t *testing.T) {
	pr, pw := io.Pipe()

	go func() {
		fmt.Println("写入数据...")
		_, err := pw.Write([]byte("Hello"))
		if err != nil {
			log.Fatal(err)
		}
		pw.Close()
		fmt.Println("关闭写端...")

		fmt.Println("尝试继续写入数据...")
		_, err = pw.Write([]byte("World"))
		if err != nil {
			fmt.Printf("继续写入失败：%v\n", err)
			return
		}
		fmt.Println("继续写入成功")

	}()

	data, err := io.ReadAll(pr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("读取数据:", string(data))
	pr.Close()
	fmt.Println("关闭读端...")
	time.Sleep(1 * time.Second)
	fmt.Println("尝试继续读取数据...")
	data, err = io.ReadAll(pr)
	if err != nil {
		fmt.Printf("继续读取失败：%v\n", err)
		return
	}
	fmt.Println("继续读取成功")

	// Output:
	// 写入数据...
	// 关闭写端...
	// 尝试继续写入数据...
	// 读取数据: Hello
	// 关闭读端...
	// 继续写入失败：io: read/write on closed pipe
	// 尝试继续读取数据...
	// 继续读取失败：io: read/write on closed pipe
}
