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
    props: true,
  },
  {
    path: '/issuer/transfers',
    name: 'IssuerTransfer',
    component: IssuerTransfer,
    props: true,
  },
  {
    path: '/admin',
    name: 'AdminCert',
    component: AdminCert,
    props: true,
  },
  {
    path: '/student',
    name: 'StudentCert',
    component: StudentCert,
    props: true
  },
]

const router = new VueRouter({
  routes,
})

export default router
