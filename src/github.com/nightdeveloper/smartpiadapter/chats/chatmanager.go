package chats

import (
	"github.com/nightdeveloper/smartpiadapter/settings"
	"github.com/nightdeveloper/smartpiadapter/devices"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nightdeveloper/smartpiadapter/logger"
	"fmt"
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

	cm.bot = bot;

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if (update.Message == nil) {
			continue
		}
		logger.Info(fmt.Sprintf("%d %s sends me %s",
			update.Message.From.ID, update.Message.From.UserName, update.Message.Text))
	}

}