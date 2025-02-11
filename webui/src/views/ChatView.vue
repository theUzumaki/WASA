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
						<img :src="message.sender.picture" alt="User Profile" class="rounded-circle" width="40" height="40"> {{ message.sender.name }}:<br>
						<div v-if="isBase64Image(message.content)">
							<img :src="`${message.content}`" style="width: 200px; height: 200px; object-fit: cover;"/>
						</div>
						<div v-else>
							{{ message.content }}
						</div>
					</div>
                </div>
				<div class="btn-group me-2" >
					<input type="text" class="form-control" placeholder="Type message"
					v-model="newMessageContent" @keyup.enter="newMessage(newMessageContent)" style="position: fixed; bottom: 30px; width: 30%;" >
					<button @click="triggerFileInput" style="position: fixed; bottom: 30px; right: 35%; margin-left: 10px;">
						Add Image
					</button>
					<input type="file" ref="fileInput" @change="handleFileUpload" style="display: none;" accept="image/*">
				</div>
            </div>
        </div>
    </div>
</template>

<style>
</style>