<script>
	import Changename from '@/components/change name.vue';
	import { changeusername } from '@/services/change.js';
	export default {
	components: { Changename },
	data() {
		return {
		errorMessage: ''
		};
	},
	methods: {
		async changeuser_name(credentials) {
			this.errorMessage="";
			try {
				const user = await changeusername(credentials);
				console.log('Name change successful', user);
				// Handle successful login (e.g., redirect, store token)
				this.errorMessage='';
				this.$router.push({ name: 'profile' });
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
    <h2>Change Name</h2>
    <Changename @change="changeuser_name" />
    <p v-if="errorMessage" class="text-danger">{{ errorMessage}}</p>
	</div>
</template>

<style>
</style>
