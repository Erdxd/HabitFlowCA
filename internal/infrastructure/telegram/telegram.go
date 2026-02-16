package telegram

import (
	"HabitFlow/internal/domain/service"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api          *tgbotapi.BotAPI
	HabitService *service.HabitService
	UserService  *service.UserService
}

func NewBot(token string, userService *service.UserService, habitservice *service.HabitService) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil
	}
	return &Bot{api: api, HabitService: habitservice, UserService: userService}
}

func (tg *Bot) SendHabitNotification(chatId int64) error {
	habits, err := tg.HabitService.GetHabitsByTgId(chatId)
	if err != nil {
		return err
	}
	if len(habits) == 0 {

		message := "🎉 У тебя пока нет активных привычек! Добавь первую привычку на сайте."
		msg := tgbotapi.NewMessage(chatId, message)
		_, err := tg.api.Send(msg)
		return err

	}
	message := "📋 Твои привычки на сегодня:\n\n"
	for _, habit := range habits {
		status := "❌"
		if habit.Status_Today {

			status = "✅"
		}
		message += fmt.Sprintf("%s %s (серия: %d)\n", status, habit.Habit_Name, habit.Streak)

	}

	msg := tgbotapi.NewMessage(chatId, message)
	_, err = tg.api.Send(msg)
	return err
}
func (tg *Bot) HandleMessages() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Command() == "start" {
				args := update.Message.CommandArguments()
				userID, err := strconv.Atoi(args)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверный формат. Используй: /start ТВОЙ_ID_ПОЛЬЗОВАТЕЛЯ")
					tg.api.Send(msg)
					continue
				}

				err = tg.UserService.SaveTelegramChatID(userID, update.Message.Chat.ID)
				if err != nil {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка привязки чата")
					tg.api.Send(msg)
					continue
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "✅ Чат успешно привязан! Теперь ты будешь получать уведомления о привычках.")
				tg.api.Send(msg)
				tg.SendKeyboard(update.Message.Chat.ID)
			} else if update.Message.Command() == "GetTask" {

				tg.SendHabitNotification(update.Message.Chat.ID)
				tg.SendKeyboard(update.Message.Chat.ID)
			} else if update.Message.Text == "📋 Мои привычки" {

				tg.SendHabitNotification(update.Message.Chat.ID)
				tg.SendKeyboard(update.Message.Chat.ID)
			}

		}

	}
}

func (tg *Bot) SendKeyboard(ChatId int64) error {
	keyboard := [][]tgbotapi.KeyboardButton{
		{
			{Text: "📋 Мои привычки"},
		},
	}
	replyMarkup := tgbotapi.NewReplyKeyboard(keyboard...)
	replyMarkup.OneTimeKeyboard = false
	replyMarkup.ResizeKeyboard = true
	msg := tgbotapi.NewMessage(ChatId, "Выберите действие:")
	msg.ReplyMarkup = replyMarkup
	_, err := tg.api.Send(msg)
	return err
}
