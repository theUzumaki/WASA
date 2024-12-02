package database

var user = `CREATE TABLE IF NOT EXIST user(
	userId INTEGER NOT NULL,
	userName TEXT NOT NULL,
	profileImage TEXT,
	groupsCollectionId INTEGER NOT NULL,
	chatsCollectionId INTEGER NOT NULL,
	PRIMARY KEY (userId),
	FOREIGN KEY (groupsCollectionId) REFERENCES groupsCollection(collectionId),
	FOREIGN KEY (chatsCollectionId) REFERENCES chats(collectionId),
)`

var groupsCollection = `CREATE TABLE IF NOT EXIST groupsCollection(
	collectionId INTEGER NOT NULL,
	groupId INTEGER NOT NULL,
	PRIMARY KEY (collectionId, groupId),
	FOREIGN KEY (groupId) REFERENCES group(groupId),
)`

var group = `CREATE TABLE IF NOT EXIST group(
	groupId INTEGER NOT NULL,
	groupName TEXT NOT NULL,
	groupImage TEXT,
	PRIMARY KEY (groupId),
)`

var chatsCollection = `CREATE TABLE IF NOT EXIST chatsCollection(
	collectionId INTEGER NOT NULL,
	chatId INTEGER NOT NULL,
	PRIMARY KEY (collectionId, chatId),
	FOREIGN KEY (chatId) REFERENCES chat(chatId),
)`

var chat = `CREATE TABLE IF NOT EXIST chat(
	chatId INTEGER NOT NULL,
	PRIMARY KEY (chatId),
)`
