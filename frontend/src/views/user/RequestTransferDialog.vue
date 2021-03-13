<template>
  <v-dialog
      v-model="dialog"
      persistent
      max-width="600px"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn
          text
          dark
          class="white--text"
          v-bind="attrs"
          v-on="on"
      >
        Request Transfer
      </v-btn>
    </template>
    <v-card>
      <v-card-title>
        <span class="headline">Request Transfer</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-text-field
              label="Organization Name"
              hint="Organization to transfer to"
          ></v-text-field>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
            color="blue darken-1"
            text
            @click="dialog = false"
        >
          Close
        </v-btn>
        <v-btn
            color="blue darken-1"
            text
            @click="dialog = false"
        >
          Request
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "RequestTransferDialog",
  data: () => ({
    dialog: false,
  }),
}
</script>

<style scoped>

</style>