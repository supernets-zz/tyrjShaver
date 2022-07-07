package main

import (
	"log"
	"time"

	. "github.com/electricbubble/gwda"
	extOpenCV "github.com/electricbubble/gwda-ext-opencv"
)

var (
	driver        WebDriver
	driverExt     *extOpenCV.DriverExt
	landUpgradePt = [9][2]int{
		{325, 150}, {325, 240}, {325, 330},
		{325, 420}, {325, 510}, {325, 600},
		{325, 420}, {325, 510}, {325, 600},
	}
	fieldPt = [9][2]int{
		{207, 465}, {117, 510}, {297, 510},
		{207, 555}, {117, 600}, {297, 600},
		{207, 645}, {117, 690}, {207, 735},
	}
	// landUpgradePt = [6][2]int{
	// 	{325, 230}, {325, 320},
	// 	{325, 410}, {325, 500},
	// 	{325, 590}, {325, 680},
	// }
)

const (
	luckyTasksBtn     = "./luckyTasks.PNG"
	luckyTasksPlay    = "./luckyTasksPlay.PNG"
	landManage        = "./landManage.PNG"
	landUpgradeBtn    = "./landUpgradeBtn.PNG"
	saveToAccountTips = "./saveToAccountTips.PNG"

	plantBtn      = "./plantBtn.PNG"
	hongbaoBtn    = "./hongbaoBtn.PNG"
	harvestBtn    = "./harvestBtn.PNG"
	copperCashBtn = "./copperCashBtn.PNG"
	killBugsBtn   = "./killBugsBtn.PNG"
	moleBtn       = "./moleBtn.PNG"

	bigHongbaoTips   = "./bigHongbaoTips.PNG"
	bigProductTips   = "./bigProductTips.PNG"
	killBugsTips     = "./killBugsTips.PNG"
	orderCompleteBtn = "./orderCompleteBtn.PNG"
	productProgress  = "./productProgress.PNG"
	speedUpAllTips   = "./speedUpAllTips.PNG"
)

func watchAD(driver WebDriver, timeout int64) {
	log.Println("看广告")
	defer log.Println("完成 看广告")
	startTick := time.Now()
	for {
		elementJump, err := driver.FindElement(BySelector{Predicate: `name IN {"| 跳过", "跳过"}`})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("点击 跳过")
			elementJump.Click()
			time.Sleep(3 * time.Second)
		}

		elementClose, err := driver.FindElement(BySelector{Predicate: `name IN {"endcard_close", "webview_closebutton", "closeButton_id"}`})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("点击 关闭")
			elementClose.DoubleTap()
			time.Sleep(5 * time.Second)
			break
		}

		time.Sleep(5 * time.Second)
		if time.Now().Unix()-startTick.Unix() > timeout {
			break
		}
	}
}

