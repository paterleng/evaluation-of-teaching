package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	username := ""
	password := ""
	options := append(
		chromedp.DefaultExecAllocatorOptions[:],
		//不检查默认浏览器
		chromedp.NoDefaultBrowserCheck,
		//禁用chrome的handless(禁用无窗口模式，即开启窗口模式)
		chromedp.Flag("headless", false),
		//开启图像界面
		chromedp.Flag("blink-settings", "imageEnabled=true"),
		//忽略错误
		chromedp.Flag("ignore-certificate-errors", true),
		//禁用网络安全标志
		chromedp.Flag("disable-web-security", true),
		//开启插件支持
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-popup-blocking", true),
		//设置网站不是首次运行
		chromedp.NoFirstRun,
		//设置窗口大小
		chromedp.WindowSize(1900, 1024),
	)
	allocator, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	print(cancel)
	ctx, cancel := chromedp.NewContext(
		allocator,
		chromedp.WithLogf(log.Printf),
	)
	//设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 10*time.Minute)
	//运行chromedp，操作浏览器
	chromedp.Run(
		ctx,
		//跳转到目标页面
		chromedp.Navigate("http://jwgl.hist.edu.cn/cas/login.action"),
		//输入账号密码
		//等页面加载出该元素后再执行操作
		chromedp.WaitVisible(`document.getElementById("username1")`, chromedp.ByJSPath),
		chromedp.SetValue(`document.getElementById("username1")`, username, chromedp.ByJSPath),
		chromedp.SetValue(`document.getElementById("username")`, username, chromedp.ByJSPath),
		//填入密码
		chromedp.WaitVisible(`document.getElementById("password1")`, chromedp.ByJSPath),
		chromedp.SetValue(`document.getElementById("password1")`, password, chromedp.ByJSPath),
		chromedp.SetValue(`document.getElementById("password")`, password, chromedp.ByJSPath),
		//点击登录按钮
		chromedp.Click(`document.getElementById("login")`, chromedp.ByJSPath),
	)
	//chromedp监听网页上弹出alert对话框
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if ev, ok := ev.(*page.EventJavascriptDialogOpening); ok {
			fmt.Println("closing alert:", ev.Message)
			go func() {
				//自动关闭alert对话框
				if err := chromedp.Run(ctx,
					//注释掉下一行可以更清楚地看到效果
					page.HandleJavaScriptDialog(true),
				); err != nil {
					panic(err)
				}
			}()
		}
	})
	//實現多次操作
	for i := 0; i < 36; i++ {
		chromedp.Run(ctx,
			//点击教学评价
			chromedp.WaitVisible(`document.querySelectorAll("#normal_use_menu li")[0]`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelectorAll("#normal_use_menu li")[0]`, chromedp.ByJSPath),
			//开始评教
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.getElementById("frmDesk").contentDocument.getElementById("frame_1").contentDocument.getElementById("selPJLC")`, chromedp.ByJSPath),
			chromedp.Click(`document.getElementById("frmDesk").contentDocument.getElementById("frame_1").contentDocument.getElementById("selPJLC")`, chromedp.ByJSPath),
			chromedp.Sleep(1*time.Second),
			chromedp.SetValue(`document.getElementById("frmDesk").contentDocument.getElementById("frame_1").contentDocument.querySelector("#selPJLC ")`, "{\"pjfsbz\":\"0\",\"lcjc\":\"第二阶段评价\",\"sfkpsj\":\"1\",\"sfzbpj\":\"1\",\"xn\":\"2022\",\"xq_m\":\"1\",\"lcqc\":\"2022-2023学年第二学期\",\"jsrq\":\"2023-06-06 23:59\",\"sfwjpj\":\"1\",\"lcdm\":\"2022102\",\"qsrq\":\"2023-06-01 00:00\"}", chromedp.ByJSPath),
			chromedp.Sleep(1*time.Second),
			//点击评价按钮
			chromedp.WaitVisible(`document.getElementById("frmDesk").contentDocument.getElementById("frame_1").contentDocument.getElementById("frmReport").contentDocument.querySelector("#tr0_wjdc > a")`, chromedp.ByJSPath),
			chromedp.Click(`document.getElementById("frmDesk").contentDocument.getElementById("frame_1").contentDocument.getElementById("frmReport").contentDocument.querySelector("#tr0_wjdc > a")`, chromedp.ByJSPath),
			//进入到了评分页面
			//1
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#pjxx0  #wdt_0_0_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#pjxx0  #wdt_0_0_1")`, chromedp.ByJSPath),
			//2
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#pjxx1 #wdt_0_1_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#pjxx1 #wdt_0_1_1")`, chromedp.ByJSPath),
			//3
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#pjxx2 #wdt_0_2_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#pjxx2 #wdt_0_2_1")`, chromedp.ByJSPath),
			//4
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_3_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_3_1")`, chromedp.ByJSPath),
			//5
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_4_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_4_1")`, chromedp.ByJSPath),
			//6
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_5_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_5_1")`, chromedp.ByJSPath),
			//7
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_6_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_6_1")`, chromedp.ByJSPath),
			//8
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_7_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_7_1")`, chromedp.ByJSPath),
			//9
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_8_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_8_1")`, chromedp.ByJSPath),
			//10
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_9_1")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#wdt_0_9_1")`, chromedp.ByJSPath),
			//1
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio0_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio0_0")`, chromedp.ByJSPath),
			//2
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio1_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio1_0")`, chromedp.ByJSPath),
			//3
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio2_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio2_0")`, chromedp.ByJSPath),
			//4
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio3_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio3_0")`, chromedp.ByJSPath),
			//5
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio4_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio4_0")`, chromedp.ByJSPath),
			//6
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio5_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio5_0")`, chromedp.ByJSPath),
			//7
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio6_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio6_0")`, chromedp.ByJSPath),
			//8
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio7_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio7_0")`, chromedp.ByJSPath),
			//9
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio8_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio8_0")`, chromedp.ByJSPath),
			//10
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio9_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio9_0")`, chromedp.ByJSPath),
			//11
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio10_0")`, chromedp.ByJSPath),
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#radio10_0")`, chromedp.ByJSPath),
			//给输入框赋值
			chromedp.WaitVisible(`document.querySelector("#dialog-frame").contentDocument.querySelector("#area11")`, chromedp.ByJSPath),
			chromedp.SetValue(`document.querySelector("#dialog-frame").contentDocument.querySelector("#area11")`, "无", chromedp.ByJSPath),
		)

		//提交
		chromedp.Run(ctx,
			chromedp.Click(`document.querySelector("#dialog-frame").contentDocument.querySelector("#butSave")`, chromedp.ByJSPath),
		)
		//提交成功后刷新頁面，避開提交成功按鈕
		chromedp.Run(ctx,
			chromedp.Sleep(1*time.Second),
			chromedp.Reload(),
		)
	}
}
