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
}

func checkGrammarIssues(c telebot.Context, db *gorm.DB, apiKey string) error {
    go jobs.SaveMessage(c, db)

    if c.Message().ReplyTo == nil {
        return c.Reply("You must reply to the message you want to convert.", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher just revise the sentences. Provide the revised version in formal and informal(casual) form, list corrected errors in bullet form. Your answer format should be as follows, without additional introductory text or section titles:\n\n👔[PUT THE FORMAL FORM HERE]\n\n🦦<i>[PUT THE CASUAL FORM HERE]</i>\n\n<blockquote>[PUT THE CORRECTED ERRORS HERE IN BULLET FORM]</blockquote>"
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
        return c.Reply("You must reply to the message you want to convert.", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher just revise the sentences. Provide the revised version in formal form, list corrected errors in bullet form. Your answer format should be as follows, without additional introductory text or section titles:\n\n👔[PUT THE FORMAL FORM HERE]\n\n<blockquote>[PUT THE CORRECTED ERRORS HERE IN BULLET FORM]</blockquote>"
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
        return c.Reply("You must reply to the message you want to convert.", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher just revise the sentences. Provide the revised version in informal(casual) form, list corrected errors in bullet form. Your answer format should be as follows, without additional introductory text or section titles:\n\n🦦<i>[PUT THE CASUAL FORM HERE]</i>\n\n<blockquote>[PUT THE CORRECTED ERRORS HERE IN BULLET FORM]</blockquote>"
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
        return c.Reply("You must reply to the message you want to convert.", telebot.ModeHTML)
    }

    // Request creation and API call
    systemRole := "Act as an English teacher and just translate the text that is provided in english. Your answer format should be as follows, without additional introductory text or section titles:\n\n👩🏻‍🏫[PUT THE TRANSLATED VERSION IN ENGLISH FORM HERE]\n"
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
