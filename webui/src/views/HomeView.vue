<script>

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			chats: null,
			users: null,
			search: false,
			showGroupForm: false,
			selectedUsers: [],
			username: null,
			groupName: null,
			groupPicture: null,
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
		async addUserToGroup(user){
			const members = this.selectedUsers;
			if (!(members.some(member => member.id === user.id)))
				this.selectedUsers.push(user)
		},
		async newChat(userlist, groupname, grouppic){
			this.search = false;
			this.errormsg = null;
			let chatpic= null
			let chatname= "chat"
			console.log("BEFORE: ", userlist)
			let newmembers= [...userlist]
			try {
				newmembers.push({
					name: JSON.parse(sessionStorage.user).name,
					id: parseInt(JSON.parse(sessionStorage.user).id),
					picture: JSON.parse(sessionStorage.user).picture
				});
				
				if (this.showGroupForm) {
					chatpic= grouppic
					chatname= groupname
				}

				let response = await this.$axios.post("/users/"+JSON.parse(sessionStorage.user).id+"/conversations", {
					name: chatname,
					members: newmembers,
					picture: chatpic
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				sessionStorage.chat= JSON.stringify(response.data);
				this.$router.push("/chat");
			} catch (e) {
				console.log(e.toString())
				this.errormsg = e.toString();
			}
		},
		async getUsers(name = ""){
			this.errormsg = null;
			this.username= JSON.parse(sessionStorage.user).name
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
			try {
				let response = await this.$axios.get("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+chat.id, {
						headers: {
							"Authorization": JSON.parse(sessionStorage.user).id
						}
					})
				sessionStorage.chat = JSON.stringify(response.data)
				this.$router.push("/chat");
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		onFileChange(event) {
			const file = event.target.files[0];
			const reader = new FileReader();

			reader.onload = (e) => {
				const base64String = e.target.result;
				this.groupPicture= base64String;
			};

			reader.readAsDataURL(file);
		},
		displayMemberName(chat) {
			if (chat.name == "chat") {
				if (JSON.parse(sessionStorage.user).name === chat.members[1].name) {
				  return chat.members[0].name;
				} else {
				  return chat.members[1].name;
				}
			} else {
				return chat.name
			}
		},
		displayChatPic(chat){
			if (chat.name == "chat") {
				if (JSON.parse(sessionStorage.user).name === chat.members[1].name) {
				  return chat.members[0].picture;
				} else {
				  return chat.members[1].picture;
				}
			} else {
				console.log("NAME: ", chat.name, "PICTURE: ", chat.picture)
				return chat.picture
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
	beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.startChatLoading();
    });
  },
	beforeRouteLeave(){
		this.stopChatLoading();
	},
	beforeUnmount() {
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
				<div v-if= "showGroupForm == false">
					<div v-for="chat in chats" :key="chat.id" style="position: relative">
						<button type="button" class="btn" @click="openChat(chat)">
							<img :src=displayChatPic(chat) alt="User Profile" class="rounded-circle" width="40" height="40"> {{ displayMemberName(chat) }}
						</button> <br>
					</div>
				</div>
				<div v-if="search" style="position: absolute; top: 0px; right:2%">
					<div v-for="user in users" :key="user.id">
						<button v-if="user.name != username" type="button" class="btn btn-to-the-right" @click="showGroupForm ? addUserToGroup(user) : newChat([user], null, user.picture)">
							<img :src="user.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ user.name }}
						</button> <br>
					</div>
				</div>
				<div v-if="showGroupForm" class="group-form" style="position: relative; top: 75px; width: 30%;">
					<div class="form-group">
						<label for="groupName">Group Name</label>
						<input type="text" class="form-control" id="groupName" v-model="groupName">
						<br>
					</div>
					<div class="form-group">
						<label for="groupPicture">Group Picture</label>
						<input type="file" class="form-control" id="groupPicture" @change="onFileChange">
						<br>
					</div>
					<div class="form-group">
						<label for="groupMembers">Search users in the bar on the corner</label>
						<ul class="list-group">
							<li v-for="user in selectedUsers" :key="user.id" class="list-group-item">{{ user.name }}</li>
						</ul>
						<br>
					</div>
					<button type="button" class="btn btn-primary" @click="newChat(selectedUsers, groupName, groupPicture)">Create Group</button>
					<button type="button" class="btn btn-secondary" @click="showGroupForm = false">Cancel</button>
				</div>
			</div>
			<div class="btn-toolbar mb-2 mb-md-0 right" >
				<button type="button" class="btn" @click="showGroupForm = true">
					New group
				</button>
				
				<div class="btn-group me-2">
					<input v-if="showGroupForm" type="text" class="form-control" placeholder="User to add to group" v-model="searchQuery" @keyup.enter="getUsers(searchQuery)">
					<input v-else type="text" class="form-control" placeholder="New chat" v-model="searchQuery" @keyup.enter="getUsers(searchQuery)">
				</div>
			</div>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
