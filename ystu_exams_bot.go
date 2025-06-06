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
			tgbotapi.NewInlineKeyboardButtonData("📐 Математика", "math"),
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
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("❓ Задать вопрос", "https://t.me/micsemal"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("📚 Заказать решение практики или лабы", "https://t.me/micsemal"),
		),
	)
)

var promoActive bool // Флаг для ожидания ввода промокода
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
				if v[count] == "payMath" {
					msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName + "\nТовар: Ответы по математике (преп. Ройтенберг В.М.)"
					msg.ReplyMarkup = checkBuyMath
				}
				if v[count] == "payAIP" {
					msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName + "\nТовар: Ответы по АИП (преп. Никитина Т.П.)"
					msg.ReplyMarkup = checkBuyAIP
				}
				if v[count] == "payCombo" {
					msg.Caption = "📸 Новый скриншот об оплате от пользователя: " + update.Message.From.UserName + "\nАктивирован промокод. Комбо(Математика + АИП)"
					msg.ReplyMarkup = checkBuyCombo
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
			case "math":
				edit := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Ответы на экзамен по математике (преп. Ройтенберг В.Ш.) + курс с практикой и ДЗ*\nЦена: 1500 рублей",
					payMath,
				)
				edit.ParseMode = "Markdown"

				if _, err := bot.Send(edit); err != nil {
					panic(err)
				}
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
				edit3 := tgbotapi.NewEditMessageTextAndMarkup(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"*Курс с практикой и ДЗ*\nЦена: 800 рублей",
					payMath,
				)
				edit3.ParseMode = "Markdown"

				if _, err := bot.Send(edit3); err != nil {
					panic(err)
				}

			case "checkOKMath":
				links := []string{"https://t.me/+9_Rn0QAdqD9hMTEy", "https://t.me/+Yq6Zwty1r943MzAy", "https://t.me/+c7JexxJuHAVlODky", "https://t.me/+BwacSv5It6VlNzUy", "https://t.me/+O1SNRYfjoF9kN2Ni", "https://t.me/+UCYYUt9BM0EyODdi", "https://t.me/+Jh0nWQ1xodEzYTcy", "https://t.me/+uaIwaj88fbU2MDAy", "https://t.me/+KNuHFm7uG3U0MzAy", "https://t.me/+v8Sqfo69DDo0Y2Iy", "https://t.me/+lFxT1GCzknQ0Njc6", "https://t.me/+jT0qWQMpruozYjU6", "https://t.me/+kHBPRKCyINY3MWE6", "https://t.me/+5TeER3364h9mYzVi", "https://t.me/+dR32D6DvVh8yMjcy", "https://t.me/+WhPxZkZy2mkzY2Ey", "https://t.me/+iQ1iFx1EPXVlNDIy", "https://t.me/+j2Nh5bLJzho5NGNi", "https://t.me/+FNCWjF5NDrg5Yjcy", "https://t.me/+InmN1wVG2IgzYjZi", "https://t.me/+UFPjva0c6UBiMGUy", "https://t.me/+0n69ET8g_WViMmEy", "https://t.me/+1xstVuSPcIBkYWJi", "https://t.me/+rd7KExlx8mE5YmVi", "https://t.me/+BepPaXPqoXYxYWYy", "https://t.me/+x4-lxtfhBWU2OWIy", "https://t.me/+DeLemII-9wc4N2I6", "https://t.me/+G0iJhK7sij5lNmMy", "https://t.me/+s38oZ_RlrsMxODZi", "https://t.me/+6HQAGbPSdvFjOTYy", "https://t.me/+ycHpajOJu5c5M2Uy", "https://t.me/+GRe5mz2UVfoxYmYy", "https://t.me/+iwlnhBL4f5BjYmQ6", "https://t.me/+NEH0sqmapxQ4Mjcy", "https://t.me/+rIpbDEiqLEU3ZWQ6", "https://t.me/+DJTI3d_twAkxNDVi", "https://t.me/+iHHV07JWu3o2NzZi", "https://t.me/+F4X83U31ODQ4YmZi", "https://t.me/+xsPt_J_6zYdjZDJi", "https://t.me/+qNXszpEojgdjOWEy", "https://t.me/+u-aw-qGY-k03NjQy", "https://t.me/+tsuJSfAYqTwwMmFi", "https://t.me/+8CmSyh-omipjY2Zi", "https://t.me/+vjrbgDjuWrU4YTUy", "https://t.me/+uPaV8cYQXPUyODI6", "https://t.me/+AMTKxuZtpns2ZmYy", "https://t.me/+uKW5HPal2Z4xNWYy", "https://t.me/+TKl4SlqEYmBhOTIy", "https://t.me/+9vjJPVKttbViMmRi", "https://t.me/+a71cPIWH4IMyMDYy", "https://t.me/+UdT8JddZr2xiMTI6", "https://t.me/+7ePy-wx1sZA5ZGEy", "https://t.me/+Iri6yO5mnV1jOGIy", "https://t.me/+T8bVf6Eq7vU4ZTgy", "https://t.me/+VHxJFX4tjM9lZWNi", "https://t.me/+Arm0GYXRVMY2ZDU6", "https://t.me/+T0nEVidEfnk1OTIy", "https://t.me/+CGzkc6K5GCY0Yjgy", "https://t.me/+h33aP4qxRHA3YzJi", "https://t.me/+t5ye9pfQ4nA4MDky", "https://t.me/+6pHfQpDbp7JjYWIy", "https://t.me/+fUxnfoR6-lYwNDEy", "https://t.me/+-qG9K8yHwuE0YTc6", "https://t.me/+gTCt7I_Fga03OTky", "https://t.me/+PCk0csw1pbQ0NGUy", "https://t.me/+kzU-y5dgpuc5OWIy", "https://t.me/+yzQ6LUFYrRRhMmFi", "https://t.me/+eau2YNTPGSE2NGFi", "https://t.me/+D_PpuwEL_QhhOGY6", "https://t.me/+UZJHoU7kszw0MTNi", "https://t.me/+81ck2IwQGtcxMmI6", "https://t.me/+xKkvko1wjJ5iM2Zi", "https://t.me/+SfDmXj1PB0diNzVi", "https://t.me/+oAacitUxbhMyZDgy", "https://t.me/+9hX-Th4499I3ZTcy", "https://t.me/+xEzO8RpatQ5iYzhi", "https://t.me/+ZIv6DFxrFNVkYzRi", "https://t.me/+P1nxlOfXg-0zMTYy", "https://t.me/+0gQ8ilcuaBszNGM6", "https://t.me/+ZcOwwyD46QpkN2U6", "https://t.me/+SeJ2vbADeWYwNDky", "https://t.me/+YIie7weL3qlmYmIy", "https://t.me/+MMcfy-nkA0tjZTli", "https://t.me/+RiEqoI9BkfY4ODg6"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[m])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				m++

			case "checkOKAIP":
				links := []string{"https://t.me/+ioeQifz8fyBlYWVi", "https://t.me/+ufFs04xq_0c1MTdi", "https://t.me/+xWm2k8WX6JU5ZjNi", "https://t.me/+5-8V4rgd_3w1OGRi", "https://t.me/+UfUZu6Vxt4xiNWI6", "https://t.me/+oINAOJq4_nIxZWYy", "https://t.me/+hsj2TopYdH45ZDRi", "https://t.me/+ZWhLxXk5zQNhN2U6", "https://t.me/+wYukijy79SExNzhi", "https://t.me/+MLB891yTYIs4MTc6", "https://t.me/+rKsbVwOd29g2NjQ6", "https://t.me/+Yps0cAuwtGNhYTky", "https://t.me/+FhFfQJjW37xlODFi", "https://t.me/+LgvZZv0si8g0NDgy", "https://t.me/+czRXGDNQ1nAyZmNi", "https://t.me/+OyhlsK8ZWk5kYTUy", "https://t.me/+4UM6A7-PFVMyODRi", "https://t.me/+Rbv9dslLfYcxOTli", "https://t.me/+dmOSiNR_ZXNkNTE6", "https://t.me/+tgbGIafZ9EZhNDFi", "https://t.me/+5lSMjWI9nzk0MzEy", "https://t.me/+PdDBVic4DlQ5MzNi", "https://t.me/+io2Lt2HgqSA4MjE6", "https://t.me/+c1nzTYWyNVgzNmNi", "https://t.me/+ZC7D0DWycc45N2Iy", "https://t.me/+gR8et0f-d0YzNWFi", "https://t.me/+XTe7--rXAw05Yzhi", "https://t.me/+EytKbSLUGmUyNGQy", "https://t.me/+uqAIRelqOgMyNGMy", "https://t.me/+Bv9RW1L4UVVlNTky", "https://t.me/+PTd19PY5Z3dmZDVi", "https://t.me/+I6rr0M0U4XUzZjcy", "https://t.me/+aiEEIog28is2ZGEy", "https://t.me/+h8vsB4HWy0hmY2Qy", "https://t.me/+cCFuoa9fqbg0YzYy", "https://t.me/+GkPXriVL19owOTNi", "https://t.me/+tUc9fM8skQZiMzAy", "https://t.me/+y3EqctCkgewyY2Iy", "https://t.me/+c3m96W0Ds7owZTNi", "https://t.me/+1mlkBMplDTFhZTNi", "https://t.me/+8E0NuR5M62U1MTYy", "https://t.me/+gwFVHdcTcEEyOWQy", "https://t.me/+EqAMMnE4H85iZGRi", "https://t.me/+m1D_v-tPI_U4YjBi", "https://t.me/+UwlTTjjk78k0OGRi", "https://t.me/+Xy9M77I7cbVkMmNi", "https://t.me/+qcLI841cppgxMmEy", "https://t.me/+l9_z-EGpmOszYzVi", "https://t.me/+nCrR18y8D9NjNWJi", "https://t.me/+DCzx2NRX5UFmNTNi", "https://t.me/+xs0hAWHVmIA4OWJi", "https://t.me/+h18FGcTjkARhNGUy", "https://t.me/+3HW5OK6xDDgwYzMy", "https://t.me/+6WdVpimqyIIyNzRi", "https://t.me/+dhsaWgZm0d43NDQ6", "https://t.me/+IrSDJX-lSYpjNGZi", "https://t.me/+TaG2ABTtg5c3ODg6", "https://t.me/+Zlu6UnWEcwg0ZGQ6", "https://t.me/+kqKclayN6StlZjQy", "https://t.me/+gwUCBrpqdTw2OWUy", "https://t.me/+iMeRerN2hGY1ODEy", "https://t.me/+tdHaUCar_hM2YmNi", "https://t.me/+FEUBdOLBJMtlZDQy", "https://t.me/+n4sH_ULlaEw4Njhi", "https://t.me/+ISp_QRBQlUQ3Mzc6", "https://t.me/+9Pq2wfmNkgw0MGQy", "https://t.me/+sgexPZsmK0M0ODcy", "https://t.me/+kZx7YI45CYE2Yjc6", "https://t.me/+N6TL8x48ZshiMGQy", "https://t.me/+gvudEyhtdXthMWUy", "https://t.me/+4082inrBcfpmYjli", "https://t.me/+X-21OFAhL-tjYWVi", "https://t.me/+aM9H7sfdV1wwOTFi", "https://t.me/+96xeqBp2oBg3YjE6", "https://t.me/+qf_m3PXQuTUyNzMy", "https://t.me/+N-tLEKiAMFliYmJi", "https://t.me/+KArZWFTvUjllZjFi", "https://t.me/+DXI6kcQMJoExNGEy", "https://t.me/+Zj3VzttakuIzYzky", "https://t.me/+rfBnWARq5b1lOTIy", "https://t.me/+qTIKYFIvlkE0NmNi", "https://t.me/+1Tei-w_LV8NiOWMy", "https://t.me/+HIpzzsjxQygxNmVi", "https://t.me/+lAYWxD2U8REyNWEy"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК "+links[a])
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				a++

			case "checkOKCombo":
				links1 := []string{"https://t.me/+9_Rn0QAdqD9hMTEy", "https://t.me/+Yq6Zwty1r943MzAy", "https://t.me/+c7JexxJuHAVlODky", "https://t.me/+BwacSv5It6VlNzUy", "https://t.me/+O1SNRYfjoF9kN2Ni", "https://t.me/+UCYYUt9BM0EyODdi", "https://t.me/+Jh0nWQ1xodEzYTcy", "https://t.me/+uaIwaj88fbU2MDAy", "https://t.me/+KNuHFm7uG3U0MzAy", "https://t.me/+v8Sqfo69DDo0Y2Iy", "https://t.me/+lFxT1GCzknQ0Njc6", "https://t.me/+jT0qWQMpruozYjU6", "https://t.me/+kHBPRKCyINY3MWE6", "https://t.me/+5TeER3364h9mYzVi", "https://t.me/+dR32D6DvVh8yMjcy", "https://t.me/+WhPxZkZy2mkzY2Ey", "https://t.me/+iQ1iFx1EPXVlNDIy", "https://t.me/+j2Nh5bLJzho5NGNi", "https://t.me/+FNCWjF5NDrg5Yjcy", "https://t.me/+InmN1wVG2IgzYjZi", "https://t.me/+UFPjva0c6UBiMGUy", "https://t.me/+0n69ET8g_WViMmEy", "https://t.me/+1xstVuSPcIBkYWJi", "https://t.me/+rd7KExlx8mE5YmVi", "https://t.me/+BepPaXPqoXYxYWYy", "https://t.me/+x4-lxtfhBWU2OWIy", "https://t.me/+DeLemII-9wc4N2I6", "https://t.me/+G0iJhK7sij5lNmMy", "https://t.me/+s38oZ_RlrsMxODZi", "https://t.me/+6HQAGbPSdvFjOTYy", "https://t.me/+ycHpajOJu5c5M2Uy", "https://t.me/+GRe5mz2UVfoxYmYy", "https://t.me/+iwlnhBL4f5BjYmQ6", "https://t.me/+NEH0sqmapxQ4Mjcy", "https://t.me/+rIpbDEiqLEU3ZWQ6", "https://t.me/+DJTI3d_twAkxNDVi", "https://t.me/+iHHV07JWu3o2NzZi", "https://t.me/+F4X83U31ODQ4YmZi", "https://t.me/+xsPt_J_6zYdjZDJi", "https://t.me/+qNXszpEojgdjOWEy", "https://t.me/+u-aw-qGY-k03NjQy", "https://t.me/+tsuJSfAYqTwwMmFi", "https://t.me/+8CmSyh-omipjY2Zi", "https://t.me/+vjrbgDjuWrU4YTUy", "https://t.me/+uPaV8cYQXPUyODI6", "https://t.me/+AMTKxuZtpns2ZmYy", "https://t.me/+uKW5HPal2Z4xNWYy", "https://t.me/+TKl4SlqEYmBhOTIy", "https://t.me/+9vjJPVKttbViMmRi", "https://t.me/+a71cPIWH4IMyMDYy", "https://t.me/+UdT8JddZr2xiMTI6", "https://t.me/+7ePy-wx1sZA5ZGEy", "https://t.me/+Iri6yO5mnV1jOGIy", "https://t.me/+T8bVf6Eq7vU4ZTgy", "https://t.me/+VHxJFX4tjM9lZWNi", "https://t.me/+Arm0GYXRVMY2ZDU6", "https://t.me/+T0nEVidEfnk1OTIy", "https://t.me/+CGzkc6K5GCY0Yjgy", "https://t.me/+h33aP4qxRHA3YzJi", "https://t.me/+t5ye9pfQ4nA4MDky", "https://t.me/+6pHfQpDbp7JjYWIy", "https://t.me/+fUxnfoR6-lYwNDEy", "https://t.me/+-qG9K8yHwuE0YTc6", "https://t.me/+gTCt7I_Fga03OTky", "https://t.me/+PCk0csw1pbQ0NGUy", "https://t.me/+kzU-y5dgpuc5OWIy", "https://t.me/+yzQ6LUFYrRRhMmFi", "https://t.me/+eau2YNTPGSE2NGFi", "https://t.me/+D_PpuwEL_QhhOGY6", "https://t.me/+UZJHoU7kszw0MTNi", "https://t.me/+81ck2IwQGtcxMmI6", "https://t.me/+xKkvko1wjJ5iM2Zi", "https://t.me/+SfDmXj1PB0diNzVi", "https://t.me/+oAacitUxbhMyZDgy", "https://t.me/+9hX-Th4499I3ZTcy", "https://t.me/+xEzO8RpatQ5iYzhi", "https://t.me/+ZIv6DFxrFNVkYzRi", "https://t.me/+P1nxlOfXg-0zMTYy", "https://t.me/+0gQ8ilcuaBszNGM6", "https://t.me/+ZcOwwyD46QpkN2U6", "https://t.me/+SeJ2vbADeWYwNDky", "https://t.me/+YIie7weL3qlmYmIy", "https://t.me/+MMcfy-nkA0tjZTli", "https://t.me/+RiEqoI9BkfY4ODg6"}
				links2 := []string{"https://t.me/+ioeQifz8fyBlYWVi", "https://t.me/+ufFs04xq_0c1MTdi", "https://t.me/+xWm2k8WX6JU5ZjNi", "https://t.me/+5-8V4rgd_3w1OGRi", "https://t.me/+UfUZu6Vxt4xiNWI6", "https://t.me/+oINAOJq4_nIxZWYy", "https://t.me/+hsj2TopYdH45ZDRi", "https://t.me/+ZWhLxXk5zQNhN2U6", "https://t.me/+wYukijy79SExNzhi", "https://t.me/+MLB891yTYIs4MTc6", "https://t.me/+rKsbVwOd29g2NjQ6", "https://t.me/+Yps0cAuwtGNhYTky", "https://t.me/+FhFfQJjW37xlODFi", "https://t.me/+LgvZZv0si8g0NDgy", "https://t.me/+czRXGDNQ1nAyZmNi", "https://t.me/+OyhlsK8ZWk5kYTUy", "https://t.me/+4UM6A7-PFVMyODRi", "https://t.me/+Rbv9dslLfYcxOTli", "https://t.me/+dmOSiNR_ZXNkNTE6", "https://t.me/+tgbGIafZ9EZhNDFi", "https://t.me/+5lSMjWI9nzk0MzEy", "https://t.me/+PdDBVic4DlQ5MzNi", "https://t.me/+io2Lt2HgqSA4MjE6", "https://t.me/+c1nzTYWyNVgzNmNi", "https://t.me/+ZC7D0DWycc45N2Iy", "https://t.me/+gR8et0f-d0YzNWFi", "https://t.me/+XTe7--rXAw05Yzhi", "https://t.me/+EytKbSLUGmUyNGQy", "https://t.me/+uqAIRelqOgMyNGMy", "https://t.me/+Bv9RW1L4UVVlNTky", "https://t.me/+PTd19PY5Z3dmZDVi", "https://t.me/+I6rr0M0U4XUzZjcy", "https://t.me/+aiEEIog28is2ZGEy", "https://t.me/+h8vsB4HWy0hmY2Qy", "https://t.me/+cCFuoa9fqbg0YzYy", "https://t.me/+GkPXriVL19owOTNi", "https://t.me/+tUc9fM8skQZiMzAy", "https://t.me/+y3EqctCkgewyY2Iy", "https://t.me/+c3m96W0Ds7owZTNi", "https://t.me/+1mlkBMplDTFhZTNi", "https://t.me/+8E0NuR5M62U1MTYy", "https://t.me/+gwFVHdcTcEEyOWQy", "https://t.me/+EqAMMnE4H85iZGRi", "https://t.me/+m1D_v-tPI_U4YjBi", "https://t.me/+UwlTTjjk78k0OGRi", "https://t.me/+Xy9M77I7cbVkMmNi", "https://t.me/+qcLI841cppgxMmEy", "https://t.me/+l9_z-EGpmOszYzVi", "https://t.me/+nCrR18y8D9NjNWJi", "https://t.me/+DCzx2NRX5UFmNTNi", "https://t.me/+xs0hAWHVmIA4OWJi", "https://t.me/+h18FGcTjkARhNGUy", "https://t.me/+3HW5OK6xDDgwYzMy", "https://t.me/+6WdVpimqyIIyNzRi", "https://t.me/+dhsaWgZm0d43NDQ6", "https://t.me/+IrSDJX-lSYpjNGZi", "https://t.me/+TaG2ABTtg5c3ODg6", "https://t.me/+Zlu6UnWEcwg0ZGQ6", "https://t.me/+kqKclayN6StlZjQy", "https://t.me/+gwUCBrpqdTw2OWUy", "https://t.me/+iMeRerN2hGY1ODEy", "https://t.me/+tdHaUCar_hM2YmNi", "https://t.me/+FEUBdOLBJMtlZDQy", "https://t.me/+n4sH_ULlaEw4Njhi", "https://t.me/+ISp_QRBQlUQ3Mzc6", "https://t.me/+9Pq2wfmNkgw0MGQy", "https://t.me/+sgexPZsmK0M0ODcy", "https://t.me/+kZx7YI45CYE2Yjc6", "https://t.me/+N6TL8x48ZshiMGQy", "https://t.me/+gvudEyhtdXthMWUy", "https://t.me/+4082inrBcfpmYjli", "https://t.me/+X-21OFAhL-tjYWVi", "https://t.me/+aM9H7sfdV1wwOTFi", "https://t.me/+96xeqBp2oBg3YjE6", "https://t.me/+qf_m3PXQuTUyNzMy", "https://t.me/+N-tLEKiAMFliYmJi", "https://t.me/+KArZWFTvUjllZjFi", "https://t.me/+DXI6kcQMJoExNGEy", "https://t.me/+Zj3VzttakuIzYzky", "https://t.me/+rfBnWARq5b1lOTIy", "https://t.me/+qTIKYFIvlkE0NmNi", "https://t.me/+1Tei-w_LV8NiOWMy", "https://t.me/+HIpzzsjxQygxNmVi", "https://t.me/+lAYWxD2U8REyNWEy"}
				msg := tgbotapi.NewMessage(chatId[p], "Оплата прошла успешно!✅ Держи ссылку-приглашение в ТГК по математике "+links1[m]+" и по АИП "+links2[a])
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
					"Вас приветствует бот *YSTU EXAMS*👋\nЭкзамены уже очень скоро! Поэтому, чтобы не терять время, ты можешь приобрести ответы на экзамены по Математике и АИП\nСкорее покупай! 🥰",
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
