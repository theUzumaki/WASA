package database

var users = `CREATE TABLE IF NOT EXISTS users (
	userId INTEGER NOT NULL,
	userName TEXT NOT NULL,
	picture STRING,
	PRIMARY KEY (userId)
);`

var messages = `CREATE TABLE IF NOT EXISTS messages (
	messageId INTEGER NOT NULL,
	date DATETIME DEFAULT CURRENT_TIMESTAMP,
	content TEXT NOT NULL,
	checkmark BOOLEAN DEFAULT FALSE,
	PRIMARY KEY (messageId)
);`

var comments = `CREATE TABLE IF NOT EXISTS comments (
	commentId INTEGER NOT NULL,
	content TEXT NOT NULL,
	PRIMARY KEY (commentId)
);`

var chats = `CREATE TABLE IF NOT EXISTS chats (
	chatId INTEGER NOT NULL,
	chatName TEXT NOT NULL,
	picture STRING,
	PRIMARY KEY (chatId)
);`

var chat_message = `CREATE TABLE IF NOT EXISTS chat_message (
	chatId INTEGER NOT NULL,
	messageId INTEGER NOT NULL,
	PRIMARY KEY (chatId, messageId),
	FOREIGN KEY (chatId) REFERENCES chats(chatId)
		ON DELETE CASCADE
	FOREIGN KEY (messageId) REFERENCES messages(messageId)
		ON DELETE CASCADE
);`

var chat_user = `CREATE TABLE IF NOT EXISTS chat_user (
	chatId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
	PRIMARY KEY (chatId, userId),
	FOREIGN KEY (chatId) REFERENCES chats(chatId)
		ON DELETE CASCADE
	FOREIGN KEY (userId) REFERENCES users(userId)
		ON DELETE CASCADE
);`

var comment_message = `CREATE TABLE IF NOT EXISTS comment_message (
	commentId INTEGER NOT NULL,
	messageId INTEGER NOT NULL,
	PRIMARY KEY (messageId, commentId),
	FOREIGN KEY (messageId) REFERENCES messages(messageId)
		ON DELETE CASCADE
	FOREIGN KEY (commentId) REFERENCES comments(commentId)
		ON DELETE CASCADE
);`

var comment_user = `CREATE TABLE IF NOT EXISTS comment_user (
	commentId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
	PRIMARY KEY (userId, commentId),
	FOREIGN KEY (userId) REFERENCES users(userId)
		ON DELETE CASCADE
	FOREIGN KEY (commentId) REFERENCES comments(commentId)
		ON DELETE CASCADE
);`

var message_user = `CREATE TABLE IF NOT EXISTS message_user (
	messageId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
	PRIMARY KEY (messageId, userId),
	FOREIGN KEY (messageId) REFERENCES messages(messageId)
		ON DELETE CASCADE
	FOREIGN KEY (userId) REFERENCES users(userId)
		ON DELETE CASCADE
);`

var message_viewer = `CREATE TABLE IF NOT EXISTS message_viewer (
	messageId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
		PRIMARY KEY (messageId, userId),
	FOREIGN KEY (messageId) REFERENCES messages(messageId)
		ON DELETE CASCADE
	FOREIGN KEY (userId) REFERENCES users(userId)
		ON DELETE CASCADE
);`

var message_reply = `CREATE TABLE IF NOT EXISTS message_reply (
	messageId INTEGER NOT NULL,
	replyId INTEGER NOT NULL,
	PRIMARY KEY (messageId, replyId),
	FOREIGN KEY (messageId) REFERENCES messages(messageId)
		ON DELETE CASCADE
	FOREIGN KEY (replyId) REFERENCES messages(messageId)
		ON DELETE CASCADE
);`
