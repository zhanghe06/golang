package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		//chromedp.DisableGPU,
		//chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false),
		chromedp.Flag("ignore-certificate-errors", true),
		//chromedp.Flag("window-size", "50,400"),
		chromedp.UserDataDir(dir),
		chromedp.ProxyServer("socks5://127.0.0.1:1086"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// create a timeout
	taskCtx, cancel = context.WithTimeout(taskCtx, 100*time.Second)
	defer cancel()

	// ensure that the browser process is started
	if err := chromedp.Run(taskCtx); err != nil {
		panic(err)
	}

	// listen network event
	listenForNetworkEventPH(taskCtx)

	chromedp.Run(taskCtx,
		network.Enable(),
		//设置webdriver检测反爬
		chromedp.ActionFunc(func(cxt context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(cxt)
			return err
		}),
		chromedp.Navigate(`https://cn.pornhub.com/view_video.php?viewkey=ph5f364cc4ed76b`),
		chromedp.WaitVisible(`body`, chromedp.BySearch),
	)

}

//监听
func listenForNetworkEventPH(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {

		case *network.EventResponseReceived:
			resp := ev.Response
			if len(resp.Headers) != 0 {
				// log.Printf("received headers: %s", resp.Headers)

				//if strings.Index(resp.URL, ".ts") != -1 {
				//	log.Printf("received headers: %s", resp.URL)
				//}
				if strings.Index(resp.URL, "master.m3u8") != -1 {
					log.Printf("received headers: %s", resp.URL)
				}
				if strings.Index(resp.URL, "index-f1-v1-a1.m3u8") != -1 {
					log.Printf("received headers: %s", resp.URL)
				}
			}

		}
		// other needed network Event
	})
}
