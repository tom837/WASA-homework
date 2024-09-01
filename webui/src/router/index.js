import {createRouter, createWebHashHistory} from 'vue-router'
import Login from '../views/Login.vue'
import Homepage from '../views/Homepage.vue'
import Comments from '../views/comments.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/',name:'Login', component: Login},
		{path: '/user',name:'Homepage', component: Homepage},
		{path: '/comments', name:'Comments', component: Comments}
	]
})

export default router
