# telu

Telu helps to create interactive text worlds.

## Local setup

### Telegram bot

* [Set up](https://dashboard.ngrok.com/get-started/setup) ngrok.
* Create test Telegram bot.
* Copy `.env_template` to `.env`.
* Set ngrok URL as your test Telegram bot's webhook URL:  
`curl -d"url=${BOT_URL}" "https://api.telegram.org/bot${BOT_TOKEN}/setWebhook"`
* Set ngrok host as `TELEGRAM_BOT_HOST` in `.env`.
* Set bot token in `.env`.