func closeSplashAD(driver WebDriver) {
	interstitialAD, err := driver.FindElement(BySelector{Name: "interstitial_ad"})
	if err != nil {
		log.Println("Splash广告: " + err.Error())
	} else {
		pt, err := interstitialAD.Location()
		if err != nil {
			log.Println(err)
		} else {
			size, err := interstitialAD.Size()
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("点击 关闭(%d, %d)", pt.X+size.Width-19, pt.Y+14)
				err = driver.Tap(pt.X+size.Width-19, pt.Y+14)
				if err != nil {
					log.Println(err)
				}
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func closeWebViewAD(driver WebDriver) {
	adWebView, err := driver.FindElement(BySelector{Predicate: `type == "XCUIElementTypeWebView"`})
	if err != nil {
		log.Println("WebView广告: " + err.Error())
	} else {
		pt, err := adWebView.Location()
		if err != nil {
			log.Println(err)
		} else {
			size, err := adWebView.Size()
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("点击 关闭(%d, %d)", pt.X+size.Width-23, pt.Y+23)
				err = driver.Tap(pt.X+size.Width-23, pt.Y+23)
				if err != nil {
					log.Println(err)
				}
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func acceptProducts(driver WebDriver, timeout int64) {
	startTick := time.Now()
	for {
		x, y, _, _, err := driverExt.FindImageRectInUIKit(saveToAccountTips)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("点击 开心收下(%d, %d)", int(x), int(y+50))
			err = driver.Tap(int(x), int(y+50))
			if err != nil {
				log.Println(err)
			}
			time.Sleep(3 * time.Second)
			break
		}
		time.Sleep(time.Second)
		if time.Now().Unix()-startTick.Unix() > timeout {
			log.Println("收货超时")
			break
		}
	}
}

func doLuckyTasks(driver WebDriver) {
	err := driverExt.Tap(luckyTasksBtn)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("点击 幸运任务")
	time.Sleep(5 * time.Second)
	for {
		err = driverExt.Tap(luckyTasksPlay)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("点击 立即领取")
		time.Sleep(5 * time.Second)

		watchAD(driver, 120)

		acceptProducts(driver, 10)

		closeWebViewAD(driver)
	}
}

func doLandUpgrade(driver WebDriver) {
	err := driverExt.Tap(landManage)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("点击 土地管理")
	time.Sleep(5 * time.Second)
	for {
		for i := 0; i < 9; i++ {
			err = driver.Tap(landUpgradePt[i][0], landUpgradePt[i][1])
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("点击 %d号土地 升级\n", i+1)
			time.Sleep(3 * time.Second)

			err = driverExt.Tap(landUpgradeBtn)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("点击 免费升级")
			time.Sleep(3 * time.Second)

			watchAD(driver, 120)

			acceptProducts(driver, 10)

			closeWebViewAD(driver)

			if i == 5 {
				err = driver.Swipe(landUpgradePt[4][0], landUpgradePt[4][1], landUpgradePt[1][0], landUpgradePt[1][1])
				if err != nil {
					log.Println(err)
					return
				}

				time.Sleep(3 * time.Second)
			} else if i == 8 {
				err = driver.Swipe(landUpgradePt[1][0], landUpgradePt[1][1], landUpgradePt[4][0], landUpgradePt[4][1])
				if err != nil {
					log.Println(err)
					return
				}

				time.Sleep(3 * time.Second)
			}
		}
	}
}

func doKillBugs(driver WebDriver) {
	log.Println("点击 除虫")
	err := driverExt.Tap(killBugsBtn)
	if err != nil {
		log.Println(err)
	} else {
		startTick := time.Now()
		for {
			killBugsTipsX, killBugsTipsY, killBugsTipsW, killBugsTipsH, err := driverExt.FindImageRectInUIKit(killBugsTips)
			if err != nil {
				log.Println("除虫提示: " + err.Error())
			} else {
				log.Printf("点击 花钱除虫(%0.0f, %0.0f)", killBugsTipsX+killBugsTipsW*3/4, killBugsTipsY+killBugsTipsH*3)
				err = driver.Tap(int(killBugsTipsX+killBugsTipsW*3/4), int(killBugsTipsY+killBugsTipsH*3))
				if err != nil {
					log.Println(err)
				}

				break
			}
			if time.Now().Unix()-startTick.Unix() > 5 {
				log.Println("除虫超时")
				break
			}
		}
	}
}

func doPunchMole(driver WebDriver, x, y int) {
	for i := 0; i < 9; i++ {
		log.Printf("锤 地鼠(%d, %d)", x, y)
		err := driver.Tap(x, y)
		if err != nil {
			log.Println("锤 地鼠 结束: " + err.Error())
			break
		}
		time.Sleep(200 * time.Microsecond)
	}

	for {
		err := driverExt.Tap(moleBtn)
		if err != nil {
			log.Println("锤 地鼠 结束: " + err.Error())
			break
		} else {
			log.Println("锤 地鼠")
		}
	}
}

func doHarvest(driver WebDriver) {
	for i := 0; i < 9; i++ {
		log.Printf("点击 地块%d(%d, %d)", i+1, fieldPt[i][0], fieldPt[i][1])
		err := driver.Tap(fieldPt[i][0], fieldPt[i][1])
		if err != nil {
			log.Println(err)
			break
		} else {
			time.Sleep(300 * time.Microsecond)
			bigHongbaoTipsX, bigHongbaoTipsY, bigHongbaoTipsW, bigHongbaoTipsH, err := driverExt.FindImageRectInUIKit(bigHongbaoTips)
			if err != nil {
				log.Println("没出现大红包: " + err.Error())
				closeSplashAD(driver)
				closeWebViewAD(driver)
			} else {
				log.Printf("点击 开(%0.0f, %0.0f)", bigHongbaoTipsX+bigHongbaoTipsW/2, bigHongbaoTipsY+3*bigHongbaoTipsH)
				err = driver.Tap(int(bigHongbaoTipsX+bigHongbaoTipsW/2), int(bigHongbaoTipsY+3*bigHongbaoTipsH))
				if err != nil {
					log.Println(err)
				} else {
					watchAD(driver, 120)
					acceptProducts(driver, 10)
					closeSplashAD(driver)
					closeWebViewAD(driver)
					continue
				}
			}

			bigProductTipsX, bigProductTipsY, bigProductTipsW, bigProductTipsH, err := driverExt.FindImageRectInUIKit(bigProductTips)
			if err != nil {
				log.Println("没出现变异作物: " + err.Error())
				closeSplashAD(driver)
				closeWebViewAD(driver)
			} else {
				log.Printf("点击 开心收下(%0.0f, %0.0f)", bigProductTipsX+bigProductTipsW/2, bigProductTipsY+3*bigProductTipsH)
				err = driver.Tap(int(bigProductTipsX+bigProductTipsW/2), int(bigProductTipsY+3*bigProductTipsH))
				if err != nil {
					log.Println(err)
				} else {
					watchAD(driver, 120)
					acceptProducts(driver, 10)
					closeSplashAD(driver)
					closeWebViewAD(driver)
					continue
				}
			}

			err = driverExt.Tap(orderCompleteBtn)
			if err != nil {
				log.Println("没有订单完成: " + err.Error())
			} else {
				log.Println("点击 订单完成 开心收下")
			}
		}
	}
}

func doPlant(driver WebDriver) {
	productProgressX, productProgressY, productProgressW, productProgressH, err := driverExt.FindImageRectInUIKit(productProgress)
	if err != nil {
		log.Println("未完成作物: " + err.Error())
	} else {
		i := 0
		for {
			log.Printf("种植 作物(%0.0f, %0.0f)", productProgressX-4*productProgressW, productProgressY-2*productProgressH)
			err = driver.Tap(int(productProgressX-4*productProgressW), int(productProgressY-2*productProgressH))
			if err != nil {
				log.Println("种植 结束: " + err.Error())
				break
			}
			time.Sleep(300 * time.Microsecond)
			i++
			if i >= 9 {
				x, y, w, h, err := driverExt.FindImageRectInUIKit(speedUpAllTips)
				if err != nil {
					log.Println("全体加速提示: " + err.Error())
				} else {
					err = driver.Tap(int(x+w/2), int(y+2*h))
					if err != nil {
						log.Println("点击 全体加速: " + err.Error())
					} else {
						log.Println("点击 全体加速")
						watchAD(driver, 120)
						closeSplashAD(driver)
						closeWebViewAD(driver)
						break
					}
				}
			}
		}
	}
}

func doPlantJobs(driver WebDriver) {
	for {
		hongbaoBtnX, hongbaoBtnY, _, _, errHongbao := driverExt.FindImageRectInUIKit(hongbaoBtn)
		if errHongbao != nil {
			log.Println("小红包: " + errHongbao.Error())
		} else {
			log.Printf("红包(%0.0f, %0.0f)", hongbaoBtnX, hongbaoBtnY)
		}
		harvestBtnX, harvestBtnY, _, _, errHarvest := driverExt.FindImageRectInUIKit(harvestBtn)
		if errHarvest != nil {
			log.Println("收获: " + errHarvest.Error())
		} else {
			log.Printf("收获(%0.0f, %0.0f)", harvestBtnX, harvestBtnY)
		}
		copperCashX, copperCashY, _, _, errCopperCash := driverExt.FindImageRectInUIKit(copperCashBtn)
		if errCopperCash != nil {
			log.Println("铜钱: " + errCopperCash.Error())
		} else {
			log.Printf("铜钱(%0.0f, %0.0f)", copperCashX, copperCashY)
		}
		killBugsBtnX, killBugsBtnY, _, _, errKillBugs := driverExt.FindImageRectInUIKit(killBugsBtn)
		if errKillBugs != nil {
			log.Println("除虫: " + errKillBugs.Error())
		} else {
			log.Printf("除虫(%0.0f, %0.0f)", killBugsBtnX, killBugsBtnY)
		}
		moleBtnX, moleBtnY, moleBtnW, moleBtnH, errMole := driverExt.FindImageRectInUIKit(moleBtn)
		if errMole != nil {
			log.Println("地鼠: " + errMole.Error())
		} else {
			log.Printf("地鼠(%0.0f, %0.0f)", moleBtnX, moleBtnY)
		}
		plantBtnX, plantBtnY, _, _, errPlant := driverExt.FindImageRectInUIKit(plantBtn)
		if errPlant != nil {
			log.Println("种植: " + errPlant.Error())
		} else {
			log.Printf("种植(%0.0f, %0.0f)", plantBtnX, plantBtnY)
		}

		if errKillBugs == nil {
			doKillBugs(driver)
			continue
		}

		if errMole == nil {
			doPunchMole(driver, int(moleBtnX+moleBtnW/2), int(moleBtnY+moleBtnH/2))
			continue
		}

		if errHarvest == nil {
			doHarvest(driver)
			continue
		}

		if errPlant == nil {
			doPlant(driver)
			continue
		}

		closeSplashAD(driver)
		closeWebViewAD(driver)
	}
}

func main() {
	log.Println("start")
	driver, err := NewUSBDriver(nil)
	checkErr(err)
	log.Println("connected")

	windowSize, _ := driver.WindowSize()
	log.Println(windowSize)

	driverExt, err = extOpenCV.Extend(driver, 0.95)
	checkErr(err, "扩展 driver ，指定匹配阀值为 95%（在不修改或者使用 `OnlyOnceThreshold` 的情况下）")
	log.Println("opencv ready")

	// closeSplashAD(driver)
	// closeWebViewAD(driver)
	// doLuckyTasks(driver)
	// doLandUpgrade(driver)
	doPlantJobs(driver)
	// for {
	// 	hongbaoBtn := "./hongbao.PNG"
	// 	err := driverExt.Tap(hongbaoBtn)
	// 	if err != nil {
	// 		log.Println(err)
	// 		break
	// 	}
	// }

	// for {
	// 	harvestBtn := "./harvest.PNG"
	// 	err := driverExt.Tap(harvestBtn)
	// 	if err != nil {
	// 		log.Println(err)
	// 		break
	// 	}
	// }

	// closeBtn := "./splashADCloseBtn.jpeg"
	// err = driverExt.Tap(closeBtn)
	// if err != nil {
	// 	log.Println(err)
	// }
	// element, err := driver.FindElement(BySelector{Name: "endcard_close"})
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	element.Click()
	// }

}

func checkErr(err error, msg ...string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
