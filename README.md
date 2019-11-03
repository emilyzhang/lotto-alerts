# lotto-alerts

Simple task to send text message alerts whenever a California lottery jackpot is
over a certain target number. It uses your gmail account to send SMS text
messages and is very insecure (susceptible to man in the middle attacks) because
I turned off TLS verification for connecting to the Gmail SMTP server (I didn't
want to deal with it at the time).

This is a lambda function - you can upload function.zip and set `Handler` to
`lotto-alerts` within the AWS Lambda Console. Set the following environment
variables within the lambda function:

- `GMAIL_USER` (ex: example@gmail.com) a gmail account that you own.
- `PASSWORD` (ex: password123) the password for your gmail account.
- `PHONE_NUMBER` (ex: 4081119999) must be an AT&T phone number!
- `TARGET` (ex: 40) integer amount that represents the threshold of millions
  when you would want to be texted. the example number means that you'll be
  texted when the jackpot is over 40 million dollars for a california lottery.

I then used Cloudwatch Events to set a Rule to trigger this lambda function to
run every Monday, Tuesday, Thursday, and Friday (days before drawings for three
different lotteries are due). I used the following cron expression: `0 10 ? *
MON,TUE,THU,FRI *`.

You also have to turn on IMAP in your gmail account (in Settings/Forwarding and
POP/IMAP) and turn on [less secure app
access](https://myaccount.google.com/lesssecureapps) for this to work.
