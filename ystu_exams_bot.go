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
	checkBuyCurs = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить✅", "checkOKCurs"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Отклонить❌", "checkBAD"),
		),
	)
	checkBuyMathCurs = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить✅", "checkOKMathCurs"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Отклонить❌", "checkBAD"),
		),
	)
	payMathCurs = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💸 Оплатить", "payMathCurs"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Назад", "backMenu"),
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
	payCurs = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💸 Оплатить", "payCurs"),
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
			tgbotapi.NewInlineKeyboardButtonData("🎁 Получить подарок", "podarok"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("❓ Задать вопрос", "https://t.me/micsemal"),
		),
	)
	menuMath = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✍️[LITE] Ответы на вопросы к экзамену", "otvetyMath"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎓[PRO] Курс с практикой и ДЗ", "cursMath"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🧠[VIP] Ответы на вопросы + курс", "mathcurs"),
		),
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
				if v[count] == "payMathCurs" {
					msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName + "\nТовар: Ответы по математике (преп. Ройтенберг В.М.) + курс с практикой и ДЗ"
					msg.ReplyMarkup = checkBuyMathCurs
				}
				if v[count] == "payCurs" {
					msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName + "\nТовар: Курс по матану с практикой и ДЗ"
					msg.ReplyMarkup = checkBuyCurs
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
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вас приветствует бот *YSTU EXAMS*👋\nЭкзамены уже очень скоро! Поэтому, чтобы не терять время, ты можешь приобрести ответы на экзамены по Математике + курс по всему 2 семестру с теорией и разборами заданий из билетов прошлых лет + ДЗ! 🥰")
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
					"Выбери подходящий тариф:",
					menuMath,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "mathcurs":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Ответы на экзамен по математике (преп. Ройтенберг В.Ш.) + курс с практикой и ДЗ*\nЦена: 1500 рублей",
					payMathCurs,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "otvetyMath":
				edit2 := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Ответы на экзамен по математике (преп. Ройтенберг В.Ш.)*\nЦена: 900 рублей",
					payMath,
				)
				edit2.ParseMode = "Markdown"

				if _, err := bot.Send(edit2); err != nil {
					panic(err)
				}
			case "cursMath":
				edit3 := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Курс по матану с практикой и ДЗ*\nЦена: 1000 рублей",
					payCurs,
				)
				edit3.ParseMode = "Markdown"

				if _, err := bot.Send(edit3); err != nil {
					panic(err)
				}

			case "checkOKMath":
				links := []string{"https://t.me/+Hm7JCKMS5p80YzQy","https://t.me/+i-1kWLSHbwxiMDU6","https://t.me/+Td6ZKPl55aRlMDk6","https://t.me/+qebWRLDdgrRhYjQy","https://t.me/+Ns6QuCVNsYgyMzRi","https://t.me/+fTlftaWs5IUyMjZi","https://t.me/+zkaabqZRQWAwMzdi","https://t.me/+FkrIykhzoA44Mzcy","https://t.me/+Yz-wlXzePU9jNWVi","https://t.me/+uRIynb6FGWVmNzhi","https://t.me/+ssEyrDQ4ZE4zNTgy","https://t.me/+jvVku9n0aKsxN2Ey","https://t.me/+RNSB52nh4s00YjQ6","https://t.me/+yjEejkckmiY5MmRi","https://t.me/+exfDaoih2yk3NmQy","https://t.me/+k-E6QMv_hnI1YjUy","https://t.me/+e66fLWh5l7I2MjUy","https://t.me/+Lq38ZJWWD7Y3Njk6","https://t.me/+86lS81a67Q00Nzky","https://t.me/+LMLNT4-Ka385ZmQy","https://t.me/+LwOlS5plmRdiNzgy","https://t.me/+GyOnbIQqU_Y3NmVi","https://t.me/+RadF3gvTR-M0ZGEy","https://t.me/+V0gXSsGj1eEwZGJi","https://t.me/+d1ma0Y0ykns3ZmUy","https://t.me/+WsQyULJTG2IwNzhi","https://t.me/+j2D2qjr53_c5NTFi","https://t.me/+MTBA9_IPBVk0OGUy","https://t.me/+jGJa2LXjredhNTBi","https://t.me/+oCucPaT8NDg2NzBi","https://t.me/+aSBmf3iSm-kyYmRi","https://t.me/+LtU7I9IfLY5kNTli","https://t.me/+17grbK7W6zZmYzYy","https://t.me/+0xWmHB5dP383Y2Yy","https://t.me/+chs8DyD8WrhjM2Ni","https://t.me/+Q9i_D8QOGBI5YjEy","https://t.me/+YyNsuuo_vTM5ZmNi","https://t.me/+RY5WPjJ_8P4xOWNi","https://t.me/+bpGO7BvA7SliNjBi","https://t.me/+47C6RsBZWD9kOGEy","https://t.me/+QNBoJXbdoJw5OGUy","https://t.me/+NwrKLqmbMWkxZDAy","https://t.me/+rDgB6OO55Vc0MWMy","https://t.me/+hZagTFEv4b4wYTky","https://t.me/+mfeZypIvd9U5MGUy"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[m])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				m++
			

			case "checkOKCurs":
				links := []string{"https://t.me/+khhLP4SViY4zYWVi","https://t.me/+lHnldtl2-IRjMDFi","https://t.me/+fUGO6xTFdDgwYTUy","https://t.me/+d57RhkHGYHJhZWZi","https://t.me/+1QNUY13PMK9hZDZi","https://t.me/+PvSbNvacWmpkMjYy","https://t.me/+QYz6qoYJ_iQyYjZi","https://t.me/+utItj_BLubAxOTY6","https://t.me/+sKk1Dj-YX5U2ZTNi","https://t.me/+q8YswIRdA6E1YWU6","https://t.me/+tXFO4WMM3IlhZTAy","https://t.me/+uIJ7IeO_cgljYzcy","https://t.me/+cliu8tX01i00MjAy","https://t.me/+yVPeV6uw49k3Y2Uy","https://t.me/+1fQ2u2K1839hNTU6","https://t.me/+qUdtE3pDZXA2MDky","https://t.me/+4rCSvSLNUGkxYmUy","https://t.me/+nBoMyYWghbthNTVi","https://t.me/+Vz7lsVnlT5xlYjVi","https://t.me/+r3ZUmivUnJw0YWU6","https://t.me/+rEQnA134Zlk2ZmUy","https://t.me/+lOnLgooQarllZDQy","https://t.me/+axnhKSP-WghkM2Ni","https://t.me/+RedElvwOCjs2YjAy","https://t.me/+UdkkzGAMqRU4YzM6","https://t.me/+9SznMfzebXBmN2Qy","https://t.me/+flqrenV3vBsxY2Iy","https://t.me/+-8VZid8sPR02NTZi","https://t.me/+fXeC4JU6Cm9lYjAy","https://t.me/+reb5zGMdniU4OTEy","https://t.me/+jPrPGCj57rw3Y2Ji","https://t.me/+KlQF4noTUAI4M2Ri","https://t.me/+b3W9Tb5wNxk0MjEy","https://t.me/+U60JCPpkYXY4ZmEy","https://t.me/+V7dk8CKnPxsxYjcy","https://t.me/+JynIZqBj99UzNDNi","https://t.me/+kONOvv1sqaZhYmQy","https://t.me/+OomTRpHWE1w3MmQy","https://t.me/+kDpCn-vsp0UwYTBi","https://t.me/+kMDP6qFgZzRjOTky","https://t.me/+xc5mwza8tHFkNWMy","https://t.me/+71PDJGocNkRmZjk6","https://t.me/+tS8tcysxAigzMzli","https://t.me/+6vkNfg64XMw3MzUy","https://t.me/+hKfMPCCgi3sxOGUy"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[a])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				a++

			case "checkOKMathCurs":
				links1 := []string{"https://t.me/+Hm7JCKMS5p80YzQy","https://t.me/+i-1kWLSHbwxiMDU6","https://t.me/+Td6ZKPl55aRlMDk6","https://t.me/+qebWRLDdgrRhYjQy","https://t.me/+Ns6QuCVNsYgyMzRi","https://t.me/+fTlftaWs5IUyMjZi","https://t.me/+zkaabqZRQWAwMzdi","https://t.me/+FkrIykhzoA44Mzcy","https://t.me/+Yz-wlXzePU9jNWVi","https://t.me/+uRIynb6FGWVmNzhi","https://t.me/+ssEyrDQ4ZE4zNTgy","https://t.me/+jvVku9n0aKsxN2Ey","https://t.me/+RNSB52nh4s00YjQ6","https://t.me/+yjEejkckmiY5MmRi","https://t.me/+exfDaoih2yk3NmQy","https://t.me/+k-E6QMv_hnI1YjUy","https://t.me/+e66fLWh5l7I2MjUy","https://t.me/+Lq38ZJWWD7Y3Njk6","https://t.me/+86lS81a67Q00Nzky","https://t.me/+LMLNT4-Ka385ZmQy","https://t.me/+LwOlS5plmRdiNzgy","https://t.me/+GyOnbIQqU_Y3NmVi","https://t.me/+RadF3gvTR-M0ZGEy","https://t.me/+V0gXSsGj1eEwZGJi","https://t.me/+d1ma0Y0ykns3ZmUy","https://t.me/+WsQyULJTG2IwNzhi","https://t.me/+j2D2qjr53_c5NTFi","https://t.me/+MTBA9_IPBVk0OGUy","https://t.me/+jGJa2LXjredhNTBi","https://t.me/+oCucPaT8NDg2NzBi","https://t.me/+aSBmf3iSm-kyYmRi","https://t.me/+LtU7I9IfLY5kNTli","https://t.me/+17grbK7W6zZmYzYy","https://t.me/+0xWmHB5dP383Y2Yy","https://t.me/+chs8DyD8WrhjM2Ni","https://t.me/+Q9i_D8QOGBI5YjEy","https://t.me/+YyNsuuo_vTM5ZmNi","https://t.me/+RY5WPjJ_8P4xOWNi","https://t.me/+bpGO7BvA7SliNjBi","https://t.me/+47C6RsBZWD9kOGEy","https://t.me/+QNBoJXbdoJw5OGUy","https://t.me/+NwrKLqmbMWkxZDAy","https://t.me/+rDgB6OO55Vc0MWMy","https://t.me/+hZagTFEv4b4wYTky","https://t.me/+mfeZypIvd9U5MGUy"}
				links2 := []string{"https://t.me/+khhLP4SViY4zYWVi","https://t.me/+lHnldtl2-IRjMDFi","https://t.me/+fUGO6xTFdDgwYTUy","https://t.me/+d57RhkHGYHJhZWZi","https://t.me/+1QNUY13PMK9hZDZi","https://t.me/+PvSbNvacWmpkMjYy","https://t.me/+QYz6qoYJ_iQyYjZi","https://t.me/+utItj_BLubAxOTY6","https://t.me/+sKk1Dj-YX5U2ZTNi","https://t.me/+q8YswIRdA6E1YWU6","https://t.me/+tXFO4WMM3IlhZTAy","https://t.me/+uIJ7IeO_cgljYzcy","https://t.me/+cliu8tX01i00MjAy","https://t.me/+yVPeV6uw49k3Y2Uy","https://t.me/+1fQ2u2K1839hNTU6","https://t.me/+qUdtE3pDZXA2MDky","https://t.me/+4rCSvSLNUGkxYmUy","https://t.me/+nBoMyYWghbthNTVi","https://t.me/+Vz7lsVnlT5xlYjVi","https://t.me/+r3ZUmivUnJw0YWU6","https://t.me/+rEQnA134Zlk2ZmUy","https://t.me/+lOnLgooQarllZDQy","https://t.me/+axnhKSP-WghkM2Ni","https://t.me/+RedElvwOCjs2YjAy","https://t.me/+UdkkzGAMqRU4YzM6","https://t.me/+9SznMfzebXBmN2Qy","https://t.me/+flqrenV3vBsxY2Iy","https://t.me/+-8VZid8sPR02NTZi","https://t.me/+fXeC4JU6Cm9lYjAy","https://t.me/+reb5zGMdniU4OTEy","https://t.me/+jPrPGCj57rw3Y2Ji","https://t.me/+KlQF4noTUAI4M2Ri","https://t.me/+b3W9Tb5wNxk0MjEy","https://t.me/+U60JCPpkYXY4ZmEy","https://t.me/+V7dk8CKnPxsxYjcy","https://t.me/+JynIZqBj99UzNDNi","https://t.me/+kONOvv1sqaZhYmQy","https://t.me/+OomTRpHWE1w3MmQy","https://t.me/+kDpCn-vsp0UwYTBi","https://t.me/+kMDP6qFgZzRjOTky","https://t.me/+xc5mwza8tHFkNWMy","https://t.me/+71PDJGocNkRmZjk6","https://t.me/+tS8tcysxAigzMzli","https://t.me/+6vkNfg64XMw3MzUy","https://t.me/+hKfMPCCgi3sxOGUy"}
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
					"🤑Оплата 900 руб. на +79997971960 СБП (ВТБ).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}

			case "payCurs":
				count++
				v = append(v, "payCurs")
				edit := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"🤑Оплата 1000 руб. на +79997971960 СБП (ВТБ).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "payMathCurs":
				count++
				v = append(v, "payMathCurs")
				edit := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"🤑Оплата 1500 руб. на +79997971960 СБП (ВТБ).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "podarok":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Держи по одному бесплатному ответу на каждый предмет. Убедись в качестве и забери полный комплект!🥰")
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
				// Открываем PDF-файл
				pdfFile1, err := os.Open("27 вопрос (АИП).pdf")
				if err != nil {
					log.Panic(err)
				}
				defer pdfFile1.Close()

				// Создаём документ для отправки
				doc1 := tgbotapi.NewDocument(update.CallbackQuery.Message.Chat.ID, tgbotapi.FileReader{
					Name:   "27 вопрос (АИП).pdf",
					Reader: pdfFile1,
				})

				// Отправляем файл пользователю
				if _, err := bot.Send(doc1); err != nil {
					log.Panic(err)
				}
				// Открываем PDF-файл
				pdfFile2, err := os.Open("Вопрос 2 (ВМ2).pdf")
				if err != nil {
					log.Panic(err)
				}
				defer pdfFile2.Close()

				// Создаём документ для отправки
				doc2 := tgbotapi.NewDocument(update.CallbackQuery.Message.Chat.ID, tgbotapi.FileReader{
					Name:   "Вопрос 2 (ВМ2).pdf",
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
					"Вас приветствует бот *YSTU EXAMS*👋\nЭкзамены уже очень скоро! Поэтому, чтобы не терять время, ты можешь приобрести ответы на экзамены по Математике + курс по всему 2 семестру с теорией и разборами заданий из билетов прошлых лет + ДЗ! 🥰",
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
