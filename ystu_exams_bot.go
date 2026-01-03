package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

const adminID int64 = 1283075660

var (
	checkBuyMath = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить✅", "checkOKMath"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Отклонить❌", "checkBAD"),
		),
	)
	checkBuyAIP = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить✅", "checkOKAIP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Отклонить❌", "checkBAD"),
		),
	)
	checkBuyMathAIP = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить✅", "checkOKMathAIP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Отклонить❌", "checkBAD"),
		),
	)
	payMath = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💸 Оплатить", "payMath"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Назад", "backMenu"),
		),
	)
	payAIP = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💸 Оплатить", "payAIP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Назад", "backMenu"),
		),
	)
	payMathAIP = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💸 Оплатить", "payMathAIP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Назад", "backMenu"),
		),
	)	

	menuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📐 Математика", "menuMath"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💻 Алгоритмизация и программирование", "menuAIP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💻 Комплект Математика + АИП", "menuMathAIP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎁 Получить подарок", "podarok"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("❓ Задать вопрос", "https://t.me/micsemal"),
		),
	)
	menuMath = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✍️Ответы на вопросы к экзамену", "otvetyMath"),
		)
	)
	menuAIP = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✍️Ответы на вопросы к экзамену", "otvetyAIP"),
		)
	)
	menuMathAIP = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✍️Ответы на вопросы к экзамену по Математике + АИП", "otvetyAIP"),
		)
	)	
)

var m int = 0
var a int = 0
var v []string
var count int = -1
var chatId []int64
var p int = -1

