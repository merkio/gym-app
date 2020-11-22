<template>
  <v-card
      max-width="700"
      class="mx-auto"
      outline>
    <v-img min-height="50" max-height="150" max-width="250" :src="exercise.image">IMAGE</v-img>

    <v-card-title>
      {{ exercise.title }}
    </v-card-title>
    <v-card-text>
      <v-row align="center"
             class="mx-0"
      >
        <v-textarea
            rows="1"
            multiline
            auto-grow
            :label="exercise.title"
            :value="exercise.description"
            hint="Full description of the exercise"/>
      </v-row>
    </v-card-text>
    <v-card-text>
      <v-chip-group column>
        <v-chip v-for="tag in exercise.tags.split(',')" :key="tag">
          {{ tag }}
        </v-chip>
      </v-chip-group>
      <v-chip-group column>
        <v-chip v-for="group in exercise.group.split(',')" :key="group">
          {{ group }}
        </v-chip>
      </v-chip-group>
    </v-card-text>
    <v-card-actions>
      <v-btn medium outlined text :to="exercise.videoLink">
        <v-icon medium>mdi-video-plus-outline</v-icon>
      </v-btn>
      <v-btn color="grey" text medium outlined>
        <v-icon medium>mdi-google-circles-group</v-icon>
      </v-btn>
      <v-btn color="grey" text medium outlined>
        <v-icon medium dark>mdi-tag-plus-outline</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import exercises from "@/store/exercises";

export default {
  name: "Exercise",
  data() {
    return {
      exercise: {},
    };
  },
  created() {
    this.initialize();
  },
  methods: {
    initialize() {
      const {id} = this.$route.params
      this.exercise = exercises.find(item => item.id === id)
    },
  },
};
</script>

<style scoped>
</style>
