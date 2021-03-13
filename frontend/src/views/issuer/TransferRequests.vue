<template>
  <div>
    <AppBar :authenticated=true />
    <v-content>
      <v-container>
        <v-data-table
            v-model="selected"
            :headers="headers"
            :items="requests"
            :items-per-page="5"
            show-select
            :single-select=false
            item-key="user"
            class="elevation-1 text-left">
          <template v-slot:top>
            <v-toolbar flat dense>
              <v-btn text dense color="#ea292f" v-if="selected.length > 0">
                <v-icon
                    left
                >
                  mdi-check
                </v-icon>
                Approve Selected
              </v-btn>
            </v-toolbar>
          </template>
        </v-data-table>
      </v-container>
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/AppBar';

export default {
  name: "TransferRequests",
  components: {
    AppBar,
  },
  props: {
    authenticated: Boolean,
  },
  data () {
    return {
      selected: [],
      headers: [
        {
          text: 'User',
          align: 'start',
          value: 'user',
        },
        {text: 'Organization', value: 'organization'},
      ],
      requests: [
        {
          user: 'A',
          organization: 'MUIC',
        },
        {
          user: 'B',
          organization: 'MUIC',
        },
        {
          user: 'C',
          organization: 'MUIDS',
        },
        {
          user: 'D',
          organization: 'MUIDS',
        },
        {
          user: 'A',
          organization: 'MUIC',
        },
        {
          user: 'B',
          organization: 'MUIC',
        },
        {
          user: 'C',
          organization: 'MUIDS',
        },
        {
          user: 'D',
          organization: 'MUIDS',
        },
      ],
    }
  }
}
</script>

<style scoped>
</style>