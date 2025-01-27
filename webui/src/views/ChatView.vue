<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			chat: JSON.parse(sessionStorage.chat).messages
		}
	},
	methods: {
		async newMessage(content){
			try {
				let response = await this.$axios.post("/users/"+sessionStorage.userId+"/conversations/"+sessionStorage.chat_id, {
					chat_id: parseInt(sessionStorage.chat_id),
                    senders: sessionStorage.userId,
                    date: new Date().toISOString(),
					content: content
				}, {
					headers: {
						"Authorization": sessionStorage.userId
					}
				});
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
	}
}
</script>

<template>
    <div>
        <div class="homescreen">
            <div class="list-group-item list-group-item-action">
                <div v-for="message in this.chat">
                    <div class="d-flex w-100 justify-content-between">
                        <h5 class="mb-1">{{ message.name }}</h5>
                        <small>{{ message.id }}</small>
                    </div>
                </div>
            </div>
			<div class="btn-group me-2">
				<input type="text" class="form-control" placeholder="Type message" v-model="newMessageContent" @keyup.enter="newMessage(newMessageContent)">
			</div>
        </div>
    </div>
</template>

<style>
</style>