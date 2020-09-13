<template>
  <div>
    <v-alert v-model="alert" dense tile dismissible type="error">
      I'm a dense alert with the
      <strong>outlined</strong> prop and a
      <strong>type</strong> of error
    </v-alert>
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
          <v-btn text color="primary" @click="computedModal = !computedModal">Cancel</v-btn>
          <v-btn
            @click="loadingOverlay = !loadingOverlay"
            :disabled="date === null"
            text
            color="primary"
          >OK</v-btn>
        </v-card-actions>
      </v-date-picker>
    </v-dialog>
    <v-overlay :value="loadingOverlay">
      <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay>
  </div>
</template>

<script>
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

  watch: {
    loadingOverlay(val) {
      this.computedModal = false;
      val &&
        setTimeout(() => {
          this.loadingOverlay = false;
          this.alert = true;
        }, 3000);
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