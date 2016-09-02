package gui

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Init2() {
	log.Println("初始化GUI")
	logo_img, _ := walk.NewImageFromFile("static/img/logo_img.bmp")
	window := MainWindow{
		MinSize: Size{430, 330},
		MaxSize: Size{430, 330},
		Layout:  HBox{},
		Children: []Widget{
			ImageView{MinSize: Size{415, 160}, Image: logo_img},
		},
	}
	window.Run()
}

func Init() {
	log.Println("初始化GUI")
	window, _ := walk.NewMainWindow()
	window.SetSize(walk.Size{430, 330})
	window.SetX(700)
	window.SetY(300)

	font, _ := walk.NewFont("宋体", 14, walk.FontItalic)

	logo, _ := walk.NewImageView(window)
	logo.SetX(0)
	logo.SetY(0)
	logo.SetSize(walk.Size{415, 160})
	logo_img, _ := walk.NewImageFromFile("static/img/logo_img.bmp")
	logo.SetImage(logo_img)

	face, _ := walk.NewImageView(window)
	face.SetX(30)
	face.SetY(170)
	face.SetSize(walk.Size{80, 80})
	face_img, _ := walk.NewImageFromFile("static/img/face.bmp")
	face.SetImage(face_img)

	account, _ := walk.NewLineEdit(window)
	account.SetX(130)
	account.SetY(170)
	account.SetSize(walk.Size{195, 30})
	account.SetFont(font)
	account.SetCueBanner("请输入账号")

	reg, _ := walk.NewLabel(window)
	reg.SetText("注册账号")
	reg.SetX(335)
	reg.SetY(170)
	reg.SetSize(walk.Size{50, 30})

	pwd, _ := walk.NewLineEdit(window)
	pwd.SetX(130)
	pwd.SetY(200)
	pwd.SetSize(walk.Size{195, 30})
	pwd.SetPasswordMode(true)
	pwd.SetFont(font)
	pwd.SetCueBanner("请输入密码")

	findPwd, _ := walk.NewLabel(window)
	findPwd.SetText("找回密码")
	findPwd.SetX(335)
	findPwd.SetY(200)
	findPwd.SetSize(walk.Size{50, 30})

	remberPwd, _ := walk.NewCheckBox(window)
	remberPwd.SetX(130)
	remberPwd.SetY(232)
	remberPwd.SetSize(walk.Size{100, 20})
	remberPwd.SetText("记住密码")

	autoLogin, _ := walk.NewCheckBox(window)
	autoLogin.SetX(260)
	autoLogin.SetY(232)
	autoLogin.SetSize(walk.Size{100, 20})
	autoLogin.SetText("自动登录")

	login, _ := walk.NewPushButton(window)
	login.SetText(" ")
	login.SetX(130)
	login.SetY(255)
	login.SetSize(walk.Size{195, 30})
	img, _ := walk.NewImageFromFile("D:/Go/src/github.com/louch2010/gochat-server/static/img/login_btn.bmp")
	login.SetImage(img)
	login.SetCursor(walk.CursorHand())
	login.Clicked().Attach(func() {
		log.Println("aaaaaaa")
	})

	window.Show()
	window.Run()
}
