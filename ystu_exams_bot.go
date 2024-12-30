package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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
	checkBuyCombo = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить✅", "checkOKCombo"),
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

	payIsPromo = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💸 Оплатить", "payCombo"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Назад", "backMenu"),
		),
	)

	menuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📐 Ответы на Математику", "otvetyMath"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💻 Ответы на АИП", "otvetyAIP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🏷️ Ввести промокод", "promo"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎁 Получить подарок", "podarok"),
		),
	)
)

var promoActive bool // Флаг для ожидания ввода промокода
var m int = 0
var a int = 0
var v []string
var count int = -1
var chatId []int64

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
			// Проверяем, ожидает ли бот промокод
			if promoActive {
				if update.Message.Text == "СЕССИЯ" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Промокод активирован!✅ Забирай ответы на Математику и АИП всего *за 800 рублей*😍")
					msg.ParseMode = "Markdown"
					msg.ReplyMarkup = payIsPromo
					if _, err := bot.Send(msg); err != nil {
						log.Panic(err)
					}
					promoActive = false // Сбрасываем флаг
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "❌ Неверный промокод. Попробуйте ещё раз.")
					if _, err := bot.Send(msg); err != nil {
						log.Panic(err)
					}
				}
				continue
			}

			// Обработка фото
			if update.Message.Photo != nil {
				// Пересылаем фото админу
				photo := update.Message.Photo[len(update.Message.Photo)-1] // Берем фото в максимальном разрешении
				msg := tgbotapi.NewPhoto(adminID, tgbotapi.FileID(photo.FileID))
				msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName
				if v[count] == "payMath" {
					msg.ReplyMarkup = checkBuyMath
				}
				if v[count] == "payAIP" {
					msg.ReplyMarkup = checkBuyAIP
				}
				if v[count] == "payCombo" {
					msg.ReplyMarkup = checkBuyCombo
				}
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}

				// Ответ пользователю
				chatId = append(chatId, update.Message.Chat.ID)
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
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вас приветствует бот *YSTU EXAMS*👋\nЭкзамены уже очень скоро! А у меня ты можешь приобрести ответы на экзамены по Математике и АИП\nСкорее покупай! 🥰")
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
			case "otvetyMath":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Ответы на экзамен по математике (преп. Ройтенберг В.Ш.)*\nЦена: 700 рублей",
					payMath,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}

			case "checkOKMath":
				links := []string{"https://t.me/+Si-jbMCNE15jYmY6", "https://t.me/+_gel4G0PPds1ZmEy", "https://t.me/+-RUIbD4DA9cxMzdi", "https://t.me/+deQhSJaCHr8xZTBi", "https://t.me/+jzRfJbtjTnQ3ODM6", "https://t.me/+fSj0swgA47ZkYTVi", "https://t.me/+slzyzlZpDZlmZTUy", "https://t.me/+rbQXmGAI-pMyYTQy", "https://t.me/+EYlK32rx1FM3ZmQ6", "https://t.me/+9_Rn0QAdqD9hMTEy"}
				msg := tgbotapi.NewMessage(chatId[0], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[m])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				m++

			case "checkOKAIP":
				links := []string{"https://t.me/+E1pmezWFfmk4MzY6", "https://t.me/+w3ny9xII8JoxYThi", "https://t.me/+roVKsM823n85Nzky", "https://t.me/+TodPjt8wCiVjNjcy", "https://t.me/+FVRXtzYc7JVkOTcy", "https://t.me/+ci_HDPgwA4g5YjZi", "https://t.me/+ioeQifz8fyBlYWVi", "https://t.me/+ufFs04xq_0c1MTdi", "https://t.me/+xWm2k8WX6JU5ZjNi", "https://t.me/+5-8V4rgd_3w1OGRi"}
				msg := tgbotapi.NewMessage(chatId[0], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[a])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				a++

			case "checkOKCombo":
				links1 := []string{"https://t.me/+Si-jbMCNE15jYmY6", "https://t.me/+_gel4G0PPds1ZmEy", "https://t.me/+-RUIbD4DA9cxMzdi", "https://t.me/+deQhSJaCHr8xZTBi", "https://t.me/+jzRfJbtjTnQ3ODM6", "https://t.me/+fSj0swgA47ZkYTVi", "https://t.me/+slzyzlZpDZlmZTUy", "https://t.me/+rbQXmGAI-pMyYTQy", "https://t.me/+EYlK32rx1FM3ZmQ6", "https://t.me/+9_Rn0QAdqD9hMTEy"}
				links2 := []string{"https://t.me/+E1pmezWFfmk4MzY6", "https://t.me/+w3ny9xII8JoxYThi", "https://t.me/+roVKsM823n85Nzky", "https://t.me/+TodPjt8wCiVjNjcy", "https://t.me/+FVRXtzYc7JVkOTcy", "https://t.me/+ci_HDPgwA4g5YjZi", "https://t.me/+ioeQifz8fyBlYWVi", "https://t.me/+ufFs04xq_0c1MTdi", "https://t.me/+xWm2k8WX6JU5ZjNi", "https://t.me/+5-8V4rgd_3w1OGRi"}
				msg := tgbotapi.NewMessage(chatId[0], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК по математике "+links1[m]+" и по АИП "+links2[a])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				m++
				a++

			case "checkBAD":
				msg := tgbotapi.NewMessage(chatId[0], "Оплата была отклонена❌ Попробуйте снова")
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			case "payMath":
				count++
				v = append(v, "payMath")
				edit := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"🤑Оплата 700 руб. на +79536424194 СБП (Сбер).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}

			case "payAIP":
				count++
				v = append(v, "payAIP")
				edit := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"🤑Оплата 500 руб. на +79536424194 СБП (Сбер).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "payCombo":
				count++
				v = append(v, "payCombo")
				edit := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"🤑Оплата 800 руб. на +79536424194 СБП (Сбер).\nОБЯЗАТЕЛЬНО отправь скрин об оплате!")

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}

			case "otvetyAIP":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Ответы на экзамен по Алгоритмизации и программированию (преп. Никитина Т.П.)*\nЦена: 500 рублей",
					payAIP,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
			case "promo":
				promoActive = true // Устанавливаем флаг ожидания промокода
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите промокод:")
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			case "backMenu":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Вас приветствует бот *YSTU EXAMS*👋\nЭкзамены уже очень скоро! А у меня ты можешь приобрести ответы на экзамены по Математике и АИП\nСкорее покупай! 🥰",
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
