package database

var users = `CREATE TABLE IF NOT EXISTS users (
	userId INTEGER NOT NULL,
	userName TEXT NOT NULL,
	PRIMARY KEY (userId)
);`

var messages = `CREATE TABLE IF NOT EXISTS messages (
	messageId INTEGER NOT NULL,
	date DATETIME DEFAULT CURRENT_TIMESTAMP,
	content TEXT,
	PRIMARY KEY (messageId)
);`

var chats = `CREATE TABLE IF NOT EXISTS chats (
	chatId INTEGER NOT NULL,
	chatName TEXT NOT NULL,
	PRIMARY KEY (chatId)
);`

var chat_message = `CREATE TABLE IF NOT EXISTS chat_message (
	chatId INTEGER NOT NULL,
	messageId INTEGER NOT NULL,
	PRIMARY KEY (chatId, messageId),
	FOREIGN KEY (chatId) REFERENCES chats(chatId)
	FOREIGN KEY (messageId) REFERENCES messages(messageId)
);`

var chat_user = `CREATE TABLE IF NOT EXISTS chat_user (
	userId INTEGER NOT NULL,
	chatId INTEGER NOT NULL,
	PRIMARY KEY (userId, chatId),
	FOREIGN KEY (chatId) REFERENCES chats(chatId)
	FOREIGN KEY (userId) REFERENCES users(userId)
);`

var message_user = `CREATE TABLE IF NOT EXISTS message_user (
	messageId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
	PRIMARY KEY (messageId, userId),
	FOREIGN KEY (messageId) REFERENCES messages(messageId)
	FOREIGN KEY (userId) REFERENCES users(userId)
);`
