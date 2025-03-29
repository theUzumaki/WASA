<script>
export default {
	data: function() {
		return {
			intervalId: null,
			errormsg: null,
			loading: false,
			search: false,
			isGroup: false,
			forward: false,
			message_operations: false,
			message: null,
			reply_id: -1,
			messages: JSON.parse(sessionStorage.chat).messages,
			members: JSON.parse(sessionStorage.chat).members,
			users: [],
			chats: [],
		}
	},
	methods: {
		isBase64Image(content) {
			const base64Pattern = /^data:image\/(png|jpg|jpeg|gif);base64,/;
			return base64Pattern.test(content);
		},
		triggerFileInput() {
			this.$refs.fileInput.click();
		},
		handleFileUpload(event) {
			const file = event.target.files[0];
			const reader = new FileReader();

			reader.onload = (e) => {
				const base64String = e.target.result;
				this.newMessage(base64String);
			};

			reader.readAsDataURL(file);
		},
		async newMessage(content){
			try {
				let formData = new FormData();
                formData.append('chat_id', JSON.parse(sessionStorage.chat).id);
                formData.append('sender_id', JSON.parse(sessionStorage.user).id);
				formData.append('sender_name', JSON.parse(sessionStorage.user).name);
				formData.append('sender_pic', JSON.parse(sessionStorage.user).picture);
                formData.append('date', new Date().toISOString());
                formData.append('content', content);
				formData.append('reply_id', this.reply_id);

                let response = await this.$axios.post("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+JSON.parse(sessionStorage.chat).id,
					formData, {
                    headers: {
                        "Authorization": JSON.parse(sessionStorage.user).id,
                        "Content-Type": "multipart/form-data"
                    }
                });
				sessionStorage.chat = JSON.stringify(response.data)
				this.messages = JSON.parse(sessionStorage.chat).messages
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async forwardMessage(chat) {
			let chatpic= null
			let chatname= "chat"
			
			try {
				let response = await this.$axios.post("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+JSON.parse(sessionStorage.chat).id+"/message/"+this.message.id, {
					id: chat.id.toString()
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				let chat_obt = response.data;

				alert("Message forwarded successfully!");

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.forward= false;
		},
		async newChat(newuser){
			try {
				let response = await this.$axios.post("/users/"+JSON.parse(sessionStorage.user).id+"/conversations", {
					name: "chat",
					members: [JSON.parse(sessionStorage.user), newuser]
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				this.forwardMessage(response.data)
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.search= false
		},
		setMessage(message){
			if ( this.message_operations ) {
				this.message_operations= false;
				this.forward= false;
			}
			else this.message_operations= true;
			this.search= false;
			this.message= message;
		},
		deleteMessage(){
			try {
				let response = this.$axios.delete("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+JSON.parse(sessionStorage.chat).id+"/message/"+this.message.id, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				})
				this.loadMessages();
				this.message_operations= false;
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		commentMessage(emoji){
			try {
				if (this.message.sender.id == JSON.parse(sessionStorage.user).id) throw "You can not comment your own messages"
				let response = this.$axios.put("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+JSON.parse(sessionStorage.chat).id+"/message/"+this.message.id+"/comment", {
					content: emoji
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				})
				this.loadMessages();
				this.message_operations= false;
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		uncommentMessage(){
			try {
				let response = this.$axios.delete("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+JSON.parse(sessionStorage.chat).id+"/message/"+this.message.id+"/comment", {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				})
				this.loadMessages();
				this.message_operations= false;
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async loadMessages(){
			try {
				let response = await this.$axios.get("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+JSON.parse(sessionStorage.chat).id, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				})
				sessionStorage.chat = JSON.stringify(response.data)
				this.messages = JSON.parse(sessionStorage.chat).messages
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async getUsers(name = ""){
			this.errormsg = null;
			try {
				if (name.length < 1) throw "It has to have at least 1 character"
				let response = await this.$axios.get("/users/"+JSON.parse(sessionStorage.user).id+"/search/"+name, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				if (response.data == null) this.users= []
				else {
					this.users = response.data.filter(user => user.id != JSON.parse(sessionStorage.user).id);
				}
				this.search = true;
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async getChats(name = ""){
			this.errormsg = null;
			try {
				if (name.length < 1) throw "It has to have at least 1 character"
				let response = await this.$axios.get("/users/"+JSON.parse(sessionStorage.user).id+"/conversations", {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				console.log(response.data)
				if (response.data == null) this.chats= []
				else {
					this.chats = response.data.filter(chat => chat.name.startsWith(name) && chat.name != "chat" && chat.id != JSON.parse(sessionStorage.chat).id);
				}
				this.search = true;

			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		leaveGroup(){
			try {
				let response = this.$axios.delete("/users/"+JSON.parse(sessionStorage.user).id+"/groups/"+JSON.parse(sessionStorage.chat).id, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				})
				sessionStorage.chat = JSON.stringify(response.data)
				this.messages = JSON.parse(sessionStorage.chat).messages
			} catch (e) {
				this.errormsg = e.toString();
			}
			
			this.$router.push("/home");
		},
		addToGroup(user){
			try {
				let response = this.$axios.put("/users/"+JSON.parse(sessionStorage.user).id+"/groups/"+JSON.parse(sessionStorage.chat).id, 
					user, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				})
				JSON.parse(sessionStorage.chat).members.push(user)
				this.members = this.members.concat(user)
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.search= false
		},
		checkPresence(usertosearch){
			const members = JSON.parse(sessionStorage.chat).members;
			return !(members.some(member => member.id === usertosearch.id));
		},
		checkGroup(){
			const name = JSON.parse(sessionStorage.chat).name;
			if (name == "chat"){
				this.isGroup= false
			} else {
				this.isGroup= true
			}
		},
		getComment(message){

			let c_s = [];
			if (message.comm_senders != null) {
				message.comm_senders.forEach(cs => {
					c_s.push({ comment: cs.comment.content, sender: cs.sender.name });
				});
			}
			return c_s;
		},
		checkProperty(message){
			if (JSON.parse(sessionStorage.user).id == message.sender.id) return true;
			else return false;
		},
		checkCommentProperty(message){
			console.log("user ID: ", JSON.parse(sessionStorage.user).id)
			if (message.comm_senders != null) {
				for (let cs of message.comm_senders) {
					console.log("sender ID: ", cs.sender.id)
					if (cs.sender.id == JSON.parse(sessionStorage.user).id) {
						return true;
					}
				}
			}
			return false
			},
		getChatName() {
			if (JSON.parse(sessionStorage.chat).name == "chat") {
				if (JSON.parse(sessionStorage.user).name === JSON.parse(sessionStorage.chat).members[1].name) {
				  return JSON.parse(sessionStorage.chat).members[0].name;
				} else {
				  return JSON.parse(sessionStorage.chat).members[1].name;
				}
			} else {
				return JSON.parse(sessionStorage.chat).name
			}
		},
		startMessageLoading() {
			this.intervalId = setInterval(() => {
        		this.loadMessages();
    		}, 1000);
    	},
    	stopMessageLoading() {
    		if (this.intervalId) {
        		clearInterval(this.intervalId);
    		}
    	},
		changeSearchState(){
			if (this.search) {
				if (this.forward) {
					this.search= true;
					this.forward= false;
				} else this.search= false;
			} else this.search= true;
			this.users= [];
			this.chats= [];
			return this.search;
		}
	},
	mounted() {
	    this.startMessageLoading();
		this.checkGroup();
	},
	beforeRouteLeave(){
		this.stopMessageLoading();
	},
	beforeUnmount() {
    	this.stopMessageLoading();
	},
}
</script>

<template>
    <div>
		<h1 style="position: absolute; top: 8%">{{ getChatName() }}</h1>
		<div class="homescreen">
			<div class="list-group-item list-group-item-action" style="left: 0px; margin-block-end: 70px;">
				<div v-for="message in messages" :key="message.id">
					<div class="message" style="text-align: left; font-size: medium; padding-top: 15px;">
						<img :src="message.sender.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ message.sender.name }}:<br>
						<div v-if="message.reply_id != -1" style="font-size: small; color: gray;">
							Replying to: {{ messages.find(msg => msg.id === message.reply_id)?.content.slice(0, 32) }}
						</div>
						<button class="btn" @click="setMessage(message)">
							<div v-if="isBase64Image(message.content)">
								<img :src="`${message.content}`" style="width: 200px; height: 200px; object-fit: cover;"/> 
								<div v-if=checkProperty(message) style="display: inline-block; vertical-align: middle;">
									<img src="/svgviewer-output.svg" style="width: 20px; height: 20px;"/>
									<div v-if="message.checkmark" style="display: inline-block; vertical-align: top;">
										<img src="/svgviewer-output.svg" style="width: 20px; height: 20px;"/>
									</div>
								</div>
							</div>
							<div v-else>
								{{ message.content }}
								<div v-if=checkProperty(message) style="display: inline-block; vertical-align: middle;">
									<img src="/svgviewer-output.svg" style="width: 20px; height: 20px;"/>
									<div v-if="message.checkmark" style="display: inline-block; vertical-align: top;">
										<img src="/svgviewer-output.svg" style="width: 20px; height: 20px;"/>
									</div>
								</div>
							</div>
						</button>
					</div>
					<div style="margin-left: 15px; margin-bottom: 10px;">
						<span v-for="(comment, index) in getComment(message)" :key="index">{{ comment.sender }}:  {{ comment.comment }}<span v-if="index < getComment(message).length - 1">, </span></span>
					</div>

				</div>
				<div v-if="search" style="position: absolute; right: 5%; top:10%;">
					<h1 style="font-size: medium;">Users:</h1>
					<div v-for="user in users" :key="user.id">
						<div v-if="checkPresence(user) || (isGroup && forward)">
						<button v-if="isGroup && !forward" type="button" class="btn btn-to-the-right" @click="addToGroup(user)">
							<img :src="user.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ user.name }}
						</button>
						<button v-else-if="forward" type="button" class="btn btn-to-the-right" @click="newChat(user)">
							<img :src="user.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ user.name }}
						</button> <br>
						</div>
					</div>
					<div v-if="forward">
						<h1 style="font-size: medium;">Groups:</h1>
						<div v-for="chat in chats" :key="chat.id">
							<button type="button" class="btn btn-to-the-right" @click="forwardMessage(chat)">
								<img :src="chat.picture" alt="Chat " class="rounded-circle" width="40" height="40"> {{ chat.name }}
							</button> <br>
						</div>
					</div>
				</div>
				<div v-else-if="isGroup && !message_operations" style="position: absolute; left: 80%; top: 12%;">
					<h1 style="font-size: medium;">Members:</h1>
					<div v-for="user in members" :key="user.id">
						<img :src="user.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ user.name }}<br>
					</div>
				</div>
				<div v-if="message_operations" style="position: absolute; top: 50%; left: 80%; transform: translate(-50%, -50%); background: white; border-radius: 10px;">
					<button class="btn" v-if="checkProperty(message)" @click="deleteMessage()">Delete Message</button><br>
					<div v-if="!checkProperty(message)">
						<button class="btn" @click="commentMessage('😊'); message_operations = false">😊</button>
						<button class="btn" @click="commentMessage('😂'); message_operations = false">😂</button>
						<button class="btn" @click="commentMessage('😢'); message_operations = false">😢</button>
						<button class="btn" @click="commentMessage('😡'); message_operations = false">😡</button>
					</div>
					<div v-if="checkCommentProperty(message)">
						<button class="btn" @click="uncommentMessage(); message_operations = false">Uncomment Message</button>
					</div>
					<div>
						<button class="btn" @click="changeSearchState(); forward = true; message_operations = false">Forward Message</button>
					</div>
					<div>
						<button class="btn" @click="reply_id = message.id; message_operations = false">Reply</button>
					</div>
				</div>
				<div class="btn-group me-2" >
					<div style="position: fixed; bottom: 30px; width: 30%;">
						<button v-if="reply_id != -1" @click="reply_id = -1" type="button" class="btn" style="font-size: small; color: gray; margin-bottom: 5px;">
							Click to cancel reply: " {{ messages.find(msg => msg.id === reply_id)?.content.slice(0, 32) }} "
						</button>
						<input type="text" class="form-control" placeholder="Type message"
						v-model="newMessageContent" @keyup.enter="newMessage(newMessageContent)">
					</div>
					<div>
						<button class="btn" @click="triggerFileInput()" style="position: fixed; bottom: 30px; left: 50%; margin-left: 10px;">
							Send Image
						</button>
						<button class="btn" @click=leaveGroup() style="position: fixed; bottom: 30px; left: 55%; margin-left: 5%;">
							Leave group/chat
						</button>
					</div>
					<input v-if="search && (isGroup || forward)" type="text" class="form-control" placeholder="Find users"
						v-model="searchQuery" @keyup.enter="getUsers(searchQuery); getChats(searchQuery); searchQuery = ''" style="position: fixed; top: 80px; left: 80%; width: 15%">
					<button class="btn" v-if="isGroup" @click="changeSearchState()" style="position: fixed; bottom: 30px; left: 70%; margin-left: 5%;">
						Add to group
					</button>
					<button class="btn" v-if="isGroup" @click="$router.push('/settings')" style="position: fixed; bottom: 30px; left: 85%; margin-left: 10px;">
						Group Settings
					</button>
					<input type="file" ref="fileInput" @change="handleFileUpload" style="display: none;" accept="image/*">
				</div>
			</div>
        </div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    </div>
</template>

<style>
</style>
