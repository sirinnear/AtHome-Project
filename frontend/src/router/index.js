import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '../views/Login.vue'
import IssuerCert from '../views/issuer/IssuerCertificateTable.vue'
import IssuerTransfer from '../views/issuer/TransferRequests.vue'
import AdminCert from '../views/admin/IssueCertificateRequests'
import StudentCert from '../views/student/StudentCertificateTable'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  // {
  //   path: '/about',
  //   name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/issuer/certificates',
    name: 'IssuerCert',
    component: IssuerCert,
  },
  {
    path: '/issuer/transfers',
    name: 'IssuerTransfer',
    component: IssuerTransfer,
  },
  {
    path: '/admin',
    name: 'AdminCert',
    component: AdminCert,
  },
  {
    path: '/student',
    name: 'StudentCert',
    component: StudentCert,
  },
]

const router = new VueRouter({
  mode: 'history',
  routes,
})

router.beforeEach((to, from, next) => {
  const publicPages = ['/login'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem('user');

  if (authRequired && !loggedIn) {
    next('/login');
  } else {
    next();
  }
});

export default router
