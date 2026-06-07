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
					"*Ответы на экзамен по математике (преп. Ройтенберг В.Ш.) + курс с практикой и ДЗ*\nЦена: 1500 рублей\nПроверка ДЗ: + 500 рублей",
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
					"*Ответы на экзамен по математике (преп. Ройтенберг В.Ш.)*\nЦена: 800 рублей",
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
					"*Курс по матану с практикой и ДЗ*\nЦена: 1000 рублей\nПроверка ДЗ: + 500 рублей",
					payCurs,
				)
				edit3.ParseMode = "Markdown"

				if _, err := bot.Send(edit3); err != nil {
					panic(err)
				}

			case "checkOKMath":
				links := []string{"https://t.me/+jGJa2LXjredhNTBi", "https://t.me/+oCucPaT8NDg2NzBi", "https://t.me/+aSBmf3iSm-kyYmRi", "https://t.me/+LtU7I9IfLY5kNTli", "https://t.me/+17grbK7W6zZmYzYy", "https://t.me/+0xWmHB5dP383Y2Yy", "https://t.me/+chs8DyD8WrhjM2Ni", "https://t.me/+Q9i_D8QOGBI5YjEy", "https://t.me/+YyNsuuo_vTM5ZmNi", "https://t.me/+RY5WPjJ_8P4xOWNi", "https://t.me/+bpGO7BvA7SliNjBi", "https://t.me/+47C6RsBZWD9kOGEy", "https://t.me/+QNBoJXbdoJw5OGUy", "https://t.me/+NwrKLqmbMWkxZDAy", "https://t.me/+rDgB6OO55Vc0MWMy", "https://t.me/+hZagTFEv4b4wYTky", "https://t.me/+mfeZypIvd9U5MGUy", "https://t.me/+4qqGtywEQ5s5YjQy", "https://t.me/+qeJivwdWtOxkODMy", "https://t.me/+rNRv_UgdO_E3NTcy"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[m])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				m++
			

			case "checkOKCurs":
				links := []string{"https://t.me/+yUtP_1nVe_M2NTky", "https://t.me/+u6vYeC9anKg4MWQy", "https://t.me/+DCotD4Z8dM1jMTEy", "https://t.me/+SA94NpTjguoxY2Ri", "https://t.me/+IEAY3Fqmp0U5NmZi", "https://t.me/+QnBTcG0co3A0NmUy", "https://t.me/+u49tvM-x-_1hMjk6", "https://t.me/+1zYOvKlaZ2AzOTIy", "https://t.me/+72mY8PNWwuhkMjU6", "https://t.me/+UI4IDNTC-jllZmIy", "https://t.me/+VBz6JqLcBYxhZmQy", "https://t.me/+LNUv5XSCBngyMWY6", "https://t.me/+PZZgd6A0DF4zMmU6", "https://t.me/+mua_4CcpRK9hYTky", "https://t.me/+W40cuTOPdZs3MmFi", "https://t.me/+xeEEtiQ7euRhMDY6", "https://t.me/+2pWAeYuHhU1kYjdi", "https://t.me/+ShOu5KSlJwA1MDli", "https://t.me/+gFzXGO4p-p4xNTAy", "https://t.me/+oy8pToMIt-I3YWVi"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[a])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				a++

			case "checkOKMathCurs":
				links1 := []string{"https://t.me/+jGJa2LXjredhNTBi", "https://t.me/+oCucPaT8NDg2NzBi", "https://t.me/+aSBmf3iSm-kyYmRi", "https://t.me/+LtU7I9IfLY5kNTli", "https://t.me/+17grbK7W6zZmYzYy", "https://t.me/+0xWmHB5dP383Y2Yy", "https://t.me/+chs8DyD8WrhjM2Ni", "https://t.me/+Q9i_D8QOGBI5YjEy", "https://t.me/+YyNsuuo_vTM5ZmNi", "https://t.me/+RY5WPjJ_8P4xOWNi", "https://t.me/+bpGO7BvA7SliNjBi", "https://t.me/+47C6RsBZWD9kOGEy", "https://t.me/+QNBoJXbdoJw5OGUy", "https://t.me/+NwrKLqmbMWkxZDAy", "https://t.me/+rDgB6OO55Vc0MWMy", "https://t.me/+hZagTFEv4b4wYTky", "https://t.me/+mfeZypIvd9U5MGUy", "https://t.me/+4qqGtywEQ5s5YjQy", "https://t.me/+qeJivwdWtOxkODMy", "https://t.me/+rNRv_UgdO_E3NTcy"}
				links2 := []string{"https://t.me/+yUtP_1nVe_M2NTky", "https://t.me/+u6vYeC9anKg4MWQy", "https://t.me/+DCotD4Z8dM1jMTEy", "https://t.me/+SA94NpTjguoxY2Ri", "https://t.me/+IEAY3Fqmp0U5NmZi", "https://t.me/+QnBTcG0co3A0NmUy", "https://t.me/+u49tvM-x-_1hMjk6", "https://t.me/+1zYOvKlaZ2AzOTIy", "https://t.me/+72mY8PNWwuhkMjU6", "https://t.me/+UI4IDNTC-jllZmIy", "https://t.me/+VBz6JqLcBYxhZmQy", "https://t.me/+LNUv5XSCBngyMWY6", "https://t.me/+PZZgd6A0DF4zMmU6", "https://t.me/+mua_4CcpRK9hYTky", "https://t.me/+W40cuTOPdZs3MmFi", "https://t.me/+xeEEtiQ7euRhMDY6", "https://t.me/+2pWAeYuHhU1kYjdi", "https://t.me/+ShOu5KSlJwA1MDli", "https://t.me/+gFzXGO4p-p4xNTAy", "https://t.me/+oy8pToMIt-I3YWVi"}
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
					"🤑Оплата 800 руб. на +79997971960 СБП (ВТБ).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

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
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Держи бесплатный ответ на вопрос по математике + одно бесплатное видео с курса + ДЗ. Убедись в качестве и забери полный комплект!🥰")
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
				msg1 := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "https://disk.yandex.ru/i/YuRJ6zj1hgqlMw . Ссылка на видео с курса на тему: Исследование функций на экстремум.")
				if _, err := bot.Send(msg1); err != nil {
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
