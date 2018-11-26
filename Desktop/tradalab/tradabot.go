package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	td "takedatabase"
	"time"

	_ "github.com/go-sql-driver/mysql"

	tb "gopkg.in/tucnak/telebot.v2"
)

const adminID = 751335863

// Init step
const (
	step0 = iota
	step1
	step2
	step3
	step4
	step5
)

// Student Infor
var studentInfor [6]string

// OK
var yesSir bool

var text string

var texttest string = "OK man"

// Step int
type Step int

//Init userInfor
type userInfo struct {
	displayName       string
	toAgreed          bool
	subscription      Step
	lastSigninRequest time.Time
	thisStep          int
}

//Init userMap
var userMap = make(map[int]*userInfo)

func main() {

	os.Setenv("botAPI", "700769156:AAHG9K8E_c3eTW_4Sa782CVzIh8WTkgWqlY")
	// Init b: newBot
	b, err := tb.NewBot(
		tb.Settings{
			Token:  os.Getenv("botApI"),
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		},
	)
	checkErr(err)

	// Handle with tb.Ontext
	b.Handle(tb.OnText, func(m *tb.Message) {
		handleOnText(b, m)
	})

	// Handle with /start
	b.Handle("/start", func(m *tb.Message) {
		studentInfor[0] = getFullName(b, m)
		next(b, m)
	})
	// Handle with /ls
	b.Handle("/ls", func(m *tb.Message) {
		next(b, m)
	})
	// Handle with /cancel
	b.Handle("/cancel", func(m *tb.Message) {
		cancelFunc(b, m)
	})

	// Handle with  /help
	b.Handle("/help", func(m *tb.Message) {
		helpFunc(b, m)
	})

	// Handel with /info
	b.Handle("/info", func(m *tb.Message) {
		infoFunc(b, m)
	})

	//Handle with /me
	b.Handle("/me", func(m *tb.Message) {
		meFunc(b, m)
	})

	b.Start()
}

// Handle on text

func handleOnText(b *tb.Bot, m *tb.Message) {
	if user, ok := userMap[m.Sender.ID]; ok {
		// Step1: Show List Of List
		switch user.thisStep {
		//start
		case step0:
			log.Println("step:0")
			user.thisStep = step1
			next(b, m)
		// chose lesson
		case step1:
			log.Println("step:1")
			if yesSir == false {
				user.thisStep = step1
				next(b, m)
			} else {
				user.thisStep = step2
				next(b, m)
			}
		// Question
		case step2:
			log.Printf("step:2")
			if yesSir == true {
				user.thisStep = step3
				next(b, m)
			} else {
				user.thisStep = step4
			}

		case step3:
			user.thisStep = step4
			next(b, m)

		// Nother info
		// Fill name
		// case step3:
		// 	log.Printf("step:3")
		// 	if yesSir == false {
		// 		user.thisStep = step3
		// 		next(b, m)
		// 	} else {
		// 		user.thisStep = step4
		// 		next(b, m)
		// 	}
		// case step4:
		// 	log.Printf("step:4")
		// 	if yesSir == false {
		// 		user.thisStep = step4
		// 		next(b, m)
		// 	} else {
		// 		user.thisStep = step5
		// 		next(b, m)
		// 	}
		// case step5:
		// 	log.Printf("step:5")
		// 	if yesSir == false {
		// 		user.thisStep = step5
		// 		next(b, m)
		// 	} else {
		// 		user.thisStep = step6
		// 		next(b, m)
		// 	}
		case step4:
			log.Printf("step:3")
			next(b, m)
		case step5:
			user.thisStep = step5
			next(b, m)
		default:
			user.thisStep = step1
		}
	}
}

// Function with next
func next(b *tb.Bot, m *tb.Message) {
	if user, ok := userMap[m.Sender.ID]; ok {
		// Array Function
		funcArray := []func(*tb.Bot, *tb.Message){
			startAsk,
			// choseLesson,
			startSignUp,
			choseOther,
			anotherInfor,
			end,
			// startConFirmName,
			// startConFirmNumber,
			// startConFirmYearOld,
		}
		log.Println("user.thisStep:", user.thisStep)
		funcArray[user.thisStep](b, m)
	} else {
		checkUser(b, m)
	}

}

// Step1: start and Chose ID Lesson
func startAsk(b *tb.Bot, m *tb.Message) {
	send(b, m, "Cảm ơn bạn đã liên lạc với Trada Tech. Đây là bot trả lời tự động hỗ trợ các lệnh sau\n /register - đăng kí khoá học, dịch vụ, hoặc đặt câu hỏi cho Trade Tech\n/me - xem danh sách khoá học đã đăng kí\n/cancel - thông báo muốn huỷ đăng kí\n/help - hiện hướng dẫn.")
	log.Println("now m.Text:", m.Text)
}

