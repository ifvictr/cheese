# 🧀 Cheese

<img alt="Moldy Cheese" width="128" src="https://files.ifvictr.com/2020/07/cheese.jpg" />

The Cheese Touch, but over Slack. Idea taken from [Diary of a Wimpy Kid](https://diary-of-a-wimpy-kid.fandom.com/wiki/Cheese_Touch). Why did I make this? Besides having a reason to practice Go, I don’t know. 🤷‍♂

## Setup

### Creating the Slack app

You’ll need to create a Slack app (not a classic one) with at least the following bot token scopes. The reasons each scope is required are also included:

- `channels:history`: Used to check if a new user sent a message to pass the cheese touch. Only works if the app has been invited to the channel.
- `chat:write`: Used for sending messages when the cheese touch starts or gets passed to someone else.
- `groups:history`: Used to check if a new user sent a message to pass the cheese touch. Only works if the app has been invited to the channel.
- `reactions:read`: For listening to users that “touch” the cheese via reaction.

Then you’ll need to subscribe the app to a few events. The server has an endpoint open at `/slack/events`, so when you’re asked for a request URL, just put `https://<SERVER>/slack/events`. Only the following events are needed:

- `message.channels`
- `message.groups`
- `reaction_added`

### Environment variables

Here are all the variables you need to set up, with hints.

```bash
# Port to run the app server on
PORT=3000
# Redis database to record the user with the cheese touch
REDIS_URL=redis://…
# App config. Obtained from the "Basic Information" page of your app.
SLACK_CLIENT_BOT_TOKEN=xoxb-…
SLACK_VERIFICATION_TOKEN=xxxx…
```

### Deploying

```bash
# Run it
make

# Run it explicitly from a binary
make build
./bin/cheese
```

After you’ve followed all the above steps, you should see something like this:

```bash
Starting Cheese…
Listening on port 3000
```

Enjoy! 🧀

## License

[MIT License](LICENSE.txt)
