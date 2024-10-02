<template>
  <v-card>
    <v-layout>
      <v-app-bar
        color="#0096a9"
        prominent
      >
        <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer" />

        <v-toolbar-title>OPTIMUS</v-toolbar-title>

        <v-spacer />

        <template v-if="$vuetify.display.mdAndUp">
          <v-btn icon="mdi-magnify" variant="text" />

          <v-btn icon="mdi-filter" variant="text" />
        </template>

        <div class="text-center">
          <v-select
            v-model="location"
            :items="locations"
            label="Location"
          />
          <v-menu :location="location">
            <template #activator="{ props }">
              <v-btn
                color="primary"
                dark
                v-bind="props"
              >
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
        </div>
      </v-app-bar>

    </v-layout>
  </v-card>
</template>
  <script>
  export default {
    data: () => ({
      drawer: false,
      group: null,
      items: [
        {
          title: 'Foo',
          value: 'foo',
        },
        {
          title: 'Bar',
          value: 'bar',
        },
        {
          title: 'Fizz',
          value: 'fizz',
        },
        {
          title: 'Buzz',
          value: 'buzz',
        },
      ],
    }),

    watch: {
      group () {
        this.drawer = false
      },
    },
  }
  </script>