// Chose Button
func sendCaseButton(b *tb.Bot, m *tb.Message, text string) (*tb.Message, error) {
	button1 := tb.ReplyButton{Text: "DApp"}
	button2 := tb.ReplyButton{Text: "Another"}

	caseButton := [][]tb.ReplyButton{[]tb.ReplyButton{button1, button2}}
	return b.Send(m.Sender, text, &tb.ReplyMarkup{
		ReplyKeyboard:       caseButton,
		ResizeReplyKeyboard: true,
		OneTimeKeyboard:     true,
		ReplyKeyboardRemove: true,
	})
}

// Remove Button

func sendAndRemoveButton(b *tb.Bot, m *tb.Message, text string) (*tb.Message, error) {
	return b.Send(m.Sender, text, &tb.ReplyMarkup{
		ReplyKeyboardRemove: true,
	})
}

// Chose and remove keyBoard

// Chose lesson and show to telegram
// Step1: chose Lesson
// func choseLesson(b *tb.Bot, m *tb.Message) {
// 	send(b, m, "Thông tin khác:")
// 	log.Println("now m.Text:", m.Text)
// 	studentInfor[0] = m.Text
// }

// Step 0	 : Start Sign Up
func startSignUp(b *tb.Bot, m *tb.Message) {
	checkRegister(b, m, text)
	log.Println(" m.Text:", m.Text)
}

// Step 1: Infor mre
func checkRegister(b *tb.Bot, m *tb.Message, text string) (*tb.Message, error) {
	if m.Text == "/register" {
		yesSir = true
		text = "Vui lòng chọn.\nDApp - Khoá học Ethereum DApp Development\nKhác - Các yêu cầu khác"
		return sendCaseButton(b, m, text)
	} else {
		yesSir = false
		return send(b, m, " Xin lỗi, bạn vui lòng gõ /register")
	}
}

// // Step 3 : get Name
// func startConFirmName(b *tb.Bot, m *tb.Message) {
// 	studentInfor[1] = m.Text
// 	sendAndCheck(b, m, checkSpellName(m))
// 	log.Println("getName m.Text:", m.Text)

// }

//
// func checkSpellName(m *tb.Message) string {
// 	if m.Text != "" {
// 		yesSir = true
// 		return "Vui lòng nhập số điện thoại của bạn:"
// 	} else {
// 		yesSir = false
// 		return "Xin lỗi bạn đã gõ sai, vui lòng nhập lại"
// 	}

// }

// // Step 4 : get Phone Number
// func startConFirmNumber(b *tb.Bot, m *tb.Message) {
// 	studentInfor[2] = m.Text
// 	sendAndCheck(b, m, checkSpellNumber(m))
// 	log.Println("now m.Text:", m.Text)
// }

// // check phone
// func checkSpellNumber(m *tb.Message) string {
// 	if checkPhoneTrue(m) == true {
// 		yesSir = true
// 		return "Vui lòng nhập năm sinh của bạn:"
// 	} else {
// 		yesSir = false
// 		return "Xin lỗi bạn đã gõ sai, vui lòng nhập lại"
// 	}
// }

// // Step 5 : get yearold
// func startConFirmYearOld(b *tb.Bot, m *tb.Message) {
// 	studentInfor[3] = m.Text
// 	sendAndCheck(b, m, checkConfirmYear(m))
// 	log.Println("now m.Text:", m.Text)

// }

// // check yearold
// func checkConfirmYear(m *tb.Message) string {

// 	if checkYearOld(m) == true {
// 		yesSir = true
// 		return "Cảm ơn bạn đã đăng kí, vui lòng bấm OK để tiếp tục"
// 	} else {
// 		yesSir = false
// 		return "Xin lỗi bạn đã gõ sai, vui lòng gõ lại"
// 	}

// }

// Step 2 : Chose other
func choseOther(b *tb.Bot, m *tb.Message) {
	studentInfor[1] = m.Text
	if m.Text == "Another" {
		yesSir = true
		send(b, m, "Vui lòng nhập thông tin thêm:")
	} else {
		yesSir = false
		sendInfo(b, m)
	}
}

// Step 3: another Infor
func anotherInfor(b *tb.Bot, m *tb.Message) {
	sendInfo(b, m)
}

