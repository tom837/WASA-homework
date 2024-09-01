import {createRouter, createWebHashHistory} from 'vue-router'
import Login from '../views/Login.vue'
import Homepage from '../views/Homepage.vue'
import Comments from '../views/comments.vue'
import Profile from '../views/profile.vue'
import uploadpic from '../views/uploadpic.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/',name:'Login', component: Login},
		{path: '/user',name:'Homepage', component: Homepage},
		{path: '/comments/:photoId', name:'Comments', component: Comments},
		{path: '/user/profile', name:"profile", component: Profile},
		{path: '/user/ipload', name:'upload', component: uploadpic}
	]
})

export default router
