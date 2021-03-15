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
              <v-btn text dense color="#ea292f" v-if="selected.length > 0" @click="deleteSelected">
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
          text: 'Student',
          align: 'start',
          value: 'student',
        },
        {text: 'Organization', value: 'organization'},
      ],
      requests: [
        {
          student: 'Camelia Anders',
          userID: 1,
          organization: 'MU',
        },
        {
          student: 'John Lee',
          userID: 2,
          organization: 'MU',
        },
        {
          student: 'Zaynah Ahmad',
          userID: 3,
          organization: 'MUIDS',
        },
        {
          student: 'Fedrick Hams',
          userID: 4,
          organization: 'MUIDS',
        },
        {
          student: 'Bay Leaf',
          userID: 5,
          organization: 'MU',
        },
        {
          student: 'Douglas Stein',
          userID: 6,
          organization: 'MU',
        },
        {
          student: 'Sand Witch',
          userID: 7,
          organization: 'MUIDS',
        },
        {
          student: 'Abra Cadabra',
          userID: 8,
          organization: 'MUIDS',
        },
      ],
    }
  },
  methods: {
    deleteSelected() {
      this.requests = this.requests.filter(r => this.selected.indexOf(r) < 0);
      this.selected = [];
    }
  }
}
</script>

<style scoped>
</style>