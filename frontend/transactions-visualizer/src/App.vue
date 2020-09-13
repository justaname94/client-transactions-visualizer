<template>
  <v-app>
    <v-app-bar dense app flat dark>
      <v-toolbar-title>Client transactions visualizer</v-toolbar-title>

      <v-spacer></v-spacer>

      <div class="calendar">
        <v-btn target="_blank" @click="modal = !modal" text>
          <span class="mr-2">Load Date</span>
          <v-icon>mdi-calendar</v-icon>
        </v-btn>
      </div>
    </v-app-bar>
    <v-main>
      <calendar-overlay :modal="modal" v-on:update:modal="updateModal" />
    </v-main>
  </v-app>
</template>

<script>
import CalendarOverlay from "./components/CalendarOverlay";

export default {
  name: "App",

  components: {
    CalendarOverlay
  },

  data: () => ({
    modal: false
  }),

  methods: {
    loadDate() {
      this.loadingOverlay = true;
    },
    updateModal(value) {
      this.modal = value;
    }
  }
};
</script>
