<script>
export default {
	data: function() {
		return {
			intervalId: null,
			errormsg: null,
			loading: false,
			messages: JSON.parse(sessionStorage.chat).messages
		}
	},
	methods: {
		async newMessage(content){
			try {
				let response = await this.$axios.post("/users/"+JSON.parse(sessionStorage.user).id+"/conversations/"+JSON.parse(sessionStorage.chat).id, {
					chat_id: JSON.parse(sessionStorage.chat).id,
                    sender: JSON.parse(sessionStorage.user),
                    date: new Date().toISOString(),
					content: content
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
				sessionStorage.chat = JSON.stringify(response.data)
				this.messages = JSON.parse(sessionStorage.chat).messages
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
	},
	mounted() {
	    this.startMessageLoading();
	},
	beforeRouteLeave(){
		this.stopMessageLoading();
	},
	beforeDestroy() {
    	this.stopMessageLoading();
	},
}
</script>

<template>
    <div>
        <div class="homescreen">
			<div class="list-group-item list-group-item-action" style="left: 0px; margin-block-end: 70px;">
				<div v-for="message in this.messages">
					<div class="message" style="text-align: left; font-size: medium; padding-bottom: 10px;">
						{{ message.sender.name }}	:    {{ message.content }}
					</div>
                </div>
				<div class="btn-group me-2" >
					<input type="text" class="form-control" placeholder="Type message"
					v-model="newMessageContent" @keyup.enter="newMessage(newMessageContent)" style="position: fixed; bottom: 30px; width: 30%;" >
				</div>
            </div>
        </div>
    </div>
</template>

<style>
</style>