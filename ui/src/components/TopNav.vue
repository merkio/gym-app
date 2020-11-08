<template>
  <div>
    <v-app-bar
        color="grey"
    >
      <v-toolbar-title>GYM</v-toolbar-title>

      <v-spacer/>

      <span class="hidden-sm-and-up">
        <v-btn @click.stop="drawer = !drawer">Menu</v-btn>
      </span>

      <v-toolbar-items class="hidden-xs-only">
        <v-btn v-for="item in items" :key="item.title" :to="item.link" text>
          <v-icon small left>fa-{{item.icon}}</v-icon>
          {{ item.title }}
        </v-btn>
      </v-toolbar-items>
    </v-app-bar>
    <v-navigation-drawer
        v-model="drawer"
        absolute
        temporary
        right
    >
      <v-list>
        <v-list-item
            v-for="item in items"
            :key="item.title"
            :to="item.link"
            link>
          <v-list-item-icon>
            <v-icon small>fa-{{ item.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>
              {{ item.title }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  name: "TopNav",
  data() {
    return {
      drawer: false,
      items: [
        {title: 'Home', link: '/', icon: 'home'},
        {title: 'Programs', link: 'programs', icon: 'address-book'},
        {title: 'Exercises', link: 'exercises', icon: 'list'},
        {title: 'Results', link: 'results', icon: 'book'}
      ]
    }
  }
}
</script>
