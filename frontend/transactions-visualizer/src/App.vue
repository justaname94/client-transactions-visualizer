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
      <template>
        <calendar-overlay :modal="modal" v-on:update:modal="updateModal" />
        <v-tabs>
          <v-tab>All Clients</v-tab>
          <v-tab>Client info</v-tab>
          <v-tab-item>
            <v-container>
              <all-clients-tab />
            </v-container>
          </v-tab-item>
          <v-tab-item>
            <v-container>
              <client-tab />
            </v-container>
          </v-tab-item>
        </v-tabs>
      </template>
    </v-main>
    <v-card height="100">
      <v-footer dark tile absolute class="font-weight-medium">
        <v-col class="text-center" cols="12">
          <strong>2020 - Roniel Valdez</strong>
        </v-col>
      </v-footer>
    </v-card>
  </v-app>
</template>

<script>
import CalendarOverlay from "@/components/CalendarOverlay";
import AllClientsTab from "@/components/AllClientsTab";
import ClientTab from "@/components/ClientTab";

export default {
  name: "App",

  components: {
    CalendarOverlay,
    AllClientsTab,
    ClientTab
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
