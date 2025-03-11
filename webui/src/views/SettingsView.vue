<script>
export default {
    data() {
        return {
            username: '',
            groupname: '',
            profilePicture: null,
            profilePictureUrl: '',
            errormsg: null,
            isGroup: false,
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
                if (this.isGroup){
                    let response = await this.$axios.put("/users/"+JSON.parse(sessionStorage.user).id+"/groups/"+JSON.parse(sessionStorage.chat).id+"/name", {
                        name: this.groupname,
                    }, {
                        headers: {
                            "Authorization": JSON.parse(sessionStorage.user).id
                        }
                    });
                    let newpic= JSON.parse(sessionStorage.user).picture
                    if (this.profilePicture != null) {
                        response = await this.$axios.put("/users/"+JSON.parse(sessionStorage.user).id+"/groups/"+JSON.parse(sessionStorage.chat).id+"/picture", {
                            picture: this.profilePicture
                        }, {
                            headers: {
                                "Authorization": JSON.parse(sessionStorage.user).id
                            }
                        });
                        newpic= this.profilePicture
                    }
                    sessionStorage.chat = JSON.stringify({
                        id: JSON.parse(sessionStorage.chat).id,
                        name: this.groupname,
                        members: JSON.parse(sessionStorage.chat).members,
                        messages: JSON.parse(sessionStorage.chat).messages,
                        picture: newpic
                    })
                    this.$router.push("/chat");
                } else {
                    let response = await this.$axios.put("/users/"+JSON.parse(sessionStorage.user).id+"/name", {
                        id: JSON.parse(sessionStorage.user).id,
                        name: this.username,
                        picture: this.profilePicture
                    }, {
                        headers: {
                            "Authorization": JSON.parse(sessionStorage.user).id
                        }
                    });
                    let newpic= JSON.parse(sessionStorage.user).picture
                    if (this.profilePicture != null) {
                        response = await this.$axios.put("/users/"+JSON.parse(sessionStorage.user).id+"/picture", {
                            id: JSON.parse(sessionStorage.user).id,
                            name: this.username,
                            picture: this.profilePicture
                        }, {
                            headers: {
                                "Authorization": JSON.parse(sessionStorage.user).id
                            }
                        });
                        newpic= this.profilePicture
                    }
                    sessionStorage.user = JSON.stringify({
                        id: JSON.parse(sessionStorage.user).id,
                        name: this.username,
                        picture: newpic
                    })
                }
			} catch (e) {
                if (e.toString() == "AxiosError: Request failed with status code 400") {
                    this.errormsg = "Name already taken"
                }
                else {
                    this.errormsg = e.toString();
                }
            }
            this.username = '';
            this.groupname = '';
            this.profilePicture = null;
            this.profilePictureUrl = '';
        },
        checkGroup(){
			const chat_name = JSON.parse(sessionStorage.chat).name;
			if (chat_name == "chat"){
				this.isGroup= false
			} else if (chat_name == null){
                this.isGroup= false
            } else {
				this.isGroup= true
			}
		},
        setGroupFalse(){
            this.isGroup= false;
        }
    },
    beforeRouteEnter(to, from, next) {
        console.log("BEFORE ROUTE")
        if (from.path === '/chat') {
            next(vm => {
                vm.checkGroup();
            });
        } else {
            next(vm => {
                vm.setGroupFalse();
            })
        }
    },
};
</script>

<template>
    <div class="settings-view">
        <h1>Settings</h1>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <form @submit.prevent="saveSettings">
            <div class="form-group">
                <div v-if="isGroup">
                    <label for="groupname">Group name:</label>
                    <input type="text" id="groupname" v-model="groupname" required />
                </div>
                <div v-else>
                    <label for="username">Username:</label>
                    <input type="text" id="username" v-model="username" required />
                </div>
            </div>
            <div class="form-group">
                <div v-if="isGroup">
                    <label for="profilePicture">Group Picture:</label>
                    <input type="file" id="groupPicture" @change="handleFileUpload" />
                </div>
                <div v-else>
                    <label for="profilePicture">Profile Picture:</label>
                    <input type="file" id="profilePicture" @change="handleFileUpload" />
                </div>
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