func main() {
	bot, err := tgbotapi.NewBotAPI("8140603482:AAFYnRZdxm-QOzTK5AOSJZ3l2ouBQCZJJsA")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			// Обработка фото
			if update.Message.Photo != nil {
				// Пересылаем фото админу
				photo := update.Message.Photo[len(update.Message.Photo)-1] // Берем фото в максимальном разрешении
				msg := tgbotapi.NewPhoto(adminID, tgbotapi.FileID(photo.FileID))
				if v[count] == "payMath" {
					msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName + "\nТовар: Ответы по математике (преп. Ройтенберг В.М.)"
					msg.ReplyMarkup = checkBuyMath
				}
				if v[count] == "payAIP" {
					msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName + "\nТовар: Ответы по АИП (преп. Никитина Т.П.)"
					msg.ReplyMarkup = checkBuyAIP
				}
				if v[count] == "payMathAIP" {
					msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName + "\nТовар: Ответы по Математике + АИП"
					msg.ReplyMarkup = checkBuyMathAIP
				}
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}

				// Ответ пользователю
				chatId = append(chatId, update.Message.Chat.ID)
				p++
				reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Скриншот отправлен администратору✅. Ждите подтверждения...")
				if _, err := bot.Send(reply); err != nil {
					log.Panic(err)
				}
			} else {
				if update.Message.Command() != "start" {
					// Ответ на другие сообщения
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, отправьте скриншот об оплате.")
					if _, err := bot.Send(msg); err != nil {
						log.Panic(err)
					}
				}
			}

			// Обработка команды /start
			if update.Message.Command() == "start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вас приветствует бот *YSTU EXAMS*👋\nЭкзамены уже очень скоро! Поэтому, чтобы не терять время, ты можешь приобрести ответы на экзамены по Математике, а также Алгоритмизации и программированию! 🥰")
				msg.ParseMode = "Markdown"
				msg.ReplyMarkup = menuKeyboard

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			switch update.CallbackQuery.Data {
			case "menuMath":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Выбери действие:",
					menuMath,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}

			case "menuAIP":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Выбери действие:",
					menuAIP,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "menuMathAIP":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Выбери действие:",
					menuMathAIP,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}				
			case "otvetyMath":
				edit2 := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Ответы на экзамен по математике (преп. Ройтенберг В.Ш.)*\nЦена: 800 рублей",
					payMath,
				)
				edit2.ParseMode = "Markdown"

				if _, err := bot.Send(edit2); err != nil {
					panic(err)
				}
			case "otvetyAIP":
				edit3 := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Ответы на экзамен по Алгоритмизации и программированию (преп. Никитина Т.П.)*\nЦена: 600 рублей",
					payAIP,
				)
				edit3.ParseMode = "Markdown"

				if _, err := bot.Send(edit3); err != nil {
					panic(err)
				}

			case "checkOKMath":
				links := []string{"https://t.me/+MMcfy-nkA0tjZTli","https://t.me/+YIie7weL3qlmYmIy","https://t.me/+SeJ2vbADeWYwNDky","https://t.me/+ZcOwwyD46QpkN2U6","https://t.me/+0gQ8ilcuaBszNGM6","https://t.me/+P1nxlOfXg-0zMTYy","https://t.me/+ZIv6DFxrFNVkYzRi","https://t.me/+xEzO8RpatQ5iYzhi","https://t.me/+9hX-Th4499I3ZTcy","https://t.me/+oAacitUxbhMyZDgy","https://t.me/+SfDmXj1PB0diNzVi","https://t.me/+xKkvko1wjJ5iM2Zi","https://t.me/+81ck2IwQGtcxMmI6","https://t.me/+UZJHoU7kszw0MTNi","https://t.me/+D_PpuwEL_QhhOGY6","https://t.me/+eau2YNTPGSE2NGFi","https://t.me/+yzQ6LUFYrRRhMmFi","https://t.me/+kzU-y5dgpuc5OWIy","https://t.me/+PCk0csw1pbQ0NGUy","https://t.me/+gTCt7I_Fga03OTky","https://t.me/+-qG9K8yHwuE0YTc6","https://t.me/+fUxnfoR6-lYwNDEy","https://t.me/+6pHfQpDbp7JjYWIy","https://t.me/+t5ye9pfQ4nA4MDky","https://t.me/+h33aP4qxRHA3YzJi","https://t.me/+CGzkc6K5GCY0Yjgy","https://t.me/+T0nEVidEfnk1OTIy","https://t.me/+Arm0GYXRVMY2ZDU6","https://t.me/+VHxJFX4tjM9lZWNi","https://t.me/+T8bVf6Eq7vU4ZTgy","https://t.me/+Iri6yO5mnV1jOGIy","https://t.me/+7ePy-wx1sZA5ZGEy","https://t.me/+UdT8JddZr2xiMTI6","https://t.me/+a71cPIWH4IMyMDYy","https://t.me/+9vjJPVKttbViMmRi","https://t.me/+TKl4SlqEYmBhOTIy","https://t.me/+uKW5HPal2Z4xNWYy","https://t.me/+AMTKxuZtpns2ZmYy","https://t.me/+uPaV8cYQXPUyODI6","https://t.me/+vjrbgDjuWrU4YTUy","https://t.me/+8CmSyh-omipjY2Zi","https://t.me/+tsuJSfAYqTwwMmFi","https://t.me/+u-aw-qGY-k03NjQy","https://t.me/+qNXszpEojgdjOWEy","https://t.me/+xsPt_J_6zYdjZDJi"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[m])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				m++
			

			case "checkOKAIP":
				links := []string{"https://t.me/+lAYWxD2U8REyNWEy","https://t.me/+HIpzzsjxQygxNmVi","https://t.me/+1Tei-w_LV8NiOWMy","https://t.me/+qTIKYFIvlkE0NmNi","https://t.me/+rfBnWARq5b1lOTIy","https://t.me/+Zj3VzttakuIzYzky","https://t.me/+DXI6kcQMJoExNGEy","https://t.me/+KArZWFTvUjllZjFi","https://t.me/+N-tLEKiAMFliYmJi","https://t.me/+qf_m3PXQuTUyNzMy","https://t.me/+96xeqBp2oBg3YjE6","https://t.me/+aM9H7sfdV1wwOTFi","https://t.me/+X-21OFAhL-tjYWVi","https://t.me/+4082inrBcfpmYjli","https://t.me/+gvudEyhtdXthMWUy","https://t.me/+N6TL8x48ZshiMGQy","https://t.me/+kZx7YI45CYE2Yjc6","https://t.me/+sgexPZsmK0M0ODcy","https://t.me/+9Pq2wfmNkgw0MGQy","https://t.me/+ISp_QRBQlUQ3Mzc6","https://t.me/+n4sH_ULlaEw4Njhi","https://t.me/+FEUBdOLBJMtlZDQy","https://t.me/+tdHaUCar_hM2YmNi","https://t.me/+iMeRerN2hGY1ODEy","https://t.me/+gwUCBrpqdTw2OWUy","https://t.me/+kqKclayN6StlZjQy","https://t.me/+Zlu6UnWEcwg0ZGQ6","https://t.me/+TaG2ABTtg5c3ODg6","https://t.me/+IrSDJX-lSYpjNGZi","https://t.me/+dhsaWgZm0d43NDQ6","https://t.me/+6WdVpimqyIIyNzRi","https://t.me/+3HW5OK6xDDgwYzMy","https://t.me/+h18FGcTjkARhNGUy","https://t.me/+xs0hAWHVmIA4OWJi","https://t.me/+DCzx2NRX5UFmNTNi","https://t.me/+nCrR18y8D9NjNWJi","https://t.me/+l9_z-EGpmOszYzVi","https://t.me/+qcLI841cppgxMmEy","https://t.me/+Xy9M77I7cbVkMmNi","https://t.me/+UwlTTjjk78k0OGRi","https://t.me/+m1D_v-tPI_U4YjBi","https://t.me/+EqAMMnE4H85iZGRi","https://t.me/+gwFVHdcTcEEyOWQy","https://t.me/+8E0NuR5M62U1MTYy","https://t.me/+8E0NuR5M62U1MTYy"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[a])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				a++

			case "checkOKMathAIP":
				links1 := []string{"https://t.me/+MMcfy-nkA0tjZTli","https://t.me/+YIie7weL3qlmYmIy","https://t.me/+SeJ2vbADeWYwNDky","https://t.me/+ZcOwwyD46QpkN2U6","https://t.me/+0gQ8ilcuaBszNGM6","https://t.me/+P1nxlOfXg-0zMTYy","https://t.me/+ZIv6DFxrFNVkYzRi","https://t.me/+xEzO8RpatQ5iYzhi","https://t.me/+9hX-Th4499I3ZTcy","https://t.me/+oAacitUxbhMyZDgy","https://t.me/+SfDmXj1PB0diNzVi","https://t.me/+xKkvko1wjJ5iM2Zi","https://t.me/+81ck2IwQGtcxMmI6","https://t.me/+UZJHoU7kszw0MTNi","https://t.me/+D_PpuwEL_QhhOGY6","https://t.me/+eau2YNTPGSE2NGFi","https://t.me/+yzQ6LUFYrRRhMmFi","https://t.me/+kzU-y5dgpuc5OWIy","https://t.me/+PCk0csw1pbQ0NGUy","https://t.me/+gTCt7I_Fga03OTky","https://t.me/+-qG9K8yHwuE0YTc6","https://t.me/+fUxnfoR6-lYwNDEy","https://t.me/+6pHfQpDbp7JjYWIy","https://t.me/+t5ye9pfQ4nA4MDky","https://t.me/+h33aP4qxRHA3YzJi","https://t.me/+CGzkc6K5GCY0Yjgy","https://t.me/+T0nEVidEfnk1OTIy","https://t.me/+Arm0GYXRVMY2ZDU6","https://t.me/+VHxJFX4tjM9lZWNi","https://t.me/+T8bVf6Eq7vU4ZTgy","https://t.me/+Iri6yO5mnV1jOGIy","https://t.me/+7ePy-wx1sZA5ZGEy","https://t.me/+UdT8JddZr2xiMTI6","https://t.me/+a71cPIWH4IMyMDYy","https://t.me/+9vjJPVKttbViMmRi","https://t.me/+TKl4SlqEYmBhOTIy","https://t.me/+uKW5HPal2Z4xNWYy","https://t.me/+AMTKxuZtpns2ZmYy","https://t.me/+uPaV8cYQXPUyODI6","https://t.me/+vjrbgDjuWrU4YTUy","https://t.me/+8CmSyh-omipjY2Zi","https://t.me/+tsuJSfAYqTwwMmFi","https://t.me/+u-aw-qGY-k03NjQy","https://t.me/+qNXszpEojgdjOWEy","https://t.me/+xsPt_J_6zYdjZDJi"}
				links2 := []string{"https://t.me/+lAYWxD2U8REyNWEy","https://t.me/+HIpzzsjxQygxNmVi","https://t.me/+1Tei-w_LV8NiOWMy","https://t.me/+qTIKYFIvlkE0NmNi","https://t.me/+rfBnWARq5b1lOTIy","https://t.me/+Zj3VzttakuIzYzky","https://t.me/+DXI6kcQMJoExNGEy","https://t.me/+KArZWFTvUjllZjFi","https://t.me/+N-tLEKiAMFliYmJi","https://t.me/+qf_m3PXQuTUyNzMy","https://t.me/+96xeqBp2oBg3YjE6","https://t.me/+aM9H7sfdV1wwOTFi","https://t.me/+X-21OFAhL-tjYWVi","https://t.me/+4082inrBcfpmYjli","https://t.me/+gvudEyhtdXthMWUy","https://t.me/+N6TL8x48ZshiMGQy","https://t.me/+kZx7YI45CYE2Yjc6","https://t.me/+sgexPZsmK0M0ODcy","https://t.me/+9Pq2wfmNkgw0MGQy","https://t.me/+ISp_QRBQlUQ3Mzc6","https://t.me/+n4sH_ULlaEw4Njhi","https://t.me/+FEUBdOLBJMtlZDQy","https://t.me/+tdHaUCar_hM2YmNi","https://t.me/+iMeRerN2hGY1ODEy","https://t.me/+gwUCBrpqdTw2OWUy","https://t.me/+kqKclayN6StlZjQy","https://t.me/+Zlu6UnWEcwg0ZGQ6","https://t.me/+TaG2ABTtg5c3ODg6","https://t.me/+IrSDJX-lSYpjNGZi","https://t.me/+dhsaWgZm0d43NDQ6","https://t.me/+6WdVpimqyIIyNzRi","https://t.me/+3HW5OK6xDDgwYzMy","https://t.me/+h18FGcTjkARhNGUy","https://t.me/+xs0hAWHVmIA4OWJi","https://t.me/+DCzx2NRX5UFmNTNi","https://t.me/+nCrR18y8D9NjNWJi","https://t.me/+l9_z-EGpmOszYzVi","https://t.me/+qcLI841cppgxMmEy","https://t.me/+Xy9M77I7cbVkMmNi","https://t.me/+UwlTTjjk78k0OGRi","https://t.me/+m1D_v-tPI_U4YjBi","https://t.me/+EqAMMnE4H85iZGRi","https://t.me/+gwFVHdcTcEEyOWQy","https://t.me/+8E0NuR5M62U1MTYy","https://t.me/+8E0NuR5M62U1MTYy"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК по математике "+links1[m]+" и ТГК с курсом по матану"+links2[a])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				m++
				a++

			case "checkBAD":
				msg := tgbotapi.NewMessage(chatId[p], "Оплата была отклонена❌ Попробуйте снова")
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			case "payMath":
				count++
				v = append(v, "payMath")
				edit := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"🤑Оплата 700 руб. на +79997971960 СБП (ВТБ).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}

			case "payAIP":
				count++
				v = append(v, "payAIP")
				edit := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"🤑Оплата 600 руб. на +79997971960 СБП (ВТБ).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "payMathAIP":
				count++
				v = append(v, "payMathAIP")
				edit := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"🤑Оплата 500 руб. на +79997971960 СБП (ВТБ).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "podarok":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Держи бесплатный ответ на вопрос по математике + по АИП. Убедись в качестве и забери полный комплект!🥰")
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
				// Открываем PDF-файл
				pdfFile1, err := os.Open("Домашнее задание по исследованию функций на экстремум.pdf")
				if err != nil {
					log.Panic(err)
				}
				defer pdfFile1.Close()

				// Создаём документ для отправки
				doc1 := tgbotapi.NewDocument(update.CallbackQuery.Message.Chat.ID, tgbotapi.FileReader{
					Name:   "Домашнее задание по исследованию функций на экстремум.pdf",
					Reader: pdfFile1,
				})

				// Отправляем файл пользователю
				if _, err := bot.Send(doc1); err != nil {
					log.Panic(err)
				}
				// Открываем PDF-файл
				pdfFile2, err := os.Open("32. Интегральный признак сходимости рядов с положительными членами. Ряд Дирихле и условия его сходимости..pdf")
				if err != nil {
					log.Panic(err)
				}
				defer pdfFile2.Close()

				// Создаём документ для отправки
				doc2 := tgbotapi.NewDocument(update.CallbackQuery.Message.Chat.ID, tgbotapi.FileReader{
					Name:   "32. Интегральный признак сходимости рядов с положительными членами. Ряд Дирихле и условия его сходимости..pdf",
					Reader: pdfFile2,
				})

				// Отправляем файл пользователю
				if _, err := bot.Send(doc2); err != nil {
					log.Panic(err)
				}
			case "backMenu":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Вас приветствует бот *YSTU EXAMS*👋\nЭкзамены уже очень скоро! Поэтому, чтобы не терять время, ты можешь приобрести ответы на экзамены по Математике + Алгоритмизации и программированию! 🥰",
					menuKeyboard,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			}
		}
	}
}