func sendInfo(b *tb.Bot, m *tb.Message) {
	studentInfor[5] = m.Sender.Username
	log.Println("now m.Text:", m.Sender.ID)

	if yesSir == true {
		studentInfor[2] = m.Text
	} else {
		studentInfor[2] = "Không có"
	}

	studentInfor[3] = strconv.Itoa(m.Sender.ID)
	text = "Có đăng kí mới!" + "\nTừ:" + "@" + studentInfor[5] + "\n" + "Khóa học:" + studentInfor[1] + "\nThông tin thêm:" + studentInfor[2]
	sendAndRemoveButton(b, m, "Cảm ơn bạn đã đăng kí khóa học tại trà đá lab, nhân viên chúng tôi sẽ liên lạc với bạn sau ít phút.")
	// td.InsertToDatbase(studentInfor)
	sendToMe(b, m, text)

	for i := 0; i < 6; i++ {
		studentInfor[i] = ""
	}
}

// Send to admin
func sendToMe(b *tb.Bot, m *tb.Message, text string) (*tb.Message, error) {
	// ID anh Thy
	m.Sender.ID = 144242180

	for i := 0; i < len(studentInfor); i++ {
		fmt.Println(" student info : ", i, studentInfor[i])
	}

	m.Sender.ID = adminID
	return b.Send(m.Sender, text)
}

//End
func end(b *tb.Bot, m *tb.Message) {
	var text string = "hãy chờ Admin trong ít phút"
	b.Send(m.Sender, text)
}

// func checkSpellSendAdmin(m *tb.Message) string {
// 	return "Cảm ơn bạn. Nhân viên của Trada Tech sẽ liên lạc lại với bạn để hỏi thêm chi tiết."
// }

// // Check phone true
// func checkPhoneTrue(m *tb.Message) bool {

// 	var text string
// 	text = m.Text

// 	var tran int
// 	tran, err := strconv.Atoi(text)
// 	if err != nil {
// 		return false
// 	} else {
// 		log.Println(tran)
// 		return true
// 	}

// }

// Check yearold true
// func checkYearOld(m *tb.Message) bool {
// 	var text string
// 	text = m.Text

// 	yearOld, err := strconv.Atoi(text)
// 	checkErr(err)

// 	if yearOld > 2000 || yearOld < 1920 {
// 		return false
// 	} else {
// 		return true
// 	}

// }

// Get Mess and rerturn to m.Text
func sendAndCheck(b *tb.Bot, m *tb.Message, text string) (*tb.Message, error) {
	return b.Send(m.Sender, text)
}

// Các lệnh ghi chú:
// Handle with /help
func helpFunc(b *tb.Bot, m *tb.Message) {
	send(b, m, "/register - đăng kí khoá học, dịch vụ, hoặc đặt câu hỏi cho Trada Tech\n/me - xem danh sách khoá học đã đăng kí\n/cancel - thông báo muốn huỷ đăng kí\n/help - hiện hướng dẫn.")
}

// Handle with /me
func meFunc(b *tb.Bot, m *tb.Message) {
	k := td.CancelLesson(strconv.Itoa(m.Sender.ID))

	if k == "" {
		b.Send(m.Sender, "Bạn chưa đăng kí khóa học nào")
	} else {
		b.Send(m.Sender, "Khóa học mọi bạn đăng kí là : ")
		b.Send(m.Sender, k)
	}

}

// Handle with /infor
func infoFunc(b *tb.Bot, m *tb.Message) {
	b.Send(m.Sender, "Thi măng cụt: \n 1. Đào tạo Ethereum Dapp Developer \n2. Đào tạo theo yêu cầu riêng của công ty \n3. Tư vấn về phát triển và ứng dụng blockchain")
}

// Handle with /Cancel
func cancelFunc(b *tb.Bot, m *tb.Message) {
	k := td.CancelLesson(strconv.Itoa(m.Sender.ID))

	if k == "" {
		b.Send(m.Sender, "Bạn chưa đăng kí khóa học nào nên không cần hủy")
	} else {
		b.Send(m.Sender, "Yêu cầu của bạn được xác nhận , chúng tôi sẽ gửi tới Admin")
	}

}

// Các lệnh hệ thống.
// Send Mess
func send(b *tb.Bot, m *tb.Message, text string) (*tb.Message, error) {
	return b.Send(m.Sender, text)
}

// Check User
func checkUser(b *tb.Bot, m *tb.Message) {
	newUserInfo := userInfo{thisStep: step0}
	userMap[m.Sender.ID] = &newUserInfo

	startAsk(b, m)
}

// Check err
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetFullName
func getFullName(b *tb.Bot, m *tb.Message) string {
	if user, ok := userMap[m.Sender.ID]; ok {
		if user.displayName == "" {
			return user.displayName
		}
	}
	return fmt.Sprintf("%s %s", m.Sender.FirstName, m.Sender.LastName)
}
