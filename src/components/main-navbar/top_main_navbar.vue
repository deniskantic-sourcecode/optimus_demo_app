<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" app>
      <v-list>
        <v-list-item v-for="(item, index) in items" :key="index">
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app color="#0096a9">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />

      <v-toolbar-title>OPTIMUS</v-toolbar-title>

      <v-spacer />

      <template v-if="$vuetify.display.mdAndUp">
        <v-btn icon @click="search">
          <v-icon>mdi-magnify</v-icon>
        </v-btn>

        <v-btn icon @click="filter">
          <v-icon>mdi-filter</v-icon>
        </v-btn>
      </template>
      <v-menu>
        <template #activator="{ on }">
          <v-btn color="primary" dark v-on="on">
            Dropdown
          </v-btn>
        </template>

        <v-list>
          <v-list-item
            v-for="(item, index) in items"
            :key="index"
          >
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
  </v-app>
</template>

<script>
  export default {
    data: () => ({
      drawer: false,
      location: null,
      locations: ['Location 1', 'Location 2', 'Location 3'], // Example locations
      items: [
        { title: 'Start', value: 'start' },
        { title: 'Bar', value: 'bar' },
        { title: 'Fizz', value: 'fizz' },
        { title: 'Buzz', value: 'buzz' },
      ],
    }),

    methods: {
      search () {
      // Handle search action
      },
      filter () {
      // Handle filter action
      },
    },
  }
</script>
