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

```bash
# Port to run the app server on
PORT=
# Redis database to record the user with the cheese touch
REDIS_URL=
# App config
SLACK_CLIENT_BOT_TOKEN=
SLACK_CLIENT_SIGNING_SECRET=
```

### Deploying

```bash
# If you just want to run it
make

# Running from a binary
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
