<template>
  <v-alert
    transition="scale-transition"
    v-model="alert"
    dense
    tile
    dismissible
    :type="type"
    >{{ message }}</v-alert
  >
</template>

<script>
import EventBus from "../event-bus";
import Capitalize from "lodash.capitalize";

export default {
  data: () => ({
    alert: false,
    type: "success",
    message: ""
  }),

  mounted() {
    /**
     * @typedef payload
     * @type {object}
     * @property {string} alert - wether show alert.
     * @property {string} type - alert type.
     * @property {number} message - alert message.
     */

    EventBus.$on("alert", payload => {
      this.alert = payload.alert;
      this.type = payload.type;
      this.message = Capitalize(payload.message);
    });
  }
};
</script>

<style scoped>
.v-alert {
  margin-bottom: 0;
}
</style>
