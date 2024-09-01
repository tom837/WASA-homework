<template>
    <div class="container mt-5">
        <div class="row justify-content-center">
            <div class="col-md-6">
                <h2 class="text-center">Login</h2>
                <form @submit.prevent="submitLogin">
                    <div class="form-group">
                        <label for="username">Username</label>
                        <input
                            type="text"
                            v-model="name"
                            class="form-control"
                            id="name"
                            placeholder="Enter your username"
                            required
                        />
                    </div>
                    <button type="submit" class="btn btn-primary btn-block">Login</button>
                </form>
                <p v-if="errorMessage" class="text-danger mt-3">{{ errorMessage }}</p>
            </div>
        </div>
    </div>
</template>

<script>
    export default{
        data(){
            return {
                UserName: '',
                errorMessage: ''
            };
        },
        methods: {
            submitLogin() {
                //console.log('Button pressed!');
                this.$emit('login', { name: this.name});
                this.errorMessage = '';
            }
        }
    };
</script>