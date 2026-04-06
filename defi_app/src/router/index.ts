import { createRouter, createWebHistory } from 'vue-router'
import YieldsPage from '../pages/YieldsPage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [{ path: '/', component: YieldsPage }],
})

export default router
