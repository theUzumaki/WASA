<script>
export default {
    data() {
        return {
            username: '',
            profilePicture: null,
            profilePictureUrl: '',
            errormsg: null,
        };
    },
    methods: {
        handleFileUpload(event) {
			const file = event.target.files[0];
			const reader = new FileReader();

			reader.onload = (e) => {
				const base64String = e.target.result;
				this.profilePicture= base64String;
			};

            this.profilePictureUrl= URL.createObjectURL(file)
			reader.readAsDataURL(file);
        },
        async saveSettings() {
            try {
				let response = await this.$axios.put("/users/"+JSON.parse(sessionStorage.user).id+"/name", {
                    id: JSON.parse(sessionStorage.user).id,
                    name: this.username,
                    picture: this.profilePicture
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
                response = await this.$axios.put("/users/"+JSON.parse(sessionStorage.user).id+"/picture", {
                    id: JSON.parse(sessionStorage.user).id,
                    name: this.username,
                    picture: this.profilePicture
				}, {
					headers: {
						"Authorization": JSON.parse(sessionStorage.user).id
					}
				});
                sessionStorage.user = JSON.stringify({
                    id: JSON.parse(sessionStorage.user).id,
                    name: this.username,
                    picture: this.profilePicture
                })
			} catch (e) {
                this.errormsg = e.toString();
			}
            this.username = '';
            this.profilePicture = null;
            this.profilePictureUrl = '';
        }
    }
};
</script>

<template>
    <div class="settings-view">
        <h1>Settings</h1>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <form @submit.prevent="saveSettings">
            <div class="form-group">
                <label for="username">Username:</label>
                <input type="text" id="username" v-model="username" required />
            </div>
            <div class="form-group">
                <label for="profilePicture">Profile Picture:</label>
                <input type="file" id="profilePicture" @change="handleFileUpload" />
            </div>
            <button type="submit">Save</button>
        </form>
        <div v-if="profilePictureUrl">
            <h2>Preview:</h2>
            <img :src="profilePictureUrl" alt="Profile Picture Preview" />
        </div>
    </div>
</template>


<style scoped>
.settings-view {
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 4px;
}
.form-group {
    margin-bottom: 15px;
}
label {
    display: block;
    margin-bottom: 5px;
}
input[type="text"],
input[type="file"] {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
}
button {
    padding: 10px 15px;
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}
button:hover {
    background-color: #0056b3;
}
img {
    max-width: 100%;
    height: auto;
    margin-top: 15px;
}
</style>