<script>
	import submitLogin from '@/components/login.vue';
	import { login } from '@/services/auth.js';
	export default {
	components: { submitLogin },
	data() {
		return {
		errorMessage: '',
		};
	},
	methods: {
		async handleLogin(credentials) {
			this.errorMessage="";
			try {
				const user = await login(credentials);
				console.log('Login successful', user);
				// Handle successful login (e.g., redirect, store token)
				this.errorMessage='';
				this.$router.push({ name: 'Homepage' });
			} catch (error) {
				this.errorMessage = error.message;
				console.error('Caught an error: '+ error.message);
			}
		}
	}
};
</script>

<template>
	<div>
    <h2>Login Page</h2>
    <submitLogin @login="handleLogin" />
    <p v-if="errorMessage" class="text-danger">{{ errorMessage}}</p>
	</div>
</template>

<style>
</style>
