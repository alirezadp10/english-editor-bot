package bot

import (
    "english-editor-bot/api"
    "english-editor-bot/jobs"
    "gopkg.in/telebot.v4"
    "gorm.io/gorm"
)

func SetupHandlers(bot *telebot.Bot, db *gorm.DB, apiKey string) {
    bot.Handle("/check", func(c telebot.Context) error {
        return checkGrammarIssues(c, db, apiKey)
    })

    bot.Handle("/formal", func(c telebot.Context) error {
        return convertToFormal(c, db, apiKey)
    })

    bot.Handle("/informal", func(c telebot.Context) error {
        return convertToInformal(c, db, apiKey)
    })

    bot.Handle("/en", func(c telebot.Context) error {
        return translateToEnglish(c, db, apiKey)
    })

    bot.Handle("/fa", func(c telebot.Context) error {
        return translateToFarsi(c, db, apiKey)
    })
}

func checkGrammarIssues(c telebot.Context, db *gorm.DB, apiKey string) error {
    go jobs.SaveMessage(c, db)

    if c.Message().ReplyTo == nil {
        return c.Reply("باید رو مسیجی که می‌خوای تصحیح شه ریپلای کنی", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher just revise the sentenceses. Provide the revised version in formal and informal(casual) form, list corrected errors in bullet form. Your answer format should be as follows, without additional introductory text or section titles:\n\n👔[PUT THE FORMAL FORM HERE]\n\n🦦<i>[PUT THE CASUAL FORM HERE]</i>\n\n<blockquote>[PUT THE CORRECTED ERRORS HERE IN BULLET FORM]</blockquote>"
    requestBody := api.CreateRequestBody(c.Message().ReplyTo.Text, systemRole)
    responseBody, err := api.SendRequest(requestBody, apiKey)
    if err != nil {
        return err
    }

    // Parse and reply
    content, err := api.ParseResponse(responseBody)
    if err != nil {
        return err
    }
    return c.Reply(content, telebot.ModeHTML)
}

func convertToFormal(c telebot.Context, db *gorm.DB, apiKey string) error {
    go jobs.SaveMessage(c, db)

    if c.Message().ReplyTo == nil {
        return c.Reply("باید رو مسیجی که می‌خوای تبدیل شه ریپلای کنی", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher to check and correct the sentences in your responses. Provide the revised version in formal form, list corrected errors in bullet form. For example, your answer format should be as follows, without additional introductory text or section titles:\n\n<b>👔 رسمی:</b>\n\n[Corrected formal form]\n\n<b>🚧 توضیحات: </b>\n<blockquote>[Corrected errors in bullet list]</blockquote>"
    requestBody := api.CreateRequestBody(c.Message().ReplyTo.Text, systemRole)
    responseBody, err := api.SendRequest(requestBody, apiKey)
    if err != nil {
        return err
    }

    // Parse and reply
    content, err := api.ParseResponse(responseBody)
    if err != nil {
        return err
    }
    return c.Reply(content, telebot.ModeHTML)
}

func convertToInformal(c telebot.Context, db *gorm.DB, apiKey string) error {
    go jobs.SaveMessage(c, db)

    if c.Message().ReplyTo == nil {
        return c.Reply("باید رو مسیجی که می‌خوای تبدیل شه ریپلای کنی", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher to check and correct the sentences in your responses. Provide the revised version in informal(casual) form, list corrected errors in bullet form. For example, your answer format should be as follows, without additional introductory text or section titles:\n\n<b>🦦 Informal Form:</b>\n\n[Corrected informal form]\n\n<b>🚧 Details: </b>\n<blockquote>[Corrected errors in bullet list]</blockquote>"
    requestBody := api.CreateRequestBody(c.Message().ReplyTo.Text, systemRole)
    responseBody, err := api.SendRequest(requestBody, apiKey)
    if err != nil {
        return err
    }

    // Parse and reply
    content, err := api.ParseResponse(responseBody)
    if err != nil {
        return err
    }
    return c.Reply(content, telebot.ModeHTML)
}

func translateToEnglish(c telebot.Context, db *gorm.DB, apiKey string) error {
    go jobs.SaveMessage(c, db)

    if c.Message().ReplyTo == nil {
        return c.Reply("باید رو مسیجی که می‌خوای تبدیل شه ریپلای کنی", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher and translate the text that is provided in english. For example, your answer format should be as follows, without additional introductory text or section titles:\n\n<b>👩🏻‍🏫 You Should Say:</b>\n\n[translated version in english]"
    requestBody := api.CreateRequestBody(c.Message().ReplyTo.Text, systemRole)
    responseBody, err := api.SendRequest(requestBody, apiKey)
    if err != nil {
        return err
    }

    // Parse and reply
    content, err := api.ParseResponse(responseBody)
    if err != nil {
        return err
    }
    return c.Reply(content, telebot.ModeHTML)
}

func translateToFarsi(c telebot.Context, db *gorm.DB, apiKey string) error {
    go jobs.SaveMessage(c, db)

    if c.Message().ReplyTo == nil {
        return c.Reply("باید رو مسیجی که می‌خوای تبدیل شه ریپلای کنی", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher and translate the text that is provided in Farsi(persian). For example, your answer format should be as follows, without additional introductory text or section titles:\n\n<b>👨🏻‍🏫 یعنی:</b>\n\n[translated version in farsi(persian)]"
    requestBody := api.CreateRequestBody(c.Message().ReplyTo.Text, systemRole)
    responseBody, err := api.SendRequest(requestBody, apiKey)
    if err != nil {
        return err
    }

    // Parse and reply
    content, err := api.ParseResponse(responseBody)
    if err != nil {
        return err
    }
    return c.Reply(content, telebot.ModeHTML)
}
