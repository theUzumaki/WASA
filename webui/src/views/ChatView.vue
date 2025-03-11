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
			messages: JSON.parse(sessionStorage.chat).messages
		}
	},
	methods: {
		isBase64Image(content) {
			const base64Pattern = /^data:image\/(png|jpg|jpeg);base64,/;
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
		async forwardMessage(user) {
			let chatpic= null
			let chatname= "chat"
			
			try {
				let response = await this.$axios.post("/users/"+JSON.parse(sessionStorage.user).id+"/conversations", {
					name: chatname,
					members: [{
						name: JSON.parse(sessionStorage.user).name,
						id: parseInt(JSON.parse(sessionStorage.user).id),
						picture: JSON.parse(sessionStorage.user).picture
					}, {
						name: user.name,
						id: user.id,
						picture: user.picture
					}],
					picture: chatpic
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				let chat_obt = response.data;
				let formData = new FormData();
                formData.append('chat_id', chat_obt.id);
                formData.append('sender_id', JSON.parse(sessionStorage.user).id);
				formData.append('sender_name', JSON.parse(sessionStorage.user).name);
				formData.append('sender_pic', JSON.parse(sessionStorage.user).picture);
                formData.append('date', new Date().toISOString());
                formData.append('content', this.message.content);

                response = this.$axios.post("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+chat_obt.id,
					formData, {
                    headers: {
                        "Authorization": JSON.parse(sessionStorage.user).id,
                        "Content-Type": "multipart/form-data"
                    }
                });

				alert("Message forwarded successfully!");

			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		setMessage(message){
			if ( this.message_operations ) {
				this.message_operations= false;
				this.forward= false;
			}
			else this.message_operations= true;
			this.message= message;
			console.log("MESSAGE: ", this.message)
			console.log(" operations: ", this.message_operations)
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
				this.users = response.data;
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
			if (this.search == true) this.search= false;
			else this.search= true;
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
					<div class="message" style="text-align: left; font-size: medium; padding-bottom: 10px;">
						<img :src="message.sender.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ message.sender.name }}:<br>
						<button @click="setMessage(message)">
							<div v-if="isBase64Image(message.content)">
								<img :src="`${message.content}`" style="width: 200px; height: 200px; object-fit: cover;"/> 
							</div>
							<div v-else>
								{{ message.content }}
							</div>
						</button>
					</div>
					<div>
						<span v-for="(comment, index) in getComment(message)" :key="index">{{ comment.sender }} :{{ comment.comment }}<span v-if="index < getComment(message).length - 1">, </span></span>
					</div>

                </div>
				<div v-if="search" style="position: absolute; left:77%">
					<div v-for="user in users" :key="user.id">
						<div v-if="checkPresence(user)">
						<button v-if="isGroup" type="button" class="btn btn-to-the-right" @click="addToGroup(user)">
							<img :src="user.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ user.name }}
						</button>
						<button v-else-if="forward" type="button" class="btn btn-to-the-right" @click="forwardMessage(user)">
							<img :src="user.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ user.name }}
						</button> <br>
						</div>
					</div>
				</div>
				<div v-if="message_operations" style="position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); background: white; padding: 20px; border: 1px solid #ccc; border-radius: 10px;">
					<button v-if="checkProperty(message)" @click="deleteMessage()">Delete Message</button><br>
					<div v-if="!checkProperty(message)">
						<button @click="commentMessage('😊')">😊</button>
						<button @click="commentMessage('😂')">😂</button>
						<button @click="commentMessage('😢')">😢</button>
						<button @click="commentMessage('😡')">😡</button>
					</div>
					<div v-if="checkCommentProperty(message)">
						<button @click="uncommentMessage">Uncomment Message</button>
					</div>
					<div>
						<button @click="changeSearchState(); forward = true">Forward Message</button>
					</div>
				</div>
				<div class="btn-group me-2" >
					<input type="text" class="form-control" placeholder="Type message"
					v-model="newMessageContent" @keyup.enter="newMessage(newMessageContent)" style="position: fixed; bottom: 30px; width: 30%;" >
					<button @click="triggerFileInput" style="position: fixed; bottom: 30px; left: 50%; margin-left: 10px;">
						Send Image
					</button>
					<button @click=leaveGroup() style="position: fixed; bottom: 30px; right: 32%; margin-left: 5%;">
						Leave group/chat
					</button>
					<input v-if="search && (isGroup || forward)" type="text" class="form-control" placeholder="Find user to add"
					v-model="searchQuery" @keyup.enter="getUsers(searchQuery)" style="position: fixed; top: 80px; left: 80%; width: 15%">
					<button v-if="isGroup" @click="changeSearchState()" style="position: fixed; bottom: 30px; right: 25%; margin-left: 5%;">
						Add to group
					</button>
					<button v-if="isGroup" @click="$router.push('/settings')" style="position: fixed; bottom: 30px; right: 15%; margin-left: 10px;">
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