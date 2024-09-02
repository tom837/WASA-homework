import {createRouter, createWebHashHistory} from 'vue-router'
import Login from '../views/Login.vue'
import Homepage from '../views/Homepage.vue'
import Comments from '../views/comments.vue'
import Profile from '../views/profile.vue'
import uploadpic from '../views/uploadpic.vue'
import changename from '../views/changename.vue'
import search from '../views/search.vue'
import Userprofile from '../views/userprofile.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/',name:'Login', component: Login},
		{path: '/user',name:'Homepage', component: Homepage},
		{path: '/comments/:photoId', name:'Comments', component: Comments},
		{path: '/user/profile', name:"profile", component: Profile},
		{path: '/user/upload', name:'upload', component: uploadpic},
		{path: '/user/changename', name:"changename", component: changename},
		{path: '/user/search', name:"search", component:search},
		{path: '/user/:id/:username/profile', name:"userprofile", component: Userprofile}
	]
})

export default router
