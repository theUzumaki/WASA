<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			chats: null,
			users: null,
			search: false,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async newChat(userId, username, userpicture){
			this.search = false;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/users/"+JSON.parse(sessionStorage.user).id+"/conversations", {
					name: "chat",
					members: [{
						name: JSON.parse(sessionStorage.user).name,
						id: parseInt(JSON.parse(sessionStorage.user).id),
						picture: JSON.parse(sessionStorage.user).picture
					}, {
						name: username,
						id: parseInt(userId),
						picture: userpicture
					}]
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				sessionStorage.chat= JSON.stringify(response.data);
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.$router.push("/chat");
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
		async getMyConversations(){
			try {
				let response = await this.$axios.get("/users/"+JSON.parse(sessionStorage.user).id+"/conversations", {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				this.chats = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async openChat(chat){
			sessionStorage.chat= JSON.stringify(chat)
			this.$router.push("/chat");
		},
		displayMemberName(chat) {
		  if (JSON.parse(sessionStorage.user).name === chat.members[1].name) {
			return chat.members[0].name;
		  } else {
			return chat.members[1].name;
		  }
		},
		startChatLoading() {
			this.intervalId = setInterval(() => {
        		this.getMyConversations();
    		}, 1000);
    	},
    	stopChatLoading() {
    		if (this.intervalId) {
        		clearInterval(this.intervalId);
    		}
    	},
	},
	mounted() {
	    this.startChatLoading();
	},
	beforeRouteLeave(){
		this.stopChatLoading();
	},
	beforeDestroy() {
    	this.stopChatLoading();
	},
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class= "homescreen">
				<div v-for="chat in chats" style="position: absolute; top:50px;">
					<button type="button" class="btn" @click="openChat(chat)">
						{{ displayMemberName(chat) }}
					</button>
				</div>
				<div v-if="search" style="position: absolute; top:50px; left:77%">
					<div v-for="user in users">
						<button v-if="user.name != username" type="button" class="btn btn-to-the-right" @click="newChat(user.id, user.name, user.picture)">
							<img :src="user.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ user.name }}
						</button> <br>
					</div>
				</div>
			</div>
			<div class="btn-toolbar mb-2 mb-md-0 right" >
				<div class="btn-group me-2">
					<button type="button" class="btn" @click="getMyConversations">
						New group
					</button>
				</div>
				<div class="btn-group me-2">
					<input type="text" class="form-control" placeholder="Find user to chat" v-model="searchQuery" @keyup.enter="getUsers(searchQuery)">
				</div>
			</div>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
