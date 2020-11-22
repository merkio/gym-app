<template>
  <v-card
    outlined
  >
    <v-card-title>
      {{ program.text.substring(0, 30) }}
    </v-card-title>


    <v-card-subtitle>
      ${{ product.text }}
    </v-card-subtitle>
    <v-card-text>
      <v-row>
        <v-text-field
          outlined
          readonly
        >{{program}}</v-text-field>
      </v-row>
    </v-card-text>
    <v-card-actions>
      <v-btn
          color="green"
          outlined
          @click="editProgram(i)"
      >
        <v-icon>fa-edit</v-icon>
        Update
      </v-btn>
      <v-btn
          color="warning"
          outlined
          @click="removeProgram(i)"
      >
        <v-icon>fa-cross</v-icon>
        Cancel
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: "ProgramDetails",
  props: {
    program: {},
  }
}
</script>

<style scoped>

</style>
