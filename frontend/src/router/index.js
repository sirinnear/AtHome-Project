import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '../views/Login.vue'
import IssuerCert from '../views/issuer/IssuerCertificateTable.vue'
import IssuerTransfer from '../views/issuer/TransferRequests.vue'
import AdminCert from '../views/admin/IssueCertificateRequests'
import UserCert from '../views/user/UserCertificateTable'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
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
    component: IssuerCert
  },
  {
    path: '/issuer/transfers',
    name: 'IssuerTransfer',
    component: IssuerTransfer
  },
  {
    path: '/admin',
    name: 'AdminCert',
    component: AdminCert
  },
  {
    path: '/user',
    name: 'UserCert',
    component: UserCert
  },
]

const router = new VueRouter({
  routes
})

export default router
