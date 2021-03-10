<template>
  <div>
    <AppBar :authenticated=true />
  </div>
</template>

<script>
import AppBar from '@/components/AppBar';

export default {
  name: "Dashboard",
  components: {
    AppBar,
  },
  props: {
    authenticated: Boolean,
  },
}
</script>

<style scoped>

</style>