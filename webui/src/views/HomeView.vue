<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			chats: null,
			users: null,
			search: false,
			username: sessionStorage.username,
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
		async newChat(userId, username){
			this.search = false;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/users/"+sessionStorage.userId+"/conversations", {
					name: "chat",
					members: [{
						name: sessionStorage.username,
						id: parseInt(sessionStorage.userId)
					}, {
						name: username,
						id: parseInt(userId)
					}]
				}, {
					headers: {
						"Authorization": sessionStorage.userId
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
				if (name.length < 3 || name.length > 16) throw "It has to be between 3 and 16 characters long"
				let response = await this.$axios.get("/users/"+sessionStorage.userId+"/search/"+name, {
					headers: {
						"Authorization": sessionStorage.userId
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
				let response = await this.$axios.get("/users/"+sessionStorage.userId+"/conversations", {
					headers: {
						"Authorization": sessionStorage.userId
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
		}
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0 right" >
				<div class="btn-group me-2">
					<button type="button" class="btn" @click="getMyConversations">
						Load chats
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn" @click="newChat">
						New chat
					</button>
				</div>
				<div class="btn-group me-2">
					<input type="text" class="form-control" placeholder="Search user" v-model="searchQuery" @keyup.enter="getUsers(searchQuery)">
				</div>
			</div>
			<div class= "homescreen">
				<div v-for="chat in chats">
					<button type="button" class="btn" @click="openChat(chat)">
						{{ chat.members[1].name }}
					</button>
				</div>
				<div v-if="search" style="position: absolute; top:0px; left:77%">
					<div v-for="user in users">
						<button v-if="user.name != username" type="button" class="btn btn-to-the-right" @click="newChat(user.id, user.name)">
							{{ user.name }}
						</button> <br>
					</div>
				</div>
			</div>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
