<template>
  <v-app>
    <v-app-bar dense app flat dark>
      <v-spacer></v-spacer>

      <div class="calendar">
        <v-btn target="_blank" @click="calendarModal = !calendarModal" text>
          <span class="mr-2">Load Date</span>
          <v-icon>mdi-calendar</v-icon>
        </v-btn>
      </div>
    </v-app-bar>
    <v-main>
      <div>
        <v-alert :value="alert" dense dismissible type="error">
          I'm a dense alert with the <strong>outlined</strong> prop and a
          <strong>type</strong> of error
        </v-alert>
        <v-dialog dark max-width="290px" persistent v-model="calendarModal">
          <v-date-picker v-model="date" scrollable flat actions>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                text
                color="primary"
                @click="calendarModal = !calendarModal"
                >Cancel</v-btn
              >
              <v-btn
                @click="loadingOverlay = !loadingOverlay"
                :disabled="date === null"
                text
                color="primary"
                >OK</v-btn
              >
            </v-card-actions>
          </v-date-picker>
        </v-dialog>
        <v-overlay :value="loadingOverlay">
          <v-progress-circular indeterminate size="64"></v-progress-circular>
        </v-overlay>
      </div>
    </v-main>
  </v-app>
</template>

<script>
// import HelloWorld from "./components/HelloWorld";

export default {
  name: "App",

  components: {
    // HelloWorld,
  },

  data: () => ({
    calendarModal: false,
    loadingOverlay: false,
    alert: false,
    date: null,
  }),

  methods: {
    loadDate() {
      this.loadingOverlay = true;
    },
  },

  watch: {
    loadingOverlay(val) {
      this.calendarModal = false;
      val &&
        setTimeout(() => {
          this.loadingOverlay = false;
          this.alert = true;
        }, 3000);
    },
  },
};
</script>
