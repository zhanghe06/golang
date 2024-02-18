package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
)

// MouseDragNode 模拟滑动
func MouseDragNode(n *cdp.Node, cxt context.Context) error {
	boxes, err := dom.GetContentQuads().WithNodeID(n.NodeID).Do(cxt)
	if err != nil {
		return err
	}
	if len(boxes) == 0 {
		return chromedp.ErrInvalidDimensions
	}
	content := boxes[0]
	c := len(content)
	if c%2 != 0 || c < 1 {
		return chromedp.ErrInvalidDimensions
	}
	var x, y float64
	for i := 0; i < c; i += 2 {
		x += content[i]
		y += content[i+1]
	}
	x /= float64(c / 2)
	y /= float64(c / 2)
	p := &input.DispatchMouseEventParams{
		Type:       input.MousePressed,
		X:          x,
		Y:          y,
		Button:     input.Left,
		ClickCount: 1,
	}
	// 鼠标左键按下
	if err := p.Do(cxt); err != nil {
		return err
	}
	// 拖动
	p.Type = input.MouseMoved
	max := 380.0
	for {
		if p.X > max {
			break
		}
		rt := rand.Intn(20) + 20
		chromedp.Run(cxt, chromedp.Sleep(time.Millisecond*time.Duration(rt)))
		x := rand.Intn(2) + 15
		y := rand.Intn(2)
		p.X = p.X + float64(x)
		p.Y = p.Y + float64(y)
		//fmt.Println("X坐标：",p.X)
		if err := p.Do(cxt); err != nil {
			return err
		}
	}
	// 鼠标松开
	p.Type = input.MouseReleased
	return p.Do(cxt)
}

func main() {
	// 参数设置
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		//chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
		// 设置代理
		chromedp.ProxyServer("socks5://127.0.0.1:1086"),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	// 创建chrome示例
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var (
		buf   []byte
		value string
	)
	err := chromedp.Run(ctx,
		chromedp.Tasks{
			// 打开导航
			chromedp.Navigate("https://google.com/"),
			// 等待元素加载完成
			chromedp.WaitVisible("body", chromedp.ByQuery),
			// 输入chromedp
			chromedp.SendKeys("APjFqb", "chromedp", chromedp.ByID),
			// 打印输入框的值
			chromedp.Value("APjFqb", &value),
			// 提交
			chromedp.Submit("APjFqb", chromedp.ByID),
			chromedp.Sleep(3 * time.Second),
			// 截图
			chromedp.CaptureScreenshot(&buf),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("value: ", value)
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		fmt.Println(err)
	}
}
