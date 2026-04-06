import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('../layouts/MainLayout.vue'),
      children: [
        {
          path: '', // correspond à http://localhost:5173/
          component: () => import('../pages/YieldsPage.vue'),
        },
      ],
    },
  ],
})

export default router
