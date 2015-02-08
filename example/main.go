package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	// Should be called at the very beginning of main().
	systray.Run(onReady)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("quit", "Quit", "Quit the whole app")
	go func() {
		<-mQuit.Ch
		systray.Quit()
		fmt.Println("Quit now...")
	}()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetIcon(icon.Data)
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Pretty awesome棒棒嗒")
		mChange := systray.AddMenuItem("change", "Change Me", "Change Me")
		mChecked := systray.AddMenuItem("check", "Unchecked", "Check Me")
		mEnabled := systray.AddMenuItem("enable", "Enabled", "Enabled")
		systray.AddMenuItem("ignore", "Ignored", "Ignored")
		mUrl := systray.AddMenuItem("lantern", "Open Lantern.org", "my home")
		mQuit := systray.AddMenuItem("quit2", "退出", "Quit the whole app")
		for {
			select {
			case <-mChange.Ch:
				mChange.Title = "I've Changed"
				systray.Update(mChange)
			case <-mChecked.Ch:
				mChecked.Checked = !mChecked.Checked
				if mChecked.Checked {
					mChecked.Title = "Checked"
				} else {
					mChecked.Title = "Unchecked"
				}
				systray.Update(mChecked)
			case <-mEnabled.Ch:
				mEnabled.Disabled = !mEnabled.Disabled
				mEnabled.Title = "Disabled"
				systray.Update(mEnabled)
			case <-mUrl.Ch:
				open.Run("https://www.getlantern.org")
			case <-mQuit.Ch:
				systray.Quit()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()
}
