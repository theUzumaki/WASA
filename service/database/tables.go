package database

var usersCollection = `CREATE TABLE IF NOT EXIST usersCollection(
	collectionId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
	PRIMARY KEY (collectionId, userId),
)`

var users = `CREATE TABLE IF NOT EXIST users(
	userId INTEGER NOT NULL,
	userName TEXT NOT NULL,
	chatsCollectionId INTEGER NOT NULL,
	PRIMARY KEY (userId),
	FOREIGN KEY (groupsCollectionId) REFERENCES groupsCollection(collectionId),
	FOREIGN KEY (chatsCollectionId) REFERENCES chats(collectionId),
)`

/*
var groupsCollection = `CREATE TABLE IF NOT EXIST groupsCollection(
	collectionId INTEGER NOT NULL,
	groupId INTEGER NOT NULL,
	PRIMARY KEY (collectionId, groupId),
	FOREIGN KEY (groupId) REFERENCES group(groupId),
)`

var group = `CREATE TABLE IF NOT EXIST group(
	groupId INTEGER NOT NULL,
	groupName TEXT NOT NULL,
	PRIMARY KEY (groupId),
)`
*/

var messagesCollection = `CREATE TABLE IF NOT EXIST messagesCollection(
	collectionId INTEGER NOT NULL,
	messageId INTEGER NOT NULL,
	PRIMARY KEY (collectionId, messageId),
	FOREIGN KEY (messageId) REFERENCES message(messageId),
	)`

var message = `CREATE TABLE IF NOT EXIST message(
	messageId INTEGER NOT NULL,
	sender INTEGER NOT NULL,
	chat INTEGER NOT NULL,
	date DATETIME DEFAULT CURRENT_TIMESTAMP,
	content TEXT,
	PRIMARY KEY (messageId),
	FOREIGN KEY sender REFERENCES user(userId),
	FOREIGN KEY chat REFERENCES chat(chatId),
	)`

var chatsCollection = `CREATE TABLE IF NOT EXIST chatsCollection(
	collectionId INTEGER NOT NULL,
	chatId INTEGER NOT NULL,
	PRIMARY KEY (collectionId, chatId),
	FOREIGN KEY (chatId) REFERENCES chat(chatId),
)`

var chat = `CREATE TABLE IF NOT EXIST chat(
	chatId INTEGER NOT NULL,
	chatName TEXT,
	members INTEGER NOT NULL,
	messages INTEGER NOT NULL,
	PRIMARY KEY (chatId),
	FOREIGN KEY (members) REFERENCES chatsCollection(collectionId),
	FOREIGN KEY (messages) REFERENCES messagesCollection(collectionId)
)`
