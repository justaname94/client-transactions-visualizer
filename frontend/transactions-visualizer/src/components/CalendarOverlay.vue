<template>
  <div>
    <v-dialog dark max-width="290px" persistent v-model="computedModal">
      <v-date-picker
        v-model="date"
        scrollable
        flat
        actions
        :max="new Date().toISOString().substr(0, 10)"
      >
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="computedModal = !computedModal"
            >Cancel</v-btn
          >
          <v-btn
            @click="loadDate"
            :disabled="date === null"
            text
            color="primary"
            >Load</v-btn
          >
        </v-card-actions>
      </v-date-picker>
    </v-dialog>
    <v-overlay :value="loadingOverlay">
      <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay>
  </div>
</template>

<script>
import EventBus from "../event-bus";
import Endpoints from "../services/EndpointsService";

export default {
  name: "CalendarOverlay",

  props: {
    modal: {
      type: Boolean,
      required: true
    }
  },

  data: () => ({
    loadingOverlay: false,
    alert: false,
    date: null
  }),

  methods: {
    async loadDate() {
      this.computedModal = false;
      this.loadingOverlay = true;
      const { data, success } = await Endpoints.loadDate(this.date);
      this.loadingOverlay = false;

      let payload;

      if (success) {
        payload = {
          alert: true,
          type: "success",
          message: "Date successfully loaded"
        };
      } else {
        payload = {
          alert: true,
          type: "error",
          message: data
        };
      }

      EventBus.$emit("alert", payload);
    }
  },

  computed: {
    computedModal: {
      get() {
        return this.modal;
      },
      set(value) {
        this.$emit("update:modal", value);
      }
    }
  }
};
</script>

<style scoped>
.v-alert {
  margin-bottom: 0;
}
</style>
