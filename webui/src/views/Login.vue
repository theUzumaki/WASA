import ErrorMsg from '@/components/ErrorMsg.vue';

<template>
    <ErrorMsg v-if="error" :msg="error"></ErrorMsg>
    <div class="login">
        <h1>Login</h1>
        <form @submit.prevent="handleSubmit">
            <div>
                <label for="username">Username:</label>
                <input type="text" id="username" v-model="username" @keyup.enter="handleSubmit" required/>
            </div>
        </form>
    </div>
</template>
<script>


export default {
    data() {
        return {
            username: '',
            error: null
        };
    },
    methods: {
        async handleSubmit() {
            
            try {
                if (this.username.length < 3 || this.username.length > 16) throw "It has to be between 3 and 16 characters long"
                let response = await this.$axios.post('/session', {
                    name: this.username,
                });
                sessionStorage.user= JSON.stringify(response.data);
            } catch (e) {
                this.error = e.toString();
                return;
            }

            this.$router.push("/home");
        }
    }
};
</script>

<style scoped>
.login {
    margin-top: 50px;
    max-width: 300px;
    margin: 0 auto;
    padding: 1em;
    border: 0px solid #ccc;
}

label {
    display: block;
    margin-bottom: 0.5em;
}

input {
    width: 100%;
    padding: 0.5em;
    margin-bottom: 1em;
}

</style>