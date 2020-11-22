<template>
    <v-card outline sm3>
      <v-card-title>
        {{ exercise.title }}
      </v-card-title>
      <v-card-subtitle>
        <v-item-group multiple>
          <v-subheader>Tags</v-subheader>
          <v-item v-for="tag in exercise.tags.split(',')" :key="tag">
            <v-chip>
              {{ tag }}
            </v-chip>
          </v-item>
        </v-item-group>
      </v-card-subtitle>
      <v-card-text>
        <v-item-group multiple>
          <v-subheader>Groups</v-subheader>
          <v-item v-for="group in exercise.group.split(',')" :key="group">
            <v-chip>
              {{ group }}
            </v-chip>
          </v-item>
        </v-item-group>
      </v-card-text>
      <v-card-actions>
        <v-btn
            color="grey"
            text
            :to="{ name: 'exercise', params: { id: exercise.id }}"
        >
          Edit
        </v-btn>
      </v-card-actions>
    </v-card>
</template>

<script>

export default {
  name: "Exercise",
  props: ["exercise"],
};
</script>

<style scoped></style>
