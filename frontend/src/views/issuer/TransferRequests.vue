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
            item-key="userID"
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
import AppBar from '@/components/IssuerAppBar';

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
          userID: 1,
          organization: 'MUIC',
        },
        {
          user: 'B',
          userID: 2,
          organization: 'MUIC',
        },
        {
          user: 'C',
          userID: 3,
          organization: 'MUIDS',
        },
        {
          user: 'D',
          userID: 4,
          organization: 'MUIDS',
        },
        {
          user: 'A',
          userID: 5,
          organization: 'MUIC',
        },
        {
          user: 'B',
          userID: 6,
          organization: 'MUIC',
        },
        {
          user: 'C',
          userID: 7,
          organization: 'MUIDS',
        },
        {
          user: 'D',
          userID: 8,
          organization: 'MUIDS',
        },
      ],
    }
  }
}
</script>

<style scoped>
</style>