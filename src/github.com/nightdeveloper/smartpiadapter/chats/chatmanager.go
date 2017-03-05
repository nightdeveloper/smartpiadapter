package chats

import (
	"github.com/nightdeveloper/smartpiadapter/settings"
	"github.com/nightdeveloper/smartpiadapter/devices"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nightdeveloper/smartpiadapter/logger"
	"fmt"
	"strconv"
	"time"
)

type ChatManager struct {
	c *settings.Config
	dm *devices.DeviceManager
	bot *tgbotapi.BotAPI;
}

func (cm *ChatManager) Init(c *settings.Config, dm *devices.DeviceManager) {
	cm.c = c;
	cm.dm = dm;
}

func (cm *ChatManager) Start() {

	bot, err := tgbotapi.NewBotAPI(cm.c.TelegramKey)

	if err != nil {
		logger.Error("telegram bot creating failed", err)
		return;
	}

	logger.Info("telegram bot created");

	cm.bot = bot;

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		time.Sleep(time.Second)

		if (update.Message == nil) {
			continue
		}

		if (strconv.Itoa(update.Message.From.ID) != cm.c.TelegramOpId) {
			logger.Info(fmt.Sprintf("[%d %s] sends me unauth message: %s",
				update.Message.From.ID, update.Message.From.UserName, update.Message.Text))
			continue
		}

		logger.Info(fmt.Sprintf("operator sends me %s", update.Message.Text))
	}

}