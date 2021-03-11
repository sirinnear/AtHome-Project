<template>
  <div>
    <AppBar :authenticated=true />
    <v-content>
      <DataTable>
<!--        <v-simple-table class="text-left">-->
<!--          <thead>-->
<!--          <tr>-->
<!--            <th>-->
<!--              Name-->
<!--            </th>-->
<!--            <th>-->
<!--              Issuer-->
<!--            </th>-->
<!--            <th colspan="2" class="text-center">-->
<!--              Actions-->
<!--            </th>-->
<!--          </tr>-->
<!--          </thead>-->
<!--          <tbody>-->
<!--          <tr-->
<!--              v-for="item in certificates"-->
<!--              :key="item.name"-->
<!--          >-->
<!--            <td>{{ item.name }}</td>-->
<!--            <td>{{ item.issuer }}</td>-->
<!--            <td class="text-center">-->
<!--              <div v-if="item.signed">-->
<!--                Signed-->
<!--              </div>-->
<!--              <div v-else>-->
<!--                <v-btn>Request Signing</v-btn>-->
<!--              </div>-->
<!--            </td>-->
<!--            <td class="text-center">-->
<!--              <v-btn class="ma-0">Export</v-btn>-->
<!--            </td>-->
<!--          </tr>-->
<!--          </tbody>-->
<!--        </v-simple-table>-->
      </DataTable>
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/AppBar';
import DataTable from "@/components/DataTable";

export default {
  name: "Dashboard",
  components: {
    AppBar,
    DataTable,
  },
  props: {
    authenticated: Boolean,
  },
  data () {
    return {
      certificates: [
        {
          name: 'A',
          issuer: 'MUIC',
          signed: false,
        },
        {
          name: 'B',
          issuer: 'MUIC',
          signed: true,
        },
        {
          name: 'C',
          issuer: 'MUIDS',
          signed: false,
        },
        {
          name: 'D',
          issuer: 'MUIDS',
          signed: false,
        },
      ],
    }
  },
}
</script>

<style scoped>
</style>